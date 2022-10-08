package resolver

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/open-boardgame-stats/backend/internal/auth"
)

func AuthenticatedDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user, err := auth.UserFromContext(ctx)
	if (err != nil) || (user == nil) {
		return nil, fmt.Errorf("not authenticated")
	}
	return next(ctx)
}
