package api

import (
	"camilla/constants"
	"camilla/ds"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

type Claims struct {
	UserID   ds.Snowflake `json:"user_id"`
	Username string       `json:"username"`
	jwt.RegisteredClaims
}

func (rt *_router) genJWT(userID ds.Snowflake, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.TokenMaxDuration)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(rt.jwtSecret))
}

func (rt *_router) parseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(rt.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token.Claims.(*Claims), nil
}

type contextKey string

const claimsKey contextKey = "claims"

func (rt *_router) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims, err := rt.parseJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rt *_router) AuthLogin(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	state, err := generateState()

	if err != nil {
		http.Error(w, "Failed to generate state", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HttpOnly: true,
		Secure:   rt.isProd,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   300,
	})

	params := url.Values{
		"client_id":     {rt.oauth2ClientID},
		"redirect_uri":  {rt.oauth2RedirectURI},
		"response_type": {"code"},
		"scope":         {"identify"},
		"state":         {state},
	}

	http.Redirect(w, r, ds.Routes.OAuth2Authorize(params), http.StatusTemporaryRedirect)
}

func (rt *_router) AuthLogout(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) AuthCallback(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	stateCookie, err := r.Cookie("oauth_state")

	if err != nil || r.URL.Query().Get("state") != stateCookie.Value {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "oauth_state", Value: "", MaxAge: -1, Path: "/"})

	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	_, endpoint := ds.Routes.OAuth2Token()
	res, err := http.PostForm(ds.ApiURL+endpoint, url.Values{
		"client_id":     {rt.oauth2ClientID},
		"client_secret": {rt.oauth2Secret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {rt.oauth2RedirectURI},
	})

	if err != nil {
		rt.logger.Error("Token exchange failed", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	var accessTokenData ds.AccessTokenResponse

	err = json.NewDecoder(res.Body).Decode(&accessTokenData)

	if err != nil {
		rt.logger.Error("Token exchange failed", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	method, endpoint := ds.Routes.GetMe()
	req, err := http.NewRequest(method, ds.ApiURL+endpoint, nil)

	if err != nil {
		rt.logger.Error("Failed to create request", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessTokenData.AccessToken)

	userRes, err := http.DefaultClient.Do(req)

	if err != nil {
		rt.logger.Error("Failed to get user info", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	defer userRes.Body.Close()

	var user ds.User
	err = json.NewDecoder(userRes.Body).Decode(&user)

	if err != nil {
		rt.logger.Error("Failed to decode user info", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	err = rt.db.UpsertUser(&user, r.Context())

	if err != nil {
		rt.logger.Error("Failed to upsert user", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	token, err := rt.genJWT(user.ID, user.Username)

	if err != nil {
		rt.logger.Error("Failed to generate token", "err", err)
		http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		HttpOnly: true,
		Secure:   rt.isProd,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   int(constants.TokenMaxDuration),
	})

	http.Redirect(w, r, rt.frontendURL, http.StatusTemporaryRedirect)
}
