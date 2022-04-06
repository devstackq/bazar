package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type ColorUseCase struct {
	colorRepo admin.ColorRepoInterface
}

func ColorUseCaseInit(r admin.ColorRepoInterface) admin.ColorUseCaseInterface {
	return ColorUseCase{colorRepo: r}
}

func (cuc ColorUseCase) CreateColor(c *models.Color) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.Create(ctx, c)
}

func (cuc ColorUseCase) GetColorByID(id int) (*models.Color, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.GetByID(ctx, id)
}

func (cuc ColorUseCase) GetListColor() ([]*models.Color, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.GetList(ctx)
}
