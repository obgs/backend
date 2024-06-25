package migrations

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

type Migration interface {
	Name() string
	Run(ctx context.Context, client *ent.Client) error
}
