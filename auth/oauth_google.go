package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oAuthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"

func newOAuthGoogleConfig(googleClientID, googleClientSecret, serverHost, serverPort string) oauth2.Config {
	return oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  fmt.Sprintf("http://%s:%s/auth/google/callback", serverHost, serverPort),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

type oAuthGoogleUserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (d *oAuthGoogleUserData) GetEmail() string {
	return d.Email
}

func (d *oAuthGoogleUserData) GetName() string {
	return d.Name
}

func (a *AuthService) OAuthGoogleSignIn(w http.ResponseWriter, r *http.Request) {
	data, err := a.oAuthGoogleGetData(r.FormValue("token"))
	if err != nil {
		internalServerError(w, fmt.Sprintf("failed to get user data: %v", err))
		return
	}

	a.oAuthSignUp(data, w)
}

func (a *AuthService) oAuthGoogleGetData(token string) (*oAuthGoogleUserData, error) {
	response, err := http.Get(fmt.Sprintf("%s?access_token=%s", oAuthGoogleUrlAPI, token))
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %v", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %v", err)
	}

	userData := &oAuthGoogleUserData{}
	jsonErr := json.Unmarshal(contents, userData)
	if jsonErr != nil {
		return nil, fmt.Errorf("failed to parse user info: %v", err)
	}

	return userData, nil
}
