package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// EnumStatDescription holds the schema definition for the EnumStatDescription entity.
type EnumStatDescription struct {
	ent.Schema
}

// Fields of the EnumStatDescription.
func (EnumStatDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.EnumStatDescription)),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Strings("possible_values").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the EnumStatDescription.
func (EnumStatDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("enum_stats", EnumStat.Type).Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}
