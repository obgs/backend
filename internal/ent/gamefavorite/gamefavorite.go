// Code generated by ent, DO NOT EDIT.

package gamefavorite

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the gamefavorite type in the database.
	Label = "game_favorite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the gamefavorite in the database.
	Table = "game_favorites"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "game_favorites"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_favorites"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "game_favorites"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_favorite_games"
)

// Columns holds all SQL columns for gamefavorite fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "game_favorites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_favorites",
	"user_favorite_games",
}

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// OrderOption defines the ordering options for the GameFavorite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGameField orders the results by game field.
func ByGameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newGameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GameTable, GameColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
