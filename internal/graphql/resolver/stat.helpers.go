package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

const (
	MIN_SUM_STATS = 2
)

func marshalStatMetadata(t stat.StatType, input *model.StatMetadataInput) (metadata string, err error) {
	if input == nil {
		return "", nil
	}

	switch t {
	case stat.Numeric:
		return "", nil
	case stat.Enum:
		if input.EnumMetadata == nil {
			return "", fmt.Errorf("enum metadata is required for enum stat")
		}
		if len(input.EnumMetadata.PossibleValues) < MIN_ENUM_VALUES {
			return "", fmt.Errorf("enum stat must have at least %d possible values", MIN_ENUM_VALUES)
		}

		bytes, err := json.Marshal(input.EnumMetadata)
		if err != nil {
			return "", err
		}

		metadata = string(bytes)
	case stat.Aggregate:
		if input.AggregateMetadata == nil {
			return "", fmt.Errorf("aggregate metadata is required for aggregate stat")
		}
		if len(input.AggregateMetadata.StatOrderNumbers) < MIN_SUM_STATS {
			return "", fmt.Errorf("aggregate stat must have at least two stats")
		}

		bytes, err := json.Marshal(input.AggregateMetadata)
		if err != nil {
			return "", err
		}

		metadata = string(bytes)
	default:
		return "", fmt.Errorf("unknown stat type: %s", t)
	}

	return metadata, err
}
