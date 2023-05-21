// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GroupMembershipQuery is the builder for querying GroupMembership entities.
type GroupMembershipQuery struct {
	config
	ctx        *QueryContext
	order      []groupmembership.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupMembership
	withGroup  *GroupQuery
	withUser   *UserQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GroupMembership) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupMembershipQuery builder.
func (gmq *GroupMembershipQuery) Where(ps ...predicate.GroupMembership) *GroupMembershipQuery {
	gmq.predicates = append(gmq.predicates, ps...)
	return gmq
}

// Limit the number of records to be returned by this query.
func (gmq *GroupMembershipQuery) Limit(limit int) *GroupMembershipQuery {
	gmq.ctx.Limit = &limit
	return gmq
}

// Offset to start from.
func (gmq *GroupMembershipQuery) Offset(offset int) *GroupMembershipQuery {
	gmq.ctx.Offset = &offset
	return gmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gmq *GroupMembershipQuery) Unique(unique bool) *GroupMembershipQuery {
	gmq.ctx.Unique = &unique
	return gmq
}

// Order specifies how the records should be ordered.
func (gmq *GroupMembershipQuery) Order(o ...groupmembership.OrderOption) *GroupMembershipQuery {
	gmq.order = append(gmq.order, o...)
	return gmq
}

