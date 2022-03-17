package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type BodyTypeRepository struct {
	db *sql.DB
}

func BodyTypeRepoInit(db *sql.DB) gallery.BodyTypeRepoInterface {
	return &BodyTypeRepository{
		db: db,
	}
}

func (cr BodyTypeRepository) GetByID(ctx context.Context, id int) (*models.BodyType,  error) {
	
	var result models.BodyType
	var err error

	query := `SELECT name FROM bazar_body_type WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr BodyTypeRepository) Create(ctx context.Context, cat *models.BodyType) (int,  error) {
	sqlQuery := `INSERT INTO bazar_body_type(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr BodyTypeRepository) GetList(ctx context.Context) ([]*models.BodyType,  error) {

query := `SELECT id, name FROM bazar_body_type`
result := []*models.BodyType{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.BodyType{}
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