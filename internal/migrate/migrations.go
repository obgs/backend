package migrate

import "github.com/open-boardgame-stats/backend/internal/migrate/migrations"

var existingMigrations = map[string]migrations.Migration{
	"game_versions": migrations.NewGameVersionMigration(),
}
