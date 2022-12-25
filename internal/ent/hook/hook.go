// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

// The GameFunc type is an adapter to allow the use of ordinary
// function as Game mutator.
type GameFunc func(context.Context, *ent.GameMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GameFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GameMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GameMutation", m)
	}
	return f(ctx, mv)
}

// The GameFavoriteFunc type is an adapter to allow the use of ordinary
// function as GameFavorite mutator.
type GameFavoriteFunc func(context.Context, *ent.GameFavoriteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GameFavoriteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GameFavoriteMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GameFavoriteMutation", m)
	}
	return f(ctx, mv)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
	}
	return f(ctx, mv)
}

// The GroupMembershipFunc type is an adapter to allow the use of ordinary
// function as GroupMembership mutator.
type GroupMembershipFunc func(context.Context, *ent.GroupMembershipMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupMembershipMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMembershipMutation", m)
	}
	return f(ctx, mv)
}

// The GroupMembershipApplicationFunc type is an adapter to allow the use of ordinary
// function as GroupMembershipApplication mutator.
type GroupMembershipApplicationFunc func(context.Context, *ent.GroupMembershipApplicationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipApplicationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupMembershipApplicationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMembershipApplicationMutation", m)
	}
	return f(ctx, mv)
}

// The GroupSettingsFunc type is an adapter to allow the use of ordinary
// function as GroupSettings mutator.
type GroupSettingsFunc func(context.Context, *ent.GroupSettingsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupSettingsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupSettingsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupSettingsMutation", m)
	}
	return f(ctx, mv)
}

// The MatchFunc type is an adapter to allow the use of ordinary
// function as Match mutator.
type MatchFunc func(context.Context, *ent.MatchMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MatchFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MatchMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MatchMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerFunc type is an adapter to allow the use of ordinary
// function as Player mutator.
type PlayerFunc func(context.Context, *ent.PlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerSupervisionRequestFunc type is an adapter to allow the use of ordinary
// function as PlayerSupervisionRequest mutator.
type PlayerSupervisionRequestFunc func(context.Context, *ent.PlayerSupervisionRequestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerSupervisionRequestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerSupervisionRequestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerSupervisionRequestMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerSupervisionRequestApprovalFunc type is an adapter to allow the use of ordinary
// function as PlayerSupervisionRequestApproval mutator.
type PlayerSupervisionRequestApprovalFunc func(context.Context, *ent.PlayerSupervisionRequestApprovalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerSupervisionRequestApprovalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerSupervisionRequestApprovalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerSupervisionRequestApprovalMutation", m)
	}
	return f(ctx, mv)
}

// The StatDescriptionFunc type is an adapter to allow the use of ordinary
// function as StatDescription mutator.
type StatDescriptionFunc func(context.Context, *ent.StatDescriptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatDescriptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatDescriptionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatDescriptionMutation", m)
	}
	return f(ctx, mv)
}

// The StatisticFunc type is an adapter to allow the use of ordinary
// function as Statistic mutator.
type StatisticFunc func(context.Context, *ent.StatisticMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatisticFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatisticMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatisticMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
