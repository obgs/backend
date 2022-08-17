package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"

	"github.com/open-boardgame-stats/backend/ent"
	"github.com/open-boardgame-stats/backend/ent/migrate"
	"github.com/open-boardgame-stats/backend/graphql/resolver"

	_ "github.com/lib/pq"
)

var serverPort string

// server represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the gql server",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBAddress, config.DBPort, config.DBUser, config.DBPass, config.DBName))
		if err != nil {
			log.Fatalf("failed to open connection to postgres: %v", err)
		}
		if err := client.Schema.Create(
			context.Background(),
			migrate.WithGlobalUniqueID(true),
		); err != nil {
			log.Fatalf("failed to migrate schema: %v", err)
		}

		srv := handler.NewDefaultServer(resolver.NewSchema(client))
		srv.Use(entgql.Transactioner{TxOpener: client})

		http.Handle("/", playground.Handler("OBGS", "/graphql"))
		http.Handle("/graphql", srv)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "which port to serve the schema on (default: 8080)")
}
