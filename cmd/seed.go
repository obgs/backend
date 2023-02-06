package cmd

import (
	"context"
	"log"

	"github.com/open-boardgame-stats/backend/internal/ent/migrate/seed"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := createEntClient()
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()

		err = seed.SeedDB(ctx, client)
		if err != nil {
			log.Fatal(err)
		}
	},
}
