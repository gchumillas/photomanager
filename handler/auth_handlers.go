package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// GetAuthURL gets the authentication URL.
func (env *Env) GetAuthURL(w http.ResponseWriter, r *http.Request, redirectURI, appKey string) {
	u, _ := url.Parse("https://www.dropbox.com/oauth2/authorize")
	q := u.Query()
	q.Add("redirect_uri", redirectURI)
	q.Add("client_id", appKey)
	q.Add("response_type", "token")
	u.RawQuery = q.Encode()

	json.NewEncoder(w).Encode(u.String())
}
