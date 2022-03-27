package gallery

import "github.com/devstackq/bazar/internal/models"

type MachineUseCaseInterface interface {
	Create(*models.Machine) (int, error)
	GetMachineByID(int) (*models.Machine, error)
	GetRelevantMachines(int) ([]*models.Machine, error)
	GetListMachineByUserID(int) ([]*models.Machine, error)
}

type FilterUseCaseInterface interface {
	GetListMachineByFilter(map[string]string) ([]*models.Machine, error)
}
type FileManagerUseCaseInterface interface {
	CreateSrc([]string, int) error
	GetListSrc(int) ([]string, error)
}

type SortUseCaseInterface interface {
	SortByType(key string, filter map[string]string) ([]*models.Machine, error) // default, getFiltered data/ asc/desc
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
