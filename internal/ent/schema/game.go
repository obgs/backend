package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// Game is the schema for the Game entity.
type Game struct {
	ent.Schema
}

const (
	MIN_PLAYERS_DEFAULT = 1
	MAX_PLAYERS_DEFAULT = 10
)

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.Game)),
		field.String("name").NotEmpty(),
		field.Int("min_players").Default(MIN_PLAYERS_DEFAULT),
		field.Int("max_players").Default(MAX_PLAYERS_DEFAULT),
		field.String("description").Optional().Default("").Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		field.String("boardgamegeek_url").Optional().Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("games").Required().Unique(),
		edge.To("favorites", GameFavorite.Type),
		edge.From("stat_descriptions", StatDescription.Type).Ref("game").Required().Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		edge.To("matches", Match.Type).Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Annotations of the Game.
func (Game) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
