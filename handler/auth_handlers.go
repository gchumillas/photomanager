package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const authURL = "https://www.dropbox.com/oauth2/authorize"

// GetAuthURL gets the authentication URL.
func (env *Env) GetAuthURL(w http.ResponseWriter, r *http.Request, appKey string) {
	redirectURI := getParam(r, "redirect_uri", "")
	if len(redirectURI) == 0 {
		httpError(w, badParamsError)
		return
	}

	u, _ := url.Parse(authURL)
	q := u.Query()
	q.Add("redirect_uri", redirectURI)
	q.Add("client_id", appKey)
	q.Add("response_type", "code")
	u.RawQuery = q.Encode()

	json.NewEncoder(w).Encode(u.String())
}
