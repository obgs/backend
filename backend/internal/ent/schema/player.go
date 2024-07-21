package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.Player)),
		field.String("name").Default(""),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("main_player").Unique(),
		edge.From("supervisors", User.Type).Ref("players"),
		edge.To("supervision_requests", PlayerSupervisionRequest.Type),
		edge.From("matches", Match.Type).Ref("players").Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		edge.To("stats", Statistic.Type).Annotations(
			entgql.Skip(entgql.SkipAll),
		),
	}
}

func (Player) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
