package repository

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/repository/psql"
)

// each use case - have all access - to repo layer
func MachineReposInit(db *sql.DB) gallery.Repositories {
	return gallery.Repositories{
		MachineRepoInterface:     psql.MachineRepoInit(db),
		SearchRepoInterface:      psql.SearchRepoInit(db),
		FilterRepoInterface:      psql.FilterRepoInit(db),
		FileManagerRepoInterface: psql.FileManagerRepoInit(db),
	}
}
