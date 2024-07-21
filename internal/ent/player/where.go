// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/obgs/backend/internal/ent/predicate"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Player {
	return predicate.Player(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Player {
	return predicate.Player(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Player {
	return predicate.Player(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Player {
	return predicate.Player(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Player {
	return predicate.Player(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Player {
	return predicate.Player(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Player {
	return predicate.Player(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Player {
	return predicate.Player(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Player {
	return predicate.Player(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Player {
	return predicate.Player(sql.FieldContainsFold(FieldName, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSupervisors applies the HasEdge predicate on the "supervisors" edge.
func HasSupervisors() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SupervisorsTable, SupervisorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSupervisorsWith applies the HasEdge predicate on the "supervisors" edge with a given conditions (other predicates).
func HasSupervisorsWith(preds ...predicate.User) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newSupervisorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSupervisionRequests applies the HasEdge predicate on the "supervision_requests" edge.
func HasSupervisionRequests() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SupervisionRequestsTable, SupervisionRequestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSupervisionRequestsWith applies the HasEdge predicate on the "supervision_requests" edge with a given conditions (other predicates).
func HasSupervisionRequestsWith(preds ...predicate.PlayerSupervisionRequest) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newSupervisionRequestsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatches applies the HasEdge predicate on the "matches" edge.
func HasMatches() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MatchesTable, MatchesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchesWith applies the HasEdge predicate on the "matches" edge with a given conditions (other predicates).
func HasMatchesWith(preds ...predicate.Match) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newMatchesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStats applies the HasEdge predicate on the "stats" edge.
func HasStats() predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatsTable, StatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatsWith applies the HasEdge predicate on the "stats" edge with a given conditions (other predicates).
func HasStatsWith(preds ...predicate.Statistic) predicate.Player {
	return predicate.Player(func(s *sql.Selector) {
		step := newStatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Player) predicate.Player {
	return predicate.Player(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Player) predicate.Player {
	return predicate.Player(sql.NotPredicates(p))
}
