package repository

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/repository/psql"
)

//each use case - have all access - to repo layer
func MachineReposInit(db *sql.DB) gallery.Repositories {
	return  gallery.Repositories {
		MachineRepoInterface: psql.MachineRepoInit(db),
		CategoryRepoInterface: psql.CategoryRepoInit(db),
		SearchRepoInterface: psql.SearchRepoInit(db),
		TransmissionRepoInterface: psql.TransRepoInit(db),
		CountryRepoInterface: psql.CountryRepoInit(db),
		CityRepoInterface: psql.CityRepoInit(db),
		StateRepoInterface: psql.StateRepoInit(db),
		BrandRepoInterface: psql.BrandRepoInit(db),
		ModelRepoInterface: psql.ModelRepoInit(db),
		FuelRepoInterface: psql.FuelRepoInit(db),
		DriveUnitRepoInterface: psql.DriveUnitRepoInit(db),
		BodyTypeRepoInterface: psql.BodyTypeRepoInit(db),
		ColorRepoInterface: psql.ColorRepoInit(db),
		FilterRepoInterface: psql.FilterRepoInit(db),
		// UserRepoInterface :psql.UserRepoInit(db),
	}
}