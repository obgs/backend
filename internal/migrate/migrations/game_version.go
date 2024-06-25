package migrations

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

type GameVersionMigration struct{}

func NewGameVersionMigration() *GameVersionMigration {
	return &GameVersionMigration{}
}

func (m *GameVersionMigration) Name() string {
	return "game_version"
}

func moveMatches(ctx context.Context, client *ent.Client, gameID guidgql.GUID, versionID guidgql.GUID) error {
	_, err := client.ExecContext(
		ctx,
		`UPDATE matches SET game_version_matches = $1 WHERE game_matches = $2`,
		versionID.String(),
		gameID.String(),
	)

	return err
}

func moveStatDescriptions(ctx context.Context, client *ent.Client, gameID guidgql.GUID, versionID guidgql.GUID) error {
	_, err := client.ExecContext(
		ctx,
		`INSERT INTO stat_description_game_version (stat_description_id, game_version_id)
			SELECT stat_description_id, $1 FROM stat_description_game WHERE stat_description_game.game_id = $2`,
		versionID.String(),
		gameID.String(),
	)

	return err
}

func (m *GameVersionMigration) Run(ctx context.Context, client *ent.Client) error {
	games, err := client.Game.Query().All(ctx)
	if err != nil {
		return err
	}

	// create a game version for each game
	gameToVersionMap := make(map[guidgql.GUID]guidgql.GUID)
	for _, game := range games {
		id := guidgql.New(guidgql.GameVersion)()
		gameToVersionMap[game.ID] = id
		_, err = client.ExecContext(ctx, `INSERT INTO game_versions VALUES ($1, $2, $3)`, id, 1, game.ID.String())
		if err != nil {
			return err
		}
	}

	// move all matches and stat descriptions from their games to respective newly created versions
	for gameID, versionID := range gameToVersionMap {
		err = moveMatches(ctx, client, gameID, versionID)
		if err != nil {
			return err
		}

		err = moveStatDescriptions(ctx, client, gameID, versionID)
		if err != nil {
			return err
		}
	}

	return nil
}
