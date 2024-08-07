package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

type GroupMembershipApplication struct {
	ent.Schema
}

func (GroupMembershipApplication) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.GroupMembershipApplication)),
		field.String("message").Default(""),
	}
}

func (GroupMembershipApplication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("group_membership_applications").Required().Unique(),
		edge.From("group", Group.Type).Ref("applications").Required().Unique(),
	}
}

func (GroupMembershipApplication) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
	}
}
