package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

type AuthService struct {
	client      *ent.Client
	secret      string
	oAuthConfig *oAuthConfig
}

type oAuthConfig struct {
	google oauth2.Config
}

func NewOAuthConfig(serverHost, serverPort, googleClientID, googleClientSecret string) *oAuthConfig {
	return &oAuthConfig{
		google: newOAuthGoogleConfig(googleClientID, googleClientSecret, serverHost, serverPort),
	}
}

type oAuthData interface {
	GetEmail() string
	GetName() string
}

// NewAuthService returns a new AuthService
func NewAuthService(client *ent.Client, ctx context.Context, secret string, oAuthConfig *oAuthConfig) *AuthService {
	return &AuthService{
		client,
		secret,
		oAuthConfig,
	}
}

func internalServerError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}

func invalidRefreshToken(w http.ResponseWriter) {
	http.Error(w, "invalid refresh token", http.StatusUnauthorized)
}

const DAY = 24 * time.Hour

const RANDOM_PASSWORD_LENGTH = 32

// create and sign access and refresh tokens
func (a *AuthService) generateTokens(w http.ResponseWriter, userId uuid.UUID, statusCode int) {
	now := time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": now.Add(DAY).Unix(),
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

func (a *AuthService) createUser(ctx context.Context, email, password string) (*ent.User, error) {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// create the user
	u, err := a.client.User.Create().SetEmail(email).SetPassword(string(hashedPassword)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// create the user's player
	_, err = a.client.Player.Create().SetOwner(u).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create the user's player: %v", err)
	}

	return u, nil
}

// SignUp creates a new user with email and password from form data
func (a *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	// get form data
	email := r.FormValue("email")
	password := r.FormValue("password")

	u, err := a.createUser(context.Background(), email, password)
	if err != nil {
		internalServerError(w, err.Error())
		return
	}

	a.generateTokens(w, u.ID, http.StatusCreated)
}

// SignIn authenticates a user with email and password from form data
func (a *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	// get form data
	email := r.FormValue("email")
	password := r.FormValue("password")

	ctx := context.Background()
	// get the user
	u, err := a.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
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

func (a *AuthService) oAuthSignUp(data oAuthData, w http.ResponseWriter) {
	email := data.GetEmail()

	ctx := context.Background()
	u, findErr := a.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if ent.IsNotFound(findErr) {
		newU, err := a.createUser(ctx, email, data.GetName())
		if err != nil {
			internalServerError(w, err.Error())
			return
		}

		a.generateTokens(w, newU.ID, http.StatusCreated)

		return
	}

	if findErr != nil {
		internalServerError(w, fmt.Sprintf("failed to find user: %v", findErr))
		return
	}

	a.generateTokens(w, u.ID, http.StatusOK)
}

func (a *AuthService) randomPassword(n int) string {
	b := make([]byte, n)
	rand.Read(b)

	return hex.EncodeToString(b)
}

// encrypt the password
func (a *AuthService) encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
