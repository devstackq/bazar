package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type BodyTypeUseCase struct {
	bodyTypeRepo gallery.BodyTypeRepoInterface
}

func BodyTypeUseCaseInit(r gallery.BodyTypeRepoInterface) gallery.BodyTypeUseCaseInterface {
	return BodyTypeUseCase{bodyTypeRepo: r}
}

func (cuc BodyTypeUseCase) CreateBodyType(c *models.BodyType) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.bodyTypeRepo.Create(ctx, c)
}

func (cuc BodyTypeUseCase) GetBodyTypeByID(id int) (*models.BodyType, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.bodyTypeRepo.GetByID(ctx, id)
}

func (cuc BodyTypeUseCase) GetListBodyType() ([]*models.BodyType, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.bodyTypeRepo.GetList(ctx)
}
