package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

const (
	MAX_FAVS        = 5
	MIN_ENUM_VALUES = 2
)

func addReferencesForAggregateStats(ctx context.Context, client *ent.Client, stats []*ent.StatDescription, input []*model.StatDescriptionInput) (err error) {
	for i, s := range stats {
		if s.Type != stat.Aggregate {
			continue
		}

		metadata, err := stat.UnmarshalMetadata(s.Type, s.Metadata)
		if err != nil {
			return err
		}

		aggregateMetadata, ok := metadata.(*stat.AggregateStatMetadata)
		if !ok {
			return fmt.Errorf("invalid metadata for aggregate stat")
		}

		orderNumbers := input[i].Metadata.AggregateMetadata.StatOrderNumbers
		aggregateMetadata.StatIds = make([]guidgql.GUID, len(orderNumbers))
		for j, orderNumber := range orderNumbers {
			aggregateMetadata.StatIds[j] = stats[orderNumber].ID
		}

		metadataBytes, err := json.Marshal(aggregateMetadata)
		if err != nil {
			return err
		}

		err = client.StatDescription.UpdateOne(s).SetMetadata(string(metadataBytes)).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
