package resolvers

import (
	"context"
	"fmt"

	graphql_types "github.com/Sang-it/handy-man-api/graphql-types"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *graphql_types.CreateUserInput) (*graphql_types.User, error) {
	panic(fmt.Errorf("not implemented"))
}
