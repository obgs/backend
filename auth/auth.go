package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func invalidRefreshToken(w http.ResponseWriter) {
	http.Error(w, "invalid refresh token", http.StatusUnauthorized)
}

// create and sign access and refresh tokens
func (a *AuthService) generateTokens(w http.ResponseWriter, userId uuid.UUID, statusCode int) {
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
	w.WriteHeader(statusCode)
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

	a.generateTokens(w, u.ID, http.StatusCreated)
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

	a.generateTokens(w, u.ID, http.StatusOK)
}

// Refresh refreshes the access and refresh token
func (a *AuthService) Refresh(w http.ResponseWriter, r *http.Request) {
	// get the refresh token
	refreshToken := r.FormValue("refresh_token")

	// parse the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		invalidRefreshToken(w)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			invalidRefreshToken(w)
			return
		}
		a.generateTokens(w, id, http.StatusOK)
	} else {
		invalidRefreshToken(w)
	}
}
