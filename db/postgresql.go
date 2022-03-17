package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"

	"log"
)

//set params outside
// func NewPostgresStorage(name, password, url, port, dbName string) *PostgreSql {}
// func (p *PostgreSql) InitPostgresDb() (*sql.DB, error) {}

type PostgreSqlDb struct {
	Config
}

func (p *PostgreSqlDb) InitDb() (interface{}, error) {
	log.Print(p.Config)
	db, err := sql.Open("postgres", "postgres://"+p.Config.user+":"+p.Config.password+"@"+p.Config.host+"/"+p.Config.dbName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// if err = createTables(db, p.Config.tableName); err != nil {
	// 	return nil, err
	// }
	return db, nil
}

