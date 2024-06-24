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
	"github.com/open-boardgame-stats/backend/internal/ent/gameversion"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GameCreate is the builder for creating a Game entity.
type GameCreate struct {
	config
	mutation *GameMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (gc *GameCreate) SetName(s string) *GameCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetMinPlayers sets the "min_players" field.
func (gc *GameCreate) SetMinPlayers(i int) *GameCreate {
	gc.mutation.SetMinPlayers(i)
	return gc
}

// SetNillableMinPlayers sets the "min_players" field if the given value is not nil.
func (gc *GameCreate) SetNillableMinPlayers(i *int) *GameCreate {
	if i != nil {
		gc.SetMinPlayers(*i)
	}
	return gc
}

// SetMaxPlayers sets the "max_players" field.
func (gc *GameCreate) SetMaxPlayers(i int) *GameCreate {
	gc.mutation.SetMaxPlayers(i)
	return gc
}

// SetNillableMaxPlayers sets the "max_players" field if the given value is not nil.
func (gc *GameCreate) SetNillableMaxPlayers(i *int) *GameCreate {
	if i != nil {
		gc.SetMaxPlayers(*i)
	}
	return gc
}

// SetDescription sets the "description" field.
func (gc *GameCreate) SetDescription(s string) *GameCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GameCreate) SetNillableDescription(s *string) *GameCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetBoardgamegeekURL sets the "boardgamegeek_url" field.
func (gc *GameCreate) SetBoardgamegeekURL(s string) *GameCreate {
	gc.mutation.SetBoardgamegeekURL(s)
	return gc
}

