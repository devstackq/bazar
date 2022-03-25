package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type BrandUseCase struct {
	brandRepo admin.BrandRepoInterface
}

func BrandUseCaseInit(r admin.BrandRepoInterface) admin.BrandUseCaseInterface {
	return BrandUseCase{brandRepo: r}
}

func (cuc BrandUseCase) CreateBrand(c *models.Brand) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.brandRepo.Create(ctx, c)
}

func (cuc BrandUseCase) GetBrandByID(id int) (*models.Brand, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.brandRepo.GetByID(ctx, id)
}

func (cuc BrandUseCase) GetListBrand() ([]*models.Brand, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.brandRepo.GetList(ctx)
}
