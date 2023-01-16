// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
)

// StatDescriptionUpdate is the builder for updating StatDescription entities.
type StatDescriptionUpdate struct {
	config
	hooks    []Hook
	mutation *StatDescriptionMutation
}

// Where appends a list predicates to the StatDescriptionUpdate builder.
func (sdu *StatDescriptionUpdate) Where(ps ...predicate.StatDescription) *StatDescriptionUpdate {
	sdu.mutation.Where(ps...)
	return sdu
}

// SetType sets the "type" field.
func (sdu *StatDescriptionUpdate) SetType(st stat.StatType) *StatDescriptionUpdate {
	sdu.mutation.SetType(st)
	return sdu
}

// SetName sets the "name" field.
func (sdu *StatDescriptionUpdate) SetName(s string) *StatDescriptionUpdate {
	sdu.mutation.SetName(s)
	return sdu
}

// SetDescription sets the "description" field.
func (sdu *StatDescriptionUpdate) SetDescription(s string) *StatDescriptionUpdate {
	sdu.mutation.SetDescription(s)
	return sdu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sdu *StatDescriptionUpdate) SetNillableDescription(s *string) *StatDescriptionUpdate {
	if s != nil {
		sdu.SetDescription(*s)
	}
	return sdu
}

// ClearDescription clears the value of the "description" field.
func (sdu *StatDescriptionUpdate) ClearDescription() *StatDescriptionUpdate {
	sdu.mutation.ClearDescription()
	return sdu
}

// SetMetadata sets the "metadata" field.
func (sdu *StatDescriptionUpdate) SetMetadata(s string) *StatDescriptionUpdate {
	sdu.mutation.SetMetadata(s)
	return sdu
}

// SetNillableMetadata sets the "metadata" field if the given value is not nil.
func (sdu *StatDescriptionUpdate) SetNillableMetadata(s *string) *StatDescriptionUpdate {
	if s != nil {
		sdu.SetMetadata(*s)
	}
	return sdu
}

// ClearMetadata clears the value of the "metadata" field.
func (sdu *StatDescriptionUpdate) ClearMetadata() *StatDescriptionUpdate {
	sdu.mutation.ClearMetadata()
	return sdu
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (sdu *StatDescriptionUpdate) AddGameIDs(ids ...guidgql.GUID) *StatDescriptionUpdate {
	sdu.mutation.AddGameIDs(ids...)
	return sdu
}

// AddGame adds the "game" edges to the Game entity.
func (sdu *StatDescriptionUpdate) AddGame(g ...*Game) *StatDescriptionUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sdu.AddGameIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the Statistic entity by IDs.
func (sdu *StatDescriptionUpdate) AddStatIDs(ids ...guidgql.GUID) *StatDescriptionUpdate {
	sdu.mutation.AddStatIDs(ids...)
	return sdu
}

// AddStats adds the "stats" edges to the Statistic entity.
func (sdu *StatDescriptionUpdate) AddStats(s ...*Statistic) *StatDescriptionUpdate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdu.AddStatIDs(ids...)
}

// Mutation returns the StatDescriptionMutation object of the builder.
func (sdu *StatDescriptionUpdate) Mutation() *StatDescriptionMutation {
	return sdu.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (sdu *StatDescriptionUpdate) ClearGame() *StatDescriptionUpdate {
	sdu.mutation.ClearGame()
	return sdu
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (sdu *StatDescriptionUpdate) RemoveGameIDs(ids ...guidgql.GUID) *StatDescriptionUpdate {
	sdu.mutation.RemoveGameIDs(ids...)
	return sdu
}

// RemoveGame removes "game" edges to Game entities.
func (sdu *StatDescriptionUpdate) RemoveGame(g ...*Game) *StatDescriptionUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sdu.RemoveGameIDs(ids...)
}

// ClearStats clears all "stats" edges to the Statistic entity.
func (sdu *StatDescriptionUpdate) ClearStats() *StatDescriptionUpdate {
	sdu.mutation.ClearStats()
	return sdu
}

// RemoveStatIDs removes the "stats" edge to Statistic entities by IDs.
func (sdu *StatDescriptionUpdate) RemoveStatIDs(ids ...guidgql.GUID) *StatDescriptionUpdate {
	sdu.mutation.RemoveStatIDs(ids...)
	return sdu
}

// RemoveStats removes "stats" edges to Statistic entities.
func (sdu *StatDescriptionUpdate) RemoveStats(s ...*Statistic) *StatDescriptionUpdate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdu.RemoveStatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *StatDescriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdu.hooks) == 0 {
		if err = sdu.check(); err != nil {
			return 0, err
		}
		affected, err = sdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdu.check(); err != nil {
				return 0, err
			}
			sdu.mutation = mutation
			affected, err = sdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdu.hooks) - 1; i >= 0; i-- {
			if sdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *StatDescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *StatDescriptionUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *StatDescriptionUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdu *StatDescriptionUpdate) check() error {
	if v, ok := sdu.mutation.GetType(); ok {
		if err := statdescription.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "StatDescription.type": %w`, err)}
		}
	}
	if v, ok := sdu.mutation.Name(); ok {
		if err := statdescription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "StatDescription.name": %w`, err)}
		}
	}
	return nil
}

