package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)


type DriveUnitRepository struct {
	db *sql.DB
}

func DriveUnitRepoInit(db *sql.DB) gallery.DriveUnitRepoInterface {
	return &DriveUnitRepository{
		db: db,
	}
}

func (cr DriveUnitRepository) GetByID(ctx context.Context, id int) (*models.DriveUnit,  error) {
	
	var result models.DriveUnit
	var err error

	query := `SELECT name FROM bazar_drive_unit WHERE id = $1`
	err = cr.db.QueryRowContext(ctx, query, id).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr DriveUnitRepository) Create(ctx context.Context, cat *models.DriveUnit) (int,  error) {
	sqlQuery := `INSERT INTO bazar_drive_unit(name) VALUES($1) RETURNING id`
		var id int
		var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, cat.Name)
		err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cr DriveUnitRepository) GetList(ctx context.Context) ([]*models.DriveUnit,  error) {

query := `SELECT id, name FROM bazar_drive_unit`
result := []*models.DriveUnit{}
rows, err := cr.db.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

for rows.Next() {
	temp := models.DriveUnit{}
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