package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// Statistic holds the schema definition for the Statistic entity.
type Statistic struct {
	ent.Schema
}

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.Statistic)),
		field.String("value").Default(""),
	}
}

// Edges of the Stat.
func (Statistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("match", Match.Type).Ref("stats").Unique().Required(),
		edge.From("stat_description", StatDescription.Type).Ref("stats").Unique().Required(),
		edge.From("player", Player.Type).Ref("stats").Unique().Required(),
	}
}

// Annotations of the Statistic.
func (Statistic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
	}
}
