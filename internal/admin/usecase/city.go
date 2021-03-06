package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type CityUseCase struct {
	cityRepo admin.CityRepoInterface
}

func CityUseCaseInit(r admin.CityRepoInterface) admin.CityUseCaseInterface {
	return CityUseCase{cityRepo: r}
}

func (cuc CityUseCase) CreateCity(c *models.City) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.cityRepo.Create(ctx, c)
}

func (cuc CityUseCase) GetCityByID(id int) (*models.City, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.cityRepo.GetByID(ctx, id)
}

func (cuc CityUseCase) GetListCityByCountryID(countryID int) ([]*models.City, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.cityRepo.GetList(ctx, countryID)
}
