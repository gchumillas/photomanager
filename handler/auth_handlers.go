package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gchumillas/photomanager/dbox"
	"github.com/gchumillas/photomanager/manager"
)

const (
	authURL  = "https://api.dropbox.com/oauth2/authorize"
	tokenURL = "https://api.dropbox.com/oauth2/token"
)

// GetAuthURL gets the authentication URL.
func (env *Env) GetAuthURL(w http.ResponseWriter, r *http.Request, appKey, redirectURI string) {
	uri := getParam(r, "redirect_uri", redirectURI)
	if len(uri) == 0 {
		httpError(w, badParamsError)
		return
	}

	u := dbox.GetAuthURL(uri, appKey)
	json.NewEncoder(w).Encode(u)
}

// Login logs into the system.
func (env *Env) Login(w http.ResponseWriter, r *http.Request, appKey, appSecret, redirectURI string) {
	code := getParam(r, "code", "")
	uri := getParam(r, "redirect_uri", redirectURI)
	if len(code) == 0 || len(uri) == 0 {
		httpError(w, badParamsError)
		return
	}

	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", uri)
	body := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", tokenURL, body)
	req.SetBasicAuth(appKey, appSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		httpError(w, badParamsError)
		return
	}

	var target struct {
		AccessToken string `json:"access_token"`
		AccountID   string `json:"account_id"`
	}
	json.NewDecoder(resp.Body).Decode(&target)

	u := manager.NewUser()
	u.AccountID = target.AccountID
	if !u.ReadUserByAccountID(env.DB) {
		u.AccessToken = target.AccessToken
		u.CreateUser(env.DB)
	} else {
		u.AccessToken = target.AccessToken
		u.UpdateUser(env.DB)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"uid":   target.AccountID,
		"token": target.AccessToken,
	})
}
