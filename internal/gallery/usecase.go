package gallery

import "github.com/devstackq/bazar/internal/gallery/models"

type UseCases struct {
	MachineUseCase
	FilterUseCase
	SearchUseCase
	CategoryUseCase
	TransUseCase

}
type CategoryUseCase interface {
	CreateCategory(*models.Category) (int, error)
	GetCategoryByID(int)(*models.Category, error)
}

type MachineUseCase interface {
	Create(*models.Machine)(int, error)
	GetMachineByID(int)(*models.Machine, error)
	GetRelevantMachines()([]*models.Machine,  error)
}

type FilterUseCase interface {
	GetListMachineByFilter()
	// GetListCarBy() // 1 params?
}

type SearchUseCase interface {
	SearchByKeyWord(string)([]*models.Machine, error)
}

type TransUseCase interface {
	CreateTransmission(t *models.Transmission) (int, error)
	GetTransByID(id int)(*models.Transmission, error)
}