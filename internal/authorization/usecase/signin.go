package usecase

import "context"

func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (int, error) {
	// dbPassword, err := auth.authRepo.GetUserPassword(ctx, username)
	inputHashedPwd := auth.hashPassword(password)
	// todo: add email check
	id, err := auth.authRepo.GetUser(ctx, username, inputHashedPwd)
	if err != nil {
		return 0, err
	}
	return id, nil
}
