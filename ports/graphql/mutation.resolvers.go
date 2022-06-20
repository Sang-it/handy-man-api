package resolvers

import (
	"context"
	"fmt"

	"github.com/Sang-it/handy-man-api/generated"
)

func (r *mutationResolver) Hello(ctx context.Context, name string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