func (sdu *StatDescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statdescription.Table,
			Columns: statdescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: statdescription.FieldID,
			},
		},
	}
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.GetType(); ok {
		_spec.SetField(statdescription.FieldType, field.TypeEnum, value)
	}
	if value, ok := sdu.mutation.Name(); ok {
		_spec.SetField(statdescription.FieldName, field.TypeString, value)
	}
	if value, ok := sdu.mutation.Description(); ok {
		_spec.SetField(statdescription.FieldDescription, field.TypeString, value)
	}
	if sdu.mutation.DescriptionCleared() {
		_spec.ClearField(statdescription.FieldDescription, field.TypeString)
	}
	if value, ok := sdu.mutation.Metadata(); ok {
		_spec.SetField(statdescription.FieldMetadata, field.TypeString, value)
	}
	if sdu.mutation.MetadataCleared() {
		_spec.ClearField(statdescription.FieldMetadata, field.TypeString)
	}
	if sdu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.RemovedGameIDs(); len(nodes) > 0 && !sdu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sdu.mutation.StatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.RemovedStatsIDs(); len(nodes) > 0 && !sdu.mutation.StatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statdescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StatDescriptionUpdateOne is the builder for updating a single StatDescription entity.
type StatDescriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatDescriptionMutation
}

// SetType sets the "type" field.
func (sduo *StatDescriptionUpdateOne) SetType(st stat.StatType) *StatDescriptionUpdateOne {
	sduo.mutation.SetType(st)
	return sduo
}

// SetName sets the "name" field.
func (sduo *StatDescriptionUpdateOne) SetName(s string) *StatDescriptionUpdateOne {
	sduo.mutation.SetName(s)
	return sduo
}

// SetDescription sets the "description" field.
func (sduo *StatDescriptionUpdateOne) SetDescription(s string) *StatDescriptionUpdateOne {
	sduo.mutation.SetDescription(s)
	return sduo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sduo *StatDescriptionUpdateOne) SetNillableDescription(s *string) *StatDescriptionUpdateOne {
	if s != nil {
		sduo.SetDescription(*s)
	}
	return sduo
}

// ClearDescription clears the value of the "description" field.
func (sduo *StatDescriptionUpdateOne) ClearDescription() *StatDescriptionUpdateOne {
	sduo.mutation.ClearDescription()
	return sduo
}

// SetMetadata sets the "metadata" field.
func (sduo *StatDescriptionUpdateOne) SetMetadata(s string) *StatDescriptionUpdateOne {
	sduo.mutation.SetMetadata(s)
	return sduo
}

// SetNillableMetadata sets the "metadata" field if the given value is not nil.
func (sduo *StatDescriptionUpdateOne) SetNillableMetadata(s *string) *StatDescriptionUpdateOne {
	if s != nil {
		sduo.SetMetadata(*s)
	}
	return sduo
}

