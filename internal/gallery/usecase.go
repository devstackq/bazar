package gallery

import "github.com/devstackq/bazar/internal/models"

type MachineUseCaseInterface interface {
	Create(*models.Machine) (int, error)
	GetMachineByID(int) (*models.Machine, error)
	GetRelevantMachines() ([]*models.Machine, error)
	GetListMachineByUserID(int) ([]*models.Machine, error)
}

type FilterUseCaseInterface interface {
	GetListMachineByFilter(map[string]string) ([]*models.Machine, error)
	// GetListMachineByRangePrice(map[string]string)([]*models.Machine, error) // price=from=to
}

type SortUseCaseInterface interface {
	SortByType(key string, filter map[string]string) ([]*models.Machine, error) // default, getFiltered data/ asc/desc
	// priceAsc, dateAsc, yearAsc, kmsAsc, filter
}

type SearchUseCaseInterface interface {
	SearchByKeyWord(string) ([]*models.Machine, error)
}

type UseCases struct {
	MachineUseCaseInterface
	FilterUseCaseInterface
	SearchUseCaseInterface
}
