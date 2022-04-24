package psql

import (
	"database/sql"
	"log"

	"github.com/devstackq/bazar/internal/config"
)

func InitDb(cfg config.Config) (*sql.DB, error) {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=require",
	// 	cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)

	psqlInfo := "postgres://jdukmdmdikvyup:92b83e192256d2b9c4d4173dfb66bedf264c78bd4a602a8972d987ac501f2cc8@ec2-99-80-170-190.eu-west-1.compute.amazonaws.com:5432/d448svjcam10be"
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
