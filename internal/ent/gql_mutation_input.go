// Code generated by ent, DO NOT EDIT.

package ent

import "github.com/google/uuid"

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name            *string
	Email           *string
	AvatarURL       *string
	AddPlayerIDs    []uuid.UUID
	RemovePlayerIDs []uuid.UUID
	ClearMainPlayer bool
	MainPlayerID    *uuid.UUID
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if v := i.AddPlayerIDs; len(v) > 0 {
		m.AddPlayerIDs(v...)
	}
	if v := i.RemovePlayerIDs; len(v) > 0 {
		m.RemovePlayerIDs(v...)
	}
	if i.ClearMainPlayer {
		m.ClearMainPlayer()
	}
	if v := i.MainPlayerID; v != nil {
		m.SetMainPlayerID(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
