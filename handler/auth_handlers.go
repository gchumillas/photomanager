package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAuthURL gets the authentication URL.
func (env *Env) GetAuthURL(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(
		"https://www.dropbox.com/oauth2/authorize?redirect_uri=%s&client_id=%s&response_type=token",
		"http://localhost:8080/v1/auth/login",
		env.DropboxKey,
	)
	json.NewEncoder(w).Encode(url)
}
