package seed

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

func SeedDB(ctx context.Context, client *ent.Client) (err error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return
	}

	player1, player2, player3 := createPlayers(ctx, tx)

	user1, user2 := createUsers(ctx, tx, player1, player2)

	tx.Player.UpdateOne(player3).AddSupervisors(user1, user2).ExecX(ctx)

	terraformingMars := createTerraformingMars(ctx, tx, user1)

	addGameToFavorites(ctx, tx, user1, terraformingMars)
	addGameToFavorites(ctx, tx, user2, terraformingMars)

	v1 := terraformingMars.QueryVersions().AllX(ctx)[0]
	createMatch(ctx, tx, v1, []*ent.Player{player1, player2, player3})
	createMatch(ctx, tx, v1, []*ent.Player{player1, player3})
	createMatch(ctx, tx, v1, []*ent.Player{player1, player2})
	createMatch(ctx, tx, v1, []*ent.Player{player2, player3})

	return tx.Commit()
}
