// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeSupervisors holds the string denoting the supervisors edge name in mutations.
	EdgeSupervisors = "supervisors"
	// EdgeSupervisionRequests holds the string denoting the supervision_requests edge name in mutations.
	EdgeSupervisionRequests = "supervision_requests"
	// EdgeMatches holds the string denoting the matches edge name in mutations.
	EdgeMatches = "matches"
	// EdgeStats holds the string denoting the stats edge name in mutations.
	EdgeStats = "stats"
	// Table holds the table name of the player in the database.
	Table = "players"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "players"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_main_player"
	// SupervisorsTable is the table that holds the supervisors relation/edge. The primary key declared below.
	SupervisorsTable = "user_players"
	// SupervisorsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SupervisorsInverseTable = "users"
	// SupervisionRequestsTable is the table that holds the supervision_requests relation/edge.
	SupervisionRequestsTable = "player_supervision_requests"
	// SupervisionRequestsInverseTable is the table name for the PlayerSupervisionRequest entity.
	// It exists in this package in order to avoid circular dependency with the "playersupervisionrequest" package.
	SupervisionRequestsInverseTable = "player_supervision_requests"
	// SupervisionRequestsColumn is the table column denoting the supervision_requests relation/edge.
	SupervisionRequestsColumn = "player_supervision_requests"
	// MatchesTable is the table that holds the matches relation/edge. The primary key declared below.
	MatchesTable = "match_players"
	// MatchesInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchesInverseTable = "matches"
	// StatsTable is the table that holds the stats relation/edge.
	StatsTable = "statistics"
	// StatsInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatsInverseTable = "statistics"
	// StatsColumn is the table column denoting the stats relation/edge.
	StatsColumn = "player_stats"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_main_player",
}

var (
	// SupervisorsPrimaryKey and SupervisorsColumn2 are the table columns denoting the
	// primary key for the supervisors relation (M2M).
	SupervisorsPrimaryKey = []string{"user_id", "player_id"}
	// MatchesPrimaryKey and MatchesColumn2 are the table columns denoting the
	// primary key for the matches relation (M2M).
	MatchesPrimaryKey = []string{"match_id", "player_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// OrderOption defines the ordering options for the Player queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// BySupervisorsCount orders the results by supervisors count.
func BySupervisorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSupervisorsStep(), opts...)
	}
}

// BySupervisors orders the results by supervisors terms.
func BySupervisors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSupervisorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySupervisionRequestsCount orders the results by supervision_requests count.
func BySupervisionRequestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSupervisionRequestsStep(), opts...)
	}
}

// BySupervisionRequests orders the results by supervision_requests terms.
func BySupervisionRequests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSupervisionRequestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMatchesCount orders the results by matches count.
func ByMatchesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMatchesStep(), opts...)
	}
}

// ByMatches orders the results by matches terms.
func ByMatches(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMatchesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatsCount orders the results by stats count.
func ByStatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatsStep(), opts...)
	}
}

// ByStats orders the results by stats terms.
func ByStats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
	)
}
func newSupervisorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SupervisorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SupervisorsTable, SupervisorsPrimaryKey...),
	)
}
func newSupervisionRequestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SupervisionRequestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SupervisionRequestsTable, SupervisionRequestsColumn),
	)
}
func newMatchesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MatchesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MatchesTable, MatchesPrimaryKey...),
	)
}
func newStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StatsTable, StatsColumn),
	)
}