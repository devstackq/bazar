package gallery

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery/models"
)

type MachineRepoInterface interface {
	Create(context.Context, *models.Machine)(int, error)
	GetByID(context.Context, int)(*models.Machine, error)
	GetList(context.Context)([]*models.Machine,  error)
	GetListByUserID(context.Context, int)([]*models.Machine,  error)

}
type CategoryRepoInterface interface {
	// GetyID(context.Context, int) (*models.Category, error)
	Create(context.Context, *models.Category)(int, error)
	GetByID(context.Context, int)(*models.Category, error)
	GetList(context.Context) ([]*models.Category, error)
}

type SearchRepoInterface interface {
	Search(context.Context, string)([]*models.Machine, error)
}

type FilterRepoInterface interface {
	//search all key, prepare quesery
	GetListMachineByFilter(context.Context, map[string]string)([]*models.Machine, error)
}

type TransmissionRepoInterface interface {
	Create(context.Context,  *models.Transmission) (int, error) 
	GetByID(context.Context, int)(*models.Transmission, error)
	GetList(context.Context)([]*models.Transmission, error)
}

type CountryRepoInterface interface {
	Create(context.Context, *models.Country) (int, error)
	GetByID(context.Context, int)(*models.Country, error)
	GetList(context.Context ) ([]*models.Country, error)
}

type CityRepoInterface interface {
	Create(context.Context, *models.City) (int, error)
	GetByID(context.Context,int)(*models.City, error)
	GetList(context.Context) ([]*models.City, error)
}

type StateRepoInterface interface {
	Create(context.Context, *models.State) (int, error)
	GetByID(context.Context,int)(*models.State, error)
	GetList(context.Context) ([]*models.State, error)
}
type BrandRepoInterface interface {
	Create(context.Context, *models.Brand) (int, error)
	GetByID(context.Context,int)(*models.Brand, error)
	GetList(context.Context) ([]*models.Brand, error)
}
type ModelRepoInterface interface {
	Create(context.Context, *models.Model) (int, error)
	GetByID(context.Context,int)(*models.Model, error)
	GetList(context.Context, int) ([]*models.Model, error)
}
type FuelRepoInterface interface {
	Create(context.Context, *models.Fuel) (int, error)
	GetByID(context.Context,int)(*models.Fuel, error)
	GetList(context.Context) ([]*models.Fuel, error)
}

type DriveUnitRepoInterface interface {
	Create(context.Context, *models.DriveUnit) (int, error)
	GetByID(context.Context,int)(*models.DriveUnit, error)
	GetList(context.Context) ([]*models.DriveUnit, error)
}
type BodyTypeRepoInterface interface {
	Create(context.Context, *models.BodyType) (int, error)
	GetByID(context.Context,int)(*models.BodyType, error)
	GetList(context.Context) ([]*models.BodyType, error)
}
type ColorRepoInterface interface {
	Create(context.Context, *models.Color) (int, error)
	GetByID(context.Context,int)(*models.Color, error)
	GetList(context.Context) ([]*models.Color, error)
}
//temp
type UserRepoInterface interface {
	GetUserID(context.Context, int) (int, error)
	GetByID(context.Context, int) (models.Machine, error)
}


type Repositories struct {
	UserRepoInterface
	ColorRepoInterface
	BodyTypeRepoInterface
	DriveUnitRepoInterface
	FuelRepoInterface
	ModelRepoInterface
	BrandRepoInterface
	MachineRepoInterface
	SearchRepoInterface
	FilterRepoInterface
	CategoryRepoInterface
	TransmissionRepoInterface
	CountryRepoInterface
	CityRepoInterface
	StateRepoInterface
}
