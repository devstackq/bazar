package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)

type TransUseCase struct {
	transRepo gallery.TransmissionRepo
}

func TransUseCaseInit(r gallery.TransmissionRepo) gallery.TransUseCase {
	return TransUseCase{transRepo: r}
}

func (tuc TransUseCase) CreateTransmission(t *models.Transmission) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.Create(ctx, t)
}

func (tuc TransUseCase) GetTransByID(id int)(*models.Transmission, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return tuc.transRepo.GetTransByID(ctx, id)
}