// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/gameversion"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
)

// MatchCreate is the builder for creating a Match entity.
type MatchCreate struct {
	config
	mutation *MatchMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MatchCreate) SetCreatedAt(t time.Time) *MatchCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MatchCreate) SetNillableCreatedAt(t *time.Time) *MatchCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MatchCreate) SetID(gu guidgql.GUID) *MatchCreate {
	mc.mutation.SetID(gu)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MatchCreate) SetNillableID(gu *guidgql.GUID) *MatchCreate {
	if gu != nil {
		mc.SetID(*gu)
	}
	return mc
}

// SetGameVersionID sets the "game_version" edge to the GameVersion entity by ID.
func (mc *MatchCreate) SetGameVersionID(id guidgql.GUID) *MatchCreate {
	mc.mutation.SetGameVersionID(id)
	return mc
}

// SetGameVersion sets the "game_version" edge to the GameVersion entity.
func (mc *MatchCreate) SetGameVersion(g *GameVersion) *MatchCreate {
	return mc.SetGameVersionID(g.ID)
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (mc *MatchCreate) AddPlayerIDs(ids ...guidgql.GUID) *MatchCreate {
	mc.mutation.AddPlayerIDs(ids...)
	return mc
}

// AddPlayers adds the "players" edges to the Player entity.
func (mc *MatchCreate) AddPlayers(p ...*Player) *MatchCreate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPlayerIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the Statistic entity by IDs.
func (mc *MatchCreate) AddStatIDs(ids ...guidgql.GUID) *MatchCreate {
	mc.mutation.AddStatIDs(ids...)
	return mc
}

// AddStats adds the "stats" edges to the Statistic entity.
func (mc *MatchCreate) AddStats(s ...*Statistic) *MatchCreate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddStatIDs(ids...)
}

// Mutation returns the MatchMutation object of the builder.
func (mc *MatchCreate) Mutation() *MatchMutation {
	return mc.mutation
}

// Save creates the Match in the database.
func (mc *MatchCreate) Save(ctx context.Context) (*Match, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MatchCreate) SaveX(ctx context.Context) *Match {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MatchCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MatchCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MatchCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := match.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := match.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MatchCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Match.created_at"`)}
	}
	if _, ok := mc.mutation.GameVersionID(); !ok {
		return &ValidationError{Name: "game_version", err: errors.New(`ent: missing required edge "Match.game_version"`)}
	}
	if len(mc.mutation.PlayersIDs()) == 0 {
		return &ValidationError{Name: "players", err: errors.New(`ent: missing required edge "Match.players"`)}
	}
	return nil
}

func (mc *MatchCreate) sqlSave(ctx context.Context) (*Match, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MatchCreate) createSpec() (*Match, *sqlgraph.CreateSpec) {
	var (
		_node = &Match{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(match.Table, sqlgraph.NewFieldSpec(match.FieldID, field.TypeString))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(match.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := mc.mutation.GameVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   match.GameVersionTable,
			Columns: []string{match.GameVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameversion.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.game_version_matches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   match.PlayersTable,
			Columns: match.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   match.StatsTable,
			Columns: []string{match.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statistic.FieldID, field.TypeString),
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
//	client.Match.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MatchUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MatchCreate) OnConflict(opts ...sql.ConflictOption) *MatchUpsertOne {
	mc.conflict = opts
	return &MatchUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Match.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MatchCreate) OnConflictColumns(columns ...string) *MatchUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MatchUpsertOne{
		create: mc,
	}
}

type (
	// MatchUpsertOne is the builder for "upsert"-ing
	//  one Match node.
	MatchUpsertOne struct {
		create *MatchCreate
	}

	// MatchUpsert is the "OnConflict" setter.
	MatchUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Match.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(match.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MatchUpsertOne) UpdateNewValues() *MatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(match.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(match.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Match.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MatchUpsertOne) Ignore() *MatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MatchUpsertOne) DoNothing() *MatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MatchCreate.OnConflict
// documentation for more info.
func (u *MatchUpsertOne) Update(set func(*MatchUpsert)) *MatchUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MatchUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *MatchUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MatchCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MatchUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MatchUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MatchUpsertOne.ID is not supported by MySQL driver. Use MatchUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MatchUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MatchCreateBulk is the builder for creating many Match entities in bulk.
type MatchCreateBulk struct {
	config
	err      error
	builders []*MatchCreate
	conflict []sql.ConflictOption
}

// Save creates the Match entities in the database.
func (mcb *MatchCreateBulk) Save(ctx context.Context) ([]*Match, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Match, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MatchMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MatchCreateBulk) SaveX(ctx context.Context) []*Match {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MatchCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MatchCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Match.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MatchUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MatchCreateBulk) OnConflict(opts ...sql.ConflictOption) *MatchUpsertBulk {
	mcb.conflict = opts
	return &MatchUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Match.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MatchCreateBulk) OnConflictColumns(columns ...string) *MatchUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MatchUpsertBulk{
		create: mcb,
	}
}

// MatchUpsertBulk is the builder for "upsert"-ing
// a bulk of Match nodes.
type MatchUpsertBulk struct {
	create *MatchCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Match.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(match.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MatchUpsertBulk) UpdateNewValues() *MatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(match.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(match.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Match.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MatchUpsertBulk) Ignore() *MatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MatchUpsertBulk) DoNothing() *MatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MatchCreateBulk.OnConflict
// documentation for more info.
func (u *MatchUpsertBulk) Update(set func(*MatchUpsert)) *MatchUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MatchUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *MatchUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MatchCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MatchCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MatchUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
