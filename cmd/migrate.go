/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/obgs/backend/internal/migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the postMigrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migrations that require extra work rather than just SQL.",
	Long: `Runs migrations that require extra work rather than just SQL.
			Each migration requires a specific state of the database, determined by which atlas migrations where run,
			marked by "_prepare" and "_final" suffixes, meaning that this command must be run between applications of those.
			Only suitable for migrating existing data, the development databases do all the required staff in the seed scripts.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := createEntClient()
		if err != nil {
			log.Fatalf("%v", err)
		}

		err = migrate.Migrate(ctx, client, args[0])
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
