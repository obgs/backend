// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/enums"
	"github.com/obgs/backend/internal/ent/group"
	"github.com/obgs/backend/internal/ent/groupmembership"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/user"
)

// GroupMembershipCreate is the builder for creating a GroupMembership entity.
type GroupMembershipCreate struct {
	config
	mutation *GroupMembershipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRole sets the "role" field.
func (gmc *GroupMembershipCreate) SetRole(e enums.Role) *GroupMembershipCreate {
	gmc.mutation.SetRole(e)
	return gmc
}

// SetID sets the "id" field.
func (gmc *GroupMembershipCreate) SetID(gu guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetID(gu)
	return gmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableID(gu *guidgql.GUID) *GroupMembershipCreate {
	if gu != nil {
		gmc.SetID(*gu)
	}
	return gmc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmc *GroupMembershipCreate) SetGroupID(id guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetGroupID(id)
	return gmc
}

// SetGroup sets the "group" edge to the Group entity.
func (gmc *GroupMembershipCreate) SetGroup(g *Group) *GroupMembershipCreate {
	return gmc.SetGroupID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmc *GroupMembershipCreate) SetUserID(id guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetUserID(id)
	return gmc
}

// SetUser sets the "user" edge to the User entity.
func (gmc *GroupMembershipCreate) SetUser(u *User) *GroupMembershipCreate {
	return gmc.SetUserID(u.ID)
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmc *GroupMembershipCreate) Mutation() *GroupMembershipMutation {
	return gmc.mutation
}

// Save creates the GroupMembership in the database.
func (gmc *GroupMembershipCreate) Save(ctx context.Context) (*GroupMembership, error) {
	gmc.defaults()
	return withHooks(ctx, gmc.sqlSave, gmc.mutation, gmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMembershipCreate) SaveX(ctx context.Context) *GroupMembership {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMembershipCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMembershipCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmc *GroupMembershipCreate) defaults() {
	if _, ok := gmc.mutation.ID(); !ok {
		v := groupmembership.DefaultID()
		gmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMembershipCreate) check() error {
	if _, ok := gmc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "GroupMembership.role"`)}
	}
	if v, ok := gmc.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupMembership.group"`)}
	}
	if _, ok := gmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GroupMembership.user"`)}
	}
	return nil
}

func (gmc *GroupMembershipCreate) sqlSave(ctx context.Context) (*GroupMembership, error) {
	if err := gmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*guidgql.GUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	gmc.mutation.id = &_node.ID
	gmc.mutation.done = true
	return _node, nil
}

func (gmc *GroupMembershipCreate) createSpec() (*GroupMembership, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMembership{config: gmc.config}
		_spec = sqlgraph.NewCreateSpec(groupmembership.Table, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	)
	_spec.OnConflict = gmc.conflict
	if id, ok := gmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gmc.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := gmc.mutation.GroupIDs(); len(nodes) > 0 {
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
		_node.group_members = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_group_memberships = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupMembership.Create().
//		SetRole(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupMembershipUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (gmc *GroupMembershipCreate) OnConflict(opts ...sql.ConflictOption) *GroupMembershipUpsertOne {
	gmc.conflict = opts
	return &GroupMembershipUpsertOne{
		create: gmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gmc *GroupMembershipCreate) OnConflictColumns(columns ...string) *GroupMembershipUpsertOne {
	gmc.conflict = append(gmc.conflict, sql.ConflictColumns(columns...))
	return &GroupMembershipUpsertOne{
		create: gmc,
	}
}

type (
	// GroupMembershipUpsertOne is the builder for "upsert"-ing
	//  one GroupMembership node.
	GroupMembershipUpsertOne struct {
		create *GroupMembershipCreate
	}

	// GroupMembershipUpsert is the "OnConflict" setter.
	GroupMembershipUpsert struct {
		*sql.UpdateSet
	}
)

// SetRole sets the "role" field.
func (u *GroupMembershipUpsert) SetRole(v enums.Role) *GroupMembershipUpsert {
	u.Set(groupmembership.FieldRole, v)
	return u
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *GroupMembershipUpsert) UpdateRole() *GroupMembershipUpsert {
	u.SetExcluded(groupmembership.FieldRole)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(groupmembership.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupMembershipUpsertOne) UpdateNewValues() *GroupMembershipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(groupmembership.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupMembershipUpsertOne) Ignore() *GroupMembershipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupMembershipUpsertOne) DoNothing() *GroupMembershipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupMembershipCreate.OnConflict
// documentation for more info.
func (u *GroupMembershipUpsertOne) Update(set func(*GroupMembershipUpsert)) *GroupMembershipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupMembershipUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *GroupMembershipUpsertOne) SetRole(v enums.Role) *GroupMembershipUpsertOne {
	return u.Update(func(s *GroupMembershipUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *GroupMembershipUpsertOne) UpdateRole() *GroupMembershipUpsertOne {
	return u.Update(func(s *GroupMembershipUpsert) {
		s.UpdateRole()
	})
}

// Exec executes the query.
func (u *GroupMembershipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupMembershipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupMembershipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupMembershipUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GroupMembershipUpsertOne.ID is not supported by MySQL driver. Use GroupMembershipUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupMembershipUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupMembershipCreateBulk is the builder for creating many GroupMembership entities in bulk.
type GroupMembershipCreateBulk struct {
	config
	err      error
	builders []*GroupMembershipCreate
	conflict []sql.ConflictOption
}

// Save creates the GroupMembership entities in the database.
func (gmcb *GroupMembershipCreateBulk) Save(ctx context.Context) ([]*GroupMembership, error) {
	if gmcb.err != nil {
		return nil, gmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMembership, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMembershipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) SaveX(ctx context.Context) []*GroupMembership {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupMembership.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupMembershipUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (gmcb *GroupMembershipCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupMembershipUpsertBulk {
	gmcb.conflict = opts
	return &GroupMembershipUpsertBulk{
		create: gmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gmcb *GroupMembershipCreateBulk) OnConflictColumns(columns ...string) *GroupMembershipUpsertBulk {
	gmcb.conflict = append(gmcb.conflict, sql.ConflictColumns(columns...))
	return &GroupMembershipUpsertBulk{
		create: gmcb,
	}
}

// GroupMembershipUpsertBulk is the builder for "upsert"-ing
// a bulk of GroupMembership nodes.
type GroupMembershipUpsertBulk struct {
	create *GroupMembershipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(groupmembership.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupMembershipUpsertBulk) UpdateNewValues() *GroupMembershipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(groupmembership.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupMembership.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupMembershipUpsertBulk) Ignore() *GroupMembershipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupMembershipUpsertBulk) DoNothing() *GroupMembershipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupMembershipCreateBulk.OnConflict
// documentation for more info.
func (u *GroupMembershipUpsertBulk) Update(set func(*GroupMembershipUpsert)) *GroupMembershipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupMembershipUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *GroupMembershipUpsertBulk) SetRole(v enums.Role) *GroupMembershipUpsertBulk {
	return u.Update(func(s *GroupMembershipUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *GroupMembershipUpsertBulk) UpdateRole() *GroupMembershipUpsertBulk {
	return u.Update(func(s *GroupMembershipUpsert) {
		s.UpdateRole()
	})
}

// Exec executes the query.
func (u *GroupMembershipUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupMembershipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupMembershipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupMembershipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
