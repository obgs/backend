// Code generated by ent, DO NOT EDIT.

package statdescription

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v stat.StatType) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v stat.StatType) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...stat.StatType) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...stat.StatType) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.StatDescription {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// PossibleValuesIsNil applies the IsNil predicate on the "possible_values" field.
func PossibleValuesIsNil() predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPossibleValues)))
	})
}

// PossibleValuesNotNil applies the NotNil predicate on the "possible_values" field.
func PossibleValuesNotNil() predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPossibleValues)))
	})
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GameTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GameTable, GamePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GameInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GameTable, GamePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StatDescription) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StatDescription) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StatDescription) predicate.StatDescription {
	return predicate.StatDescription(func(s *sql.Selector) {
		p(s.Not())
	})
}
