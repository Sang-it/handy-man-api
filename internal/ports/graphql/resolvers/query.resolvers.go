package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sang-it/handy-man-api/internal/ports/graphql/generated"
)

func (r *queryResolver) Hi(ctx context.Context, name string) (*string, error) {
	result := "hi" + name
	return &result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
