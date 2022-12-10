package guidgql

type Table int

const (
	User Table = iota
	Player
	PlayerSupervisionRequest
	PlayerSupervisionRequestApproval
	Group
	GroupMembership
	GroupMembershipApplication
	GroupSettings
	Game
	GameFavorites
)

var TableNames = map[Table]string{
	User:                             "users",
	Player:                           "players",
	PlayerSupervisionRequest:         "player_supervision_requests",
	PlayerSupervisionRequestApproval: "player_supervision_request_approvals",
	Group:                            "groups",
	GroupMembership:                  "group_memberships",
	GroupMembershipApplication:       "group_membership_applications",
	GroupSettings:                    "group_settings",
	Game:                             "games",
	GameFavorites:                    "game_favorites",
}
