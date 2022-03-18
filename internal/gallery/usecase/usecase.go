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
		MachineUseCaseInterface: MachineUseCaseInit(
			r.MachineRepoInterface, 
			r.CategoryRepoInterface, 
			r.TransmissionRepoInterface, 
			r.CountryRepoInterface,
			r.CityRepoInterface,
			r.StateRepoInterface,
			r.BrandRepoInterface,
			r.ModelRepoInterface,
			r.FuelRepoInterface,
			r.DriveUnitRepoInterface,
			r.BodyTypeRepoInterface,
			r.ColorRepoInterface,
		),
		
		CategoryUseCaseInterface: CategoryUseCaseInit(r.CategoryRepoInterface),
		TransUseCaseInterface: TransUseCaseInit(r.TransmissionRepoInterface),
		SearchUseCaseInterface: SearchUseCaseInit(r.SearchRepoInterface),
		CountryUseCaseInterface: CountryUseCaseInit(r.CountryRepoInterface),
		CityUseCaseInterface: CityUseCaseInit(r.CityRepoInterface),
		StateUseCaseInterface: StateUseCaseInit(r.StateRepoInterface),
		BrandUseCaseInterface: BrandUseCaseInit(r.BrandRepoInterface),
		ModelUseCaseInterface: ModelUseCaseInit(r.ModelRepoInterface),
		FuelUseCaseInterface: FuelUseCaseInit(r.FuelRepoInterface),
		DriveUnitUseCaseInterface: DriveUnitUseCaseInit(r.DriveUnitRepoInterface),
		BodyTypeUseCaseInterface: BodyTypeUseCaseInit(r.BodyTypeRepoInterface),
		ColorUseCaseInterface: ColorUseCaseInit(r.ColorRepoInterface),
		FilterUseCaseInterface: FilterUseCaseInit(r.FilterRepoInterface),
		// FilterMachineInterface: FilterMachineUseCaseInit(r.FilterMachineRepo),
	}
}