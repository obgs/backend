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
	"github.com/obgs/backend/internal/ent/groupmembershipapplication"
	"github.com/obgs/backend/internal/ent/predicate"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/user"
)

// GroupMembershipApplicationQuery is the builder for querying GroupMembershipApplication entities.
type GroupMembershipApplicationQuery struct {
	config
	ctx        *QueryContext
	order      []groupmembershipapplication.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupMembershipApplication
	withUser   *UserQuery
	withGroup  *GroupQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GroupMembershipApplication) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupMembershipApplicationQuery builder.
func (gmaq *GroupMembershipApplicationQuery) Where(ps ...predicate.GroupMembershipApplication) *GroupMembershipApplicationQuery {
	gmaq.predicates = append(gmaq.predicates, ps...)
	return gmaq
}

// Limit the number of records to be returned by this query.
func (gmaq *GroupMembershipApplicationQuery) Limit(limit int) *GroupMembershipApplicationQuery {
	gmaq.ctx.Limit = &limit
	return gmaq
}

// Offset to start from.
func (gmaq *GroupMembershipApplicationQuery) Offset(offset int) *GroupMembershipApplicationQuery {
	gmaq.ctx.Offset = &offset
	return gmaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gmaq *GroupMembershipApplicationQuery) Unique(unique bool) *GroupMembershipApplicationQuery {
	gmaq.ctx.Unique = &unique
	return gmaq
}

// Order specifies how the records should be ordered.
func (gmaq *GroupMembershipApplicationQuery) Order(o ...groupmembershipapplication.OrderOption) *GroupMembershipApplicationQuery {
	gmaq.order = append(gmaq.order, o...)
	return gmaq
}

