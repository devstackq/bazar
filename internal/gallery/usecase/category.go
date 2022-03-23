package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type CategoryUseCase struct {
	categoryRepo gallery.CategoryRepoInterface
}

func CategoryUseCaseInit(r gallery.CategoryRepoInterface) gallery.CategoryUseCaseInterface {
	return CategoryUseCase{categoryRepo: r}
}

func (cuc CategoryUseCase) CreateCategory(c *models.Category) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.categoryRepo.Create(ctx, c)
}

func (cuc CategoryUseCase) GetByID(id int) (*models.Category, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.categoryRepo.GetByID(ctx, id)
}

func (cuc CategoryUseCase) GetListCategories() ([]*models.Category, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.categoryRepo.GetList(ctx)
}
