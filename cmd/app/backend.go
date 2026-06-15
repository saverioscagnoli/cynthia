package main

import (
	"cynthia/ds"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	DiscordID ds.Snowflake `json:"discord_id"`
	Username  string       `json:"username"`
	jwt.RegisteredClaims
}

func (a *App) generateJWT(dsID ds.Snowflake, username string) (string, error) {
	claims := Claims{
		DiscordID: dsID,
		Username:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *App) parseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token.Claims.(*Claims), nil
}

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		claims, err := a.parseJWT(tokenStr)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-Discord-ID", claims.DiscordID)
		r.Header.Set("X-Username", claims.Username)

		next.ServeHTTP(w, r)
	})
}

func (a *App) SetupMuxHandler() {
	a.mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		if err := a.DB.Ping(a.DB.Context); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error("Healthcheck", "err", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	a.mux.HandleFunc("GET /auth/login", func(w http.ResponseWriter, r *http.Request) {
		params := url.Values{
			"client_id":     {os.Getenv("OAUTH2_CLIENT_ID")},
			"redirect_uri":  {os.Getenv("OAUTH2_REDIRECT_URI")},
			"response_type": {"code"},
			"scope":         {"identify"},
		}

		http.Redirect(w, r, "https://discord.com/oauth2/authorize?"+params.Encode(), http.StatusTemporaryRedirect)
	})

	a.mux.HandleFunc("GET /auth/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
			return
		}

		res, err := http.PostForm("https://discord.com/api/oauth2/token", url.Values{
			"client_id":     {os.Getenv("OAUTH2_CLIENT_ID")},
			"client_secret": {os.Getenv("OAUTH2_SECRET")},
			"grant_type":    {"authorization_code"},
			"code":          {code},
			"redirect_uri":  {os.Getenv("OAUTH2_REDIRECT_URI")},
		})

		if err != nil {
			slog.Error("Token exchange", "err", err)
			http.Error(w, "token exchange failed", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()

		var tokenData struct {
			AccessToken string `json:"access_token"`
		}

		json.NewDecoder(res.Body).Decode(&tokenData)

		method, endpoint := ds.Routes.GetMe()
		req, _ := http.NewRequest(method, ds.ApiURL+endpoint, nil)
		req.Header.Set("Authorization", "Bearer "+tokenData.AccessToken)
		userRes, err := http.DefaultClient.Do(req)

		if err != nil {
			slog.Error("Fetch Discord user", "err", err)
			http.Error(w, "failed to fetch user", http.StatusInternalServerError)
		}

		defer userRes.Body.Close()

		var dsUser ds.User
		json.NewDecoder(userRes.Body).Decode(&dsUser)

		_, err = a.DB.Exec(a.DB.Context, `
            INSERT INTO users (discord_id, username, avatar)
            VALUES ($1, $2, $3)
            ON CONFLICT (discord_id) DO UPDATE
            SET username = EXCLUDED.username, avatar = EXCLUDED.avatar
        `, dsUser.ID, dsUser.Username, dsUser.Avatar)

		if err != nil {
			slog.Error("Upsert user", "err", err)
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		token, err := a.generateJWT(dsUser.ID, dsUser.Username)

		if err != nil {
			http.Error(w, "jwt error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL")+"/trainers", token), http.StatusTemporaryRedirect)
	})

	protected := http.NewServeMux()
	protected.HandleFunc("GET /user/me", func(w http.ResponseWriter, r *http.Request) {
		dsID := r.Header.Get("X-Discord-ID")
		row := a.DB.QueryRow(a.DB.Context, `SELECT discord_id, username, avatar, trainer_id FROM users WHERE discord_id = $1`, dsID)

		var u struct {
			DiscordID string `json:"discordId"`
			Username  string `json:"username"`
			Avatar    string `json:"avatar"`
			TrainerID int    `json:"trainerId"`
		}

		if err := row.Scan(&u.DiscordID, &u.Username, &u.Avatar, &u.TrainerID); err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	})

	protected.HandleFunc("PUT /user/trainer", func(w http.ResponseWriter, r *http.Request) {
		dsID := r.Header.Get("X-Discord-ID")

		var body struct {
			TrainerID int `json:"trainer_id"`
		}

		json.NewDecoder(r.Body).Decode(&body)

		_, err := a.DB.Exec(a.DB.Context, `UPDATE users SET trainer_id = $1 WHERE discord_id = $2`, body.TrainerID, dsID)

		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	a.mux.Handle("/user/", a.authMiddleware(protected))
}
