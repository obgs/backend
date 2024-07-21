package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

// GameFavorite is the schema for the GameFavorite entity.
type GameFavorite struct {
	ent.Schema
}

// Fields of the GameFavorite.
func (GameFavorite) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.GameFavorites)),
	}
}

// Edges of the GameFavorite.
func (GameFavorite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).Ref("favorites").Required().Unique(),
		edge.From("user", User.Type).Ref("favorite_games").Required().Unique(),
	}
}

// Annotations of the GameFavorite.
func (GameFavorite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(
			entgql.SkipAll,
		),
	}
}
