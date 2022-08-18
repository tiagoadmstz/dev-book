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

func test() {

	//ben := Person{"Ben", 16}

	/*_, err = collection.InsertOne(context.TODO(), john)
	CheckError(err)

	persons := []interface{}{jane, ben}
	_, err = collection.InsertMany(context.TODO(), persons)
	CheckError(err)

	// update
	filter := bson.D{{"name", "John"}}

	update := bson.D{
		{"$set", bson.D{{"age", 26}}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	CheckError(err)

	// find
	var res Person
	err = collection.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	// delete
	_, err = collection.DeleteMany(context.TODO(), bson.D{{}})
	CheckError(err)*/
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
