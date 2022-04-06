package psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/devstackq/bazar/internal/gallery"
)

type FileManagerRepository struct {
	db *sql.DB
}

func FileManagerRepoInit(db *sql.DB) gallery.FileManagerRepoInterface {
	return &FileManagerRepository{
		db: db,
	}
}

func BulkInsert(rows []string, id int) (string, []interface{}, error) {

	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*2)
	i := 0

	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, post)
		valueArgs = append(valueArgs, id)
		i++
	}
	sqlQuery := fmt.Sprintf("INSERT INTO bazar_machine_image (path, machine_id) VALUES %s", strings.Join(valueStrings, ","))

	return sqlQuery, valueArgs, nil
}

func (ur FileManagerRepository) CreateSrc(ctx context.Context, listSrc []string, machineID int) error {

	query, args, err := BulkInsert(listSrc, machineID)
	if err != nil {
		return err
	}
	log.Println(query, args)
	_, err = ur.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
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
