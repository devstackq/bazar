package admin

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type RoleRepoInterface interface {
	Create(context.Context, *models.Role) (int, error)
	GetList(context.Context) ([]*models.Role, error)
}

type CategoryRepoInterface interface {
	Create(context.Context, *models.Category) (int, error)
	GetByID(context.Context, int) (*models.Category, error)
	GetList(context.Context) ([]*models.Category, error)
}

type TransmissionRepoInterface interface {
	Create(context.Context, *models.Transmission) (int, error)
	GetByID(context.Context, int) (*models.Transmission, error)
	GetList(context.Context) ([]*models.Transmission, error)
}

type CountryRepoInterface interface {
	Create(context.Context, *models.Country) (int, error)
	GetByID(context.Context, int) (*models.Country, error)
	GetList(context.Context) ([]*models.Country, error)
}

type CityRepoInterface interface {
	Create(context.Context, *models.City) (int, error)
	GetByID(context.Context, int) (*models.City, error)
	GetList(context.Context, int) ([]*models.City, error)
}

type StateRepoInterface interface {
	Create(context.Context, *models.State) (int, error)
	GetByID(context.Context, int) (*models.State, error)
	GetList(context.Context) ([]*models.State, error)
}
type BrandRepoInterface interface {
	Create(context.Context, *models.Brand) (int, error)
	GetByID(context.Context, int) (*models.Brand, error)
	GetList(context.Context) ([]*models.Brand, error)
}
type ModelRepoInterface interface {
	Create(context.Context, *models.Model) (int, error)
	GetByID(context.Context, int) (*models.Model, error)
	GetList(context.Context, int) ([]*models.Model, error)
}
type FuelRepoInterface interface {
	Create(context.Context, *models.Fuel) (int, error)
	GetByID(context.Context, int) (*models.Fuel, error)
	GetList(context.Context) ([]*models.Fuel, error)
}

type DriveUnitRepoInterface interface {
	Create(context.Context, *models.DriveUnit) (int, error)
	GetByID(context.Context, int) (*models.DriveUnit, error)
	GetList(context.Context) ([]*models.DriveUnit, error)
}
type BodyTypeRepoInterface interface {
	Create(context.Context, *models.BodyType) (int, error)
	GetByID(context.Context, int) (*models.BodyType, error)
	GetList(context.Context) ([]*models.BodyType, error)
}
type ColorRepoInterface interface {
	Create(context.Context, *models.Color) (int, error)
	GetByID(context.Context, int) (*models.Color, error)
	GetList(context.Context) ([]*models.Color, error)
}

type Repositories struct {
	ColorRepoInterface
	BodyTypeRepoInterface
	DriveUnitRepoInterface
	FuelRepoInterface
	ModelRepoInterface
	BrandRepoInterface
	CategoryRepoInterface
	TransmissionRepoInterface
	CountryRepoInterface
	CityRepoInterface
	StateRepoInterface
	RoleRepoInterface
}
