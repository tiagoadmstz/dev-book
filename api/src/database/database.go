package database

import (
	"api/src/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Open database connection and return then
func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.StringDatabaseConnection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	CheckError(err)
	if err = client.Ping(context.TODO(), nil); err != nil {
		client.Disconnect(context.TODO())
		return nil, err
	}
	return client, err
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
