package resolver

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

func createGroup(ctx context.Context, client *ent.Client, input model.CreateOrUpdateGroupInput) (*ent.Group, error) {
	var s *ent.GroupSettings
	s, err := client.GroupSettings.Create().
		SetVisibility(input.Settings.Visibility).
		SetJoinPolicy(input.Settings.JoinPolicy).
		SetMinimumRoleToInvite(*input.Settings.MinimumRoleToInvite).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	g, err := client.Group.Create().
		SetName(input.Name).
		SetLogoURL(input.LogoURL).
		SetDescription(*input.Description).
		SetSettings(s).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	u, _ := auth.UserFromContext(ctx)
	_, err = client.GroupMembership.Create().
		SetUser(u).
		SetGroup(g).
		SetRole(enums.RoleOwner).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func updateGroup(ctx context.Context, client *ent.Client, input model.CreateOrUpdateGroupInput) (*ent.Group, error) {
	err := client.GroupSettings.Update().
		Where(groupsettings.HasGroupWith(group.ID(*input.ID))).
		SetVisibility(input.Settings.Visibility).
		SetJoinPolicy(input.Settings.JoinPolicy).
		SetMinimumRoleToInvite(*input.Settings.MinimumRoleToInvite).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = client.Group.Update().
		Where(group.ID(*input.ID)).
		SetName(input.Name).
		SetLogoURL(input.LogoURL).
		SetDescription(*input.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return client.Group.Query().Where(group.ID(*input.ID)).Only(ctx)
}
