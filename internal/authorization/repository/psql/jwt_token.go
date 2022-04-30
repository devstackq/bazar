package psql

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (ur JwtTokenRepository) CreateSession(ctx context.Context, token *models.TokenDetails) error {
	query := `INSERT INTO bazar_session(access_uuid, refresh_uuid, user_id)values($1, $2, $3)`
	row := ur.db.QueryRowContext(ctx, query, token.AccessUuid, token.RefreshUuid, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (ur JwtTokenRepository) UpdateSession(ctx context.Context, token *models.TokenDetails) error {
	query := `UPDATE bazar_session SET access_uuid=$1, refresh_uuid=$2  WHERE user_ID=$3`
	row := ur.db.QueryRowContext(ctx, query, token.AccessUuid, token.RefreshUuid, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (ur JwtTokenRepository) DeleteSession(ctx context.Context, token *models.TokenDetails) error {
	query := `DELETE bazar_session WHERE user_ID=$1`
	row := ur.db.QueryRowContext(ctx, query, token.UserID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
