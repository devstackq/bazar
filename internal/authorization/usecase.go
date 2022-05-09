package authorization

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type AuthUseCaseInterface interface {
	SignUp(ctx context.Context, user *models.User) (models.User, error)
	SignIn(ctx context.Context, username, password string) (models.User, error)
}

type JwtTokenUseCaseInterface interface {
	CreateSession(context.Context, *models.TokenDetails) error
	DeleteSession(context.Context, *models.TokenDetails) error
	UpdateSession(context.Context, *models.TokenDetails) error
}
