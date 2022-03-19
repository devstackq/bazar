package psql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/devstackq/bazar/internal/models"
)

type AuthorizationRepository struct {
	db *sql.DB
}
func AuthRepositoryInit(db *sql.DB) *AuthorizationRepository {
	return &AuthorizationRepository{
		db: db,
	}
}

func (ur AuthorizationRepository) CreateUser(ctx context.Context, user *models.User) (id int, err error) {
	query := `INSERT INTO bazar_user (email, username, password, phone, first_name, last_name, created_at, country_id, city_id, role_id)values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`
	row := ur.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password, user.Phone, user.FirstName, user.LaststName, time.Now(),  user.CountryID, user.CityID, user.RoleID)
	err = row.Scan(&id)
	if err != nil {
		return 0,err
	}
	return id, nil
}

func (ur AuthorizationRepository) GetUser(ctx context.Context, username, password string) ( lastID int, err  error) {
	// var user models.User
	sqlQuery := `SELECT user_id FROM bazar_user WHERE username = $1 AND password = $2`
	row := ur.db.QueryRowContext(ctx, sqlQuery, username, password)
	err = row.Scan(&lastID)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}


// func (ur AuthorizationRepository)RemoveToken(ctx context.Context , access string,  refresh string, userID  int) error {

func (ur AuthorizationRepository)CreateSession(ctx context.Context , access string,  refresh string, userID  int) error {
	query := `INSERT INTO bazar_session(access_token, refresh_token, user_id)values($1, $2, $3)`
	log.Print(refresh, "refr")
	row := ur.db.QueryRowContext(ctx, query, access, refresh, userID)
	if row.Err() != nil {
		return row.Err()
	}
	return  nil
}