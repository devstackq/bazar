package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery/models"
)


type CountryRepository struct {
	db *sql.DB
}

func CountryRepoInit(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (cr CountryRepository) GetByID(ctx context.Context, id int) (*models.Country,  error) {
	
	var result models.Country
	var err error

	query := `SELECT name FROM bazar_country WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr CountryRepository) Create(ctx context.Context, cat *models.Country) (int,  error) {
	sqlQuery := `INSERT INTO bazar_country(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr CountryRepository) GetList(ctx context.Context) ([]*models.Country,  error) {

query := `SELECT id, name FROM bazar_country`
result := []*models.Country{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.Country{}
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
