package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type CityRepository struct {
	db *sql.DB
}

func CityRepoInit(db *sql.DB) gallery.CityRepoInterface {
	return &CityRepository{
		db: db,
	}
}

func (cr CityRepository) GetByID(ctx context.Context, id int) (*models.City, error) {
	var result models.City
	var err error

	query := `SELECT name FROM bazar_city WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr CityRepository) Create(ctx context.Context, city *models.City) (int, error) {
	sqlQuery := `INSERT INTO bazar_city(name, country_id) VALUES($1, $2) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, city.Name, city.ID)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr CityRepository) GetList(ctx context.Context, countryID int) ([]*models.City, error) {
	query := `SELECT id, name FROM bazar_city WHERE country_id = $1`
	result := []*models.City{}

	rows, err := cr.db.QueryContext(ctx, query, countryID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.City{}
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