// QueryUser chains the current query on the "user" edge.
func (gmaq *GroupMembershipApplicationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gmaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gmaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gmaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmembershipapplication.Table, groupmembershipapplication.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupmembershipapplication.UserTable, groupmembershipapplication.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gmaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (gmaq *GroupMembershipApplicationQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gmaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gmaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gmaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmembershipapplication.Table, groupmembershipapplication.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupmembershipapplication.GroupTable, groupmembershipapplication.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gmaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupMembershipApplication entity from the query.
// Returns a *NotFoundError when no GroupMembershipApplication was found.
func (gmaq *GroupMembershipApplicationQuery) First(ctx context.Context) (*GroupMembershipApplication, error) {
	nodes, err := gmaq.Limit(1).All(setContextOp(ctx, gmaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupmembershipapplication.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) FirstX(ctx context.Context) *GroupMembershipApplication {
	node, err := gmaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupMembershipApplication ID from the query.
// Returns a *NotFoundError when no GroupMembershipApplication ID was found.
func (gmaq *GroupMembershipApplicationQuery) FirstID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gmaq.Limit(1).IDs(setContextOp(ctx, gmaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupmembershipapplication.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) FirstIDX(ctx context.Context) guidgql.GUID {
	id, err := gmaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupMembershipApplication entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupMembershipApplication entity is found.
// Returns a *NotFoundError when no GroupMembershipApplication entities are found.
func (gmaq *GroupMembershipApplicationQuery) Only(ctx context.Context) (*GroupMembershipApplication, error) {
	nodes, err := gmaq.Limit(2).All(setContextOp(ctx, gmaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupmembershipapplication.Label}
	default:
		return nil, &NotSingularError{groupmembershipapplication.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) OnlyX(ctx context.Context) *GroupMembershipApplication {
	node, err := gmaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupMembershipApplication ID in the query.
// Returns a *NotSingularError when more than one GroupMembershipApplication ID is found.
// Returns a *NotFoundError when no entities are found.
func (gmaq *GroupMembershipApplicationQuery) OnlyID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gmaq.Limit(2).IDs(setContextOp(ctx, gmaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupmembershipapplication.Label}
	default:
		err = &NotSingularError{groupmembershipapplication.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) OnlyIDX(ctx context.Context) guidgql.GUID {
	id, err := gmaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupMembershipApplications.
func (gmaq *GroupMembershipApplicationQuery) All(ctx context.Context) ([]*GroupMembershipApplication, error) {
	ctx = setContextOp(ctx, gmaq.ctx, "All")
	if err := gmaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupMembershipApplication, *GroupMembershipApplicationQuery]()
	return withInterceptors[[]*GroupMembershipApplication](ctx, gmaq, qr, gmaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) AllX(ctx context.Context) []*GroupMembershipApplication {
	nodes, err := gmaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupMembershipApplication IDs.
func (gmaq *GroupMembershipApplicationQuery) IDs(ctx context.Context) (ids []guidgql.GUID, err error) {
	if gmaq.ctx.Unique == nil && gmaq.path != nil {
		gmaq.Unique(true)
	}
	ctx = setContextOp(ctx, gmaq.ctx, "IDs")
	if err = gmaq.Select(groupmembershipapplication.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) IDsX(ctx context.Context) []guidgql.GUID {
	ids, err := gmaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gmaq *GroupMembershipApplicationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gmaq.ctx, "Count")
	if err := gmaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gmaq, querierCount[*GroupMembershipApplicationQuery](), gmaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) CountX(ctx context.Context) int {
	count, err := gmaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gmaq *GroupMembershipApplicationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gmaq.ctx, "Exist")
	switch _, err := gmaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gmaq *GroupMembershipApplicationQuery) ExistX(ctx context.Context) bool {
	exist, err := gmaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupMembershipApplicationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gmaq *GroupMembershipApplicationQuery) Clone() *GroupMembershipApplicationQuery {
	if gmaq == nil {
		return nil
	}
	return &GroupMembershipApplicationQuery{
		config:     gmaq.config,
		ctx:        gmaq.ctx.Clone(),
		order:      append([]groupmembershipapplication.OrderOption{}, gmaq.order...),
		inters:     append([]Interceptor{}, gmaq.inters...),
		predicates: append([]predicate.GroupMembershipApplication{}, gmaq.predicates...),
		withUser:   gmaq.withUser.Clone(),
		withGroup:  gmaq.withGroup.Clone(),
		// clone intermediate query.
		sql:  gmaq.sql.Clone(),
		path: gmaq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gmaq *GroupMembershipApplicationQuery) WithUser(opts ...func(*UserQuery)) *GroupMembershipApplicationQuery {
	query := (&UserClient{config: gmaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gmaq.withUser = query
	return gmaq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gmaq *GroupMembershipApplicationQuery) WithGroup(opts ...func(*GroupQuery)) *GroupMembershipApplicationQuery {
	query := (&GroupClient{config: gmaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gmaq.withGroup = query
	return gmaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupMembershipApplication.Query().
//		GroupBy(groupmembershipapplication.FieldMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gmaq *GroupMembershipApplicationQuery) GroupBy(field string, fields ...string) *GroupMembershipApplicationGroupBy {
	gmaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupMembershipApplicationGroupBy{build: gmaq}
	grbuild.flds = &gmaq.ctx.Fields
	grbuild.label = groupmembershipapplication.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//	}
//
//	client.GroupMembershipApplication.Query().
//		Select(groupmembershipapplication.FieldMessage).
//		Scan(ctx, &v)
func (gmaq *GroupMembershipApplicationQuery) Select(fields ...string) *GroupMembershipApplicationSelect {
	gmaq.ctx.Fields = append(gmaq.ctx.Fields, fields...)
	sbuild := &GroupMembershipApplicationSelect{GroupMembershipApplicationQuery: gmaq}
	sbuild.label = groupmembershipapplication.Label
	sbuild.flds, sbuild.scan = &gmaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupMembershipApplicationSelect configured with the given aggregations.
func (gmaq *GroupMembershipApplicationQuery) Aggregate(fns ...AggregateFunc) *GroupMembershipApplicationSelect {
	return gmaq.Select().Aggregate(fns...)
}

func (gmaq *GroupMembershipApplicationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gmaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gmaq); err != nil {
				return err
			}
		}
	}
	for _, f := range gmaq.ctx.Fields {
		if !groupmembershipapplication.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gmaq.path != nil {
		prev, err := gmaq.path(ctx)
		if err != nil {
			return err
		}
		gmaq.sql = prev
	}
	return nil
}

func (gmaq *GroupMembershipApplicationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupMembershipApplication, error) {
	var (
		nodes       = []*GroupMembershipApplication{}
		withFKs     = gmaq.withFKs
		_spec       = gmaq.querySpec()
		loadedTypes = [2]bool{
			gmaq.withUser != nil,
			gmaq.withGroup != nil,
		}
	)
	if gmaq.withUser != nil || gmaq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembershipapplication.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupMembershipApplication).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupMembershipApplication{config: gmaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gmaq.modifiers) > 0 {
		_spec.Modifiers = gmaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gmaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gmaq.withUser; query != nil {
		if err := gmaq.loadUser(ctx, query, nodes, nil,
			func(n *GroupMembershipApplication, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := gmaq.withGroup; query != nil {
		if err := gmaq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupMembershipApplication, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	for i := range gmaq.loadTotal {
		if err := gmaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gmaq *GroupMembershipApplicationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GroupMembershipApplication, init func(*GroupMembershipApplication), assign func(*GroupMembershipApplication, *User)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*GroupMembershipApplication)
	for i := range nodes {
		if nodes[i].user_group_membership_applications == nil {
			continue
		}
		fk := *nodes[i].user_group_membership_applications
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_group_membership_applications" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gmaq *GroupMembershipApplicationQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupMembershipApplication, init func(*GroupMembershipApplication), assign func(*GroupMembershipApplication, *Group)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*GroupMembershipApplication)
	for i := range nodes {
		if nodes[i].group_applications == nil {
			continue
		}
		fk := *nodes[i].group_applications
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
			return fmt.Errorf(`unexpected foreign-key "group_applications" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gmaq *GroupMembershipApplicationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gmaq.querySpec()
	if len(gmaq.modifiers) > 0 {
		_spec.Modifiers = gmaq.modifiers
	}
	_spec.Node.Columns = gmaq.ctx.Fields
	if len(gmaq.ctx.Fields) > 0 {
		_spec.Unique = gmaq.ctx.Unique != nil && *gmaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gmaq.driver, _spec)
}

func (gmaq *GroupMembershipApplicationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupmembershipapplication.Table, groupmembershipapplication.Columns, sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString))
	_spec.From = gmaq.sql
	if unique := gmaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gmaq.path != nil {
		_spec.Unique = true
	}
	if fields := gmaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembershipapplication.FieldID)
		for i := range fields {
			if fields[i] != groupmembershipapplication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gmaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gmaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gmaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gmaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gmaq *GroupMembershipApplicationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gmaq.driver.Dialect())
	t1 := builder.Table(groupmembershipapplication.Table)
	columns := gmaq.ctx.Fields
	if len(columns) == 0 {
		columns = groupmembershipapplication.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gmaq.sql != nil {
		selector = gmaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gmaq.ctx.Unique != nil && *gmaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gmaq.predicates {
		p(selector)
	}
	for _, p := range gmaq.order {
		p(selector)
	}
	if offset := gmaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gmaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupMembershipApplicationGroupBy is the group-by builder for GroupMembershipApplication entities.
type GroupMembershipApplicationGroupBy struct {
	selector
	build *GroupMembershipApplicationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gmagb *GroupMembershipApplicationGroupBy) Aggregate(fns ...AggregateFunc) *GroupMembershipApplicationGroupBy {
	gmagb.fns = append(gmagb.fns, fns...)
	return gmagb
}

// Scan applies the selector query and scans the result into the given value.
func (gmagb *GroupMembershipApplicationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gmagb.build.ctx, "GroupBy")
	if err := gmagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMembershipApplicationQuery, *GroupMembershipApplicationGroupBy](ctx, gmagb.build, gmagb, gmagb.build.inters, v)
}

func (gmagb *GroupMembershipApplicationGroupBy) sqlScan(ctx context.Context, root *GroupMembershipApplicationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gmagb.fns))
	for _, fn := range gmagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gmagb.flds)+len(gmagb.fns))
		for _, f := range *gmagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gmagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupMembershipApplicationSelect is the builder for selecting fields of GroupMembershipApplication entities.
type GroupMembershipApplicationSelect struct {
	*GroupMembershipApplicationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gmas *GroupMembershipApplicationSelect) Aggregate(fns ...AggregateFunc) *GroupMembershipApplicationSelect {
	gmas.fns = append(gmas.fns, fns...)
	return gmas
}

// Scan applies the selector query and scans the result into the given value.
func (gmas *GroupMembershipApplicationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gmas.ctx, "Select")
	if err := gmas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMembershipApplicationQuery, *GroupMembershipApplicationSelect](ctx, gmas.GroupMembershipApplicationQuery, gmas, gmas.inters, v)
}

func (gmas *GroupMembershipApplicationSelect) sqlScan(ctx context.Context, root *GroupMembershipApplicationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gmas.fns))
	for _, fn := range gmas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gmas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
