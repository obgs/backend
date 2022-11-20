package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

type PlayerSupervisionRequestApproval struct {
	ent.Schema
}

func (PlayerSupervisionRequestApproval) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.PlayerSupervisionRequestApproval)),
		field.Bool("approved").Optional().Nillable(),
	}
}

func (PlayerSupervisionRequestApproval) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("approver", User.Type).Ref("supervision_request_approvals").Required().Unique(),
		edge.From("supervision_request", PlayerSupervisionRequest.Type).Ref("approvals").Required().Unique(),
	}
}
