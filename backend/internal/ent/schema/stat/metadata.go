package stat

import (
	"encoding/json"
	"fmt"

	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

type StatMetadata interface {
	New() StatMetadata
}

type EnumStatMetadata struct {
	PossibleValues []string `json:"possibleValues"`
}

func (m *EnumStatMetadata) New() StatMetadata {
	return &EnumStatMetadata{
		PossibleValues: []string{},
	}
}

type NumericStatMetadata struct{}

func (m *NumericStatMetadata) New() StatMetadata {
	return &NumericStatMetadata{}
}

type AggregateStatMetadata struct {
	Type    AggregateType  `json:"type"`
	StatIds []guidgql.GUID `json:"statIds"`
}

func (m *AggregateStatMetadata) New() StatMetadata {
	return &AggregateStatMetadata{
		Type:    AggregateSum,
		StatIds: []guidgql.GUID{},
	}
}

func UnmarshalMetadata(t StatType, s string) (StatMetadata, error) {
	switch t {
	case Numeric:
		return &NumericStatMetadata{}, nil
	case Enum:
		m := &EnumStatMetadata{}
		if err := json.Unmarshal([]byte(s), m); err != nil {
			return nil, err
		}

		return m, nil
	case Aggregate:
		m := &AggregateStatMetadata{}
		if err := json.Unmarshal([]byte(s), m); err != nil {
			return nil, err
		}

		return m, nil
	default:
		return nil, fmt.Errorf("invalid StatType %q", t)
	}
}
