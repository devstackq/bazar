package psql

import (
	"database/sql"
	"log"

	"github.com/devstackq/bazar/internal/config"
)

func InitDb(cfg config.Config) (*sql.DB, error) {
	log.Println(cfg.DB, "db config")

	// psqlInfo := "host=postgresdb dbname=testdb user=postgres password=postgres port=5432"

	psqlInfo := "postgres://rwecsnywflfryx:04d64cea97a6d17d7b4de6b5e9b0755e2ab090bbdc24b1be7e11943d618c777c@ec2-44-199-143-43.compute-1.amazonaws.com:5432/dffqhfqt0ef62l"

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// db, err := sql.Open("postgres", "postgres://"+cfg.DB.Username+":"+cfg.DB.Password+"@"+cfg.DB.Host+"/"+cfg.DB.DBName+"?sslmode=disable")

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// dial tcp: lookup postgresdb: Temporary failure in todo