// QueryGroup chains the current query on the "group" edge.
func (gmq *GroupMembershipQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmembership.Table, groupmembership.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupmembership.GroupTable, groupmembership.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (gmq *GroupMembershipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmembership.Table, groupmembership.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupmembership.UserTable, groupmembership.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupMembership entity from the query.
// Returns a *NotFoundError when no GroupMembership was found.
func (gmq *GroupMembershipQuery) First(ctx context.Context) (*GroupMembership, error) {
	nodes, err := gmq.Limit(1).All(setContextOp(ctx, gmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupmembership.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gmq *GroupMembershipQuery) FirstX(ctx context.Context) *GroupMembership {
	node, err := gmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupMembership ID from the query.
// Returns a *NotFoundError when no GroupMembership ID was found.
func (gmq *GroupMembershipQuery) FirstID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gmq.Limit(1).IDs(setContextOp(ctx, gmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupmembership.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gmq *GroupMembershipQuery) FirstIDX(ctx context.Context) guidgql.GUID {
	id, err := gmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupMembership entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupMembership entity is found.
// Returns a *NotFoundError when no GroupMembership entities are found.
func (gmq *GroupMembershipQuery) Only(ctx context.Context) (*GroupMembership, error) {
	nodes, err := gmq.Limit(2).All(setContextOp(ctx, gmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupmembership.Label}
	default:
		return nil, &NotSingularError{groupmembership.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gmq *GroupMembershipQuery) OnlyX(ctx context.Context) *GroupMembership {
	node, err := gmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupMembership ID in the query.
// Returns a *NotSingularError when more than one GroupMembership ID is found.
// Returns a *NotFoundError when no entities are found.
func (gmq *GroupMembershipQuery) OnlyID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = gmq.Limit(2).IDs(setContextOp(ctx, gmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupmembership.Label}
	default:
		err = &NotSingularError{groupmembership.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gmq *GroupMembershipQuery) OnlyIDX(ctx context.Context) guidgql.GUID {
	id, err := gmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupMemberships.
func (gmq *GroupMembershipQuery) All(ctx context.Context) ([]*GroupMembership, error) {
	ctx = setContextOp(ctx, gmq.ctx, "All")
	if err := gmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupMembership, *GroupMembershipQuery]()
	return withInterceptors[[]*GroupMembership](ctx, gmq, qr, gmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gmq *GroupMembershipQuery) AllX(ctx context.Context) []*GroupMembership {
	nodes, err := gmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupMembership IDs.
func (gmq *GroupMembershipQuery) IDs(ctx context.Context) (ids []guidgql.GUID, err error) {
	if gmq.ctx.Unique == nil && gmq.path != nil {
		gmq.Unique(true)
	}
	ctx = setContextOp(ctx, gmq.ctx, "IDs")
	if err = gmq.Select(groupmembership.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gmq *GroupMembershipQuery) IDsX(ctx context.Context) []guidgql.GUID {
	ids, err := gmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gmq *GroupMembershipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gmq.ctx, "Count")
	if err := gmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gmq, querierCount[*GroupMembershipQuery](), gmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gmq *GroupMembershipQuery) CountX(ctx context.Context) int {
	count, err := gmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gmq *GroupMembershipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gmq.ctx, "Exist")
	switch _, err := gmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gmq *GroupMembershipQuery) ExistX(ctx context.Context) bool {
	exist, err := gmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupMembershipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gmq *GroupMembershipQuery) Clone() *GroupMembershipQuery {
	if gmq == nil {
		return nil
	}
	return &GroupMembershipQuery{
		config:     gmq.config,
		ctx:        gmq.ctx.Clone(),
		order:      append([]groupmembership.OrderOption{}, gmq.order...),
		inters:     append([]Interceptor{}, gmq.inters...),
		predicates: append([]predicate.GroupMembership{}, gmq.predicates...),
		withGroup:  gmq.withGroup.Clone(),
		withUser:   gmq.withUser.Clone(),
		// clone intermediate query.
		sql:  gmq.sql.Clone(),
		path: gmq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gmq *GroupMembershipQuery) WithGroup(opts ...func(*GroupQuery)) *GroupMembershipQuery {
	query := (&GroupClient{config: gmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gmq.withGroup = query
	return gmq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gmq *GroupMembershipQuery) WithUser(opts ...func(*UserQuery)) *GroupMembershipQuery {
	query := (&UserClient{config: gmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gmq.withUser = query
	return gmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role enums.Role `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupMembership.Query().
//		GroupBy(groupmembership.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gmq *GroupMembershipQuery) GroupBy(field string, fields ...string) *GroupMembershipGroupBy {
	gmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupMembershipGroupBy{build: gmq}
	grbuild.flds = &gmq.ctx.Fields
	grbuild.label = groupmembership.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role enums.Role `json:"role,omitempty"`
//	}
//
//	client.GroupMembership.Query().
//		Select(groupmembership.FieldRole).
//		Scan(ctx, &v)
func (gmq *GroupMembershipQuery) Select(fields ...string) *GroupMembershipSelect {
	gmq.ctx.Fields = append(gmq.ctx.Fields, fields...)
	sbuild := &GroupMembershipSelect{GroupMembershipQuery: gmq}
	sbuild.label = groupmembership.Label
	sbuild.flds, sbuild.scan = &gmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupMembershipSelect configured with the given aggregations.
func (gmq *GroupMembershipQuery) Aggregate(fns ...AggregateFunc) *GroupMembershipSelect {
	return gmq.Select().Aggregate(fns...)
}

func (gmq *GroupMembershipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gmq); err != nil {
				return err
			}
		}
	}
	for _, f := range gmq.ctx.Fields {
		if !groupmembership.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gmq.path != nil {
		prev, err := gmq.path(ctx)
		if err != nil {
			return err
		}
		gmq.sql = prev
	}
	return nil
}

func (gmq *GroupMembershipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupMembership, error) {
	var (
		nodes       = []*GroupMembership{}
		withFKs     = gmq.withFKs
		_spec       = gmq.querySpec()
		loadedTypes = [2]bool{
			gmq.withGroup != nil,
			gmq.withUser != nil,
		}
	)
	if gmq.withGroup != nil || gmq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembership.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupMembership).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupMembership{config: gmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gmq.modifiers) > 0 {
		_spec.Modifiers = gmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gmq.withGroup; query != nil {
		if err := gmq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupMembership, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := gmq.withUser; query != nil {
		if err := gmq.loadUser(ctx, query, nodes, nil,
			func(n *GroupMembership, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range gmq.loadTotal {
		if err := gmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gmq *GroupMembershipQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupMembership, init func(*GroupMembership), assign func(*GroupMembership, *Group)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*GroupMembership)
	for i := range nodes {
		if nodes[i].group_members == nil {
			continue
		}
		fk := *nodes[i].group_members
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
			return fmt.Errorf(`unexpected foreign-key "group_members" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gmq *GroupMembershipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GroupMembership, init func(*GroupMembership), assign func(*GroupMembership, *User)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*GroupMembership)
	for i := range nodes {
		if nodes[i].user_group_memberships == nil {
			continue
		}
		fk := *nodes[i].user_group_memberships
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
			return fmt.Errorf(`unexpected foreign-key "user_group_memberships" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gmq *GroupMembershipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gmq.querySpec()
	if len(gmq.modifiers) > 0 {
		_spec.Modifiers = gmq.modifiers
	}
	_spec.Node.Columns = gmq.ctx.Fields
	if len(gmq.ctx.Fields) > 0 {
		_spec.Unique = gmq.ctx.Unique != nil && *gmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gmq.driver, _spec)
}

func (gmq *GroupMembershipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupmembership.Table, groupmembership.Columns, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	_spec.From = gmq.sql
	if unique := gmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gmq.path != nil {
		_spec.Unique = true
	}
	if fields := gmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembership.FieldID)
		for i := range fields {
			if fields[i] != groupmembership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gmq *GroupMembershipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gmq.driver.Dialect())
	t1 := builder.Table(groupmembership.Table)
	columns := gmq.ctx.Fields
	if len(columns) == 0 {
		columns = groupmembership.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gmq.sql != nil {
		selector = gmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gmq.ctx.Unique != nil && *gmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gmq.predicates {
		p(selector)
	}
	for _, p := range gmq.order {
		p(selector)
	}
	if offset := gmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupMembershipGroupBy is the group-by builder for GroupMembership entities.
type GroupMembershipGroupBy struct {
	selector
	build *GroupMembershipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gmgb *GroupMembershipGroupBy) Aggregate(fns ...AggregateFunc) *GroupMembershipGroupBy {
	gmgb.fns = append(gmgb.fns, fns...)
	return gmgb
}

// Scan applies the selector query and scans the result into the given value.
func (gmgb *GroupMembershipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gmgb.build.ctx, "GroupBy")
	if err := gmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMembershipQuery, *GroupMembershipGroupBy](ctx, gmgb.build, gmgb, gmgb.build.inters, v)
}

func (gmgb *GroupMembershipGroupBy) sqlScan(ctx context.Context, root *GroupMembershipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gmgb.fns))
	for _, fn := range gmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gmgb.flds)+len(gmgb.fns))
		for _, f := range *gmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupMembershipSelect is the builder for selecting fields of GroupMembership entities.
type GroupMembershipSelect struct {
	*GroupMembershipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gms *GroupMembershipSelect) Aggregate(fns ...AggregateFunc) *GroupMembershipSelect {
	gms.fns = append(gms.fns, fns...)
	return gms
}

// Scan applies the selector query and scans the result into the given value.
func (gms *GroupMembershipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gms.ctx, "Select")
	if err := gms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupMembershipQuery, *GroupMembershipSelect](ctx, gms.GroupMembershipQuery, gms, gms.inters, v)
}

func (gms *GroupMembershipSelect) sqlScan(ctx context.Context, root *GroupMembershipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gms.fns))
	for _, fn := range gms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
