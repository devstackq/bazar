package psql

import "context"

func (ur AuthorizationRepository) GetUser(ctx context.Context, username, password string) (lastID int, err error) {
	sqlQuery := `SELECT user_id FROM bazar_user WHERE username = $1 AND password = $2`
	row := ur.db.QueryRowContext(ctx, sqlQuery, username, password)
	err = row.Scan(&lastID)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}
