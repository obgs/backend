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
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerSupervisionRequestCreate is the builder for creating a PlayerSupervisionRequest entity.
type PlayerSupervisionRequestCreate struct {
	config
	mutation *PlayerSupervisionRequestMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetMessage sets the "message" field.
func (psrc *PlayerSupervisionRequestCreate) SetMessage(s string) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetMessage(s)
	return psrc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (psrc *PlayerSupervisionRequestCreate) SetNillableMessage(s *string) *PlayerSupervisionRequestCreate {
	if s != nil {
		psrc.SetMessage(*s)
	}
	return psrc
}

// SetID sets the "id" field.
func (psrc *PlayerSupervisionRequestCreate) SetID(gu guidgql.GUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetID(gu)
	return psrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psrc *PlayerSupervisionRequestCreate) SetNillableID(gu *guidgql.GUID) *PlayerSupervisionRequestCreate {
	if gu != nil {
		psrc.SetID(*gu)
	}
	return psrc
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (psrc *PlayerSupervisionRequestCreate) SetSenderID(id guidgql.GUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetSenderID(id)
	return psrc
}

// SetSender sets the "sender" edge to the User entity.
func (psrc *PlayerSupervisionRequestCreate) SetSender(u *User) *PlayerSupervisionRequestCreate {
	return psrc.SetSenderID(u.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (psrc *PlayerSupervisionRequestCreate) SetPlayerID(id guidgql.GUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetPlayerID(id)
	return psrc
}

// SetPlayer sets the "player" edge to the Player entity.
func (psrc *PlayerSupervisionRequestCreate) SetPlayer(p *Player) *PlayerSupervisionRequestCreate {
	return psrc.SetPlayerID(p.ID)
}

// AddApprovalIDs adds the "approvals" edge to the PlayerSupervisionRequestApproval entity by IDs.
func (psrc *PlayerSupervisionRequestCreate) AddApprovalIDs(ids ...guidgql.GUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.AddApprovalIDs(ids...)
	return psrc
}

// AddApprovals adds the "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psrc *PlayerSupervisionRequestCreate) AddApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestCreate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psrc.AddApprovalIDs(ids...)
}

// Mutation returns the PlayerSupervisionRequestMutation object of the builder.
func (psrc *PlayerSupervisionRequestCreate) Mutation() *PlayerSupervisionRequestMutation {
	return psrc.mutation
}

// Save creates the PlayerSupervisionRequest in the database.
func (psrc *PlayerSupervisionRequestCreate) Save(ctx context.Context) (*PlayerSupervisionRequest, error) {
	psrc.defaults()
	return withHooks(ctx, psrc.sqlSave, psrc.mutation, psrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psrc *PlayerSupervisionRequestCreate) SaveX(ctx context.Context) *PlayerSupervisionRequest {
	v, err := psrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psrc *PlayerSupervisionRequestCreate) Exec(ctx context.Context) error {
	_, err := psrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psrc *PlayerSupervisionRequestCreate) ExecX(ctx context.Context) {
	if err := psrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psrc *PlayerSupervisionRequestCreate) defaults() {
	if _, ok := psrc.mutation.ID(); !ok {
		v := playersupervisionrequest.DefaultID()
		psrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psrc *PlayerSupervisionRequestCreate) check() error {
	if _, ok := psrc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "PlayerSupervisionRequest.sender"`)}
	}
	if _, ok := psrc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "player", err: errors.New(`ent: missing required edge "PlayerSupervisionRequest.player"`)}
	}
	return nil
}

func (psrc *PlayerSupervisionRequestCreate) sqlSave(ctx context.Context) (*PlayerSupervisionRequest, error) {
	if err := psrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psrc.driver, _spec); err != nil {
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
	psrc.mutation.id = &_node.ID
	psrc.mutation.done = true
	return _node, nil
}

func (psrc *PlayerSupervisionRequestCreate) createSpec() (*PlayerSupervisionRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerSupervisionRequest{config: psrc.config}
		_spec = sqlgraph.NewCreateSpec(playersupervisionrequest.Table, sqlgraph.NewFieldSpec(playersupervisionrequest.FieldID, field.TypeString))
	)
	_spec.OnConflict = psrc.conflict
	if id, ok := psrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psrc.mutation.Message(); ok {
		_spec.SetField(playersupervisionrequest.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if nodes := psrc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequest.SenderTable,
			Columns: []string{playersupervisionrequest.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_sent_supervision_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psrc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequest.PlayerTable,
			Columns: []string{playersupervisionrequest.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_supervision_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psrc.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playersupervisionrequest.ApprovalsTable,
			Columns: []string{playersupervisionrequest.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playersupervisionrequestapproval.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerSupervisionRequest.Create().
//		SetMessage(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerSupervisionRequestUpsert) {
//			SetMessage(v+v).
//		}).
//		Exec(ctx)
func (psrc *PlayerSupervisionRequestCreate) OnConflict(opts ...sql.ConflictOption) *PlayerSupervisionRequestUpsertOne {
	psrc.conflict = opts
	return &PlayerSupervisionRequestUpsertOne{
		create: psrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psrc *PlayerSupervisionRequestCreate) OnConflictColumns(columns ...string) *PlayerSupervisionRequestUpsertOne {
	psrc.conflict = append(psrc.conflict, sql.ConflictColumns(columns...))
	return &PlayerSupervisionRequestUpsertOne{
		create: psrc,
	}
}

type (
	// PlayerSupervisionRequestUpsertOne is the builder for "upsert"-ing
	//  one PlayerSupervisionRequest node.
	PlayerSupervisionRequestUpsertOne struct {
		create *PlayerSupervisionRequestCreate
	}

	// PlayerSupervisionRequestUpsert is the "OnConflict" setter.
	PlayerSupervisionRequestUpsert struct {
		*sql.UpdateSet
	}
)

// SetMessage sets the "message" field.
func (u *PlayerSupervisionRequestUpsert) SetMessage(v string) *PlayerSupervisionRequestUpsert {
	u.Set(playersupervisionrequest.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *PlayerSupervisionRequestUpsert) UpdateMessage() *PlayerSupervisionRequestUpsert {
	u.SetExcluded(playersupervisionrequest.FieldMessage)
	return u
}

// ClearMessage clears the value of the "message" field.
func (u *PlayerSupervisionRequestUpsert) ClearMessage() *PlayerSupervisionRequestUpsert {
	u.SetNull(playersupervisionrequest.FieldMessage)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playersupervisionrequest.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerSupervisionRequestUpsertOne) UpdateNewValues() *PlayerSupervisionRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(playersupervisionrequest.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlayerSupervisionRequestUpsertOne) Ignore() *PlayerSupervisionRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerSupervisionRequestUpsertOne) DoNothing() *PlayerSupervisionRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerSupervisionRequestCreate.OnConflict
// documentation for more info.
func (u *PlayerSupervisionRequestUpsertOne) Update(set func(*PlayerSupervisionRequestUpsert)) *PlayerSupervisionRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerSupervisionRequestUpsert{UpdateSet: update})
	}))
	return u
}

// SetMessage sets the "message" field.
func (u *PlayerSupervisionRequestUpsertOne) SetMessage(v string) *PlayerSupervisionRequestUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *PlayerSupervisionRequestUpsertOne) UpdateMessage() *PlayerSupervisionRequestUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *PlayerSupervisionRequestUpsertOne) ClearMessage() *PlayerSupervisionRequestUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.ClearMessage()
	})
}

// Exec executes the query.
func (u *PlayerSupervisionRequestUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerSupervisionRequestCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerSupervisionRequestUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlayerSupervisionRequestUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlayerSupervisionRequestUpsertOne.ID is not supported by MySQL driver. Use PlayerSupervisionRequestUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlayerSupervisionRequestUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlayerSupervisionRequestCreateBulk is the builder for creating many PlayerSupervisionRequest entities in bulk.
type PlayerSupervisionRequestCreateBulk struct {
	config
	builders []*PlayerSupervisionRequestCreate
	conflict []sql.ConflictOption
}

// Save creates the PlayerSupervisionRequest entities in the database.
func (psrcb *PlayerSupervisionRequestCreateBulk) Save(ctx context.Context) ([]*PlayerSupervisionRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(psrcb.builders))
	nodes := make([]*PlayerSupervisionRequest, len(psrcb.builders))
	mutators := make([]Mutator, len(psrcb.builders))
	for i := range psrcb.builders {
		func(i int, root context.Context) {
			builder := psrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerSupervisionRequestMutation)
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
					_, err = mutators[i+1].Mutate(root, psrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = psrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, psrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psrcb *PlayerSupervisionRequestCreateBulk) SaveX(ctx context.Context) []*PlayerSupervisionRequest {
	v, err := psrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psrcb *PlayerSupervisionRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := psrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psrcb *PlayerSupervisionRequestCreateBulk) ExecX(ctx context.Context) {
	if err := psrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerSupervisionRequest.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerSupervisionRequestUpsert) {
//			SetMessage(v+v).
//		}).
//		Exec(ctx)
func (psrcb *PlayerSupervisionRequestCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlayerSupervisionRequestUpsertBulk {
	psrcb.conflict = opts
	return &PlayerSupervisionRequestUpsertBulk{
		create: psrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psrcb *PlayerSupervisionRequestCreateBulk) OnConflictColumns(columns ...string) *PlayerSupervisionRequestUpsertBulk {
	psrcb.conflict = append(psrcb.conflict, sql.ConflictColumns(columns...))
	return &PlayerSupervisionRequestUpsertBulk{
		create: psrcb,
	}
}

// PlayerSupervisionRequestUpsertBulk is the builder for "upsert"-ing
// a bulk of PlayerSupervisionRequest nodes.
type PlayerSupervisionRequestUpsertBulk struct {
	create *PlayerSupervisionRequestCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playersupervisionrequest.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerSupervisionRequestUpsertBulk) UpdateNewValues() *PlayerSupervisionRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(playersupervisionrequest.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequest.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlayerSupervisionRequestUpsertBulk) Ignore() *PlayerSupervisionRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerSupervisionRequestUpsertBulk) DoNothing() *PlayerSupervisionRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerSupervisionRequestCreateBulk.OnConflict
// documentation for more info.
func (u *PlayerSupervisionRequestUpsertBulk) Update(set func(*PlayerSupervisionRequestUpsert)) *PlayerSupervisionRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerSupervisionRequestUpsert{UpdateSet: update})
	}))
	return u
}

// SetMessage sets the "message" field.
func (u *PlayerSupervisionRequestUpsertBulk) SetMessage(v string) *PlayerSupervisionRequestUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *PlayerSupervisionRequestUpsertBulk) UpdateMessage() *PlayerSupervisionRequestUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *PlayerSupervisionRequestUpsertBulk) ClearMessage() *PlayerSupervisionRequestUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestUpsert) {
		s.ClearMessage()
	})
}

// Exec executes the query.
func (u *PlayerSupervisionRequestUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlayerSupervisionRequestCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerSupervisionRequestCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerSupervisionRequestUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
