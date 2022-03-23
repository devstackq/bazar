package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

type PostgreSqlDb struct {
	Config
}

func (p *PostgreSqlDb) InitDb() (interface{}, error) {
	log.Println(p.Config, "config")
	db, err := sql.Open("postgres", "postgres://"+p.Config.user+":"+p.Config.password+"@"+p.Config.host+"/"+p.Config.dbName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
