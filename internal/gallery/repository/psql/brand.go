package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type BrandRepository struct {
	db *sql.DB
}

func BrandRepoInit(db *sql.DB) gallery.BrandRepoInterface {
	return &BrandRepository{
		db: db,
	}
}

func (cr BrandRepository) GetByID(ctx context.Context, id int) (*models.Brand,  error) {
	
	var result models.Brand
	var err error

	query := `SELECT name FROM bazar_brand WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr BrandRepository) Create(ctx context.Context, cat *models.Brand) (int,  error) {
	sqlQuery := `INSERT INTO bazar_brand(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr BrandRepository) GetList(ctx context.Context) ([]*models.Brand,  error) {

query := `SELECT id, name FROM bazar_brand`
result := []*models.Brand{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.Brand{}
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