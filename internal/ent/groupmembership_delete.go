// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// GroupMembershipDelete is the builder for deleting a GroupMembership entity.
type GroupMembershipDelete struct {
	config
	hooks    []Hook
	mutation *GroupMembershipMutation
}

// Where appends a list predicates to the GroupMembershipDelete builder.
func (gmd *GroupMembershipDelete) Where(ps ...predicate.GroupMembership) *GroupMembershipDelete {
	gmd.mutation.Where(ps...)
	return gmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gmd *GroupMembershipDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gmd.hooks) == 0 {
		affected, err = gmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMembershipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gmd.mutation = mutation
			affected, err = gmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gmd.hooks) - 1; i >= 0; i-- {
			if gmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmd *GroupMembershipDelete) ExecX(ctx context.Context) int {
	n, err := gmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gmd *GroupMembershipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: groupmembership.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmembership.FieldID,
			},
		},
	}
	if ps := gmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GroupMembershipDeleteOne is the builder for deleting a single GroupMembership entity.
type GroupMembershipDeleteOne struct {
	gmd *GroupMembershipDelete
}

// Exec executes the deletion query.
func (gmdo *GroupMembershipDeleteOne) Exec(ctx context.Context) error {
	n, err := gmdo.gmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupmembership.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gmdo *GroupMembershipDeleteOne) ExecX(ctx context.Context) {
	gmdo.gmd.ExecX(ctx)
}