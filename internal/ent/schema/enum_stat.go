package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// EnumStat holds the schema definition for the EnumStat entity.
type EnumStat struct {
	ent.Schema
}

// Fields of the EnumStat.
func (EnumStat) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.EnumStat)),
		field.String("value"),
	}
}

// Edges of the EnumStat.
func (EnumStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("match", Match.Type).Ref("enum_stats").Unique().Required(),
		edge.From("enum_stat_description", EnumStatDescription.Type).Ref("enum_stats").Unique().Required(),
		edge.From("player", Player.Type).Ref("enum_stats").Unique().Required(),
	}
}
