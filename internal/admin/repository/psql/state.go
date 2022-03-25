package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type StateRepository struct {
	db *sql.DB
}

func StateRepoInit(db *sql.DB) admin.StateRepoInterface {
	return &StateRepository{
		db: db,
	}
}

func (cr StateRepository) GetByID(ctx context.Context, id int) (*models.State, error) {
	var result models.State
	var err error

	query := `SELECT name FROM bazar_state WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr StateRepository) Create(ctx context.Context, cat *models.State) (int, error) {
	sqlQuery := `INSERT INTO bazar_state(name) VALUES($1) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr StateRepository) GetList(ctx context.Context) ([]*models.State, error) {
	query := `SELECT id, name FROM bazar_state`
	result := []*models.State{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.State{}
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
