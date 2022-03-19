package auth

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type AuthRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) (int, error)
	GetUser(ctx context.Context, username, password string) (int, error)//todo: rename
	CreateSession(context.Context , string, string, int) error //todo : remove token , end time 
}
