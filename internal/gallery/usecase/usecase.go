package usecase

import (
	"github.com/devstackq/bazar/internal/gallery"
)

//Gallery -> search; filter; list cars

//all cars; || cars -> paymanet ?
//default last added - cars show; with pagination
//filter - by category; new/bu/avarinyi, model, etc

func UseCasesInit(r  gallery.Repositories) gallery.UseCases {
	return gallery.UseCases {
		MachineUseCase: MachineUseCaseInit(r.MachineRepo, r.CategoryRepo, r.TransmissionRepo),
		CategoryUseCase: CategoryUseCaseInit(r.CategoryRepo),
		TransUseCase: TransUseCaseInit(r.TransmissionRepo),
		SearchUseCase: SearchUseCaseInit(r.SearchRepo),
		// FilterMachineInterface: FilterMachineUseCaseInit(r.FilterMachineRepo),
	}
}
