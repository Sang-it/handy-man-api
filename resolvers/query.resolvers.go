package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sang-it/handy-man-api/generated"
)

func (r *queryResolver) Hello(ctx context.Context, name *string) (*string, error) {
	output := "Hello " + *name
	return &output, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
