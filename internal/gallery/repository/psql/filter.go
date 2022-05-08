package psql

import (
	"context"
	"database/sql"
	"log"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type FilterRepository struct {
	db *sql.DB
}

func FilterRepoInit(db *sql.DB) gallery.FilterRepoInterface {
	return &FilterRepository{
		db: db,
	}
}

// sort & filter ? good practice
func (fr FilterRepository) GetListMachineByFilter(ctx context.Context, keys *models.QueryParams, pageNum int) ([]*models.Machine, error) {
	var result []*models.Machine
	var err error
	// limit := 9

	query := `SELECT
		machine_id,
		vin,
		title,
		description, 
		year,
		price,
		odometer,
		created_at
	FROM bazar_machine  `

	query += prepareQuery(keys)

	// query += ` LIMIT $1 OFFSET $2 `
	// pageNum = limit * (pageNum - 1)

	rows, err := fr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
			&temp.VIN,
			&temp.Title,
			&temp.Description,
			&temp.Year,
			&temp.Price,
			&temp.Odometer,
			&temp.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, &temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	log.Println(len(result), "len res", pageNum)

	return result, nil
}
