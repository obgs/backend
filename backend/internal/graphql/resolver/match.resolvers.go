package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/obgs/backend/internal/ent"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/graphql/model"
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
		SetGameVersionID(input.GameVersionID).
		AddPlayerIDs(playerIds...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stats := make([]*ent.StatisticCreate, len(input.Stats))

	for i, stat := range input.Stats {
		// check that the player is in the player list
		playerIsPresent := false
		for _, player := range input.PlayerIds {
			if player.ID == stat.PlayerID.ID {
				playerIsPresent = true
				break
			}
		}
		if !playerIsPresent {
			return nil, fmt.Errorf("player %s is not in the player list", stat.PlayerID.ID)
		}

		stats[i] = client.Statistic.Create().
			SetPlayerID(stat.PlayerID).
			SetValue(stat.Value).
			SetStatDescriptionID(stat.StatID).
			SetMatch(m)
	}

	aggregateStats, err := calculateAllAggregateStats(ctx, client, input.Stats, playerIds, m)
	if err != nil {
		return nil, err
	}

	stats = append(stats, aggregateStats...)

	_, err = client.Statistic.CreateBulk(stats...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return m, nil
}