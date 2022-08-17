package cmd

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/spf13/cobra"

	"github.com/open-boardgame-stats/backend/graphql/generated"
)

func getIntrospection(schema graphql.ExecutableSchema) *graphql.Response {
	executor := executor.New(schema)
	ctx := graphql.StartOperationTrace(context.Background())
	now := graphql.Now()
	rc, err := executor.CreateOperationContext(ctx, &graphql.RawParams{
		Query:         introspection.Query,
		OperationName: "IntrospectionQuery",
		ReadTime: graphql.TraceTiming{
			Start: now,
			End:   now,
		},
	})
	if err != nil {
		return executor.DispatchError(ctx, err)
	}

	rc.DisableIntrospection = false
	resp, ctx2 := executor.DispatchOperation(ctx, rc)
	return resp(ctx2)
}

func printJSON(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// dumpSchemaCmd represents the dumpSchema command
var dumpSchemaCmd = &cobra.Command{
	Use:   "dumpSchema",
	Short: "Dump graphql schema introspection",
	Run: func(cmd *cobra.Command, args []string) {
		schema := generated.NewExecutableSchema(generated.Config{})
		res := getIntrospection(schema)
		if err := printJSON(os.Stdout, res); err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpSchemaCmd)
}
