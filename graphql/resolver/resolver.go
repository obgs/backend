package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/open-boardgame-stats/backend/ent"
	"github.com/open-boardgame-stats/backend/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	config := generated.Config{
		Resolvers: &Resolver{client},
	}
	config.Directives.Authenticated = AuthenticatedDirective
	return generated.NewExecutableSchema(config)
}
