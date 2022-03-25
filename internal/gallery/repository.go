package gallery

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type MachineRepoInterface interface {
	Create(context.Context, *models.Machine) (int, error)
	GetByID(context.Context, int) (*models.Machine, error)
	GetList(context.Context) ([]*models.Machine, error)
	GetListByUserID(context.Context, int) ([]*models.Machine, error)
}

type SearchRepoInterface interface {
	Search(context.Context, string) ([]*models.Machine, error)
}

type FilterRepoInterface interface {
	GetListMachineByFilter(context.Context, map[string]string) ([]*models.Machine, error)
}

type Repositories struct {
	MachineRepoInterface
	SearchRepoInterface
	FilterRepoInterface
}
