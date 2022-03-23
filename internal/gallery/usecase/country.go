package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type CountryUseCase struct {
	countryRepo gallery.CountryRepoInterface
}

func CountryUseCaseInit(r gallery.CountryRepoInterface) gallery.CountryUseCaseInterface {
	return CountryUseCase{countryRepo: r}
}

func (cuc CountryUseCase) CreateCountry(c *models.Country) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.countryRepo.Create(ctx, c)
}

func (cuc CountryUseCase) GetCountryByID(id int) (*models.Country, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.countryRepo.GetByID(ctx, id)
}

func (cuc CountryUseCase) GetListCountry() ([]*models.Country, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.countryRepo.GetList(ctx)
}
