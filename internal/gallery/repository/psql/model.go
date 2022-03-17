package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type ModelRepository struct {
	db *sql.DB
}

func ModelRepoInit(db *sql.DB) gallery.ModelRepoInterface {
	return &ModelRepository{
		db: db,
	}
}

func (cr ModelRepository) GetByID(ctx context.Context, id int) (*models.Model,  error) {
	
	var result models.Model
	var err error

	query := `SELECT name FROM bazar_model WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr ModelRepository) Create(ctx context.Context, cat *models.Model) (int,  error) {
	sqlQuery := `INSERT INTO bazar_model(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr ModelRepository) GetList(ctx context.Context) ([]*models.Model,  error) {

query := `SELECT id, name FROM bazar_model`
result := []*models.Model{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.Model{}
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