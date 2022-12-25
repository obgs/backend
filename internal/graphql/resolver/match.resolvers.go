package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// CreateMatch is the resolver for the createMatch field.
func (r *mutationResolver) CreateMatch(ctx context.Context, match model.CreateMatchInput) (*ent.Match, error) {
	playerIds := make([]guidgql.GUID, len(match.PlayerIds))
	for i, player := range match.PlayerIds {
		playerIds[i] = *player
	}

	// TODO: add checks for the current user to have access to players
	m, err := r.client.Match.Create().
		SetGameID(match.GameID).
		AddPlayerIDs(playerIds...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stats := make([]*ent.StatisticCreate, len(match.Stats))

	for i, stat := range match.Stats {
		// check that the player is in the player list
		for _, player := range match.PlayerIds {
			if player.ID != stat.PlayerID.ID {
				return nil, fmt.Errorf("player %s is not in the player list", stat.PlayerID.String())
			}
		}

		stats[i] = r.client.Statistic.Create().
			SetPlayerID(stat.PlayerID).
			SetValue(stat.Value).
			SetStatDescriptionID(stat.StatID).
			SetMatch(m)
	}

	_, err = r.client.Statistic.CreateBulk(stats...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return m, nil
}
