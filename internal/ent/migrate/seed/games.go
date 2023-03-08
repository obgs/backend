package seed

import (
	"context"
	"encoding/json"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
)

const (
	MIN_PLAYERS = 1
	MAX_PLAYERS = 5
)

func createTerraformingMars(ctx context.Context, tx *ent.Tx, author *ent.User) *ent.Game {
	stats := tx.StatDescription.CreateBulk(
		tx.StatDescription.Create().
			SetName("Terraforming Rating").
			SetType(stat.Numeric).
			SetOrderNumber(1),
		tx.StatDescription.Create().
			SetName("Milestones").
			SetType(stat.Numeric).
			SetOrderNumber(2),
		tx.StatDescription.Create().
			SetName("Awards").
			SetType(stat.Numeric).
			SetOrderNumber(3),
		tx.StatDescription.Create().
			SetName("Greeneries").
			SetType(stat.Numeric).
			SetOrderNumber(4),
		tx.StatDescription.Create().
			SetName("Cities").
			SetType(stat.Numeric).
			SetOrderNumber(5),
		tx.StatDescription.Create().
			SetName("Mars Map").
			SetType(stat.Aggregate).
			SetOrderNumber(6),
		tx.StatDescription.Create().
			SetName("Cards").
			SetType(stat.Numeric).
			SetOrderNumber(7),
		tx.StatDescription.Create().
			SetName("Crisis").
			SetType(stat.Numeric).
			SetOrderNumber(8),
		tx.StatDescription.Create().
			SetName("Overall score").
			SetType(stat.Aggregate).
			SetOrderNumber(9),
	).SaveX(ctx)

	// set references for mars map aggregate stat
	greeneries, cities := stats[3], stats[4]
	marsMapMetadata := stat.AggregateStatMetadata{
		Type:    stat.AggregateSum,
		StatIds: []guidgql.GUID{greeneries.ID, cities.ID},
	}
	marsMapMetadataBytes, _ := json.Marshal(marsMapMetadata)
	tx.StatDescription.UpdateOne(stats[5]).SetMetadata(string(marsMapMetadataBytes)).ExecX(ctx)

	// set references for overall score aggregate stat
	overallScoreMetadata := stat.AggregateStatMetadata{
		Type: stat.AggregateSum,
		StatIds: []guidgql.GUID{
			stats[0].ID, stats[1].ID, stats[2].ID, stats[5].ID, stats[6].ID, stats[7].ID,
		},
	}
	overallScoreMetadataBytes, _ := json.Marshal(overallScoreMetadata)
	tx.StatDescription.UpdateOne(stats[8]).SetMetadata(string(overallScoreMetadataBytes)).ExecX(ctx)

	return tx.Game.Create().
		SetName("Terraforming Mars").
		SetMinPlayers(MIN_PLAYERS).
		SetMaxPlayers(MAX_PLAYERS).
		SetBoardgamegeekURL("https://boardgamegeek.com/boardgame/167791/terraforming-mars").
		SetAuthor(author).
		AddStatDescriptions(stats...).
		SaveX(ctx)
}

func addGameToFavorites(ctx context.Context, tx *ent.Tx, user *ent.User, game *ent.Game) {
	tx.GameFavorite.Create().
		SetUser(user).
		SetGame(game).
		SaveX(ctx)
}
