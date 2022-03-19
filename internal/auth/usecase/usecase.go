package usecase

import (
	"time"

	"github.com/devstackq/bazar/internal/auth"
)

//repo;
//func constructor
//DI - each db - own realize; - condition - interface
//AuthUseCase struct - for relation between - layers; interface  - poly, DI;
type AuthUseCase struct {
	authRepo       auth.AuthRepositoryInterface
	HashSalt       string
	// signinKey      string
	expireDuration time.Duration
}

func AuthUseCaseInit(authRepo auth.AuthRepositoryInterface, hashSalt string, tokenTTLSecond time.Duration) *AuthUseCase {
	return &AuthUseCase{
		authRepo:       authRepo,
		HashSalt:       hashSalt,
		// signinKey:      signinKey,
		expireDuration: time.Second * tokenTTLSecond,
	}
}
