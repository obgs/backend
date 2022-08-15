package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/graphql/model"
)

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.PlayerInput) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MyPlayers(ctx context.Context) ([]*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}