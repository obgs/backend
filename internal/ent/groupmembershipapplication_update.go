// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GroupMembershipApplicationUpdate is the builder for updating GroupMembershipApplication entities.
type GroupMembershipApplicationUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMembershipApplicationMutation
}

// Where appends a list predicates to the GroupMembershipApplicationUpdate builder.
func (gmau *GroupMembershipApplicationUpdate) Where(ps ...predicate.GroupMembershipApplication) *GroupMembershipApplicationUpdate {
	gmau.mutation.Where(ps...)
	return gmau
}

// SetMessage sets the "message" field.
func (gmau *GroupMembershipApplicationUpdate) SetMessage(s string) *GroupMembershipApplicationUpdate {
	gmau.mutation.SetMessage(s)
	return gmau
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (gmau *GroupMembershipApplicationUpdate) SetNillableMessage(s *string) *GroupMembershipApplicationUpdate {
	if s != nil {
		gmau.SetMessage(*s)
	}
	return gmau
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmau *GroupMembershipApplicationUpdate) SetUserID(id guidgql.GUID) *GroupMembershipApplicationUpdate {
	gmau.mutation.SetUserID(id)
	return gmau
}

// SetUser sets the "user" edge to the User entity.
func (gmau *GroupMembershipApplicationUpdate) SetUser(u *User) *GroupMembershipApplicationUpdate {
	return gmau.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmau *GroupMembershipApplicationUpdate) SetGroupID(id guidgql.GUID) *GroupMembershipApplicationUpdate {
	gmau.mutation.SetGroupID(id)
	return gmau
}

// SetGroup sets the "group" edge to the Group entity.
func (gmau *GroupMembershipApplicationUpdate) SetGroup(g *Group) *GroupMembershipApplicationUpdate {
	return gmau.SetGroupID(g.ID)
}

// Mutation returns the GroupMembershipApplicationMutation object of the builder.
func (gmau *GroupMembershipApplicationUpdate) Mutation() *GroupMembershipApplicationMutation {
	return gmau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gmau *GroupMembershipApplicationUpdate) ClearUser() *GroupMembershipApplicationUpdate {
	gmau.mutation.ClearUser()
	return gmau
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmau *GroupMembershipApplicationUpdate) ClearGroup() *GroupMembershipApplicationUpdate {
	gmau.mutation.ClearGroup()
	return gmau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmau *GroupMembershipApplicationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gmau.hooks) == 0 {
		if err = gmau.check(); err != nil {
			return 0, err
		}
		affected, err = gmau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMembershipApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmau.check(); err != nil {
				return 0, err
			}
			gmau.mutation = mutation
			affected, err = gmau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gmau.hooks) - 1; i >= 0; i-- {
			if gmau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gmau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gmau *GroupMembershipApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := gmau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmau *GroupMembershipApplicationUpdate) Exec(ctx context.Context) error {
	_, err := gmau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmau *GroupMembershipApplicationUpdate) ExecX(ctx context.Context) {
	if err := gmau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmau *GroupMembershipApplicationUpdate) check() error {
	if _, ok := gmau.mutation.UserID(); gmau.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembershipApplication.user"`)
	}
	if _, ok := gmau.mutation.GroupID(); gmau.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembershipApplication.group"`)
	}
	return nil
}

