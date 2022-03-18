package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)


type ColorUseCase struct {
	colorRepo gallery.ColorRepoInterface
}

func ColorUseCaseInit(r gallery.ColorRepoInterface) gallery.ColorUseCaseInterface {
	return ColorUseCase{colorRepo: r}
}

func (cuc ColorUseCase) CreateColor(c *models.Color) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.Create(ctx, c)
}

func (cuc ColorUseCase) GetColorByID(id int)(*models.Color, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.GetByID(ctx, id)
}

func (cuc ColorUseCase) GetListColor()([]*models.Color, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.colorRepo.GetList(ctx)
}