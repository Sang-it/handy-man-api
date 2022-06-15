package resolvers

import (
	"context"
	"fmt"

	"github.com/Sang-it/handy-man-api/generated"
	graphql_types "github.com/Sang-it/handy-man-api/graphql-types"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *graphql_types.CreateUserInput) (*graphql_types.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
