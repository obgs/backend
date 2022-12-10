package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id guidgql.GUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithNodeType(getNodeType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []*guidgql.GUID) ([]ent.Noder, error) {
	actual := make([]guidgql.GUID, len(ids))
	for i, id := range ids {
		actual[i] = *id
	}

	return r.client.Noders(ctx, actual, ent.WithNodeType(getNodeType))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	u, _ := auth.UserFromContext(ctx)

	return r.client.Group.Query().Paginate(ctx, after, first, before, last,
		ent.WithGroupFilter(where.Filter),
		ent.WithGroupFilter(func(q *ent.GroupQuery) (*ent.GroupQuery, error) {
			// we need to show only public groups, or, if the request is authenticated, the groups the user is a member of
			p := group.HasSettingsWith(
				groupsettings.VisibilityEQ(groupsettings.VisibilityPublic),
			)
			if u != nil {
				p = group.Or(p, group.HasMembersWith(groupmembership.HasUserWith(user.ID(u.ID))))
			}

			return q.Where(p), nil
		}),
	)
}

// Players is the resolver for the players field.
func (r *queryResolver) Players(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PlayerWhereInput) (*ent.PlayerConnection, error) {
	return r.client.Player.Query().Paginate(ctx, after, first, before, last, ent.WithPlayerFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last)
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type gameResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
