package usecase

import (
	"time"

	"github.com/devstackq/bazar/internal/authorization"
)

type AuthUseCase struct {
	authRepo       authorization.AuthRepositoryInterface
	HashSalt       string
	expireDuration time.Duration
}

func AuthUseCaseInit(authRepo authorization.AuthRepositoryInterface, hashSalt string, tokenTTLSecond time.Duration) authorization.AuthUseCaseInterface {
	return &AuthUseCase{
		authRepo:       authRepo,
		HashSalt:       hashSalt,
		expireDuration: time.Second * tokenTTLSecond,
	}
}

type JwtTokenUseCase struct {
	jwtRepo authorization.JwtTokenRepositoryInterface
}

func JwtTokenUseCaseInit(jwtRepo authorization.JwtTokenRepositoryInterface) authorization.JwtTokenRepositoryInterface {
	return &JwtTokenUseCase{jwtRepo: jwtRepo}
}
