package repository

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/admin/repository/psql"
)

func AdminReposInit(db *sql.DB) admin.Repositories {
	return admin.Repositories{
		CategoryRepoInterface:     psql.CategoryRepoInit(db),
		TransmissionRepoInterface: psql.TransRepoInit(db),
		CountryRepoInterface:      psql.CountryRepoInit(db),
		CityRepoInterface:         psql.CityRepoInit(db),
		StateRepoInterface:        psql.StateRepoInit(db),
		BrandRepoInterface:        psql.BrandRepoInit(db),
		ModelRepoInterface:        psql.ModelRepoInit(db),
		FuelRepoInterface:         psql.FuelRepoInit(db),
		DriveUnitRepoInterface:    psql.DriveUnitRepoInit(db),
		BodyTypeRepoInterface:     psql.BodyTypeRepoInit(db),
		ColorRepoInterface:        psql.ColorRepoInit(db),
		RoleRepoInterface:         psql.RoleRepoInit(db),
	}
}
