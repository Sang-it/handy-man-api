package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/Sang-it/handy-man-api/internal-new/adapters/framework/left/graphql/generated"
)

func (r *queryResolver) Hi(ctx context.Context, name string) (*string, error) {
	result := "hi" + name
	return &result, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
