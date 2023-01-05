package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// CreateMatch is the resolver for the createMatch field.
func (r *mutationResolver) CreateMatch(ctx context.Context, input model.CreateMatchInput) (*ent.Match, error) {
	client := ent.FromContext(ctx)

	playerIds := make([]guidgql.GUID, len(input.PlayerIds))
	for i, player := range input.PlayerIds {
		playerIds[i] = *player
	}

	// add checks for the current user to have access to players
	m, err := client.Match.Create().
		SetGameID(input.GameID).
		AddPlayerIDs(playerIds...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	numericalStats := make([]*ent.NumericalStatCreate, len(input.NumericalStats))
	for i, stat := range input.NumericalStats {
		err = checkThatPlayerIsPresent(stat.PlayerID, input.PlayerIds)
		if err != nil {
			return nil, err
		}

		numericalStats[i] = client.NumericalStat.Create().
			SetPlayerID(stat.PlayerID).
			SetValue(stat.Value).
			SetNumericalStatDescriptionID(stat.StatID).
			SetMatch(m)
	}

	enumStats := make([]*ent.EnumStatCreate, len(input.EnumStats))
	for i, stat := range input.EnumStats {
		err = checkThatPlayerIsPresent(stat.PlayerID, input.PlayerIds)
		if err != nil {
			return nil, err
		}

		enumStats[i] = client.EnumStat.Create().
			SetPlayerID(stat.PlayerID).
			SetEnumStatDescriptionID(stat.StatID).
			SetValue(stat.Value).
			SetMatch(m)
	}

	_, err = client.NumericalStat.CreateBulk(numericalStats...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return m, nil
}
