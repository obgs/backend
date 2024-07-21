package seed

import (
	"context"

	"github.com/obgs/backend/internal/ent"
)

func createPlayers(ctx context.Context, tx *ent.Tx) (*ent.Player, *ent.Player, *ent.Player) {
	player1 := tx.Player.Create().
		SetName("test-player-1").
		SaveX(ctx)

	player2 := tx.Player.Create().
		SetName("test-player-2").
		SaveX(ctx)

	player3 := tx.Player.Create().
		SetName("test-player-3").
		SaveX(ctx)

	return player1, player2, player3
}
