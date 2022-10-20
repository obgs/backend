package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
)

type GroupMembership struct {
	ent.Schema
}

func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("role").GoType(enums.RoleMember),
	}
}

func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("members").Required().Unique(),
		edge.From("user", User.Type).Ref("group_memberships").Required().Unique(),
	}
}
