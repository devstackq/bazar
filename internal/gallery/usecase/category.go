package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)

type CategoryUseCase struct {
	categoryRepo gallery.CategoryRepo
}

func CategoryUseCaseInit(r gallery.CategoryRepo) gallery.CategoryUseCase {
	return CategoryUseCase{categoryRepo: r}
}

func (muc CategoryUseCase) CreateCategory(c *models.Category) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return muc.categoryRepo.CreateCategory(ctx, c)
}

func (muc CategoryUseCase) GetCategoryByID(id int)(*models.Category, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return muc.categoryRepo.GetCategoryByID(ctx, id)
}

