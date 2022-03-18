package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery/models"
)

type SearchRepository struct {
	db *sql.DB
}

func SearchRepoInit(db *sql.DB) *SearchRepository {
	return &SearchRepository{
		db: db,
	}
}

func (sr SearchRepository) Search(ctx context.Context, keyword string) ([]*models.Machine,  error) {
	
	var result []*models.Machine

	query := "SELECT machine_id, title, vin, description, year, price FROM bazar_machine WHERE title LIKE $1 OR description LIKE $1"

	rows, err := sr.db.Query(query, "%" + keyword + "%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.Machine{}
		err = rows.Scan(&temp.ID, &temp.Title, &temp.VIN,  &temp.Description,  &temp.Year, &temp.Price )
		if err != nil {
			return nil, err
		}
		result = append(result, &temp)
	}
	return result, nil
}