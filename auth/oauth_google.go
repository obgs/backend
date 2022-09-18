package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const oAuthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"

type oAuthGoogleUserData struct {
	Email string `json:"email"`
}

func (d *oAuthGoogleUserData) GetEmail() string {
	return d.Email
}

func (a *AuthService) OAuthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oAuthState := a.generateOAuthState()
	redirectUrl := a.oAuthConfig.google.AuthCodeURL(oAuthState)
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (a *AuthService) OAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// get state
	oAuthState := r.FormValue("state")
	defer a.clearOAuthState(oAuthState)

	// validate state
	t, ok := a.oAuthStates[oAuthState]
	if !ok || t.Add(oAuthStateTTL).Before(time.Now()) {
		internalServerError(w, "invalid oauth state")
		return
	}

	data, err := a.oAuthGoogleGetData(r.FormValue("code"))
	if err != nil {
		internalServerError(w, fmt.Sprintf("failed to get user data: %v", err))
		return
	}

	a.oAuthSignUp(data, w)
}

func (a *AuthService) oAuthGoogleGetData(code string) (*oAuthGoogleUserData, error) {
	token, err := a.oAuthConfig.google.Exchange(a.ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %v", err)
	}
	response, err := http.Get(fmt.Sprintf("%s?access_token=%s", oAuthGoogleUrlAPI, token.AccessToken))
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
