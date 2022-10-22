package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// IsMember is the resolver for the isMember field.
func (r *groupResolver) IsMember(ctx context.Context, obj *ent.Group) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, nil
	}

	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(obj.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	return membership != nil, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*ent.Group, error) {
	// create group settings
	s, err := r.client.GroupSettings.Create().
		SetVisibility(input.Visibility).
		SetJoinPolicy(input.JoinPolicy).
		SetMinimumRoleToInvite(*input.MinimumRoleToInvite).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// create the group
	g, err := r.client.Group.Create().
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
	_, err = r.client.GroupMembership.Create().
		SetGroup(g).
		SetUser(u).
		SetRole(enums.RoleOwner).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
