package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"

	"github.com/open-boardgame-stats/backend/graphql/generated"
	"github.com/open-boardgame-stats/backend/graphql/resolver"
)

var serverPort string

// server represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the gql server",
	Run: func(cmd *cobra.Command, args []string) {
		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolver.Resolver{},
		})
		srv := handler.NewDefaultServer(schema)

		http.Handle("/", playground.Handler("OBGS", "/graphql"))
		http.Handle("/graphql", srv)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "which port to serve the schema on (default: 8080)")
}
