// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ga *GameQuery) CollectFields(ctx context.Context, satisfies ...string) (*GameQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ga, nil
	}
	if err := ga.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ga, nil
}

func (ga *GameQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "author":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: ga.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ga.withAuthor = query
		}
	}
	return nil
}

type gamePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GamePaginateOption
}

func newGamePaginateArgs(rv map[string]interface{}) *gamePaginateArgs {
	args := &gamePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GameWhereInput); ok {
		args.opts = append(args.opts, WithGameFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "settings":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupSettingsQuery{config: gr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gr.withSettings = query
		case "members":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupMembershipQuery{config: gr.config}
			)
			args := newGroupMembershipPaginateArgs(fieldArgs(ctx, new(GroupMembershipWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newGroupMembershipPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					gr.loadTotal = append(gr.loadTotal, func(ctx context.Context, nodes []*Group) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID guidgql.GUID `sql:"group_members"`
							Count  int          `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(group.MembersColumn, ids...))
						})
						if err := query.GroupBy(group.MembersColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[guidgql.GUID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					gr.loadTotal = append(gr.loadTotal, func(_ context.Context, nodes []*Group) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Members)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(group.MembersColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			gr.WithNamedMembers(alias, func(wq *GroupMembershipQuery) {
				*wq = *query
			})
		case "applications":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupMembershipApplicationQuery{config: gr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gr.WithNamedApplications(alias, func(wq *GroupMembershipApplicationQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]interface{}) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gm *GroupMembershipQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupMembershipQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gm, nil
	}
	if err := gm.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gm, nil
}

func (gm *GroupMembershipQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "group":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupQuery{config: gm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gm.withGroup = query
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: gm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gm.withUser = query
		}
	}
	return nil
}

type groupmembershipPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupMembershipPaginateOption
}

func newGroupMembershipPaginateArgs(rv map[string]interface{}) *groupmembershipPaginateArgs {
	args := &groupmembershipPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupMembershipWhereInput); ok {
		args.opts = append(args.opts, WithGroupMembershipFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gma *GroupMembershipApplicationQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupMembershipApplicationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gma, nil
	}
	if err := gma.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gma, nil
}

func (gma *GroupMembershipApplicationQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: gma.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gma.withUser = query
		case "group":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupQuery{config: gma.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			gma.withGroup = query
		}
	}
	return nil
}

type groupmembershipapplicationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupMembershipApplicationPaginateOption
}

func newGroupMembershipApplicationPaginateArgs(rv map[string]interface{}) *groupmembershipapplicationPaginateArgs {
	args := &groupmembershipapplicationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gs *GroupSettingsQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupSettingsQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gs, nil
	}
	if err := gs.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gs, nil
}

func (gs *GroupSettingsQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type groupsettingsPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupSettingsPaginateOption
}

func newGroupSettingsPaginateArgs(rv map[string]interface{}) *groupsettingsPaginateArgs {
	args := &groupsettingsPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupSettingsWhereInput); ok {
		args.opts = append(args.opts, WithGroupSettingsFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MatchQuery) CollectFields(ctx context.Context, satisfies ...string) (*MatchQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MatchQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "game":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GameQuery{config: m.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			m.withGame = query
		case "players":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: m.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			m.WithNamedPlayers(alias, func(wq *PlayerQuery) {
				*wq = *query
			})
		case "stats":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StatisticQuery{config: m.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			m.WithNamedStats(alias, func(wq *StatisticQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type matchPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MatchPaginateOption
}

func newMatchPaginateArgs(rv map[string]interface{}) *matchPaginateArgs {
	args := &matchPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*MatchWhereInput); ok {
		args.opts = append(args.opts, WithMatchFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pl *PlayerQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pl, nil
	}
	if err := pl.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pl, nil
}

func (pl *PlayerQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.withOwner = query
		case "supervisors":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedSupervisors(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "supervisionRequests":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedSupervisionRequests(alias, func(wq *PlayerSupervisionRequestQuery) {
				*wq = *query
			})
		case "matches":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &MatchQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedMatches(alias, func(wq *MatchQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type playerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerPaginateOption
}

func newPlayerPaginateArgs(rv map[string]interface{}) *playerPaginateArgs {
	args := &playerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerWhereInput); ok {
		args.opts = append(args.opts, WithPlayerFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (psr *PlayerSupervisionRequestQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerSupervisionRequestQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return psr, nil
	}
	if err := psr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return psr, nil
}

func (psr *PlayerSupervisionRequestQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "sender":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.withSender = query
		case "player":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.withPlayer = query
		case "approvals":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestApprovalQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.WithNamedApprovals(alias, func(wq *PlayerSupervisionRequestApprovalQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type playersupervisionrequestPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerSupervisionRequestPaginateOption
}

func newPlayerSupervisionRequestPaginateArgs(rv map[string]interface{}) *playersupervisionrequestPaginateArgs {
	args := &playersupervisionrequestPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerSupervisionRequestWhereInput); ok {
		args.opts = append(args.opts, WithPlayerSupervisionRequestFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (psra *PlayerSupervisionRequestApprovalQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerSupervisionRequestApprovalQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return psra, nil
	}
	if err := psra.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return psra, nil
}

func (psra *PlayerSupervisionRequestApprovalQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "approver":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: psra.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psra.withApprover = query
		case "supervisionRequest":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestQuery{config: psra.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psra.withSupervisionRequest = query
		}
	}
	return nil
}

type playersupervisionrequestapprovalPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerSupervisionRequestApprovalPaginateOption
}

func newPlayerSupervisionRequestApprovalPaginateArgs(rv map[string]interface{}) *playersupervisionrequestapprovalPaginateArgs {
	args := &playersupervisionrequestapprovalPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerSupervisionRequestApprovalWhereInput); ok {
		args.opts = append(args.opts, WithPlayerSupervisionRequestApprovalFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sd *StatDescriptionQuery) CollectFields(ctx context.Context, satisfies ...string) (*StatDescriptionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sd, nil
	}
	if err := sd.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *StatDescriptionQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type statdescriptionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StatDescriptionPaginateOption
}

func newStatDescriptionPaginateArgs(rv map[string]interface{}) *statdescriptionPaginateArgs {
	args := &statdescriptionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *StatisticQuery) CollectFields(ctx context.Context, satisfies ...string) (*StatisticQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *StatisticQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "match":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &MatchQuery{config: s.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			s.withMatch = query
		case "statDescription":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StatDescriptionQuery{config: s.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			s.withStatDescription = query
		case "player":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: s.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			s.withPlayer = query
		}
	}
	return nil
}

type statisticPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StatisticPaginateOption
}

func newStatisticPaginateArgs(rv map[string]interface{}) *statisticPaginateArgs {
	args := &statisticPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "players":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedPlayers(alias, func(wq *PlayerQuery) {
				*wq = *query
			})
		case "mainPlayer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.withMainPlayer = query
		case "groupMemberships":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupMembershipQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedGroupMemberships(alias, func(wq *GroupMembershipQuery) {
				*wq = *query
			})
		case "groupMembershipApplications":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupMembershipApplicationQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedGroupMembershipApplications(alias, func(wq *GroupMembershipApplicationQuery) {
				*wq = *query
			})
		case "games":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GameQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedGames(alias, func(wq *GameQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
