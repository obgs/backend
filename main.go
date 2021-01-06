package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/open-boardgame-stats/backend/graphql/generated"
	"github.com/open-boardgame-stats/backend/graphql/resolver"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{},
	}))

	http.Handle("/", playground.Handler("OBGS", "/graphql"))
	http.Handle("/graphql", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
