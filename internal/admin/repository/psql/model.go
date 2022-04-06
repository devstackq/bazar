package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type ModelRepository struct {
	db *sql.DB
}

func ModelRepoInit(db *sql.DB) admin.ModelRepoInterface {
	return &ModelRepository{
		db: db,
	}
}

func (cr ModelRepository) GetByID(ctx context.Context, id int) (*models.Model, error) {
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

func (cr ModelRepository) Create(ctx context.Context, model *models.Model) (int, error) {
	sqlQuery := `INSERT INTO bazar_model(name, brand_id) VALUES($1, $2) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, model.Name, model.ID)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr ModelRepository) GetList(ctx context.Context, brandID int) ([]*models.Model, error) {
	query := `SELECT id, name FROM bazar_model where brand_id=$1`
	result := []*models.Model{}
	rows, err := cr.db.QueryContext(ctx, query, brandID)
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
