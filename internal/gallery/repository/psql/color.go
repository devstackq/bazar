package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type ColorRepository struct {
	db *sql.DB
}

func ColorRepoInit(db *sql.DB) gallery.ColorRepoInterface {
	return &ColorRepository{
		db: db,
	}
}

func (cr ColorRepository) GetByID(ctx context.Context, id int) (*models.Color,  error) {
	
	var result models.Color
	var err error

	query := `SELECT name FROM bazar_color WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr ColorRepository) Create(ctx context.Context, cat *models.Color) (int,  error) {
	sqlQuery := `INSERT INTO bazar_color(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr ColorRepository) GetList(ctx context.Context) ([]*models.Color,  error) {

query := `SELECT id, name FROM bazar_color`
result := []*models.Color{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.Color{}
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