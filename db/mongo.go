package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Config
}

func (m *MongoDb) InitDb() (interface{}, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.Config.host + m.Config.port))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	// m.db = client.Database(m.Config.name)
	// m.Config.tableName
	return client.Database("tablename"), nil
}
