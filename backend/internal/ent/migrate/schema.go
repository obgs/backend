// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "min_players", Type: field.TypeInt, Default: 1},
		{Name: "max_players", Type: field.TypeInt, Default: 10},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "boardgamegeek_url", Type: field.TypeString, Nullable: true},
		{Name: "user_games", Type: field.TypeString},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_users_games",
				Columns:    []*schema.Column{GamesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GameFavoritesColumns holds the columns for the "game_favorites" table.
	GameFavoritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "game_favorites", Type: field.TypeString},
		{Name: "user_favorite_games", Type: field.TypeString},
	}
	// GameFavoritesTable holds the schema information for the "game_favorites" table.
	GameFavoritesTable = &schema.Table{
		Name:       "game_favorites",
		Columns:    GameFavoritesColumns,
		PrimaryKey: []*schema.Column{GameFavoritesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_favorites_games_favorites",
				Columns:    []*schema.Column{GameFavoritesColumns[1]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "game_favorites_users_favorite_games",
				Columns:    []*schema.Column{GameFavoritesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GameVersionsColumns holds the columns for the "game_versions" table.
	GameVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "version_number", Type: field.TypeInt, Default: 1},
		{Name: "game_version_game", Type: field.TypeString},
	}
	// GameVersionsTable holds the schema information for the "game_versions" table.
	GameVersionsTable = &schema.Table{
		Name:       "game_versions",
		Columns:    GameVersionsColumns,
		PrimaryKey: []*schema.Column{GameVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_versions_games_game",
				Columns:    []*schema.Column{GameVersionsColumns[2]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "logo_url", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// GroupMembershipsColumns holds the columns for the "group_memberships" table.
	GroupMembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"owner", "admin", "member"}},
		{Name: "group_members", Type: field.TypeString},
		{Name: "user_group_memberships", Type: field.TypeString},
	}
	// GroupMembershipsTable holds the schema information for the "group_memberships" table.
	GroupMembershipsTable = &schema.Table{
		Name:       "group_memberships",
		Columns:    GroupMembershipsColumns,
		PrimaryKey: []*schema.Column{GroupMembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_memberships_groups_members",
				Columns:    []*schema.Column{GroupMembershipsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_memberships_users_group_memberships",
				Columns:    []*schema.Column{GroupMembershipsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupMembershipApplicationsColumns holds the columns for the "group_membership_applications" table.
	GroupMembershipApplicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "message", Type: field.TypeString, Default: ""},
		{Name: "group_applications", Type: field.TypeString},
		{Name: "user_group_membership_applications", Type: field.TypeString},
	}
	// GroupMembershipApplicationsTable holds the schema information for the "group_membership_applications" table.
	GroupMembershipApplicationsTable = &schema.Table{
		Name:       "group_membership_applications",
		Columns:    GroupMembershipApplicationsColumns,
		PrimaryKey: []*schema.Column{GroupMembershipApplicationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_membership_applications_groups_applications",
				Columns:    []*schema.Column{GroupMembershipApplicationsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_membership_applications_users_group_membership_applications",
				Columns:    []*schema.Column{GroupMembershipApplicationsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupSettingsColumns holds the columns for the "group_settings" table.
	GroupSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PUBLIC", "PRIVATE"}, Default: "PUBLIC"},
		{Name: "join_policy", Type: field.TypeEnum, Enums: []string{"OPEN", "INVITE_ONLY", "APPLICATION_ONLY", "INVITE_OR_APPLICATION"}, Default: "OPEN"},
		{Name: "minimum_role_to_invite", Type: field.TypeEnum, Nullable: true, Enums: []string{"owner", "admin", "member"}},
		{Name: "group_settings", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// GroupSettingsTable holds the schema information for the "group_settings" table.
	GroupSettingsTable = &schema.Table{
		Name:       "group_settings",
		Columns:    GroupSettingsColumns,
		PrimaryKey: []*schema.Column{GroupSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_settings_groups_settings",
				Columns:    []*schema.Column{GroupSettingsColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "game_version_matches", Type: field.TypeString},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_game_versions_matches",
				Columns:    []*schema.Column{MatchesColumns[2]},
				RefColumns: []*schema.Column{GameVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "user_main_player", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "players_users_main_player",
				Columns:    []*schema.Column{PlayersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PlayerSupervisionRequestsColumns holds the columns for the "player_supervision_requests" table.
	PlayerSupervisionRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "message", Type: field.TypeString, Nullable: true},
		{Name: "player_supervision_requests", Type: field.TypeString},
		{Name: "user_sent_supervision_requests", Type: field.TypeString},
	}
	// PlayerSupervisionRequestsTable holds the schema information for the "player_supervision_requests" table.
	PlayerSupervisionRequestsTable = &schema.Table{
		Name:       "player_supervision_requests",
		Columns:    PlayerSupervisionRequestsColumns,
		PrimaryKey: []*schema.Column{PlayerSupervisionRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "player_supervision_requests_players_supervision_requests",
				Columns:    []*schema.Column{PlayerSupervisionRequestsColumns[2]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "player_supervision_requests_users_sent_supervision_requests",
				Columns:    []*schema.Column{PlayerSupervisionRequestsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PlayerSupervisionRequestApprovalsColumns holds the columns for the "player_supervision_request_approvals" table.
	PlayerSupervisionRequestApprovalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "approved", Type: field.TypeBool, Nullable: true},
		{Name: "player_supervision_request_approvals", Type: field.TypeString},
		{Name: "user_supervision_request_approvals", Type: field.TypeString},
	}
	// PlayerSupervisionRequestApprovalsTable holds the schema information for the "player_supervision_request_approvals" table.
	PlayerSupervisionRequestApprovalsTable = &schema.Table{
		Name:       "player_supervision_request_approvals",
		Columns:    PlayerSupervisionRequestApprovalsColumns,
		PrimaryKey: []*schema.Column{PlayerSupervisionRequestApprovalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "player_supervision_request_approvals_player_supervision_requests_approvals",
				Columns:    []*schema.Column{PlayerSupervisionRequestApprovalsColumns[2]},
				RefColumns: []*schema.Column{PlayerSupervisionRequestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "player_supervision_request_approvals_users_supervision_request_approvals",
				Columns:    []*schema.Column{PlayerSupervisionRequestApprovalsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StatDescriptionsColumns holds the columns for the "stat_descriptions" table.
	StatDescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"numeric", "enum", "aggregate"}},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "metadata", Type: field.TypeString, Nullable: true},
		{Name: "order_number", Type: field.TypeInt},
	}
	// StatDescriptionsTable holds the schema information for the "stat_descriptions" table.
	StatDescriptionsTable = &schema.Table{
		Name:       "stat_descriptions",
		Columns:    StatDescriptionsColumns,
		PrimaryKey: []*schema.Column{StatDescriptionsColumns[0]},
	}
	// StatisticsColumns holds the columns for the "statistics" table.
	StatisticsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Default: ""},
		{Name: "match_stats", Type: field.TypeString},
		{Name: "player_stats", Type: field.TypeString},
		{Name: "stat_description_stats", Type: field.TypeString},
	}
	// StatisticsTable holds the schema information for the "statistics" table.
	StatisticsTable = &schema.Table{
		Name:       "statistics",
		Columns:    StatisticsColumns,
		PrimaryKey: []*schema.Column{StatisticsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "statistics_matches_stats",
				Columns:    []*schema.Column{StatisticsColumns[2]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "statistics_players_stats",
				Columns:    []*schema.Column{StatisticsColumns[3]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "statistics_stat_descriptions_stats",
				Columns:    []*schema.Column{StatisticsColumns[4]},
				RefColumns: []*schema.Column{StatDescriptionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString, Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// MatchPlayersColumns holds the columns for the "match_players" table.
	MatchPlayersColumns = []*schema.Column{
		{Name: "match_id", Type: field.TypeString},
		{Name: "player_id", Type: field.TypeString},
	}
	// MatchPlayersTable holds the schema information for the "match_players" table.
	MatchPlayersTable = &schema.Table{
		Name:       "match_players",
		Columns:    MatchPlayersColumns,
		PrimaryKey: []*schema.Column{MatchPlayersColumns[0], MatchPlayersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "match_players_match_id",
				Columns:    []*schema.Column{MatchPlayersColumns[0]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "match_players_player_id",
				Columns:    []*schema.Column{MatchPlayersColumns[1]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StatDescriptionGameVersionColumns holds the columns for the "stat_description_game_version" table.
	StatDescriptionGameVersionColumns = []*schema.Column{
		{Name: "stat_description_id", Type: field.TypeString},
		{Name: "game_version_id", Type: field.TypeString},
	}
	// StatDescriptionGameVersionTable holds the schema information for the "stat_description_game_version" table.
	StatDescriptionGameVersionTable = &schema.Table{
		Name:       "stat_description_game_version",
		Columns:    StatDescriptionGameVersionColumns,
		PrimaryKey: []*schema.Column{StatDescriptionGameVersionColumns[0], StatDescriptionGameVersionColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stat_description_game_version_stat_description_id",
				Columns:    []*schema.Column{StatDescriptionGameVersionColumns[0]},
				RefColumns: []*schema.Column{StatDescriptionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "stat_description_game_version_game_version_id",
				Columns:    []*schema.Column{StatDescriptionGameVersionColumns[1]},
				RefColumns: []*schema.Column{GameVersionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPlayersColumns holds the columns for the "user_players" table.
	UserPlayersColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "player_id", Type: field.TypeString},
	}
	// UserPlayersTable holds the schema information for the "user_players" table.
	UserPlayersTable = &schema.Table{
		Name:       "user_players",
		Columns:    UserPlayersColumns,
		PrimaryKey: []*schema.Column{UserPlayersColumns[0], UserPlayersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_players_user_id",
				Columns:    []*schema.Column{UserPlayersColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_players_player_id",
				Columns:    []*schema.Column{UserPlayersColumns[1]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GamesTable,
		GameFavoritesTable,
		GameVersionsTable,
		GroupsTable,
		GroupMembershipsTable,
		GroupMembershipApplicationsTable,
		GroupSettingsTable,
		MatchesTable,
		PlayersTable,
		PlayerSupervisionRequestsTable,
		PlayerSupervisionRequestApprovalsTable,
		StatDescriptionsTable,
		StatisticsTable,
		UsersTable,
		MatchPlayersTable,
		StatDescriptionGameVersionTable,
		UserPlayersTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = UsersTable
	GameFavoritesTable.ForeignKeys[0].RefTable = GamesTable
	GameFavoritesTable.ForeignKeys[1].RefTable = UsersTable
	GameVersionsTable.ForeignKeys[0].RefTable = GamesTable
	GroupMembershipsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembershipsTable.ForeignKeys[1].RefTable = UsersTable
	GroupMembershipApplicationsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembershipApplicationsTable.ForeignKeys[1].RefTable = UsersTable
	GroupSettingsTable.ForeignKeys[0].RefTable = GroupsTable
	MatchesTable.ForeignKeys[0].RefTable = GameVersionsTable
	PlayersTable.ForeignKeys[0].RefTable = UsersTable
	PlayerSupervisionRequestsTable.ForeignKeys[0].RefTable = PlayersTable
	PlayerSupervisionRequestsTable.ForeignKeys[1].RefTable = UsersTable
	PlayerSupervisionRequestApprovalsTable.ForeignKeys[0].RefTable = PlayerSupervisionRequestsTable
	PlayerSupervisionRequestApprovalsTable.ForeignKeys[1].RefTable = UsersTable
	StatisticsTable.ForeignKeys[0].RefTable = MatchesTable
	StatisticsTable.ForeignKeys[1].RefTable = PlayersTable
	StatisticsTable.ForeignKeys[2].RefTable = StatDescriptionsTable
	MatchPlayersTable.ForeignKeys[0].RefTable = MatchesTable
	MatchPlayersTable.ForeignKeys[1].RefTable = PlayersTable
	StatDescriptionGameVersionTable.ForeignKeys[0].RefTable = StatDescriptionsTable
	StatDescriptionGameVersionTable.ForeignKeys[1].RefTable = GameVersionsTable
	UserPlayersTable.ForeignKeys[0].RefTable = UsersTable
	UserPlayersTable.ForeignKeys[1].RefTable = PlayersTable
}