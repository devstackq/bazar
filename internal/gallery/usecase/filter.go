package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type FilterUseCase struct {
	filterRepo gallery.FilterRepoInterface
}

func FilterUseCaseInit(r gallery.FilterRepoInterface) gallery.FilterUseCaseInterface {
	return FilterUseCase{filterRepo: r}
}

func (cuc FilterUseCase) GetListMachineByFilter(keys map[string]string) ([]*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.filterRepo.GetListMachineByFilter(ctx, keys)
}
