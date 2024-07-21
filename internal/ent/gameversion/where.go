// Code generated by ent, DO NOT EDIT.

package gameversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/obgs/backend/internal/ent/predicate"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldLTE(FieldID, id))
}

// VersionNumber applies equality check predicate on the "version_number" field. It's identical to VersionNumberEQ.
func VersionNumber(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldEQ(FieldVersionNumber, v))
}

// VersionNumberEQ applies the EQ predicate on the "version_number" field.
func VersionNumberEQ(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldEQ(FieldVersionNumber, v))
}

// VersionNumberNEQ applies the NEQ predicate on the "version_number" field.
func VersionNumberNEQ(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldNEQ(FieldVersionNumber, v))
}

// VersionNumberIn applies the In predicate on the "version_number" field.
func VersionNumberIn(vs ...int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldIn(FieldVersionNumber, vs...))
}

// VersionNumberNotIn applies the NotIn predicate on the "version_number" field.
func VersionNumberNotIn(vs ...int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldNotIn(FieldVersionNumber, vs...))
}

// VersionNumberGT applies the GT predicate on the "version_number" field.
func VersionNumberGT(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldGT(FieldVersionNumber, v))
}

// VersionNumberGTE applies the GTE predicate on the "version_number" field.
func VersionNumberGTE(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldGTE(FieldVersionNumber, v))
}

// VersionNumberLT applies the LT predicate on the "version_number" field.
func VersionNumberLT(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldLT(FieldVersionNumber, v))
}

// VersionNumberLTE applies the LTE predicate on the "version_number" field.
func VersionNumberLTE(v int) predicate.GameVersion {
	return predicate.GameVersion(sql.FieldLTE(FieldVersionNumber, v))
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GameTable, GameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := newGameStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatDescriptions applies the HasEdge predicate on the "stat_descriptions" edge.
func HasStatDescriptions() predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StatDescriptionsTable, StatDescriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatDescriptionsWith applies the HasEdge predicate on the "stat_descriptions" edge with a given conditions (other predicates).
func HasStatDescriptionsWith(preds ...predicate.StatDescription) predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := newStatDescriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatches applies the HasEdge predicate on the "matches" edge.
func HasMatches() predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchesTable, MatchesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchesWith applies the HasEdge predicate on the "matches" edge with a given conditions (other predicates).
func HasMatchesWith(preds ...predicate.Match) predicate.GameVersion {
	return predicate.GameVersion(func(s *sql.Selector) {
		step := newMatchesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GameVersion) predicate.GameVersion {
	return predicate.GameVersion(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GameVersion) predicate.GameVersion {
	return predicate.GameVersion(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GameVersion) predicate.GameVersion {
	return predicate.GameVersion(sql.NotPredicates(p))
}
