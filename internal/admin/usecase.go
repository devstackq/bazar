package admin

import "github.com/devstackq/bazar/internal/models"

type CityUseCaseInterface interface {
	CreateCity(*models.City) (int, error)
	GetCityByID(int) (*models.City, error)
	GetListCityByCountryID(int) ([]*models.City, error)
}

type CountryUseCaseInterface interface {
	CreateCountry(*models.Country) (int, error)
	GetCountryByID(int) (*models.Country, error)
	GetListCountry() ([]*models.Country, error)
}

type CategoryUseCaseInterface interface {
	CreateCategory(*models.Category) (int, error)
	GetByID(int) (*models.Category, error)
	GetListCategories() ([]*models.Category, error)
}

type TransUseCaseInterface interface {
	CreateTransmission(t *models.Transmission) (int, error)
	GetTransmissionByID(id int) (*models.Transmission, error)
	GetListTransmission() ([]*models.Transmission, error)
}

type StateUseCaseInterface interface {
	CreateState(t *models.State) (int, error)
	GetStateByID(id int) (*models.State, error)
	GetListState() ([]*models.State, error)
}
type BrandUseCaseInterface interface {
	CreateBrand(t *models.Brand) (int, error)
	GetBrandByID(id int) (*models.Brand, error)
	GetListBrand() ([]*models.Brand, error)
}

type ModelUseCaseInterface interface {
	CreateModel(t *models.Model) (int, error)
	GetModelByID(id int) (*models.Model, error)
	GetListModelByBrandID(int) ([]*models.Model, error)
	// GetListModel()([]*models.Model, error)
}
type FuelUseCaseInterface interface {
	CreateFuel(t *models.Fuel) (int, error)
	GetFuelByID(id int) (*models.Fuel, error)
	GetListFuel() ([]*models.Fuel, error)
}
type DriveUnitUseCaseInterface interface {
	CreateDriveUnit(t *models.DriveUnit) (int, error)
	GetDriveUnitByID(id int) (*models.DriveUnit, error)
	GetListDriveUnit() ([]*models.DriveUnit, error)
}

type BodyTypeUseCaseInterface interface {
	CreateBodyType(t *models.BodyType) (int, error)
	GetBodyTypeByID(id int) (*models.BodyType, error)
	GetListBodyType() ([]*models.BodyType, error)
}

type ColorUseCaseInterface interface {
	CreateColor(t *models.Color) (int, error)
	GetColorByID(id int) (*models.Color, error)
	GetListColor() ([]*models.Color, error)
}

type RoleUseCaseInterface interface {
	CreateRole(*models.Role) (int, error)
	GetListRole() ([]*models.Role, error)
}

type UseCases struct {
	ColorUseCaseInterface
	BodyTypeUseCaseInterface
	DriveUnitUseCaseInterface
	FuelUseCaseInterface
	ModelUseCaseInterface
	BrandUseCaseInterface
	StateUseCaseInterface
	CategoryUseCaseInterface
	TransUseCaseInterface
	CountryUseCaseInterface
	CityUseCaseInterface
	RoleUseCaseInterface
}
