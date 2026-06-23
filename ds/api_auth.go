package ds

import (
	"net/http"
	"net/url"
)

func (r *routes) OAuth2Token() (string, string) {
	return http.MethodPost, "/oauth2/token"
}

func (r *routes) OAuth2Authorize(params url.Values) string {
	return "https://discord.com/oauth2/authorize?" + params.Encode()
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
