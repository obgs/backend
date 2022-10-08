package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
)

// GetFileUploadURL is the resolver for the getFileUploadURL field.
func (r *queryResolver) GetFileUploadURL(ctx context.Context) (string, error) {
	return r.filestorage.SignUploadURL(ctx)
}
