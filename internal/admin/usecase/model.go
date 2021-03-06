package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type ModelUseCase struct {
	ModelRepo admin.ModelRepoInterface
}

func ModelUseCaseInit(r admin.ModelRepoInterface) admin.ModelUseCaseInterface {
	return ModelUseCase{ModelRepo: r}
}

func (cuc ModelUseCase) CreateModel(c *models.Model) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.ModelRepo.Create(ctx, c)
}

func (cuc ModelUseCase) GetModelByID(id int) (*models.Model, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.ModelRepo.GetByID(ctx, id)
}

func (cuc ModelUseCase) GetListModelByBrandID(brandID int) ([]*models.Model, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.ModelRepo.GetList(ctx, brandID)
}
