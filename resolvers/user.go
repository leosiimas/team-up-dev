package resolvers

import (
	"context"
	"time"

	"github.com/leosiimas/team-up-dev/models"
)

func (r *mutationResolver) UserCreate(ctx context.Context, user models.UserInput) (*models.User, error) {
	usr := models.User{
		Email:     user.Email,
		FirstName: user.FirstName,
	}

	return &usr, nil
}
