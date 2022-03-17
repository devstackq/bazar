package repository

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/repository/psql"
)


func MachineReposInit(db *sql.DB) gallery.Repositories {
	return  gallery.Repositories {
		MachineRepo: psql.MachineRepoInit(db),
		CategoryRepo: psql.CategoryRepoInit(db),
		SearchRepo: psql.SearchRepoInit(db),
		TransmissionRepo: psql.TransRepoInit(db),
	}
}