package resolver

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/obgs/backend/internal/ent"
	"github.com/obgs/backend/internal/filestorage"
	"github.com/obgs/backend/internal/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client      *ent.Client
	filestorage *filestorage.FileStorageService
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, filestorage *filestorage.FileStorageService) graphql.ExecutableSchema {
	config := generated.Config{
		Resolvers: &Resolver{
			client,
			filestorage,
		},
	}
	config.Directives.Authenticated = AuthenticatedDirective

	return generated.NewExecutableSchema(config)
}
