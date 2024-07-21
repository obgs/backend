// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/obgs/backend/internal/ent/player"
	"github.com/obgs/backend/internal/ent/playersupervisionrequest"
	"github.com/obgs/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/obgs/backend/internal/ent/predicate"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/user"
)

// PlayerSupervisionRequestUpdate is the builder for updating PlayerSupervisionRequest entities.
type PlayerSupervisionRequestUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerSupervisionRequestMutation
}

// Where appends a list predicates to the PlayerSupervisionRequestUpdate builder.
func (psru *PlayerSupervisionRequestUpdate) Where(ps ...predicate.PlayerSupervisionRequest) *PlayerSupervisionRequestUpdate {
	psru.mutation.Where(ps...)
	return psru
}

// SetMessage sets the "message" field.
func (psru *PlayerSupervisionRequestUpdate) SetMessage(s string) *PlayerSupervisionRequestUpdate {
	psru.mutation.SetMessage(s)
	return psru
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (psru *PlayerSupervisionRequestUpdate) SetNillableMessage(s *string) *PlayerSupervisionRequestUpdate {
	if s != nil {
		psru.SetMessage(*s)
	}
	return psru
}

// ClearMessage clears the value of the "message" field.
func (psru *PlayerSupervisionRequestUpdate) ClearMessage() *PlayerSupervisionRequestUpdate {
	psru.mutation.ClearMessage()
	return psru
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (psru *PlayerSupervisionRequestUpdate) SetSenderID(id guidgql.GUID) *PlayerSupervisionRequestUpdate {
	psru.mutation.SetSenderID(id)
	return psru
}

// SetSender sets the "sender" edge to the User entity.
func (psru *PlayerSupervisionRequestUpdate) SetSender(u *User) *PlayerSupervisionRequestUpdate {
	return psru.SetSenderID(u.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (psru *PlayerSupervisionRequestUpdate) SetPlayerID(id guidgql.GUID) *PlayerSupervisionRequestUpdate {
	psru.mutation.SetPlayerID(id)
	return psru
}

// SetPlayer sets the "player" edge to the Player entity.
func (psru *PlayerSupervisionRequestUpdate) SetPlayer(p *Player) *PlayerSupervisionRequestUpdate {
	return psru.SetPlayerID(p.ID)
}

// AddApprovalIDs adds the "approvals" edge to the PlayerSupervisionRequestApproval entity by IDs.
func (psru *PlayerSupervisionRequestUpdate) AddApprovalIDs(ids ...guidgql.GUID) *PlayerSupervisionRequestUpdate {
	psru.mutation.AddApprovalIDs(ids...)
	return psru
}

// AddApprovals adds the "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psru *PlayerSupervisionRequestUpdate) AddApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestUpdate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psru.AddApprovalIDs(ids...)
}

// Mutation returns the PlayerSupervisionRequestMutation object of the builder.
func (psru *PlayerSupervisionRequestUpdate) Mutation() *PlayerSupervisionRequestMutation {
	return psru.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (psru *PlayerSupervisionRequestUpdate) ClearSender() *PlayerSupervisionRequestUpdate {
	psru.mutation.ClearSender()
	return psru
}

// ClearPlayer clears the "player" edge to the Player entity.
func (psru *PlayerSupervisionRequestUpdate) ClearPlayer() *PlayerSupervisionRequestUpdate {
	psru.mutation.ClearPlayer()
	return psru
}

// ClearApprovals clears all "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psru *PlayerSupervisionRequestUpdate) ClearApprovals() *PlayerSupervisionRequestUpdate {
	psru.mutation.ClearApprovals()
	return psru
}

// RemoveApprovalIDs removes the "approvals" edge to PlayerSupervisionRequestApproval entities by IDs.
func (psru *PlayerSupervisionRequestUpdate) RemoveApprovalIDs(ids ...guidgql.GUID) *PlayerSupervisionRequestUpdate {
	psru.mutation.RemoveApprovalIDs(ids...)
	return psru
}

// RemoveApprovals removes "approvals" edges to PlayerSupervisionRequestApproval entities.
func (psru *PlayerSupervisionRequestUpdate) RemoveApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestUpdate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psru.RemoveApprovalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psru *PlayerSupervisionRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psru.sqlSave, psru.mutation, psru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psru *PlayerSupervisionRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := psru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psru *PlayerSupervisionRequestUpdate) Exec(ctx context.Context) error {
	_, err := psru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psru *PlayerSupervisionRequestUpdate) ExecX(ctx context.Context) {
	if err := psru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psru *PlayerSupervisionRequestUpdate) check() error {
	if _, ok := psru.mutation.SenderID(); psru.mutation.SenderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PlayerSupervisionRequest.sender"`)
	}
	if _, ok := psru.mutation.PlayerID(); psru.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PlayerSupervisionRequest.player"`)
	}
	return nil
}

