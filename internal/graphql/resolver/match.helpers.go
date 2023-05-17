package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

func prepareAggregateStatCalculation(ctx context.Context, client *ent.Client, input []*model.StatInput, playerIds []guidgql.GUID, m *ent.Match) (map[guidgql.GUID]*ent.StatDescription, []*ent.StatDescription, map[guidgql.GUID]map[guidgql.GUID]*model.StatInput, error) {
	g, err := m.Game(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get game: %w", err)
	}

	statDescriptions, err := client.StatDescription.Query().Where(statdescription.HasGameWith(game.ID(g.ID))).All(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	// create a map of stat descriptions for faster lookup
	statDescriptionsMap := make(map[guidgql.GUID]*ent.StatDescription, len(statDescriptions))
	aggregateStats := make([]*ent.StatDescription, 0, len(statDescriptions))
	for _, s := range statDescriptions {
		statDescriptionsMap[s.ID] = s
		if s.Type == stat.Aggregate {
			aggregateStats = append(aggregateStats, s)
		}
	}

	// create a map of stat inputs
	statsPerPlayer := make(map[guidgql.GUID]map[guidgql.GUID]*model.StatInput, len(playerIds))
	for _, s := range input {
		playerStats, ok := statsPerPlayer[s.PlayerID]
		if !ok {
			playerStats = make(map[guidgql.GUID]*model.StatInput, len(statDescriptions))
			statsPerPlayer[s.PlayerID] = playerStats
		}
		playerStats[s.StatID] = s
	}

	return statDescriptionsMap, aggregateStats, statsPerPlayer, nil
}

func calculateSingleAggregateStat(metadata *model.AggregateMetadata, statDescriptionsMap map[guidgql.GUID]*ent.StatDescription, statsPerPlayer map[guidgql.GUID]map[guidgql.GUID]*model.StatInput, playerId guidgql.GUID) (*int64, error) {
	playerStats, ok := statsPerPlayer[playerId]
	if !ok || playerStats == nil {
		return nil, fmt.Errorf("player %s stats not found in input", playerId)
	}

	// calculate the aggregate stat value
	var aggregateValue int64 = 0
	for _, statID := range metadata.StatIds {
		referencedStat := statDescriptionsMap[*statID]
		if referencedStat == nil {
			return nil, fmt.Errorf("stat description %s not found", *statID)
		}

		if referencedStat.Type != stat.Numeric {
			return nil, fmt.Errorf("stat %s is not numeric", *statID)
		}

		s, ok := playerStats[*statID]
		if !ok || s == nil {
			return nil, fmt.Errorf("stat %s not found in input", *statID)
		}

		v, err := strconv.ParseInt(s.Value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse stat value: %w", err)
		}

		aggregateValue += v
	}

	return &aggregateValue, nil
}

func calculateAllAggregateStats(ctx context.Context, client *ent.Client, input []*model.StatInput, playerIds []guidgql.GUID, m *ent.Match) ([]*ent.StatisticCreate, error) {
	statDescriptionsMap, aggregateStats, statsPerPlayer, err := prepareAggregateStatCalculation(ctx, client, input, playerIds, m)
	if err != nil {
		return nil, err
	}

	createAggregateStats := make([]*ent.StatisticCreate, 0, len(aggregateStats)*len(playerIds))
	for _, statDescription := range aggregateStats {
		var metadata model.AggregateMetadata
		if err := json.Unmarshal([]byte(statDescription.Metadata), &metadata); err != nil {
			return nil, fmt.Errorf("failed to unmarshal aggregate metadata: %w", err)
		}

		for _, player := range playerIds {
			aggregateValue, err := calculateSingleAggregateStat(&metadata, statDescriptionsMap, statsPerPlayer, player)
			if err != nil {
				return nil, err
			}

			createAggregateStats = append(createAggregateStats, client.Statistic.Create().
				SetPlayerID(player).
				SetValue(strconv.FormatInt(*aggregateValue, 10)).
				SetStatDescriptionID(statDescription.ID).
				SetMatch(m))
		}
	}

	return createAggregateStats, nil
}
