package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	main "github.com/leosiimas/team-up"
	"github.com/leosiimas/team-up/graph/generated"
)

type Resolver struct{}

func (r *mutationResolver) UserCreate(ctx context.Context, user main.UserInput) (*main.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) Login(ctx context.Context, input main.Login) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input main.RefreshTokenInput) (string, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
