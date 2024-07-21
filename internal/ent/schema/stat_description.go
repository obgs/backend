package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/schema/stat"
)

// StatDescription is the schema for the StatDescription entity.
type StatDescription struct {
	ent.Schema
}

// Fields of the StatDescription.
func (StatDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.StatDescription)),
		field.Enum("type").GoType(stat.StatType(stat.Numeric)),
		field.String("name").NotEmpty(),
		field.String("description").Optional().Default(""),
		field.String("metadata").Optional().Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int("order_number").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the StatDescription.
func (StatDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_version", GameVersion.Type).Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("stats", Statistic.Type).Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Annotations of the StatDescription.
func (StatDescription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
	}
}
