package repositories

import (
	"api/src/config"
	"api/src/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	db *mongo.Client
}

// Create a user repository
func CreateUserRepository(db *mongo.Client) *Users {
	return &Users{db}
}

// Return users collection
func (repository Users) getCollection() *mongo.Collection {
	return repository.db.Database(config.DatabaseName).Collection("users")
}

// Insert a user in database
func (repository Users) Insert(user *models.User) (string, error) {
	defer repository.db.Disconnect(context.TODO())
	user.Created = time.Now()
	result, err := repository.getCollection().InsertOne(context.TODO(), user)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}
