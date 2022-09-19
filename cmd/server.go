package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"

	"github.com/open-boardgame-stats/backend/auth"
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
		ctx := context.Background()
		if err := client.Schema.Create(
			ctx,
			migrate.WithGlobalUniqueID(true),
		); err != nil {
			log.Fatalf("failed to migrate schema: %v", err)
		}

		srv := handler.NewDefaultServer(resolver.NewSchema(client))
		srv.Use(entgql.Transactioner{TxOpener: client})

		oAuthConfig := auth.NewOAuthConfig(config.ServerHost, serverPort, config.OAuthGoogleClientID, config.OAuthGoogleClientSecret)
		authService := auth.NewAuthService(client, ctx, config.JWTSecret, oAuthConfig)

		router := chi.NewRouter()

		// cors setup
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		router.Use(authService.Middleware)

		router.Get("/", playground.Handler("OBGS", "/graphql"))
		router.Handle("/graphql", srv)
		router.Post("/auth/signup", authService.SignUp)
		router.Post("/auth/signin", authService.SignIn)
		router.Post("/auth/refresh", authService.Refresh)
		router.Post("/auth/google/signin", authService.OAuthGoogleSignIn)

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "which port to serve the schema on (default: 8080)")
}
