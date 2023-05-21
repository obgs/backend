// Code generated by ent, DO NOT EDIT.

package gamefavorite

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.GameFavorite {
	return predicate.GameFavorite(sql.FieldLTE(FieldID, id))
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GameTable, GameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		step := newGameStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GameFavorite) predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GameFavorite) predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
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
func Not(p predicate.GameFavorite) predicate.GameFavorite {
	return predicate.GameFavorite(func(s *sql.Selector) {
		p(s.Not())
	})
}
