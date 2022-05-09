package authorization

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type AuthRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	GetUser(ctx context.Context, username, password string) (models.User, error) // todo: rename
}

type JwtTokenRepositoryInterface interface {
	CreateSession(context.Context, *models.TokenDetails) error // todo : remove token , end time
	UpdateSession(context.Context, *models.TokenDetails) error
	DeleteSession(context.Context, *models.TokenDetails) error
}
