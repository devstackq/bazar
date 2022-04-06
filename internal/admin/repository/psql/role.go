package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type RoleRepository struct {
	db *sql.DB
}

func RoleRepoInit(db *sql.DB) admin.RoleRepoInterface {
	return &RoleRepository{
		db: db,
	}
}

func (cr RoleRepository) GetList(ctx context.Context) ([]*models.Role, error) {

	query := `SELECT id, name FROM bazar_roles`
	result := []*models.Role{}
	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.Role{}
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

func (cr RoleRepository) Create(ctx context.Context, role *models.Role) (int, error) {

	sqlQuery := `INSERT INTO bazar_roles(name) VALUES($1) RETURNING id`
	var id int
	var err error

	row := cr.db.QueryRowContext(ctx, sqlQuery, role.Name)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
