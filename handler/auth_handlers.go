package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gchumillas/photomanager/dbox"
	"github.com/gchumillas/photomanager/manager"
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

	token, accountID := dbox.GetAuthToken(uri, code, appKey, appSecret)

	// TODO: checkout this piece of code
	u := manager.NewUser()
	u.AccountID = accountID
	if !u.ReadUserByAccountID(env.DB) {
		u.AccessToken = token
		u.CreateUser(env.DB)
	} else {
		u.AccessToken = token
		u.UpdateUser(env.DB)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"uid":   accountID,
		"token": token,
	})
}
