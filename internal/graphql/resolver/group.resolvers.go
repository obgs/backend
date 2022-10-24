package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// Role is the resolver for the role field.
func (r *groupResolver) Role(ctx context.Context, obj *ent.Group) (*enums.Role, error) {
	u, _ := auth.UserFromContext(ctx)
	if u == nil {
		return nil, nil
	}

	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(obj.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if membership == nil {
		return nil, nil
	}

	return &membership.Role, nil
}

// Applied is the resolver for the applied field.
func (r *groupResolver) Applied(ctx context.Context, obj *ent.Group) (*bool, error) {
	res := false
	u, _ := auth.UserFromContext(ctx)
	if u == nil {
		return &res, nil
	}

	a, err := r.client.GroupMembershipApplication.Query().Where(
		groupmembershipapplication.HasUserWith(user.ID(u.ID)),
		groupmembershipapplication.HasGroupWith(group.ID(obj.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if a != nil {
		res = true
	}

	return &res, nil
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

// JoinGroup is the resolver for the joinGroup field.
func (r *mutationResolver) JoinGroup(ctx context.Context, groupID guidgql.GUID) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, err
	}

	g, err := r.client.Group.Query().Where(group.ID(groupID)).Only(ctx)
	if err != nil {
		return false, err
	}

	// check if the user is already a member
	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}
	if membership != nil {
		return false, fmt.Errorf("user is already a member")
	}

	// check group's join policy
	s, err := g.QuerySettings().Where(
		groupsettings.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil {
		return false, err
	}
	if s.JoinPolicy != groupsettings.JoinPolicyOpen {
		return false, fmt.Errorf("group is not open for joining")
	}

	// create the membership
	_, err = r.client.GroupMembership.Create().
		SetGroup(g).
		SetUser(u).
		SetRole(enums.RoleMember).
		Save(ctx)

	return err == nil, err
}

// ApplyToGroup is the resolver for the applyToGroup field.
func (r *mutationResolver) ApplyToGroup(ctx context.Context, input model.GroupApplicationInput) (*ent.GroupMembershipApplication, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	g, err := r.client.Group.Query().Where(group.ID(input.GroupID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// check if there is already an application
	application, err := r.client.GroupMembershipApplication.Query().Where(
		groupmembershipapplication.HasUserWith(user.ID(u.ID)),
		groupmembershipapplication.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if application != nil {
		return nil, fmt.Errorf("user already applied to this group")
	}

	// check if the user is already a member
	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if membership != nil {
		return nil, fmt.Errorf("user is already a member")
	}

	// check group's join policy
	s, err := g.QuerySettings().Where(
		groupsettings.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	if s.JoinPolicy != groupsettings.JoinPolicyApplicationOnly && s.JoinPolicy != groupsettings.JoinPolicyInviteOrApplication {
		return nil, fmt.Errorf("group is not open for applications")
	}

	// create the application
	a, err := r.client.GroupMembershipApplication.Create().
		SetMessage(*input.Message).
		AddGroup(g).
		AddUser(u).
		Save(ctx)

	return a, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
