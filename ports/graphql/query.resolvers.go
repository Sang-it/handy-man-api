package resolvers

import (
	"context"
	"fmt"

	"github.com/Sang-it/handy-man-api/generated"
)

func (r *queryResolver) Hi(ctx context.Context, name string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
