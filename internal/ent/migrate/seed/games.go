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

var corporations = []string{
	// Default corporations
	"Credicor",
	"Ecoline",
	"Helion",
	"Mining Guild",
	"Interplanetary Cinematics",
	"Inventrix",
	"Phobolog",
	"Tharsis Republic",
	"Thorgate",
	"United Nations Mars Initiative",
	"Teractor",
	"Saturn Systems",

	// Venus Next corporations
	"Aphrodite",
	"Celestic",
	"Manutech",
	"Morning Star Inc.",
	"Viron",

	// Prelude corporations
	"Cheung Shing Mars",
	"Point Luna",
	"Robinson Industries",
	"Valley Trust",
	"Vitor",

	// Colonies corporations
	"Aridor",
	"Arklight",
	"Polyphemos",
	"Poseidon",
	"Storm Craft Incorporated",

	// Turmoil corporations
	"Lakefront Resorts",
	"Pristar",
	"Septem Tribus",
	"Terralabs Research",
	"Utopia Invest",

	"Factorum",
	"Mons Insurance",
	"Philares",
	"Arcadian Communities",
	"Recyclon",
	"Splice Tactical Genomics",
}

func createTerraformingMars(ctx context.Context, tx *ent.Tx, author *ent.User) *ent.Game {
	stats := tx.StatDescription.CreateBulk(
		tx.StatDescription.Create().
			SetName("Corporation").
			SetType(stat.Enum).
			SetOrderNumber(1),
		tx.StatDescription.Create().
			SetName("Terraforming Rating").
			SetType(stat.Numeric).
			SetOrderNumber(2),
		tx.StatDescription.Create().
			SetName("Milestones").
			SetType(stat.Numeric).
			SetOrderNumber(3),
		tx.StatDescription.Create().
			SetName("Awards").
			SetType(stat.Numeric).
			SetOrderNumber(4),
		tx.StatDescription.Create().
			SetName("Greeneries").
			SetType(stat.Numeric).
			SetOrderNumber(5),
		tx.StatDescription.Create().
			SetName("Cities").
			SetType(stat.Numeric).
			SetOrderNumber(6),
		tx.StatDescription.Create().
			SetName("Mars Map").
			SetType(stat.Aggregate).
			SetOrderNumber(7),
		tx.StatDescription.Create().
			SetName("Cards").
			SetType(stat.Numeric).
			SetOrderNumber(8),
		tx.StatDescription.Create().
			SetName("Crisis").
			SetType(stat.Numeric).
			SetOrderNumber(9),
		tx.StatDescription.Create().
			SetName("Overall score").
			SetType(stat.Aggregate).
			SetOrderNumber(10),
	).SaveX(ctx)

	// set references for mars map aggregate stat
	greeneries, cities := stats[4], stats[5]
	marsMapMetadata := stat.AggregateStatMetadata{
		Type:    stat.AggregateSum,
		StatIds: []guidgql.GUID{greeneries.ID, cities.ID},
	}
	marsMapMetadataBytes, _ := json.Marshal(marsMapMetadata)
	tx.StatDescription.UpdateOne(stats[6]).SetMetadata(string(marsMapMetadataBytes)).ExecX(ctx)

	// set references for overall score aggregate stat
	overallScoreMetadata := stat.AggregateStatMetadata{
		Type: stat.AggregateSum,
		StatIds: []guidgql.GUID{
			stats[1].ID, stats[2].ID, stats[3].ID, stats[4].ID, stats[5].ID, stats[7].ID, stats[8].ID,
		},
	}
	overallScoreMetadataBytes, _ := json.Marshal(overallScoreMetadata)
	tx.StatDescription.UpdateOne(stats[9]).SetMetadata(string(overallScoreMetadataBytes)).ExecX(ctx)

	// create corporations
	corporationsMetadata := stat.EnumStatMetadata{
		PossibleValues: corporations,
	}
	corporationsMetadataBytes, _ := json.Marshal(corporationsMetadata)
	tx.StatDescription.UpdateOne(stats[0]).SetMetadata(string(corporationsMetadataBytes)).ExecX(ctx)

	g := tx.Game.Create().
		SetName("Terraforming Mars").
		SetMinPlayers(MIN_PLAYERS).
		SetMaxPlayers(MAX_PLAYERS).
		SetBoardgamegeekURL("https://boardgamegeek.com/boardgame/167791/terraforming-mars").
		SetAuthor(author).
		SaveX(ctx)

	tx.GameVersion.Create().
		AddStatDescriptions(stats...).
		SetGame(g).
		SaveX(ctx)

	return g
}

func addGameToFavorites(ctx context.Context, tx *ent.Tx, user *ent.User, game *ent.Game) {
	tx.GameFavorite.Create().
		SetUser(user).
		SetGame(game).
		SaveX(ctx)
}
