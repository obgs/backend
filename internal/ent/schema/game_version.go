package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

type GameVersion struct {
	ent.Schema
}

// Fields of the GameVersion.
func (GameVersion) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.GameVersion)),
		field.Int("version_number").Default(1),
	}
}

// Edges of the GameVersion.
func (GameVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Required().Unique(),
		edge.From("stat_descriptions", StatDescription.Type).Ref("game_version").Required().Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		edge.To("matches", Match.Type).Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
