package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type FuelUseCase struct {
	fuelRepo admin.FuelRepoInterface
}

func FuelUseCaseInit(r admin.FuelRepoInterface) admin.FuelUseCaseInterface {
	return FuelUseCase{fuelRepo: r}
}

func (cuc FuelUseCase) CreateFuel(c *models.Fuel) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.Create(ctx, c)
}

func (cuc FuelUseCase) GetFuelByID(id int) (*models.Fuel, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.GetByID(ctx, id)
}

func (cuc FuelUseCase) GetListFuel() ([]*models.Fuel, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.fuelRepo.GetList(ctx)
}
