package dbox

import "net/url"

const (
	authURL  = "https://api.dropbox.com/oauth2/authorize"
	tokenURL = "https://api.dropbox.com/oauth2/token"
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