// SetNillableBoardgamegeekURL sets the "boardgamegeek_url" field if the given value is not nil.
func (gc *GameCreate) SetNillableBoardgamegeekURL(s *string) *GameCreate {
	if s != nil {
		gc.SetBoardgamegeekURL(*s)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GameCreate) SetID(gu guidgql.GUID) *GameCreate {
	gc.mutation.SetID(gu)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GameCreate) SetNillableID(gu *guidgql.GUID) *GameCreate {
	if gu != nil {
		gc.SetID(*gu)
	}
	return gc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (gc *GameCreate) SetAuthorID(id guidgql.GUID) *GameCreate {
	gc.mutation.SetAuthorID(id)
	return gc
}

// SetAuthor sets the "author" edge to the User entity.
func (gc *GameCreate) SetAuthor(u *User) *GameCreate {
	return gc.SetAuthorID(u.ID)
}

// AddFavoriteIDs adds the "favorites" edge to the GameFavorite entity by IDs.
func (gc *GameCreate) AddFavoriteIDs(ids ...guidgql.GUID) *GameCreate {
	gc.mutation.AddFavoriteIDs(ids...)
	return gc
}

// AddFavorites adds the "favorites" edges to the GameFavorite entity.
func (gc *GameCreate) AddFavorites(g ...*GameFavorite) *GameCreate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddFavoriteIDs(ids...)
}

// AddVersionIDs adds the "versions" edge to the GameVersion entity by IDs.
func (gc *GameCreate) AddVersionIDs(ids ...guidgql.GUID) *GameCreate {
	gc.mutation.AddVersionIDs(ids...)
	return gc
}

// AddVersions adds the "versions" edges to the GameVersion entity.
func (gc *GameCreate) AddVersions(g ...*GameVersion) *GameCreate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddVersionIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (gc *GameCreate) Mutation() *GameMutation {
	return gc.mutation
}

// Save creates the Game in the database.
func (gc *GameCreate) Save(ctx context.Context) (*Game, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GameCreate) SaveX(ctx context.Context) *Game {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GameCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GameCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GameCreate) defaults() {
	if _, ok := gc.mutation.MinPlayers(); !ok {
		v := game.DefaultMinPlayers
		gc.mutation.SetMinPlayers(v)
	}
	if _, ok := gc.mutation.MaxPlayers(); !ok {
		v := game.DefaultMaxPlayers
		gc.mutation.SetMaxPlayers(v)
	}
	if _, ok := gc.mutation.Description(); !ok {
		v := game.DefaultDescription
		gc.mutation.SetDescription(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := game.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GameCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Game.name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := game.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Game.name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.MinPlayers(); !ok {
		return &ValidationError{Name: "min_players", err: errors.New(`ent: missing required field "Game.min_players"`)}
	}
	if _, ok := gc.mutation.MaxPlayers(); !ok {
		return &ValidationError{Name: "max_players", err: errors.New(`ent: missing required field "Game.max_players"`)}
	}
	if _, ok := gc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Game.author"`)}
	}
	return nil
}

func (gc *GameCreate) sqlSave(ctx context.Context) (*Game, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
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
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GameCreate) createSpec() (*Game, *sqlgraph.CreateSpec) {
	var (
		_node = &Game{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(game.Table, sqlgraph.NewFieldSpec(game.FieldID, field.TypeString))
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(game.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.MinPlayers(); ok {
		_spec.SetField(game.FieldMinPlayers, field.TypeInt, value)
		_node.MinPlayers = value
	}
	if value, ok := gc.mutation.MaxPlayers(); ok {
		_spec.SetField(game.FieldMaxPlayers, field.TypeInt, value)
		_node.MaxPlayers = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(game.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.BoardgamegeekURL(); ok {
		_spec.SetField(game.FieldBoardgamegeekURL, field.TypeString, value)
		_node.BoardgamegeekURL = value
	}
	if nodes := gc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.AuthorTable,
			Columns: []string{game.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_games = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.FavoritesTable,
			Columns: []string{game.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamefavorite.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.VersionsTable,
			Columns: []string{game.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameversion.FieldID, field.TypeString),
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
//	client.Game.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GameUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gc *GameCreate) OnConflict(opts ...sql.ConflictOption) *GameUpsertOne {
	gc.conflict = opts
	return &GameUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Game.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GameCreate) OnConflictColumns(columns ...string) *GameUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GameUpsertOne{
		create: gc,
	}
}

type (
	// GameUpsertOne is the builder for "upsert"-ing
	//  one Game node.
	GameUpsertOne struct {
		create *GameCreate
	}

	// GameUpsert is the "OnConflict" setter.
	GameUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *GameUpsert) SetName(v string) *GameUpsert {
	u.Set(game.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GameUpsert) UpdateName() *GameUpsert {
	u.SetExcluded(game.FieldName)
	return u
}

// SetMinPlayers sets the "min_players" field.
func (u *GameUpsert) SetMinPlayers(v int) *GameUpsert {
	u.Set(game.FieldMinPlayers, v)
	return u
}

// UpdateMinPlayers sets the "min_players" field to the value that was provided on create.
func (u *GameUpsert) UpdateMinPlayers() *GameUpsert {
	u.SetExcluded(game.FieldMinPlayers)
	return u
}

// AddMinPlayers adds v to the "min_players" field.
func (u *GameUpsert) AddMinPlayers(v int) *GameUpsert {
	u.Add(game.FieldMinPlayers, v)
	return u
}

// SetMaxPlayers sets the "max_players" field.
func (u *GameUpsert) SetMaxPlayers(v int) *GameUpsert {
	u.Set(game.FieldMaxPlayers, v)
	return u
}

// UpdateMaxPlayers sets the "max_players" field to the value that was provided on create.
func (u *GameUpsert) UpdateMaxPlayers() *GameUpsert {
	u.SetExcluded(game.FieldMaxPlayers)
	return u
}

// AddMaxPlayers adds v to the "max_players" field.
func (u *GameUpsert) AddMaxPlayers(v int) *GameUpsert {
	u.Add(game.FieldMaxPlayers, v)
	return u
}

// SetDescription sets the "description" field.
func (u *GameUpsert) SetDescription(v string) *GameUpsert {
	u.Set(game.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GameUpsert) UpdateDescription() *GameUpsert {
	u.SetExcluded(game.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *GameUpsert) ClearDescription() *GameUpsert {
	u.SetNull(game.FieldDescription)
	return u
}

// SetBoardgamegeekURL sets the "boardgamegeek_url" field.
func (u *GameUpsert) SetBoardgamegeekURL(v string) *GameUpsert {
	u.Set(game.FieldBoardgamegeekURL, v)
	return u
}

// UpdateBoardgamegeekURL sets the "boardgamegeek_url" field to the value that was provided on create.
func (u *GameUpsert) UpdateBoardgamegeekURL() *GameUpsert {
	u.SetExcluded(game.FieldBoardgamegeekURL)
	return u
}

// ClearBoardgamegeekURL clears the value of the "boardgamegeek_url" field.
func (u *GameUpsert) ClearBoardgamegeekURL() *GameUpsert {
	u.SetNull(game.FieldBoardgamegeekURL)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Game.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(game.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GameUpsertOne) UpdateNewValues() *GameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(game.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Game.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GameUpsertOne) Ignore() *GameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GameUpsertOne) DoNothing() *GameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GameCreate.OnConflict
// documentation for more info.
func (u *GameUpsertOne) Update(set func(*GameUpsert)) *GameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GameUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GameUpsertOne) SetName(v string) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GameUpsertOne) UpdateName() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.UpdateName()
	})
}

// SetMinPlayers sets the "min_players" field.
func (u *GameUpsertOne) SetMinPlayers(v int) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.SetMinPlayers(v)
	})
}

// AddMinPlayers adds v to the "min_players" field.
func (u *GameUpsertOne) AddMinPlayers(v int) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.AddMinPlayers(v)
	})
}

// UpdateMinPlayers sets the "min_players" field to the value that was provided on create.
func (u *GameUpsertOne) UpdateMinPlayers() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.UpdateMinPlayers()
	})
}

// SetMaxPlayers sets the "max_players" field.
func (u *GameUpsertOne) SetMaxPlayers(v int) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.SetMaxPlayers(v)
	})
}

