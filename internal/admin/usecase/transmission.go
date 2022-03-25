package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type TransUseCase struct {
	transRepo admin.TransmissionRepoInterface
}

func TransUseCaseInit(r admin.TransmissionRepoInterface) admin.TransUseCaseInterface {
	return TransUseCase{transRepo: r}
}

func (tuc TransUseCase) CreateTransmission(t *models.Transmission) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.Create(ctx, t)
}

func (tuc TransUseCase) GetTransmissionByID(id int) (*models.Transmission, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.GetByID(ctx, id)
}

func (cuc TransUseCase) GetListTransmission() ([]*models.Transmission, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.transRepo.GetList(ctx)
}
