package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

type PlayerSupervisionRequest struct {
	ent.Schema
}

func (PlayerSupervisionRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New("player_supervision_requests")),
		field.String("message").Optional().Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}

func (PlayerSupervisionRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("sent_supervision_requests").Required().Unique(),
		edge.From("player", Player.Type).Ref("supervision_requests").Required().Unique(),
		edge.To("approvals", PlayerSupervisionRequestApproval.Type),
	}
}
