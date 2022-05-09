package psql

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (ur AuthorizationRepository) GetUser(ctx context.Context, username, password string) (user models.User, err error) {
	sqlQuery := `SELECT user_id, email FROM bazar_user WHERE username = $1 AND password = $2`
	row := ur.db.QueryRowContext(ctx, sqlQuery, username, password)
	user.Username = username
	err = row.Scan(&user.ID, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
