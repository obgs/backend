// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/obgs/backend/internal/ent/player"
	"github.com/obgs/backend/internal/ent/playersupervisionrequest"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
	"github.com/obgs/backend/internal/ent/user"
)

// PlayerSupervisionRequest is the model entity for the PlayerSupervisionRequest schema.
type PlayerSupervisionRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerSupervisionRequestQuery when eager-loading is set.
	Edges                          PlayerSupervisionRequestEdges `json:"edges"`
	player_supervision_requests    *guidgql.GUID
	user_sent_supervision_requests *guidgql.GUID
	selectValues                   sql.SelectValues
}

// PlayerSupervisionRequestEdges holds the relations/edges for other nodes in the graph.
type PlayerSupervisionRequestEdges struct {
	// Sender holds the value of the sender edge.
	Sender *User `json:"sender,omitempty"`
	// Player holds the value of the player edge.
	Player *Player `json:"player,omitempty"`
	// Approvals holds the value of the approvals edge.
	Approvals []*PlayerSupervisionRequestApproval `json:"approvals,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedApprovals map[string][]*PlayerSupervisionRequestApproval
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerSupervisionRequestEdges) SenderOrErr() (*User, error) {
	if e.Sender != nil {
		return e.Sender, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// PlayerOrErr returns the Player value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerSupervisionRequestEdges) PlayerOrErr() (*Player, error) {
	if e.Player != nil {
		return e.Player, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: player.Label}
	}
	return nil, &NotLoadedError{edge: "player"}
}

// ApprovalsOrErr returns the Approvals value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerSupervisionRequestEdges) ApprovalsOrErr() ([]*PlayerSupervisionRequestApproval, error) {
	if e.loadedTypes[2] {
		return e.Approvals, nil
	}
	return nil, &NotLoadedError{edge: "approvals"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlayerSupervisionRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playersupervisionrequest.FieldID:
			values[i] = new(guidgql.GUID)
		case playersupervisionrequest.FieldMessage:
			values[i] = new(sql.NullString)
		case playersupervisionrequest.ForeignKeys[0]: // player_supervision_requests
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		case playersupervisionrequest.ForeignKeys[1]: // user_sent_supervision_requests
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlayerSupervisionRequest fields.
func (psr *PlayerSupervisionRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playersupervisionrequest.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				psr.ID = *value
			}
		case playersupervisionrequest.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				psr.Message = value.String
			}
		case playersupervisionrequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field player_supervision_requests", values[i])
			} else if value.Valid {
				psr.player_supervision_requests = new(guidgql.GUID)
				*psr.player_supervision_requests = *value.S.(*guidgql.GUID)
			}
		case playersupervisionrequest.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_sent_supervision_requests", values[i])
			} else if value.Valid {
				psr.user_sent_supervision_requests = new(guidgql.GUID)
				*psr.user_sent_supervision_requests = *value.S.(*guidgql.GUID)
			}
		default:
			psr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlayerSupervisionRequest.
// This includes values selected through modifiers, order, etc.
func (psr *PlayerSupervisionRequest) Value(name string) (ent.Value, error) {
	return psr.selectValues.Get(name)
}

// QuerySender queries the "sender" edge of the PlayerSupervisionRequest entity.
func (psr *PlayerSupervisionRequest) QuerySender() *UserQuery {
	return NewPlayerSupervisionRequestClient(psr.config).QuerySender(psr)
}

// QueryPlayer queries the "player" edge of the PlayerSupervisionRequest entity.
func (psr *PlayerSupervisionRequest) QueryPlayer() *PlayerQuery {
	return NewPlayerSupervisionRequestClient(psr.config).QueryPlayer(psr)
}

// QueryApprovals queries the "approvals" edge of the PlayerSupervisionRequest entity.
func (psr *PlayerSupervisionRequest) QueryApprovals() *PlayerSupervisionRequestApprovalQuery {
	return NewPlayerSupervisionRequestClient(psr.config).QueryApprovals(psr)
}

// Update returns a builder for updating this PlayerSupervisionRequest.
// Note that you need to call PlayerSupervisionRequest.Unwrap() before calling this method if this PlayerSupervisionRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (psr *PlayerSupervisionRequest) Update() *PlayerSupervisionRequestUpdateOne {
	return NewPlayerSupervisionRequestClient(psr.config).UpdateOne(psr)
}

// Unwrap unwraps the PlayerSupervisionRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (psr *PlayerSupervisionRequest) Unwrap() *PlayerSupervisionRequest {
	_tx, ok := psr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlayerSupervisionRequest is not a transactional entity")
	}
	psr.config.driver = _tx.drv
	return psr
}

// String implements the fmt.Stringer.
func (psr *PlayerSupervisionRequest) String() string {
	var builder strings.Builder
	builder.WriteString("PlayerSupervisionRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", psr.ID))
	builder.WriteString("message=")
	builder.WriteString(psr.Message)
	builder.WriteByte(')')
	return builder.String()
}

// NamedApprovals returns the Approvals named value or an error if the edge was not
// loaded in eager-loading with this name.
func (psr *PlayerSupervisionRequest) NamedApprovals(name string) ([]*PlayerSupervisionRequestApproval, error) {
	if psr.Edges.namedApprovals == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := psr.Edges.namedApprovals[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (psr *PlayerSupervisionRequest) appendNamedApprovals(name string, edges ...*PlayerSupervisionRequestApproval) {
	if psr.Edges.namedApprovals == nil {
		psr.Edges.namedApprovals = make(map[string][]*PlayerSupervisionRequestApproval)
	}
	if len(edges) == 0 {
		psr.Edges.namedApprovals[name] = []*PlayerSupervisionRequestApproval{}
	} else {
		psr.Edges.namedApprovals[name] = append(psr.Edges.namedApprovals[name], edges...)
	}
}

// PlayerSupervisionRequests is a parsable slice of PlayerSupervisionRequest.
type PlayerSupervisionRequests []*PlayerSupervisionRequest
