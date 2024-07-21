package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/obgs/backend/internal/ent"
	"github.com/obgs/backend/internal/graphql/model"
)

// Metrics is the resolver for the metrics field.
func (r *gameVersionResolver) Metrics(ctx context.Context, obj *ent.GameVersion) (*model.GameVersionMetrics, error) {
	fc := graphql.GetFieldContext(ctx)
	opCtx := graphql.GetOperationContext(ctx)
	collected := graphql.CollectFields(opCtx, fc.Field.Selections, []string{"User"})
	var numericStats []*model.NumericMetric
	var enumStats []*model.EnumMetric
	var matchesCreated *model.TimeSeries
	var adoption *model.TimeFloatMetric
	for _, field := range collected {
		switch field.Name {
		case "numericStats":
			resp, err := getNumericStats(ctx, r.client, obj)
			if err != nil {
				return nil, err
			}
			numericStats = resp
		case "enumStats":
			resp, err := getEnumStats(ctx, r.client, obj)
			if err != nil {
				return nil, err
			}
			enumStats = resp
		case "matchesCreated":
			childFc, err := graphql.GetFieldContext(ctx).Child(ctx, field)
			if err != nil {
				return nil, err
			}
			input, ok := childFc.Args["input"].(model.TimeSeriesInput)
			if !ok {
				return nil, fmt.Errorf("invalid input type")
			}
			resp, err := getMatchesCreated(ctx, r.client, obj, &input)
			if err != nil {
				return nil, err
			}
			matchesCreated = resp
		case "adoption":
			childFc, err := graphql.GetFieldContext(ctx).Child(ctx, field)
			if err != nil {
				return nil, err
			}
			input, ok := childFc.Args["input"].(model.GranularityInput)
			if !ok {
				return nil, fmt.Errorf("invalid input type")
			}
			resp, err := getAdoption(ctx, r.client, obj, &input)
			if err != nil {
				return nil, err
			}
			adoption = resp
		}
	}
	return &model.GameVersionMetrics{
		NumericStats:   numericStats,
		EnumStats:      enumStats,
		MatchesCreated: matchesCreated,
		Adoption:       adoption,
	}, nil
}