func (psru *PlayerSupervisionRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := psru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(playersupervisionrequest.Table, playersupervisionrequest.Columns, sqlgraph.NewFieldSpec(playersupervisionrequest.FieldID, field.TypeString))
	if ps := psru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psru.mutation.Message(); ok {
		_spec.SetField(playersupervisionrequest.FieldMessage, field.TypeString, value)
	}
	if psru.mutation.MessageCleared() {
		_spec.ClearField(playersupervisionrequest.FieldMessage, field.TypeString)
	}
	if psru.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psru.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psru.mutation.PlayerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psru.mutation.PlayerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psru.mutation.ApprovalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psru.mutation.RemovedApprovalsIDs(); len(nodes) > 0 && !psru.mutation.ApprovalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psru.mutation.ApprovalsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playersupervisionrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psru.mutation.done = true
	return n, nil
}

// PlayerSupervisionRequestUpdateOne is the builder for updating a single PlayerSupervisionRequest entity.
type PlayerSupervisionRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerSupervisionRequestMutation
}

// SetMessage sets the "message" field.
func (psruo *PlayerSupervisionRequestUpdateOne) SetMessage(s string) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.SetMessage(s)
	return psruo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (psruo *PlayerSupervisionRequestUpdateOne) SetNillableMessage(s *string) *PlayerSupervisionRequestUpdateOne {
	if s != nil {
		psruo.SetMessage(*s)
	}
	return psruo
}

// ClearMessage clears the value of the "message" field.
func (psruo *PlayerSupervisionRequestUpdateOne) ClearMessage() *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.ClearMessage()
	return psruo
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (psruo *PlayerSupervisionRequestUpdateOne) SetSenderID(id guidgql.GUID) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.SetSenderID(id)
	return psruo
}

// SetSender sets the "sender" edge to the User entity.
func (psruo *PlayerSupervisionRequestUpdateOne) SetSender(u *User) *PlayerSupervisionRequestUpdateOne {
	return psruo.SetSenderID(u.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (psruo *PlayerSupervisionRequestUpdateOne) SetPlayerID(id guidgql.GUID) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.SetPlayerID(id)
	return psruo
}

// SetPlayer sets the "player" edge to the Player entity.
func (psruo *PlayerSupervisionRequestUpdateOne) SetPlayer(p *Player) *PlayerSupervisionRequestUpdateOne {
	return psruo.SetPlayerID(p.ID)
}

// AddApprovalIDs adds the "approvals" edge to the PlayerSupervisionRequestApproval entity by IDs.
func (psruo *PlayerSupervisionRequestUpdateOne) AddApprovalIDs(ids ...guidgql.GUID) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.AddApprovalIDs(ids...)
	return psruo
}

// AddApprovals adds the "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psruo *PlayerSupervisionRequestUpdateOne) AddApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestUpdateOne {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psruo.AddApprovalIDs(ids...)
}

