package auth

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) (int,  error) 
	SignIn(ctx context.Context, username, password string) (int, error)
	CreateSession(context.Context, string, string, int) error
}
