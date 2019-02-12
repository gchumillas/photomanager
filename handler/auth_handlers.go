package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// GetAuthURL gets the authentication URL.
func (env *Env) GetAuthURL(
	w http.ResponseWriter, r *http.Request, authURL, redirectURI, appKey string,
) {
	u, _ := url.Parse(authURL)
	q := u.Query()
	q.Add("redirect_uri", redirectURI)
	q.Add("client_id", appKey)
	q.Add("response_type", "code")
	u.RawQuery = q.Encode()

	json.NewEncoder(w).Encode(u.String())
}
