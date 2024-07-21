package seed

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/obgs/backend/internal/ent"
	"github.com/obgs/backend/internal/ent/gameversion"
	"github.com/obgs/backend/internal/ent/schema/stat"
	"github.com/obgs/backend/internal/ent/statdescription"
)

const MATCH_MAX_NUMERIC_VALUE = 100

func createPlainStats(
	tx *ent.Tx,
	descriptions []*ent.StatDescription,
	player *ent.Player,
	match *ent.Match,
	values map[string]string,
) ([]*ent.StatisticCreate, []*ent.StatDescription) {
	stats := make([]*ent.StatisticCreate, 0, len(descriptions))
	aggregateStatDescriptions := make([]*ent.StatDescription, 0, len(descriptions))
	for _, d := range descriptions {
		var value string
		metadata, err := stat.UnmarshalMetadata(d.Type, d.Metadata)
		if err != nil {
			panic(err)
		}

		switch d.Type {
		case stat.Enum:
			enumMetadata, ok := metadata.(*stat.EnumStatMetadata)
			if !ok {
				panic("invalid metadata")
			}
			value = enumMetadata.PossibleValues[rand.Intn(len(enumMetadata.PossibleValues))]
		case stat.Numeric:
			value = strconv.Itoa(rand.Intn(MATCH_MAX_NUMERIC_VALUE))
		case stat.Aggregate:
			// aggregate stats should be processed after basic ones are all done
			aggregateStatDescriptions = append(aggregateStatDescriptions, d)
			continue
		default:
			panic("invalid stat type")
		}

		values[d.ID.String()] = value
		stats = append(stats, tx.Statistic.Create().
			SetValue(value).
			SetPlayer(player).
			SetStatDescription(d).
			SetMatch(match))
	}

	return stats, aggregateStatDescriptions
}

func createAggregateStats(
	tx *ent.Tx,
	aggregateStats []*ent.StatDescription,
	player *ent.Player,
	match *ent.Match,
	values map[string]string,
) []*ent.StatisticCreate {
	stats := make([]*ent.StatisticCreate, 0, len(aggregateStats))
	for _, d := range aggregateStats {
		var value string
		metadata, err := stat.UnmarshalMetadata(d.Type, d.Metadata)
		if err != nil {
			panic(err)
		}

		switch d.Type {
		case stat.Aggregate:
			aggregateMetadata, ok := metadata.(*stat.AggregateStatMetadata)
			if !ok {
				panic("invalid metadata")
			}
			// other aggregate stats are not supported yet
			if aggregateMetadata.Type != stat.AggregateSum {
				panic("invalid aggregate metadata type")
			}
			sum := 0
			for _, id := range aggregateMetadata.StatIds {
				v, err := strconv.Atoi(values[id.String()])
				if err != nil {
					panic(err)
				}
				sum += v
			}
			value = strconv.Itoa(sum)
		default:
			panic("invalid stat type")
		}

		stats = append(stats, tx.Statistic.Create().
			SetValue(value).
			SetPlayer(player).
			SetStatDescription(d).
			SetMatch(match))
	}

	return stats
}

func createMatch(ctx context.Context, tx *ent.Tx, g *ent.GameVersion, players []*ent.Player) *ent.Match {
	descriptions := tx.StatDescription.Query().
		Where(
			statdescription.HasGameVersionWith(gameversion.ID(g.ID)),
		).
		AllX(ctx)

	stats := make([]*ent.StatisticCreate, 0, len(descriptions)*len(players))

	match := tx.Match.Create().
		SetGameVersion(g).
		AddPlayers(players...).
		SaveX(ctx)

	for _, player := range players {
		values := make(map[string]string, len(descriptions))
		plainStats, aggregateStatDescriptions := createPlainStats(tx, descriptions, player, match, values)
		aggregateStats := createAggregateStats(tx, aggregateStatDescriptions, player, match, values)
		stats = append(stats, plainStats...)
		stats = append(stats, aggregateStats...)
	}
	tx.Statistic.CreateBulk(stats...).SaveX(ctx)

	return match
}
