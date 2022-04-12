package db

// ref send  - db layer
type StorageConnecter interface {
	InitDb() (interface{}, error)
}

type Config struct {
	user     string
	password string
	host     string
	port     string
	// tableName string
	dbName string
}

func NewDbObject(typeDb, user, password, host, port, dbName string) StorageConnecter {
	if typeDb == "mongodb" {
		return &MongoDb{
			Config{user, password, host, port, dbName},
		}
	} else if typeDb == "postgresql" {
		return &PostgreSqlDb{
			Config{user, password, host, port, dbName},
		}
	}
	return nil
}
