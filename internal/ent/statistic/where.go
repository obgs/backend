// Code generated by ent, DO NOT EDIT.

package statistic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Statistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Statistic {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// HasMatch applies the HasEdge predicate on the "match" edge.
func HasMatch() predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchWith applies the HasEdge predicate on the "match" edge with a given conditions (other predicates).
func HasMatchWith(preds ...predicate.Match) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatDescription applies the HasEdge predicate on the "stat_description" edge.
func HasStatDescription() predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatDescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatDescriptionTable, StatDescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatDescriptionWith applies the HasEdge predicate on the "stat_description" edge with a given conditions (other predicates).
func HasStatDescriptionWith(preds ...predicate.StatDescription) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatDescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatDescriptionTable, StatDescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlayer applies the HasEdge predicate on the "player" edge.
func HasPlayer() predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlayerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayerWith applies the HasEdge predicate on the "player" edge with a given conditions (other predicates).
func HasPlayerWith(preds ...predicate.Player) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlayerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlayerTable, PlayerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
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
func Not(p predicate.Statistic) predicate.Statistic {
	return predicate.Statistic(func(s *sql.Selector) {
		p(s.Not())
	})
}
