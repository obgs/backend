package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/enums"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

type GroupMembership struct {
	ent.Schema
}

func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.GroupMembership)),
		field.Enum("role").GoType(enums.RoleMember),
	}
}

func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("members").Required().Unique(),
		edge.From("user", User.Type).Ref("group_memberships").Required().Unique(),
	}
}

func (GroupMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
