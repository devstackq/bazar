package psql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/devstackq/bazar/internal/config"
)

func InitDb(cfg config.Config) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(db, err, "ping err")

		return nil, err
	}
	return db, nil
}

// dial tcp: lookup postgresdb: Temporary failure in todo
