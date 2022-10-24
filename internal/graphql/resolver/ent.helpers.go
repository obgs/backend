package resolver

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

func getNodeType(_ context.Context, id guidgql.GUID) (string, error) {
	return id.Type, nil
}
