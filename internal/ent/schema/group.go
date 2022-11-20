package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.Group)),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Default("").Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		field.String("logo_url").NotEmpty().Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("settings", GroupSettings.Type).Required().Unique().Annotations(
			entgql.Mutations(
				entgql.MutationUpdate(),
			),
		),
		edge.To("members", GroupMembership.Type).Annotations(
			entgql.RelayConnection(),
		),
		edge.To("applications", GroupMembershipApplication.Type),
	}
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
