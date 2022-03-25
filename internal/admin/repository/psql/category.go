package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func CategoryRepoInit(db *sql.DB) admin.CategoryRepoInterface {
	return &CategoryRepository{
		db: db,
	}
}

func (cr CategoryRepository) GetByID(ctx context.Context, id int) (*models.Category, error) {
	var result models.Category
	var err error

	query := `SELECT name FROM bazar_category WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr CategoryRepository) Create(ctx context.Context, cat *models.Category) (int, error) {
	sqlQuery := `INSERT INTO bazar_category(name) VALUES($1) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr CategoryRepository) GetList(ctx context.Context) ([]*models.Category, error) {
	query := `SELECT id, name FROM bazar_category`
	result := []*models.Category{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.Category{}
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
