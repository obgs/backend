// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
)

// StatDescriptionDelete is the builder for deleting a StatDescription entity.
type StatDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *StatDescriptionMutation
}

// Where appends a list predicates to the StatDescriptionDelete builder.
func (sdd *StatDescriptionDelete) Where(ps ...predicate.StatDescription) *StatDescriptionDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *StatDescriptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdd.hooks) == 0 {
		affected, err = sdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdd.mutation = mutation
			affected, err = sdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdd.hooks) - 1; i >= 0; i-- {
			if sdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *StatDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *StatDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: statdescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: statdescription.FieldID,
			},
		},
	}
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// StatDescriptionDeleteOne is the builder for deleting a single StatDescription entity.
type StatDescriptionDeleteOne struct {
	sdd *StatDescriptionDelete
}

// Exec executes the deletion query.
func (sddo *StatDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{statdescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *StatDescriptionDeleteOne) ExecX(ctx context.Context) {
	sddo.sdd.ExecX(ctx)
}