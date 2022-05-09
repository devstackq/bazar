package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (auth *AuthUseCase) SignUp(ctx context.Context, user *models.User) (models.User, error) {
	// auth.HashSalt = auth.generateSalt(16) //salt, then save Db
	user.Password = auth.hashPassword(user.Password) // update password - to hash + salt
	return auth.authRepo.CreateUser(ctx, user)
}
