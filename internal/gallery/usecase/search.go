package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)

type SearchUseCase struct {
	searchRepo gallery.SearchRepo
}

func SearchUseCaseInit(r gallery.SearchRepo) gallery.SearchUseCase {
	return SearchUseCase{searchRepo: r}
}

func (suc SearchUseCase) SearchByKeyWord(key string) ([]*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return suc.searchRepo.Search(ctx, key)
}