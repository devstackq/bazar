package usecase

import (
	"context"
	"log"

	"github.com/devstackq/bazar/internal/models"
)

func (auth *AuthUseCase) SignUp(ctx context.Context, user *models.User) (int, error) {
	// auth.HashSalt = auth.generateSalt(16) //salt, then save Db
	user.Password = auth.hashPassword(user.Password) // update password - to hash + salt
	log.Print("call service auth, use case,  Signup", user)
	return auth.authRepo.CreateUser(ctx, user)
}
