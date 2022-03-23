package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type StateUseCase struct {
	stateRepo gallery.StateRepoInterface
}

func StateUseCaseInit(r gallery.StateRepoInterface) gallery.StateUseCaseInterface {
	return StateUseCase{stateRepo: r}
}

func (cuc StateUseCase) CreateState(c *models.State) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.stateRepo.Create(ctx, c)
}

func (cuc StateUseCase) GetStateByID(id int) (*models.State, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.stateRepo.GetByID(ctx, id)
}

func (cuc StateUseCase) GetListState() ([]*models.State, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.stateRepo.GetList(ctx)
}
