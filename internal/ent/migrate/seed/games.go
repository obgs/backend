package seed

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
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
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Milestones").
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Awards").
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Greeneries").
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Cities").
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Cards").
			SetType(stat.Numeric),
		tx.StatDescription.Create().
			SetName("Crisis").
			SetType(stat.Numeric),
	).SaveX(ctx)

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
