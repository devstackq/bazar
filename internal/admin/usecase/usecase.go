package usecase

import "github.com/devstackq/bazar/internal/admin"

func UseCasesInit(r admin.Repositories) admin.UseCases {
	return admin.UseCases{
		ColorUseCaseInterface:     ColorUseCaseInit(r.ColorRepoInterface),
		BodyTypeUseCaseInterface:  BodyTypeUseCaseInit(r.BodyTypeRepoInterface),
		DriveUnitUseCaseInterface: DriveUnitUseCaseInit(r.DriveUnitRepoInterface),
		FuelUseCaseInterface:      FuelUseCaseInit(r.FuelRepoInterface),
		ModelUseCaseInterface:     ModelUseCaseInit(r.ModelRepoInterface),
		BrandUseCaseInterface:     BrandUseCaseInit(r.BrandRepoInterface),
		StateUseCaseInterface:     StateUseCaseInit(r.StateRepoInterface),
		CategoryUseCaseInterface:  CategoryUseCaseInit(r.CategoryRepoInterface),
		TransUseCaseInterface:     TransUseCaseInit(r.TransmissionRepoInterface),
		CountryUseCaseInterface:   CountryUseCaseInit(r.CountryRepoInterface),
		CityUseCaseInterface:      CityUseCaseInit(r.CityRepoInterface),
		RoleUseCaseInterface:      RoleUseCaseInit(r.RoleRepoInterface),
	}
}
