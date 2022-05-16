package gallery

import (
	"github.com/devstackq/bazar/internal/models"
)

type MachineUseCaseInterface interface {
	Create(*models.Machine) (int, error)
	GetMachineByID(int) (*models.Machine, error)
	GetRelevantMachines(int) ([]*models.Machine, error)
	GetListMachineByUserID(float64, int) ([]*models.Machine, error)
}

type FilterUseCaseInterface interface {
	GetListMachineByFilter(*models.QueryParams, int) ([]*models.Machine, error)
}
type FileManagerUseCaseInterface interface {
	CreateSrc([]string, int) error
	// GetListSrc(int) ([]string, error)
}

type SearchUseCaseInterface interface {
	SearchByKeyWord(string, int) ([]*models.Machine, error)
}

type UseCases struct {
	MachineUseCaseInterface
	FilterUseCaseInterface
	SearchUseCaseInterface
	FileManagerUseCaseInterface
}
