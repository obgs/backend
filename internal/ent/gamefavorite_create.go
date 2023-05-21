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
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/gamefavorite"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GameFavoriteCreate is the builder for creating a GameFavorite entity.
type GameFavoriteCreate struct {
	config
	mutation *GameFavoriteMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (gfc *GameFavoriteCreate) SetID(gu guidgql.GUID) *GameFavoriteCreate {
	gfc.mutation.SetID(gu)
	return gfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gfc *GameFavoriteCreate) SetNillableID(gu *guidgql.GUID) *GameFavoriteCreate {
	if gu != nil {
		gfc.SetID(*gu)
	}
	return gfc
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (gfc *GameFavoriteCreate) SetGameID(id guidgql.GUID) *GameFavoriteCreate {
	gfc.mutation.SetGameID(id)
	return gfc
}

// SetGame sets the "game" edge to the Game entity.
func (gfc *GameFavoriteCreate) SetGame(g *Game) *GameFavoriteCreate {
	return gfc.SetGameID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gfc *GameFavoriteCreate) SetUserID(id guidgql.GUID) *GameFavoriteCreate {
	gfc.mutation.SetUserID(id)
	return gfc
}

// SetUser sets the "user" edge to the User entity.
func (gfc *GameFavoriteCreate) SetUser(u *User) *GameFavoriteCreate {
	return gfc.SetUserID(u.ID)
}

// Mutation returns the GameFavoriteMutation object of the builder.
func (gfc *GameFavoriteCreate) Mutation() *GameFavoriteMutation {
	return gfc.mutation
}

// Save creates the GameFavorite in the database.
func (gfc *GameFavoriteCreate) Save(ctx context.Context) (*GameFavorite, error) {
	gfc.defaults()
	return withHooks(ctx, gfc.sqlSave, gfc.mutation, gfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gfc *GameFavoriteCreate) SaveX(ctx context.Context) *GameFavorite {
	v, err := gfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gfc *GameFavoriteCreate) Exec(ctx context.Context) error {
	_, err := gfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gfc *GameFavoriteCreate) ExecX(ctx context.Context) {
	if err := gfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gfc *GameFavoriteCreate) defaults() {
	if _, ok := gfc.mutation.ID(); !ok {
		v := gamefavorite.DefaultID()
		gfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gfc *GameFavoriteCreate) check() error {
	if _, ok := gfc.mutation.GameID(); !ok {
		return &ValidationError{Name: "game", err: errors.New(`ent: missing required edge "GameFavorite.game"`)}
	}
	if _, ok := gfc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GameFavorite.user"`)}
	}
	return nil
}

func (gfc *GameFavoriteCreate) sqlSave(ctx context.Context) (*GameFavorite, error) {
	if err := gfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gfc.driver, _spec); err != nil {
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
	gfc.mutation.id = &_node.ID
	gfc.mutation.done = true
	return _node, nil
}

func (gfc *GameFavoriteCreate) createSpec() (*GameFavorite, *sqlgraph.CreateSpec) {
	var (
		_node = &GameFavorite{config: gfc.config}
		_spec = sqlgraph.NewCreateSpec(gamefavorite.Table, sqlgraph.NewFieldSpec(gamefavorite.FieldID, field.TypeString))
	)
	_spec.OnConflict = gfc.conflict
	if id, ok := gfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := gfc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamefavorite.GameTable,
			Columns: []string{gamefavorite.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.game_favorites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gfc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamefavorite.UserTable,
			Columns: []string{gamefavorite.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_favorite_games = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GameFavorite.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (gfc *GameFavoriteCreate) OnConflict(opts ...sql.ConflictOption) *GameFavoriteUpsertOne {
	gfc.conflict = opts
	return &GameFavoriteUpsertOne{
		create: gfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gfc *GameFavoriteCreate) OnConflictColumns(columns ...string) *GameFavoriteUpsertOne {
	gfc.conflict = append(gfc.conflict, sql.ConflictColumns(columns...))
	return &GameFavoriteUpsertOne{
		create: gfc,
	}
}

type (
	// GameFavoriteUpsertOne is the builder for "upsert"-ing
	//  one GameFavorite node.
	GameFavoriteUpsertOne struct {
		create *GameFavoriteCreate
	}

	// GameFavoriteUpsert is the "OnConflict" setter.
	GameFavoriteUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gamefavorite.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GameFavoriteUpsertOne) UpdateNewValues() *GameFavoriteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(gamefavorite.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GameFavoriteUpsertOne) Ignore() *GameFavoriteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GameFavoriteUpsertOne) DoNothing() *GameFavoriteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GameFavoriteCreate.OnConflict
// documentation for more info.
func (u *GameFavoriteUpsertOne) Update(set func(*GameFavoriteUpsert)) *GameFavoriteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GameFavoriteUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GameFavoriteUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GameFavoriteCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GameFavoriteUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GameFavoriteUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GameFavoriteUpsertOne.ID is not supported by MySQL driver. Use GameFavoriteUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GameFavoriteUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GameFavoriteCreateBulk is the builder for creating many GameFavorite entities in bulk.
type GameFavoriteCreateBulk struct {
	config
	builders []*GameFavoriteCreate
	conflict []sql.ConflictOption
}

// Save creates the GameFavorite entities in the database.
func (gfcb *GameFavoriteCreateBulk) Save(ctx context.Context) ([]*GameFavorite, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gfcb.builders))
	nodes := make([]*GameFavorite, len(gfcb.builders))
	mutators := make([]Mutator, len(gfcb.builders))
	for i := range gfcb.builders {
		func(i int, root context.Context) {
			builder := gfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameFavoriteMutation)
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
					_, err = mutators[i+1].Mutate(root, gfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gfcb *GameFavoriteCreateBulk) SaveX(ctx context.Context) []*GameFavorite {
	v, err := gfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gfcb *GameFavoriteCreateBulk) Exec(ctx context.Context) error {
	_, err := gfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gfcb *GameFavoriteCreateBulk) ExecX(ctx context.Context) {
	if err := gfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GameFavorite.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (gfcb *GameFavoriteCreateBulk) OnConflict(opts ...sql.ConflictOption) *GameFavoriteUpsertBulk {
	gfcb.conflict = opts
	return &GameFavoriteUpsertBulk{
		create: gfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gfcb *GameFavoriteCreateBulk) OnConflictColumns(columns ...string) *GameFavoriteUpsertBulk {
	gfcb.conflict = append(gfcb.conflict, sql.ConflictColumns(columns...))
	return &GameFavoriteUpsertBulk{
		create: gfcb,
	}
}

// GameFavoriteUpsertBulk is the builder for "upsert"-ing
// a bulk of GameFavorite nodes.
type GameFavoriteUpsertBulk struct {
	create *GameFavoriteCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gamefavorite.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GameFavoriteUpsertBulk) UpdateNewValues() *GameFavoriteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(gamefavorite.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GameFavorite.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GameFavoriteUpsertBulk) Ignore() *GameFavoriteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GameFavoriteUpsertBulk) DoNothing() *GameFavoriteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GameFavoriteCreateBulk.OnConflict
// documentation for more info.
func (u *GameFavoriteUpsertBulk) Update(set func(*GameFavoriteUpsert)) *GameFavoriteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GameFavoriteUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GameFavoriteUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GameFavoriteCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GameFavoriteCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GameFavoriteUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
