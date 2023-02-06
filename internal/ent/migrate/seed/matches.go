package seed

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
)

const MATCH_MAX_NUMERIC_VALUE = 100

func createMatch(ctx context.Context, tx *ent.Tx, g *ent.Game, players []*ent.Player) *ent.Match {
	descriptions := tx.StatDescription.Query().
		Where(
			statdescription.HasGameWith(game.ID(g.ID)),
		).
		AllX(ctx)

	stats := make([]*ent.Statistic, len(descriptions)*len(players))

	match := tx.Match.Create().
		SetGame(g).
		AddPlayers(players...).
		SaveX(ctx)

	for _, player := range players {
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
				value = enumMetadata.PossibleValues[0]
			case stat.Numeric:
				value = strconv.Itoa(rand.Intn(MATCH_MAX_NUMERIC_VALUE))
			}

			stats = append(stats, tx.Statistic.Create().
				SetValue(value).
				SetPlayer(player).
				SetStatDescription(d).
				SetMatch(match).
				SaveX(ctx),
			)
		}
	}

	return match
}
