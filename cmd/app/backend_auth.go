package main

import (
	"cynthia/ds"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	DiscordID ds.Snowflake `json:"discord_id"`
	Username  string       `json:"username"`
	jwt.RegisteredClaims
}

func (b *backend) genJWT(dsID ds.Snowflake, username string) (string, error) {
	claims := Claims{
		DiscordID: dsID,
		Username:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(b.jwtSecret))
}

func (b *backend) parseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(b.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token.Claims.(*Claims), nil

}

func (b *backend) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")

		if h == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(h, "Bearer ")
		claims, err := b.parseJWT(tokenStr)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-Discord-ID", claims.DiscordID)
		r.Header.Set("X-Username", claims.Username)

		next.ServeHTTP(w, r)
	})
}

func (b *backend) AuthLogin(w http.ResponseWriter, r *http.Request) {
	params := url.Values{
		"client_id":     {b.oauth2ClientID},
		"redirect_uri":  {b.oauth2RedirectURI},
		"response_type": {"code"},
		"scope":         {"identify"},
	}

	http.Redirect(w, r, ds.Routes.OAuth2Authorize(params), http.StatusTemporaryRedirect)
}

func (b *backend) AuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	_, endpoint := ds.Routes.OAuth2Token()
	res, err := http.PostForm(ds.ApiURL+endpoint, url.Values{
		"client_id":     {b.oauth2ClientID},
		"client_secret": {b.oauth2Secret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {b.oauth2RedirectURI},
	})

	if err != nil {
		slog.Error("Token exchange", "err", err)
		http.Error(w, "token exchange failed", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	var accessTokenRes ds.AccessTokenResponse

	err = json.NewDecoder(res.Body).Decode(&accessTokenRes)

	if err != nil {
		slog.Error("json decode error", "err", err)
		http.Error(w, "json decode error", http.StatusInternalServerError)
		return
	}

	slog.Debug("Access Token data", "res", accessTokenRes)

	method, endpoint := ds.Routes.GetMe()
	req, _ := http.NewRequest(method, ds.ApiURL+endpoint, nil)

	req.Header.Set("Authorization", "Bearer "+accessTokenRes.AccessToken)

	userRes, err := http.DefaultClient.Do(req)

	if err != nil {
		slog.Error("Failed to fetch Discord user", "err", err)
		http.Error(w, "failed to fetch user", http.StatusInternalServerError)
		return
	}

	defer userRes.Body.Close()

	var user ds.User
	err = json.NewDecoder(userRes.Body).Decode(&user)

	if err != nil {
		slog.Error("json decode error", "err", err)
		http.Error(w, "json decode error", http.StatusInternalServerError)
		return
	}

	err = b.db.UpsertUser(&user, r.Context())

	if err != nil {
		slog.Error("Upsert user", "err", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	token, err := b.genJWT(user.ID, user.Username)

	if err != nil {
		http.Error(w, "jwt error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s?token=%s", b.frontendURL, token), http.StatusTemporaryRedirect)
}
