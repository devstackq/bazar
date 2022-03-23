package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type FuelRepository struct {
	db *sql.DB
}

func FuelRepoInit(db *sql.DB) gallery.FuelRepoInterface {
	return &FuelRepository{
		db: db,
	}
}

func (cr FuelRepository) GetByID(ctx context.Context, id int) (*models.Fuel, error) {
	var result models.Fuel
	var err error

	query := `SELECT name FROM bazar_fuel WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr FuelRepository) Create(ctx context.Context, cat *models.Fuel) (int, error) {
	sqlQuery := `INSERT INTO bazar_fuel(name) VALUES($1) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr FuelRepository) GetList(ctx context.Context) ([]*models.Fuel, error) {
	query := `SELECT id, name FROM bazar_fuel`
	result := []*models.Fuel{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.Fuel{}
		if err = rows.Scan(
			&temp.ID,
			&temp.Name,
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