// Mutation returns the PlayerSupervisionRequestMutation object of the builder.
func (psruo *PlayerSupervisionRequestUpdateOne) Mutation() *PlayerSupervisionRequestMutation {
	return psruo.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (psruo *PlayerSupervisionRequestUpdateOne) ClearSender() *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.ClearSender()
	return psruo
}

// ClearPlayer clears the "player" edge to the Player entity.
func (psruo *PlayerSupervisionRequestUpdateOne) ClearPlayer() *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.ClearPlayer()
	return psruo
}

// ClearApprovals clears all "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psruo *PlayerSupervisionRequestUpdateOne) ClearApprovals() *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.ClearApprovals()
	return psruo
}

// RemoveApprovalIDs removes the "approvals" edge to PlayerSupervisionRequestApproval entities by IDs.
func (psruo *PlayerSupervisionRequestUpdateOne) RemoveApprovalIDs(ids ...guidgql.GUID) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.RemoveApprovalIDs(ids...)
	return psruo
}

// RemoveApprovals removes "approvals" edges to PlayerSupervisionRequestApproval entities.
func (psruo *PlayerSupervisionRequestUpdateOne) RemoveApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestUpdateOne {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psruo.RemoveApprovalIDs(ids...)
}

// Where appends a list predicates to the PlayerSupervisionRequestUpdate builder.
func (psruo *PlayerSupervisionRequestUpdateOne) Where(ps ...predicate.PlayerSupervisionRequest) *PlayerSupervisionRequestUpdateOne {
	psruo.mutation.Where(ps...)
	return psruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psruo *PlayerSupervisionRequestUpdateOne) Select(field string, fields ...string) *PlayerSupervisionRequestUpdateOne {
	psruo.fields = append([]string{field}, fields...)
	return psruo
}

// Save executes the query and returns the updated PlayerSupervisionRequest entity.
func (psruo *PlayerSupervisionRequestUpdateOne) Save(ctx context.Context) (*PlayerSupervisionRequest, error) {
	return withHooks(ctx, psruo.sqlSave, psruo.mutation, psruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psruo *PlayerSupervisionRequestUpdateOne) SaveX(ctx context.Context) *PlayerSupervisionRequest {
	node, err := psruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psruo *PlayerSupervisionRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := psruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psruo *PlayerSupervisionRequestUpdateOne) ExecX(ctx context.Context) {
	if err := psruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psruo *PlayerSupervisionRequestUpdateOne) check() error {
	if _, ok := psruo.mutation.SenderID(); psruo.mutation.SenderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PlayerSupervisionRequest.sender"`)
	}
	if _, ok := psruo.mutation.PlayerID(); psruo.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PlayerSupervisionRequest.player"`)
	}
	return nil
}

func (psruo *PlayerSupervisionRequestUpdateOne) sqlSave(ctx context.Context) (_node *PlayerSupervisionRequest, err error) {
	if err := psruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(playersupervisionrequest.Table, playersupervisionrequest.Columns, sqlgraph.NewFieldSpec(playersupervisionrequest.FieldID, field.TypeString))
	id, ok := psruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlayerSupervisionRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playersupervisionrequest.FieldID)
		for _, f := range fields {
			if !playersupervisionrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != playersupervisionrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psruo.mutation.Message(); ok {
		_spec.SetField(playersupervisionrequest.FieldMessage, field.TypeString, value)
	}
	if psruo.mutation.MessageCleared() {
		_spec.ClearField(playersupervisionrequest.FieldMessage, field.TypeString)
	}
	if psruo.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psruo.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psruo.mutation.PlayerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psruo.mutation.PlayerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psruo.mutation.ApprovalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psruo.mutation.RemovedApprovalsIDs(); len(nodes) > 0 && !psruo.mutation.ApprovalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psruo.mutation.ApprovalsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlayerSupervisionRequest{config: psruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playersupervisionrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psruo.mutation.done = true
	return _node, nil
}
