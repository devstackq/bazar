package psql

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/authorization"
)

type AuthorizationRepository struct {
	db *sql.DB
}

func AuthRepositoryInit(db *sql.DB) authorization.AuthRepositoryInterface {
	return &AuthorizationRepository{
		db: db,
	}
}

type JwtTokenRepository struct {
	db *sql.DB
}

func JwtTokenRepositoryInit(db *sql.DB) authorization.JwtTokenRepositoryInterface {
	return &JwtTokenRepository{
		db: db,
	}
}
