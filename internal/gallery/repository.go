package gallery

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type BridgeRepoInterface interface {
	GetByID(context.Context, int) (*models.Machine, error)
}

type MachineRepoInterface interface {
	Create(context.Context, *models.Machine) (int, error)
	GetByID(context.Context, int) (*models.Machine, error)
	GetList(context.Context, int) ([]*models.Machine, error)
	GetListByUserID(context.Context, float64, int) ([]*models.Machine, error)
}

type SearchRepoInterface interface {
	Search(context.Context, string, int) ([]*models.Machine, error)
}

type FilterRepoInterface interface {
	GetListMachineByFilter(context.Context, *models.QueryParams, int) ([]*models.Machine, error)
}

type FileManagerRepoInterface interface {
	CreateSrc(context.Context, []string, int) error
	GetListSrc(context.Context, int) ([]string, error)
}

type Repositories struct {
	MachineRepoInterface
	SearchRepoInterface
	FilterRepoInterface
	FileManagerRepoInterface
	BridgeRepoInterface
}
