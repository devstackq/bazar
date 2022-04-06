package usecase

import (
	"time"

	"github.com/devstackq/bazar/internal/auth"
)

type AuthUseCase struct {
	authRepo       auth.AuthRepositoryInterface
	HashSalt       string
	expireDuration time.Duration
}

func AuthUseCaseInit(authRepo auth.AuthRepositoryInterface, hashSalt string, tokenTTLSecond time.Duration) *AuthUseCase {
	return &AuthUseCase{
		authRepo:       authRepo,
		HashSalt:       hashSalt,
		expireDuration: time.Second * tokenTTLSecond,
	}
}
