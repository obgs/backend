package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Players is the resolver for the players field.
func (r *queryResolver) Players(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PlayerWhereInput) (*ent.PlayerConnection, error) {
	return r.client.Player.Query().Paginate(ctx, after, first, before, last, ent.WithPlayerFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
