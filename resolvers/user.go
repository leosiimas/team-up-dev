package resolvers

import (
	"context"

	"github.com/leosiimas/team-up-dev/models"
)

func (r *mutationResolver) UserCreate(ctx context.Context, user models.UserInput) (*models.User, error) {
	usr := models.User{
		Name: user.Name,
	}

	return &usr, nil
}