func (gmau *GroupMembershipApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupmembershipapplication.Table,
			Columns: groupmembershipapplication.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmembershipapplication.FieldID,
			},
		},
	}
	if ps := gmau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmau.mutation.Message(); ok {
		_spec.SetField(groupmembershipapplication.FieldMessage, field.TypeString, value)
	}
	if gmau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.UserTable,
			Columns: []string{groupmembershipapplication.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.UserTable,
			Columns: []string{groupmembershipapplication.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmau.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.GroupTable,
			Columns: []string{groupmembershipapplication.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmau.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.GroupTable,
			Columns: []string{groupmembershipapplication.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gmau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembershipapplication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GroupMembershipApplicationUpdateOne is the builder for updating a single GroupMembershipApplication entity.
type GroupMembershipApplicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMembershipApplicationMutation
}

// SetMessage sets the "message" field.
func (gmauo *GroupMembershipApplicationUpdateOne) SetMessage(s string) *GroupMembershipApplicationUpdateOne {
	gmauo.mutation.SetMessage(s)
	return gmauo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (gmauo *GroupMembershipApplicationUpdateOne) SetNillableMessage(s *string) *GroupMembershipApplicationUpdateOne {
	if s != nil {
		gmauo.SetMessage(*s)
	}
	return gmauo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmauo *GroupMembershipApplicationUpdateOne) SetUserID(id guidgql.GUID) *GroupMembershipApplicationUpdateOne {
	gmauo.mutation.SetUserID(id)
	return gmauo
}

// SetUser sets the "user" edge to the User entity.
func (gmauo *GroupMembershipApplicationUpdateOne) SetUser(u *User) *GroupMembershipApplicationUpdateOne {
	return gmauo.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmauo *GroupMembershipApplicationUpdateOne) SetGroupID(id guidgql.GUID) *GroupMembershipApplicationUpdateOne {
	gmauo.mutation.SetGroupID(id)
	return gmauo
}

// SetGroup sets the "group" edge to the Group entity.
func (gmauo *GroupMembershipApplicationUpdateOne) SetGroup(g *Group) *GroupMembershipApplicationUpdateOne {
	return gmauo.SetGroupID(g.ID)
}

// Mutation returns the GroupMembershipApplicationMutation object of the builder.
func (gmauo *GroupMembershipApplicationUpdateOne) Mutation() *GroupMembershipApplicationMutation {
	return gmauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gmauo *GroupMembershipApplicationUpdateOne) ClearUser() *GroupMembershipApplicationUpdateOne {
	gmauo.mutation.ClearUser()
	return gmauo
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmauo *GroupMembershipApplicationUpdateOne) ClearGroup() *GroupMembershipApplicationUpdateOne {
	gmauo.mutation.ClearGroup()
	return gmauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmauo *GroupMembershipApplicationUpdateOne) Select(field string, fields ...string) *GroupMembershipApplicationUpdateOne {
	gmauo.fields = append([]string{field}, fields...)
	return gmauo
}

// Save executes the query and returns the updated GroupMembershipApplication entity.
func (gmauo *GroupMembershipApplicationUpdateOne) Save(ctx context.Context) (*GroupMembershipApplication, error) {
	var (
		err  error
		node *GroupMembershipApplication
	)
	if len(gmauo.hooks) == 0 {
		if err = gmauo.check(); err != nil {
			return nil, err
		}
		node, err = gmauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMembershipApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmauo.check(); err != nil {
				return nil, err
			}
			gmauo.mutation = mutation
			node, err = gmauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gmauo.hooks) - 1; i >= 0; i-- {
			if gmauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gmauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GroupMembershipApplication)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GroupMembershipApplicationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gmauo *GroupMembershipApplicationUpdateOne) SaveX(ctx context.Context) *GroupMembershipApplication {
	node, err := gmauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmauo *GroupMembershipApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := gmauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmauo *GroupMembershipApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := gmauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmauo *GroupMembershipApplicationUpdateOne) check() error {
	if _, ok := gmauo.mutation.UserID(); gmauo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembershipApplication.user"`)
	}
	if _, ok := gmauo.mutation.GroupID(); gmauo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembershipApplication.group"`)
	}
	return nil
}

func (gmauo *GroupMembershipApplicationUpdateOne) sqlSave(ctx context.Context) (_node *GroupMembershipApplication, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupmembershipapplication.Table,
			Columns: groupmembershipapplication.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmembershipapplication.FieldID,
			},
		},
	}
	id, ok := gmauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupMembershipApplication.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembershipapplication.FieldID)
		for _, f := range fields {
			if !groupmembershipapplication.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupmembershipapplication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmauo.mutation.Message(); ok {
		_spec.SetField(groupmembershipapplication.FieldMessage, field.TypeString, value)
	}
	if gmauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.UserTable,
			Columns: []string{groupmembershipapplication.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.UserTable,
			Columns: []string{groupmembershipapplication.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmauo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.GroupTable,
			Columns: []string{groupmembershipapplication.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmauo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembershipapplication.GroupTable,
			Columns: []string{groupmembershipapplication.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupMembershipApplication{config: gmauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembershipapplication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
