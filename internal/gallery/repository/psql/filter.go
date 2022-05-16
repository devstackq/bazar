package psql

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"sync"

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

func (fr FilterRepository) GetCountMachines(ctx context.Context) (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM bazar_machine `

	rows, err := fr.db.QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err = rows.Scan(
			&count,
		); err != nil {
			return 0, err
		}
	}
	return count, nil
}

// sort & filter ? good practice
func (fr FilterRepository) GetListMachineByFilter(ctx context.Context, keys *models.QueryParams, pageNum int) ([]*models.Machine, error) {
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

	var wg sync.WaitGroup
	var seqIds []string

	for rows.Next() {
		temp := models.Machine{}

		if err = rows.Scan(
			&temp.ID,
		); err != nil {
			return nil, err
		}
		seqIds = append(seqIds, temp.ID)
		if rows.Err() != nil {
			return nil, err
		}
	}
	var idMachine int
	result := make([]*models.Machine, 0, len(seqIds))

	for _, id := range seqIds {
		wg.Add(1)

		go func(wg *sync.WaitGroup, id string, ctx context.Context) {
			idMachine, err = strconv.Atoi(id)
			if err != nil {
				log.Println(err)
			}
			car, err := fr.bridge.GetByID(ctx, idMachine)
			if err != nil {
				log.Println(err)
			}
			result = append(result, car)
			defer wg.Done()
		}(&wg, id, ctx)
	}
	wg.Wait()
	return result, nil
}
