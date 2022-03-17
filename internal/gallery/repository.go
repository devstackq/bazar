package gallery

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery/models"
)

type MachineRepo interface {
	Create(context.Context, *models.Machine)(int, error)
	GetMachineByID(context.Context, int)(*models.Machine, error)
	GetRelevantMachines(context.Context)([]*models.Machine,  error)
}
type CategoryRepo interface {
	GetCategoryByID(context.Context, int) (*models.Category, error)
	CreateCategory(context.Context, *models.Category)(int, error)
}

type SearchRepo interface {
	Search(context.Context, string)([]*models.Machine, error)
}

type FilterRepo interface {
	GetListMachineByFilter(context.Context)
}

type TransmissionRepo interface {
	Create(context.Context,  *models.Transmission) (int, error) 
	GetTransByID(context.Context, int)(*models.Transmission, error)
}

type Repositories struct {
	MachineRepo
	SearchRepo
	FilterRepo
	CategoryRepo
	TransmissionRepo
}
