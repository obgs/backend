package resolver

import (
	"context"

	"github.com/obgs/backend/internal/ent/schema/guidgql"
)

func getNodeType(_ context.Context, id guidgql.GUID) (string, error) {
	return guidgql.TableNames[id.Type], nil
}
