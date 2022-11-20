// Code generated by ent, DO NOT EDIT.

package group

import (
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLogoURL holds the string denoting the logo_url field in the database.
	FieldLogoURL = "logo_url"
	// EdgeSettings holds the string denoting the settings edge name in mutations.
	EdgeSettings = "settings"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeApplications holds the string denoting the applications edge name in mutations.
	EdgeApplications = "applications"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// SettingsTable is the table that holds the settings relation/edge.
	SettingsTable = "group_settings"
	// SettingsInverseTable is the table name for the GroupSettings entity.
	// It exists in this package in order to avoid circular dependency with the "groupsettings" package.
	SettingsInverseTable = "group_settings"
	// SettingsColumn is the table column denoting the settings relation/edge.
	SettingsColumn = "group_settings"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "group_memberships"
	// MembersInverseTable is the table name for the GroupMembership entity.
	// It exists in this package in order to avoid circular dependency with the "groupmembership" package.
	MembersInverseTable = "group_memberships"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "group_members"
	// ApplicationsTable is the table that holds the applications relation/edge. The primary key declared below.
	ApplicationsTable = "group_applications"
	// ApplicationsInverseTable is the table name for the GroupMembershipApplication entity.
	// It exists in this package in order to avoid circular dependency with the "groupmembershipapplication" package.
	ApplicationsInverseTable = "group_membership_applications"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldLogoURL,
}

var (
	// ApplicationsPrimaryKey and ApplicationsColumn2 are the table columns denoting the
	// primary key for the applications relation (M2M).
	ApplicationsPrimaryKey = []string{"group_id", "group_membership_application_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// LogoURLValidator is a validator for the "logo_url" field. It is called by the builders before save.
	LogoURLValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)