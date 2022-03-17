package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type FuelUseCase struct {
	fuelRepo gallery.FuelRepoInterface
}

func FuelUseCaseInit(r gallery.FuelRepoInterface) gallery.FuelUseCaseInterface {
	return FuelUseCase{fuelRepo: r}
}

func (cuc FuelUseCase) CreateFuel(c *models.Fuel) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.Create(ctx, c)
}

func (cuc FuelUseCase) GetFuelByID(id int)(*models.Fuel, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.GetByID(ctx, id)
}

func (cuc FuelUseCase) GetListFuel()([]*models.Fuel, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.GetList(ctx)
}