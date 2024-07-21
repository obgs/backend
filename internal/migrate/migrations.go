package migrate

import "github.com/obgs/backend/internal/migrate/migrations"

var existingMigrations = map[string]migrations.Migration{
	"game_versions": migrations.NewGameVersionMigration(),
}
