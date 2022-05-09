package psql

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type FilterRepository struct {
	db     *sql.DB
	bridge gallery.BridgeRepoInterface
}

func FilterRepoInit(db *sql.DB, bridge gallery.BridgeRepoInterface) gallery.FilterRepoInterface {
	return &FilterRepository{
		db:     db,
		bridge: bridge,
	}
}

// sort & filter ? good practice
func (fr FilterRepository) GetListMachineByFilter(ctx context.Context, keys *models.QueryParams, pageNum int) ([]*models.Machine, error) {
	var result []*models.Machine
	var err error
	// limit := 9

	query := `SELECT
		machine_id
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
		); err != nil {
			return nil, err
		}
		id, err := strconv.Atoi(temp.ID)
		if err != nil {
			return nil, err
		}

		car, err := fr.bridge.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		result = append(result, car)
	}
	if rows.Err() != nil {
		return nil, err
	}
	log.Println(len(result), "len res", pageNum)

	return result, nil
}
