// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/obgs/backend/internal/ent/enums"
	"github.com/obgs/backend/internal/ent/group"
	"github.com/obgs/backend/internal/ent/groupmembership"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/user"
)

// GroupMembership is the model entity for the GroupMembership schema.
type GroupMembership struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role enums.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMembershipQuery when eager-loading is set.
	Edges                  GroupMembershipEdges `json:"edges"`
	group_members          *guidgql.GUID
	user_group_memberships *guidgql.GUID
	selectValues           sql.SelectValues
}

// GroupMembershipEdges holds the relations/edges for other nodes in the graph.
type GroupMembershipEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMembership) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldID:
			values[i] = new(guidgql.GUID)
		case groupmembership.FieldRole:
			values[i] = new(sql.NullString)
		case groupmembership.ForeignKeys[0]: // group_members
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		case groupmembership.ForeignKeys[1]: // user_group_memberships
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMembership fields.
func (gm *GroupMembership) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gm.ID = *value
			}
		case groupmembership.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gm.Role = enums.Role(value.String)
			}
		case groupmembership.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_members", values[i])
			} else if value.Valid {
				gm.group_members = new(guidgql.GUID)
				*gm.group_members = *value.S.(*guidgql.GUID)
			}
		case groupmembership.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_group_memberships", values[i])
			} else if value.Valid {
				gm.user_group_memberships = new(guidgql.GUID)
				*gm.user_group_memberships = *value.S.(*guidgql.GUID)
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupMembership.
// This includes values selected through modifiers, order, etc.
func (gm *GroupMembership) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryGroup() *GroupQuery {
	return NewGroupMembershipClient(gm.config).QueryGroup(gm)
}

// QueryUser queries the "user" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryUser() *UserQuery {
	return NewGroupMembershipClient(gm.config).QueryUser(gm)
}

// Update returns a builder for updating this GroupMembership.
// Note that you need to call GroupMembership.Unwrap() before calling this method if this GroupMembership
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMembership) Update() *GroupMembershipUpdateOne {
	return NewGroupMembershipClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GroupMembership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMembership) Unwrap() *GroupMembership {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMembership is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMembership) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMembership(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gm.Role))
	builder.WriteByte(')')
	return builder.String()
}

// GroupMemberships is a parsable slice of GroupMembership.
type GroupMemberships []*GroupMembership
