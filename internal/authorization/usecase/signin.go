package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (models.User, error) {
	// dbPassword, err := auth.authRepo.GetUserPassword(ctx, username)
	inputHashedPwd := auth.hashPassword(password)
	// todo: add email check
	res, err := auth.authRepo.GetUser(ctx, username, inputHashedPwd)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
