// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/group"
	"github.com/obgs/backend/internal/ent/groupsettings"
	"github.com/obgs/backend/internal/ent/predicate"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

// GroupSettingsQuery is the builder for querying GroupSettings entities.
type GroupSettingsQuery struct {
	config
	ctx        *QueryContext
	order      []groupsettings.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupSettings
	withGroup  *GroupQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GroupSettings) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupSettingsQuery builder.
func (gsq *GroupSettingsQuery) Where(ps ...predicate.GroupSettings) *GroupSettingsQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit the number of records to be returned by this query.
func (gsq *GroupSettingsQuery) Limit(limit int) *GroupSettingsQuery {
	gsq.ctx.Limit = &limit
	return gsq
}

// Offset to start from.
func (gsq *GroupSettingsQuery) Offset(offset int) *GroupSettingsQuery {
	gsq.ctx.Offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GroupSettingsQuery) Unique(unique bool) *GroupSettingsQuery {
	gsq.ctx.Unique = &unique
	return gsq
}

// Order specifies how the records should be ordered.
func (gsq *GroupSettingsQuery) Order(o ...groupsettings.OrderOption) *GroupSettingsQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// QueryGroup chains the current query on the "group" edge.
func (gsq *GroupSettingsQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupsettings.Table, groupsettings.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, groupsettings.GroupTable, groupsettings.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupSettings entity from the query.
// Returns a *NotFoundError when no GroupSettings was found.
func (gsq *GroupSettingsQuery) First(ctx context.Context) (*GroupSettings, error) {
	nodes, err := gsq.Limit(1).All(setContextOp(ctx, gsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupsettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GroupSettingsQuery) FirstX(ctx context.Context) *GroupSettings {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupSettings ID from the query.
// Returns a *NotFoundError when no GroupSettings ID was found.
func (gsq *GroupSettingsQuery) FirstID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gsq.Limit(1).IDs(setContextOp(ctx, gsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupsettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GroupSettingsQuery) FirstIDX(ctx context.Context) guidgql.GUID {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupSettings entity is found.
// Returns a *NotFoundError when no GroupSettings entities are found.
func (gsq *GroupSettingsQuery) Only(ctx context.Context) (*GroupSettings, error) {
	nodes, err := gsq.Limit(2).All(setContextOp(ctx, gsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupsettings.Label}
	default:
		return nil, &NotSingularError{groupsettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GroupSettingsQuery) OnlyX(ctx context.Context) *GroupSettings {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupSettings ID in the query.
// Returns a *NotSingularError when more than one GroupSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GroupSettingsQuery) OnlyID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gsq.Limit(2).IDs(setContextOp(ctx, gsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupsettings.Label}
	default:
		err = &NotSingularError{groupsettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GroupSettingsQuery) OnlyIDX(ctx context.Context) guidgql.GUID {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupSettingsSlice.
func (gsq *GroupSettingsQuery) All(ctx context.Context) ([]*GroupSettings, error) {
	ctx = setContextOp(ctx, gsq.ctx, "All")
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupSettings, *GroupSettingsQuery]()
	return withInterceptors[[]*GroupSettings](ctx, gsq, qr, gsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GroupSettingsQuery) AllX(ctx context.Context) []*GroupSettings {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupSettings IDs.
func (gsq *GroupSettingsQuery) IDs(ctx context.Context) (ids []guidgql.GUID, err error) {
	if gsq.ctx.Unique == nil && gsq.path != nil {
		gsq.Unique(true)
	}
	ctx = setContextOp(ctx, gsq.ctx, "IDs")
	if err = gsq.Select(groupsettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GroupSettingsQuery) IDsX(ctx context.Context) []guidgql.GUID {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GroupSettingsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gsq.ctx, "Count")
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gsq, querierCount[*GroupSettingsQuery](), gsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GroupSettingsQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GroupSettingsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gsq.ctx, "Exist")
	switch _, err := gsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GroupSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GroupSettingsQuery) Clone() *GroupSettingsQuery {
	if gsq == nil {
		return nil
	}
	return &GroupSettingsQuery{
		config:     gsq.config,
		ctx:        gsq.ctx.Clone(),
		order:      append([]groupsettings.OrderOption{}, gsq.order...),
		inters:     append([]Interceptor{}, gsq.inters...),
		predicates: append([]predicate.GroupSettings{}, gsq.predicates...),
		withGroup:  gsq.withGroup.Clone(),
		// clone intermediate query.
		sql:  gsq.sql.Clone(),
		path: gsq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GroupSettingsQuery) WithGroup(opts ...func(*GroupQuery)) *GroupSettingsQuery {
	query := (&GroupClient{config: gsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGroup = query
	return gsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Visibility groupsettings.Visibility `json:"visibility,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupSettings.Query().
//		GroupBy(groupsettings.FieldVisibility).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gsq *GroupSettingsQuery) GroupBy(field string, fields ...string) *GroupSettingsGroupBy {
	gsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupSettingsGroupBy{build: gsq}
	grbuild.flds = &gsq.ctx.Fields
	grbuild.label = groupsettings.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Visibility groupsettings.Visibility `json:"visibility,omitempty"`
//	}
//
//	client.GroupSettings.Query().
//		Select(groupsettings.FieldVisibility).
//		Scan(ctx, &v)
func (gsq *GroupSettingsQuery) Select(fields ...string) *GroupSettingsSelect {
	gsq.ctx.Fields = append(gsq.ctx.Fields, fields...)
	sbuild := &GroupSettingsSelect{GroupSettingsQuery: gsq}
	sbuild.label = groupsettings.Label
	sbuild.flds, sbuild.scan = &gsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupSettingsSelect configured with the given aggregations.
func (gsq *GroupSettingsQuery) Aggregate(fns ...AggregateFunc) *GroupSettingsSelect {
	return gsq.Select().Aggregate(fns...)
}

func (gsq *GroupSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gsq); err != nil {
				return err
			}
		}
	}
	for _, f := range gsq.ctx.Fields {
		if !groupsettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	return nil
}

func (gsq *GroupSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupSettings, error) {
	var (
		nodes       = []*GroupSettings{}
		withFKs     = gsq.withFKs
		_spec       = gsq.querySpec()
		loadedTypes = [1]bool{
			gsq.withGroup != nil,
		}
	)
	if gsq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupsettings.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupSettings{config: gsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gsq.withGroup; query != nil {
		if err := gsq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupSettings, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	for i := range gsq.loadTotal {
		if err := gsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gsq *GroupSettingsQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupSettings, init func(*GroupSettings), assign func(*GroupSettings, *Group)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*GroupSettings)
	for i := range nodes {
		if nodes[i].group_settings == nil {
			continue
		}
		fk := *nodes[i].group_settings
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_settings" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gsq *GroupSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	_spec.Node.Columns = gsq.ctx.Fields
	if len(gsq.ctx.Fields) > 0 {
		_spec.Unique = gsq.ctx.Unique != nil && *gsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GroupSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupsettings.Table, groupsettings.Columns, sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString))
	_spec.From = gsq.sql
	if unique := gsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gsq.path != nil {
		_spec.Unique = true
	}
	if fields := gsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupsettings.FieldID)
		for i := range fields {
			if fields[i] != groupsettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GroupSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(groupsettings.Table)
	columns := gsq.ctx.Fields
	if len(columns) == 0 {
		columns = groupsettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gsq.ctx.Unique != nil && *gsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupSettingsGroupBy is the group-by builder for GroupSettings entities.
type GroupSettingsGroupBy struct {
	selector
	build *GroupSettingsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GroupSettingsGroupBy) Aggregate(fns ...AggregateFunc) *GroupSettingsGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the selector query and scans the result into the given value.
func (gsgb *GroupSettingsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gsgb.build.ctx, "GroupBy")
	if err := gsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupSettingsQuery, *GroupSettingsGroupBy](ctx, gsgb.build, gsgb, gsgb.build.inters, v)
}

func (gsgb *GroupSettingsGroupBy) sqlScan(ctx context.Context, root *GroupSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gsgb.fns))
	for _, fn := range gsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gsgb.flds)+len(gsgb.fns))
		for _, f := range *gsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupSettingsSelect is the builder for selecting fields of GroupSettings entities.
type GroupSettingsSelect struct {
	*GroupSettingsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gss *GroupSettingsSelect) Aggregate(fns ...AggregateFunc) *GroupSettingsSelect {
	gss.fns = append(gss.fns, fns...)
	return gss
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GroupSettingsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gss.ctx, "Select")
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupSettingsQuery, *GroupSettingsSelect](ctx, gss.GroupSettingsQuery, gss, gss.inters, v)
}

func (gss *GroupSettingsSelect) sqlScan(ctx context.Context, root *GroupSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gss.fns))
	for _, fn := range gss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
