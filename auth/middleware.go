package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/ent"
	"github.com/open-boardgame-stats/backend/ent/user"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func (a *AuthService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the access token from the request
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			next.ServeHTTP(w, r)
			return
		}
		// parse the access token
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.secret), nil
		})
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		// get the user from the database
		user, err := a.client.User.Query().Where(user.IDEQ(uuid.MustParse(claims["id"].(string)))).Only(a.ctx)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		// set the user in the context
		ctx := context.WithValue(r.Context(), userCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserFromContext finds the user from the context
func UserFromContext(ctx context.Context) (*ent.User, error) {
	raw, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	return raw, nil
}