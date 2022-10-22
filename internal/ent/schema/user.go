package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Default(""),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty().Sensitive().Annotations(
			entgql.Annotation(entgql.Skip()),
		),
		field.String("avatar_url").Default("").Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("players", Player.Type),
		edge.To("main_player", Player.Type).Unique(),
		edge.To("sent_supervision_requests", PlayerSupervisionRequest.Type).Annotations(
			entgql.Skip(entgql.SkipAll),
		),
		edge.To("supervision_request_approvals", PlayerSupervisionRequestApproval.Type).Annotations(
			entgql.Skip(entgql.SkipAll),
		),
		edge.To("group_memberships", GroupMembership.Type).Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
		edge.To("group_membership_applications", GroupMembershipApplication.Type).Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationUpdate(),
		),
	}
}
