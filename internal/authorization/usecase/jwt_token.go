package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (a JwtTokenUseCase) CreateSession(ctx context.Context, token *models.TokenDetails) error {
	return a.jwtRepo.CreateSession(ctx, token)
}

func (a JwtTokenUseCase) DeleteSession(ctx context.Context, token *models.TokenDetails) error {
	return a.jwtRepo.DeleteSession(ctx, token)
}

func (a JwtTokenUseCase) UpdateSession(ctx context.Context, token *models.TokenDetails) error {
	return a.jwtRepo.UpdateSession(ctx, token)
}
