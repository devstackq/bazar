package gallery

type Galerry struct {
	ListCars
	FilterCars
	SearchCars
}
type ListCars interface {
	GetCarByID()
	GetRelevantCars()(error)
}

type FilterCars interface {
	GetListCarByFilter()
	// GetListCarBy() // 1 params?
}

type SearchCars interface {
	GetListCarsBySearch()
}
