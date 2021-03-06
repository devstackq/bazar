package psql

import (
	"context"
	"time"

	"github.com/devstackq/bazar/internal/models"
)

func (ur AuthorizationRepository) CreateUser(ctx context.Context, user *models.User) (res models.User, err error) {
	user.Role.ID = 2 // set default saler
	query := `INSERT INTO bazar_user ( email, username, password, phone, first_name, last_name, company, created_at, country_id, city_id, role_id)values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING user_id`
	row := ur.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password, user.Phone, user.FirstName, user.LaststName, user.Company, time.Now(), user.Country.ID, user.Country.City.ID, user.Role.ID)
	err = row.Scan(&res.ID)
	res.Email = user.Email
	res.Username = user.Username

	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
