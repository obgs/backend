package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// CreatePlayer is the resolver for the createPlayer field.
func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.CreatePlayerInput) (*ent.Player, error) {
	supervisor, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return ent.FromContext(ctx).Player.Create().SetName(input.Name).AddSupervisors(supervisor).Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
