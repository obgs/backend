package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*ent.Group, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create group settings
	s, err := tx.GroupSettings.Create().
		SetVisibility(input.Visibility).
		SetJoinPolicy(input.JoinPolicy).
		SetMinimumRoleToInvite(*input.MinimumRoleToInvite).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// create the group
	g, err := tx.Group.Create().
		SetSettings(s).
		SetName(input.Name).
		SetDescription(*input.Description).
		SetLogoURL(input.LogoURL).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// add the user as the owner
	_, err = tx.GroupMembership.Create().
		SetGroup(g).
		SetUser(u).
		SetRole(enums.RoleOwner).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return g, tx.Commit()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
