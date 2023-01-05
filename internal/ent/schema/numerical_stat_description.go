package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// NumericalStatDescription holds the schema definition for the NumericalStatDescription entity.
type NumericalStatDescription struct {
	ent.Schema
}

// Fields of the NumericalStatDescription.
func (NumericalStatDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.NumericalStatDescription)),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the NumericalStatDescription.
func (NumericalStatDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("numerical_stats", NumericalStat.Type).Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}
