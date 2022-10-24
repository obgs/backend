package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
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
		ctx := context.Background()
		id, err := guidgql.UnmarshalGUID(claims["id"].(string))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		// get the user from the database
		user, err := a.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		// set the user in the context
		ctx = context.WithValue(r.Context(), userCtxKey, user)
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
