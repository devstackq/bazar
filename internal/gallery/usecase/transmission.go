package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)

type TransUseCase struct {
	transRepo gallery.TransmissionRepoInterface
}

func TransUseCaseInit(r gallery.TransmissionRepoInterface) gallery.TransUseCaseInterface {
	return TransUseCase{transRepo: r}
}

func (tuc TransUseCase) CreateTransmission(t *models.Transmission) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.Create(ctx, t)
}

func (tuc TransUseCase) GetTransmissionByID(id int)(*models.Transmission, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.GetByID(ctx, id)
}

func (cuc TransUseCase) GetListTransmission()([]*models.Transmission, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.transRepo.GetList(ctx)
}