// ClearMetadata clears the value of the "metadata" field.
func (sduo *StatDescriptionUpdateOne) ClearMetadata() *StatDescriptionUpdateOne {
	sduo.mutation.ClearMetadata()
	return sduo
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (sduo *StatDescriptionUpdateOne) AddGameIDs(ids ...guidgql.GUID) *StatDescriptionUpdateOne {
	sduo.mutation.AddGameIDs(ids...)
	return sduo
}

// AddGame adds the "game" edges to the Game entity.
func (sduo *StatDescriptionUpdateOne) AddGame(g ...*Game) *StatDescriptionUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sduo.AddGameIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the Statistic entity by IDs.
func (sduo *StatDescriptionUpdateOne) AddStatIDs(ids ...guidgql.GUID) *StatDescriptionUpdateOne {
	sduo.mutation.AddStatIDs(ids...)
	return sduo
}

// AddStats adds the "stats" edges to the Statistic entity.
func (sduo *StatDescriptionUpdateOne) AddStats(s ...*Statistic) *StatDescriptionUpdateOne {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sduo.AddStatIDs(ids...)
}

// Mutation returns the StatDescriptionMutation object of the builder.
func (sduo *StatDescriptionUpdateOne) Mutation() *StatDescriptionMutation {
	return sduo.mutation
}

// ClearGame clears all "game" edges to the Game entity.
func (sduo *StatDescriptionUpdateOne) ClearGame() *StatDescriptionUpdateOne {
	sduo.mutation.ClearGame()
	return sduo
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (sduo *StatDescriptionUpdateOne) RemoveGameIDs(ids ...guidgql.GUID) *StatDescriptionUpdateOne {
	sduo.mutation.RemoveGameIDs(ids...)
	return sduo
}

// RemoveGame removes "game" edges to Game entities.
func (sduo *StatDescriptionUpdateOne) RemoveGame(g ...*Game) *StatDescriptionUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sduo.RemoveGameIDs(ids...)
}

// ClearStats clears all "stats" edges to the Statistic entity.
func (sduo *StatDescriptionUpdateOne) ClearStats() *StatDescriptionUpdateOne {
	sduo.mutation.ClearStats()
	return sduo
}

// RemoveStatIDs removes the "stats" edge to Statistic entities by IDs.
func (sduo *StatDescriptionUpdateOne) RemoveStatIDs(ids ...guidgql.GUID) *StatDescriptionUpdateOne {
	sduo.mutation.RemoveStatIDs(ids...)
	return sduo
}

// RemoveStats removes "stats" edges to Statistic entities.
func (sduo *StatDescriptionUpdateOne) RemoveStats(s ...*Statistic) *StatDescriptionUpdateOne {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sduo.RemoveStatIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *StatDescriptionUpdateOne) Select(field string, fields ...string) *StatDescriptionUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated StatDescription entity.
func (sduo *StatDescriptionUpdateOne) Save(ctx context.Context) (*StatDescription, error) {
	var (
		err  error
		node *StatDescription
	)
	if len(sduo.hooks) == 0 {
		if err = sduo.check(); err != nil {
			return nil, err
		}
		node, err = sduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sduo.check(); err != nil {
				return nil, err
			}
			sduo.mutation = mutation
			node, err = sduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sduo.hooks) - 1; i >= 0; i-- {
			if sduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StatDescription)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StatDescriptionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *StatDescriptionUpdateOne) SaveX(ctx context.Context) *StatDescription {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *StatDescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *StatDescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sduo *StatDescriptionUpdateOne) check() error {
	if v, ok := sduo.mutation.GetType(); ok {
		if err := statdescription.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "StatDescription.type": %w`, err)}
		}
	}
	if v, ok := sduo.mutation.Name(); ok {
		if err := statdescription.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "StatDescription.name": %w`, err)}
		}
	}
	return nil
}

func (sduo *StatDescriptionUpdateOne) sqlSave(ctx context.Context) (_node *StatDescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statdescription.Table,
			Columns: statdescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: statdescription.FieldID,
			},
		},
	}
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StatDescription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statdescription.FieldID)
		for _, f := range fields {
			if !statdescription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != statdescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.GetType(); ok {
		_spec.SetField(statdescription.FieldType, field.TypeEnum, value)
	}
	if value, ok := sduo.mutation.Name(); ok {
		_spec.SetField(statdescription.FieldName, field.TypeString, value)
	}
	if value, ok := sduo.mutation.Description(); ok {
		_spec.SetField(statdescription.FieldDescription, field.TypeString, value)
	}
	if sduo.mutation.DescriptionCleared() {
		_spec.ClearField(statdescription.FieldDescription, field.TypeString)
	}
	if value, ok := sduo.mutation.Metadata(); ok {
		_spec.SetField(statdescription.FieldMetadata, field.TypeString, value)
	}
	if sduo.mutation.MetadataCleared() {
		_spec.ClearField(statdescription.FieldMetadata, field.TypeString)
	}
	if sduo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.RemovedGameIDs(); len(nodes) > 0 && !sduo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statdescription.GameTable,
			Columns: statdescription.GamePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sduo.mutation.StatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.RemovedStatsIDs(); len(nodes) > 0 && !sduo.mutation.StatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statdescription.StatsTable,
			Columns: []string{statdescription.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StatDescription{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statdescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
