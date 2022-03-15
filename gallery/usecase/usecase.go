package usecase

//Gallery -> search; filter; list cars

//all cars; || cars -> paymanet ?
//default last added - cars show; with pagination
//filter - by category; new/bu/avarinyi, model, etc

//optional field ?
type GalleryCars struct {
	filterRepo repository.FilterCars
	searchRepo repository.SearchCars
	listRepo repository.ListCars
}

func GalleryCarsUseCaseInit(filterRepo repository.FilterCars, searchRepo repository.SearchCars, listRepo repository.ListCars)usecase.Galerry {
	return &GalleryCars{
		filterRepo: filterRepo,
		searchRepo: searchRepo,
		listRepo : listRepo,
	}
}

// func (uc GalleryCarsUseCase)GetListCarsFilterBy(ctx context.Context, )

