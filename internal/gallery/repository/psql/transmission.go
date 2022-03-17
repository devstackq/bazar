package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery/models"
)


type TransRepository struct {
	db *sql.DB
}

func TransRepoInit(db *sql.DB) *TransRepository {
	return &TransRepository{
		db: db,
	}
}

func (tr TransRepository) GetByID(ctx context.Context, id int) (*models.Transmission,  error) {
	
	var result models.Transmission
	var err error

	query := `SELECT name FROM bazar_trans WHERE id = $1`
	err = tr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (tr TransRepository) Create(ctx context.Context, t *models.Transmission) (int,  error) {
	sqlQuery := `INSERT INTO bazar_trans(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := tr.db.QueryRowContext(ctx, sqlQuery, t.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (tr TransRepository) GetList(ctx context.Context) ([]*models.Transmission,  error) {

	query := `SELECT id, name FROM bazar_trans`
	
	result := []*models.Transmission{}
	
	rows, err := tr.db.QueryContext(ctx, query)
	
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		temp := models.Transmission{}
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