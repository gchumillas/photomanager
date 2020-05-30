package dbox

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	authURL   = "https://www.dropbox.com/oauth2/authorize"
	tokenURL  = "https://api.dropboxapi.com/oauth2/token"
	uploadURL = "https://content.dropboxapi.com/2/files/upload"
)

// GetAuthURL gets the authentication URL.
func GetAuthURL(redirectURI, appKey string) string {
	u, _ := url.Parse(authURL)
	q := u.Query()
	q.Add("redirect_uri", redirectURI)
	q.Add("client_id", appKey)
	q.Add("response_type", "code")
	u.RawQuery = q.Encode()

	return u.String()
}

// GetAuthToken gets a token and the user's account ID.
func GetAuthToken(redirectURI, code, appKey, appSecret string) (token, accountID string) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectURI)
	body := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", tokenURL, body)
	req.SetBasicAuth(appKey, appSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		panic(errors.New(resp.Status))
	}

	var target struct {
		AccessToken string `json:"access_token"`
		AccountID   string `json:"account_id"`
	}
	json.NewDecoder(resp.Body).Decode(&target)

	return target.AccessToken, target.AccountID
}

// UploadFile uploads a file to the user'x box.
func UploadFile(token string, file io.Reader, path string) (id string) {
	payload, _ := json.Marshal(map[string]interface{}{
		"path":       path,
		"autorename": true,
	})

	req, err := http.NewRequest("POST", uploadURL, file)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Dropbox-Api-Arg", string(payload))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		panic(resp.Status)
	}

	var target struct {
		ID string `json:"id"`
	}
	json.NewDecoder(resp.Body).Decode(&target)

	return target.ID
}
