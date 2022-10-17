package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type PlayerSupervisionRequestApproval struct {
	ent.Schema
}

func (PlayerSupervisionRequestApproval) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("approved").Optional().Nillable(),
	}
}

func (PlayerSupervisionRequestApproval) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("approver", User.Type).Ref("supervision_request_approvals").Required().Unique(),
		edge.From("supervision_request", PlayerSupervisionRequest.Type).Ref("approvals").Required().Unique(),
	}
}
