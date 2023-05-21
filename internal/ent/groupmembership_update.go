// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GroupMembershipUpdate is the builder for updating GroupMembership entities.
type GroupMembershipUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMembershipMutation
}

// Where appends a list predicates to the GroupMembershipUpdate builder.
func (gmu *GroupMembershipUpdate) Where(ps ...predicate.GroupMembership) *GroupMembershipUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetRole sets the "role" field.
func (gmu *GroupMembershipUpdate) SetRole(e enums.Role) *GroupMembershipUpdate {
	gmu.mutation.SetRole(e)
	return gmu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmu *GroupMembershipUpdate) SetGroupID(id guidgql.GUID) *GroupMembershipUpdate {
	gmu.mutation.SetGroupID(id)
	return gmu
}

// SetGroup sets the "group" edge to the Group entity.
func (gmu *GroupMembershipUpdate) SetGroup(g *Group) *GroupMembershipUpdate {
	return gmu.SetGroupID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmu *GroupMembershipUpdate) SetUserID(id guidgql.GUID) *GroupMembershipUpdate {
	gmu.mutation.SetUserID(id)
	return gmu
}

// SetUser sets the "user" edge to the User entity.
func (gmu *GroupMembershipUpdate) SetUser(u *User) *GroupMembershipUpdate {
	return gmu.SetUserID(u.ID)
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmu *GroupMembershipUpdate) Mutation() *GroupMembershipMutation {
	return gmu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmu *GroupMembershipUpdate) ClearGroup() *GroupMembershipUpdate {
	gmu.mutation.ClearGroup()
	return gmu
}

// ClearUser clears the "user" edge to the User entity.
func (gmu *GroupMembershipUpdate) ClearUser() *GroupMembershipUpdate {
	gmu.mutation.ClearUser()
	return gmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GroupMembershipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GroupMembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GroupMembershipUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GroupMembershipUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmu *GroupMembershipUpdate) check() error {
	if v, ok := gmu.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmu.mutation.GroupID(); gmu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembership.group"`)
	}
	if _, ok := gmu.mutation.UserID(); gmu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembership.user"`)
	}
	return nil
}

func (gmu *GroupMembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmembership.Table, groupmembership.Columns, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
	}
	if gmu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GroupMembershipUpdateOne is the builder for updating a single GroupMembership entity.
type GroupMembershipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMembershipMutation
}

// SetRole sets the "role" field.
func (gmuo *GroupMembershipUpdateOne) SetRole(e enums.Role) *GroupMembershipUpdateOne {
	gmuo.mutation.SetRole(e)
	return gmuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmuo *GroupMembershipUpdateOne) SetGroupID(id guidgql.GUID) *GroupMembershipUpdateOne {
	gmuo.mutation.SetGroupID(id)
	return gmuo
}

// SetGroup sets the "group" edge to the Group entity.
func (gmuo *GroupMembershipUpdateOne) SetGroup(g *Group) *GroupMembershipUpdateOne {
	return gmuo.SetGroupID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmuo *GroupMembershipUpdateOne) SetUserID(id guidgql.GUID) *GroupMembershipUpdateOne {
	gmuo.mutation.SetUserID(id)
	return gmuo
}

// SetUser sets the "user" edge to the User entity.
func (gmuo *GroupMembershipUpdateOne) SetUser(u *User) *GroupMembershipUpdateOne {
	return gmuo.SetUserID(u.ID)
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmuo *GroupMembershipUpdateOne) Mutation() *GroupMembershipMutation {
	return gmuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmuo *GroupMembershipUpdateOne) ClearGroup() *GroupMembershipUpdateOne {
	gmuo.mutation.ClearGroup()
	return gmuo
}

// ClearUser clears the "user" edge to the User entity.
func (gmuo *GroupMembershipUpdateOne) ClearUser() *GroupMembershipUpdateOne {
	gmuo.mutation.ClearUser()
	return gmuo
}

// Where appends a list predicates to the GroupMembershipUpdate builder.
func (gmuo *GroupMembershipUpdateOne) Where(ps ...predicate.GroupMembership) *GroupMembershipUpdateOne {
	gmuo.mutation.Where(ps...)
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GroupMembershipUpdateOne) Select(field string, fields ...string) *GroupMembershipUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GroupMembership entity.
func (gmuo *GroupMembershipUpdateOne) Save(ctx context.Context) (*GroupMembership, error) {
	return withHooks(ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GroupMembershipUpdateOne) SaveX(ctx context.Context) *GroupMembership {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GroupMembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GroupMembershipUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmuo *GroupMembershipUpdateOne) check() error {
	if v, ok := gmuo.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmuo.mutation.GroupID(); gmuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembership.group"`)
	}
	if _, ok := gmuo.mutation.UserID(); gmuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMembership.user"`)
	}
	return nil
}

func (gmuo *GroupMembershipUpdateOne) sqlSave(ctx context.Context) (_node *GroupMembership, err error) {
	if err := gmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmembership.Table, groupmembership.Columns, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupMembership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembership.FieldID)
		for _, f := range fields {
			if !groupmembership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupmembership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
	}
	if gmuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupMembership{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}
