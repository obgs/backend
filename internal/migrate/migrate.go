package migrate

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

func Migrate(ctx context.Context, client *ent.Client, migrationName string) error {
	migration, ok := existingMigrations[migrationName]
	if !ok {
		return fmt.Errorf("migration %v is not a valid option", migrationName)
	}

	err := validate(ctx, client, migrationName)
	if err != nil {
		return err
	}

	return migration.Run(ctx, client)
}