// AddMaxPlayers adds v to the "max_players" field.
func (u *GameUpsertOne) AddMaxPlayers(v int) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.AddMaxPlayers(v)
	})
}

// UpdateMaxPlayers sets the "max_players" field to the value that was provided on create.
func (u *GameUpsertOne) UpdateMaxPlayers() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.UpdateMaxPlayers()
	})
}

// SetDescription sets the "description" field.
func (u *GameUpsertOne) SetDescription(v string) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GameUpsertOne) UpdateDescription() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *GameUpsertOne) ClearDescription() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.ClearDescription()
	})
}

// SetBoardgamegeekURL sets the "boardgamegeek_url" field.
func (u *GameUpsertOne) SetBoardgamegeekURL(v string) *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.SetBoardgamegeekURL(v)
	})
}

// UpdateBoardgamegeekURL sets the "boardgamegeek_url" field to the value that was provided on create.
func (u *GameUpsertOne) UpdateBoardgamegeekURL() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.UpdateBoardgamegeekURL()
	})
}

// ClearBoardgamegeekURL clears the value of the "boardgamegeek_url" field.
func (u *GameUpsertOne) ClearBoardgamegeekURL() *GameUpsertOne {
	return u.Update(func(s *GameUpsert) {
		s.ClearBoardgamegeekURL()
	})
}

// Exec executes the query.
func (u *GameUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GameCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GameUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GameUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GameUpsertOne.ID is not supported by MySQL driver. Use GameUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GameUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GameCreateBulk is the builder for creating many Game entities in bulk.
type GameCreateBulk struct {
	config
	err      error
	builders []*GameCreate
	conflict []sql.ConflictOption
}

// Save creates the Game entities in the database.
func (gcb *GameCreateBulk) Save(ctx context.Context) ([]*Game, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Game, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GameCreateBulk) SaveX(ctx context.Context) []*Game {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GameCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GameCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Game.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GameUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gcb *GameCreateBulk) OnConflict(opts ...sql.ConflictOption) *GameUpsertBulk {
	gcb.conflict = opts
	return &GameUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Game.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GameCreateBulk) OnConflictColumns(columns ...string) *GameUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GameUpsertBulk{
		create: gcb,
	}
}

// GameUpsertBulk is the builder for "upsert"-ing
// a bulk of Game nodes.
type GameUpsertBulk struct {
	create *GameCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Game.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(game.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GameUpsertBulk) UpdateNewValues() *GameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(game.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Game.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GameUpsertBulk) Ignore() *GameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GameUpsertBulk) DoNothing() *GameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GameCreateBulk.OnConflict
// documentation for more info.
func (u *GameUpsertBulk) Update(set func(*GameUpsert)) *GameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GameUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GameUpsertBulk) SetName(v string) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GameUpsertBulk) UpdateName() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.UpdateName()
	})
}

// SetMinPlayers sets the "min_players" field.
func (u *GameUpsertBulk) SetMinPlayers(v int) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.SetMinPlayers(v)
	})
}

// AddMinPlayers adds v to the "min_players" field.
func (u *GameUpsertBulk) AddMinPlayers(v int) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.AddMinPlayers(v)
	})
}

// UpdateMinPlayers sets the "min_players" field to the value that was provided on create.
func (u *GameUpsertBulk) UpdateMinPlayers() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.UpdateMinPlayers()
	})
}

// SetMaxPlayers sets the "max_players" field.
func (u *GameUpsertBulk) SetMaxPlayers(v int) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.SetMaxPlayers(v)
	})
}

// AddMaxPlayers adds v to the "max_players" field.
func (u *GameUpsertBulk) AddMaxPlayers(v int) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.AddMaxPlayers(v)
	})
}

// UpdateMaxPlayers sets the "max_players" field to the value that was provided on create.
func (u *GameUpsertBulk) UpdateMaxPlayers() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.UpdateMaxPlayers()
	})
}

// SetDescription sets the "description" field.
func (u *GameUpsertBulk) SetDescription(v string) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GameUpsertBulk) UpdateDescription() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *GameUpsertBulk) ClearDescription() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.ClearDescription()
	})
}

// SetBoardgamegeekURL sets the "boardgamegeek_url" field.
func (u *GameUpsertBulk) SetBoardgamegeekURL(v string) *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.SetBoardgamegeekURL(v)
	})
}

// UpdateBoardgamegeekURL sets the "boardgamegeek_url" field to the value that was provided on create.
func (u *GameUpsertBulk) UpdateBoardgamegeekURL() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.UpdateBoardgamegeekURL()
	})
}

// ClearBoardgamegeekURL clears the value of the "boardgamegeek_url" field.
func (u *GameUpsertBulk) ClearBoardgamegeekURL() *GameUpsertBulk {
	return u.Update(func(s *GameUpsert) {
		s.ClearBoardgamegeekURL()
	})
}

// Exec executes the query.
func (u *GameUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GameCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GameCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GameUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
