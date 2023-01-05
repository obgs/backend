package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// NumericalStat holds the schema definition for the NumericalStat entity.
type NumericalStat struct {
	ent.Schema
}

// Fields of the NumericalStat.
func (NumericalStat) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.NumericalStat)),
		field.Float("value"),
	}
}

// Edges of the NumericalStat.
func (NumericalStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("match", Match.Type).Ref("numerical_stats").Unique().Required(),
		edge.From("numerical_stat_description", NumericalStatDescription.Type).Ref("numerical_stats").Unique().Required(),
		edge.From("player", Player.Type).Ref("numerical_stats").Unique().Required(),
	}
}
