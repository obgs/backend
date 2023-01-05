package resolver

import (
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

func checkThatPlayerIsPresent(playerID guidgql.GUID, players []*guidgql.GUID) error {
	playerIsPresent := false
	for _, p := range players {
		if p.ID == playerID.ID {
			playerIsPresent = true
			break
		}
	}
	if !playerIsPresent {
		return fmt.Errorf("player %s is not in the player list", playerID.ID)
	}

	return nil
}
