package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"

	"github.com/devstackq/bazar/internal/models"
)

func (auth *AuthUseCase) SignUp(ctx context.Context, user *models.User) (int, error) {
	// auth.HashSalt = auth.generateSalt(16) //salt, then save Db
	user.Password = auth.hashPassword(user.Password) // update password - to hash + salt
	log.Print("call service auth, use case,  Signup", user)
	return auth.authRepo.CreateUser(ctx, user)
}

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

func (auth *AuthUseCase) hashPassword(password string) string {
	sha1Hasher := sha1.New()
	pwdBytes := []byte(password)
	// append hased password, with salt
	pwdBytes = append(pwdBytes, []byte(auth.HashSalt)...)
	sha1Hasher.Write(pwdBytes)                    // write bytes - to hasher
	return fmt.Sprintf("%x", sha1Hasher.Sum(nil)) // hashed password
	// base64EncodingPasswordHash := base64.URLEncoding.EncodeToString(hashPasswordBytes)
}

func (a AuthUseCase) CreateSession(ctx context.Context, token *models.TokenDetails) error {
	return a.authRepo.CreateSession(ctx, token)
}

func (a AuthUseCase) DeleteSession(ctx context.Context, token *models.TokenDetails) error {
	return a.authRepo.DeleteSession(ctx, token)
}

func (a AuthUseCase) UpdateSession(ctx context.Context, token *models.TokenDetails) error {
	return a.authRepo.UpdateSession(ctx, token)
}
