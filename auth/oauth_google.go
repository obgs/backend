package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/open-boardgame-stats/backend/ent"
	"github.com/open-boardgame-stats/backend/ent/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	oAuthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"
	oAuthStateTTL     = 20 * time.Minute
)

var oAuthGoogleConfig = oauth2.Config{
	ClientID:     "43480745913-1lhqt43oe2gir7fqdv5058m5l3q1hl2p.apps.googleusercontent.com",
	ClientSecret: "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",
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:8881/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
}

type oAuthGoogleUserData struct {
	Email string `json:"email"`
}

func (a *AuthService) OAuthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oAuthState := a.generateOAuthState()
	fmt.Println(oAuthState)
	redirectUrl := oAuthGoogleConfig.AuthCodeURL(oAuthState)
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (a *AuthService) OAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// get state
	oAuthState := r.FormValue("state")

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

	a.oAuthGoogleSignUp(data, w)
}

func (a *AuthService) oAuthGoogleGetData(code string) (*oAuthGoogleUserData, error) {
	token, err := oAuthGoogleConfig.Exchange(a.ctx, code)
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

func (a *AuthService) oAuthGoogleSignUp(data *oAuthGoogleUserData, w http.ResponseWriter) {
	var u *ent.User
	u, findErr := a.client.User.Query().Where(user.EmailEQ(data.Email)).Only(a.ctx)
	if findErr != nil && !ent.IsNotFound(findErr) {
		internalServerError(w, fmt.Sprintf("failed to find user: %v", findErr))
		return
	}
	if ent.IsNotFound(findErr) {
		newU, createErr := a.client.User.Create().SetEmail(data.Email).Save(a.ctx)
		if createErr != nil {
			internalServerError(w, fmt.Sprintf("failed to create user: %v", createErr))
			return
		}
		u = newU
	}

	a.generateTokens(w, u.ID)
}
