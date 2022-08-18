package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/open-boardgame-stats/backend/ent"
	"github.com/open-boardgame-stats/backend/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	client *ent.Client
	ctx    context.Context
	secret string
}

// NewAuthService returns a new AuthService
func NewAuthService(client *ent.Client, ctx context.Context, secret string) *AuthService {
	return &AuthService{client, ctx, secret}
}

func internalServerError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}

// create and sign access and refresh tokens
func (a *AuthService) generateTokens(w http.ResponseWriter, userId int) {
	now := time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": now.Add(24 * time.Hour).Unix(),
	})
	signedAccessToken, err := accessToken.SignedString([]byte(a.secret))
	if err != nil {
		internalServerError(w, fmt.Sprintf("failed to sign access token: %v", err))
		return
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": now.Add(60 * 24 * time.Hour).Unix(),
	})
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.secret))
	if err != nil {
		internalServerError(w, fmt.Sprintf("failed to sign refresh token: %v", err))
		return
	}

	// return the response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  signedAccessToken,
		"refresh_token": signedRefreshToken,
	})
}

// SignUp creates a new user with email and password from form data
func (a *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	// get form data
	email := r.FormValue("email")
	password := r.FormValue("password")

	// encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		internalServerError(w, fmt.Sprintf("failed to hash password: %v", err))
		return
	}

	// create the user
	u, err := a.client.User.Create().SetEmail(email).SetPassword(string(hashedPassword)).Save(a.ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			http.Error(w, fmt.Sprintf("user with email %s already exists", email), http.StatusConflict)
			return
		}
		internalServerError(w, fmt.Sprintf("failed to create user: %v", err))
		return
	}

	a.generateTokens(w, u.ID)
}

// SignIn authenticates a user with email and password from form data
func (a *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	// get form data
	email := r.FormValue("email")
	password := r.FormValue("password")

	// get the user
	u, err := a.client.User.Query().Where(user.EmailEQ(email)).Only(a.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		internalServerError(w, fmt.Sprintf("failed to get user: %v", err))
		return
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	a.generateTokens(w, u.ID)
}
