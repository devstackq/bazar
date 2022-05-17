package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/lib/pq"
)

type FileManagerRepository struct {
	db *sql.DB
}

func FileManagerRepoInit(db *sql.DB) gallery.FileManagerRepoInterface {
	return &FileManagerRepository{
		db: db,
	}
}

func (ur FileManagerRepository) CreateSrc(ctx context.Context, listSrc []string, machineID int) error {

	if machineID > 0 && len(listSrc) > 0 {
		sqlQuery := "INSERT INTO bazar_machine_image(machine_id, paths_img) VALUES($1, $2)"
		_, err := ur.db.ExecContext(ctx, sqlQuery, machineID, pq.Array(listSrc))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ur FileManagerRepository) GetListSrc(ctx context.Context, machineID int) ([]string, error) {
	query := `SELECT path FROM bazar_machine_image WHERE machine_id = $1`

	result := []string{}

	rows, err := ur.db.QueryContext(ctx, query, machineID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := ""
		if err = rows.Scan(
			&temp,
		); err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}
