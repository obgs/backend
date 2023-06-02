package migrate

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

// validate that the migration is being run at the exact time
// when atlas has run the _prepare script and have not run the _final one
func validate(ctx context.Context, client *ent.Client, name string) error {
	prepare := fmt.Sprintf("%s_prepare", name)
	final := fmt.Sprintf("%s_final", name)

	// check that the prepare migration has been run
	rows, err := client.QueryContext(
		ctx,
		fmt.Sprintf(
			"SELECT description FROM atlas_schema_revisions.atlas_schema_revisions WHERE description ILIKE '%s'",
			prepare,
		),
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("migration %s_prepare has not been run", name)
	}

	// check that the final migration has not been run
	rows, err = client.QueryContext(
		ctx,
		fmt.Sprintf(
			"SELECT description FROM atlas_schema_revisions.atlas_schema_revisions WHERE description ILIKE '%s'",
			final,
		),
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		return fmt.Errorf("migration %s_final has already been run", name)
	}

	return nil
}
