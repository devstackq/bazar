package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/models"
)

type FilterRepository struct {
	db *sql.DB
}

func FilterRepoInit(db *sql.DB) *FilterRepository {
	return &FilterRepository{
		db: db,
	}
}

//sort & filter ? good practice

func (fr FilterRepository) GetListMachineByFilter(ctx context.Context, keys map[string]string) ([]*models.Machine, error) {

	var result []*models.Machine
	var err error

	query := `SELECT
		machine_id,
		vin,
		title,
		description, 
		year,
		price,
		odometer,
		created_at,
		horse_power 
	FROM bazar_machine WHERE `

	query += prepareQuery(keys)
	// query += "ORDER BY created_at ASC"

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
			&temp.HorsePower,
		); err != nil {
			return nil, err
		}
		result = append(result, &temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}
