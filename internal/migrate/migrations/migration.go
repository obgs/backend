package migrations

import (
	"context"

	"github.com/obgs/backend/internal/ent"
)

type Migration interface {
	Name() string
	Run(ctx context.Context, client *ent.Client) error
}
