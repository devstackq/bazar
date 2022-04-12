package psql

import (
	"context"
	"database/sql"

	"github.com/devstackq/bazar/internal/models"
	"github.com/devstackq/bazar/internal/profile"
)

type ProfileRepository struct {
	db *sql.DB
}

func ProfileRepositoryInit(db *sql.DB) profile.ProfileRepositoryInterface {
	return &ProfileRepository{
		db: db,
	}
}

// or use -> anothre package/service  ? countryService.GetByID()?
func (pr ProfileRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var result models.User
	var err error

	sqlQuery := `SELECT
	email, username,
	phone, first_name, 
	last_name,company,
	created_at, rl.name, ctr.name,
	ct.name
	FROM bazar_user AS u
	LEFT JOIN bazar_country AS ctr ON ctr.id = u.country_id
	LEFT JOIN bazar_city AS ct ON ct.id = u.city_id
	LEFT JOIN bazar_roles AS rl ON rl.id = u.role_id
	WHERE user_id = $1`

	row := pr.db.QueryRowContext(ctx, sqlQuery, id)
	err = row.Scan(&result.Email, &result.Username, &result.Phone, &result.FirstName, &result.LaststName, &result.Company, &result.CreatedAt, &result.Role.Name, &result.Country.Name, &result.Country.City.Name)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
