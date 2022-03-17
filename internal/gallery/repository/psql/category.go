package psql

import (
	"context"
	"database/sql"
	"log"

	"github.com/devstackq/bazar/internal/gallery/models"
)


type CategoryRepository struct {
	db *sql.DB
}

func CategoryRepoInit(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (mr CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*models.Category,  error) {
	
	var result models.Category
	var err error

	query := `SELECT name FROM bazar_category WHERE id = $1`
	err = mr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	log.Print(err,1)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (mr CategoryRepository) CreateCategory(ctx context.Context, cat *models.Category) (int,  error) {
	sqlQuery := `INSERT INTO bazar_category(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := mr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}