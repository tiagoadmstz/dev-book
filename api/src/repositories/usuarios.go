package repositories

import (
	"api/src/config"
	"api/src/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (repository Users) InsertUser(user *models.User) (string, error) {
	defer repository.db.Disconnect(context.TODO())
	user.ID = primitive.NewObjectID()
	user.Created = time.Now()
	result, err := repository.getCollection().InsertOne(context.TODO(), user)
	if err != nil {
		return "0", err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	fmt.Printf("%v", user.ID)
	return fmt.Sprintf("%v", user.ID), err
}

// Update a user in database
func (repository Users) UpdateUser(user *models.User) (string, error) {
	/*defer repository.db.Disconnect(context.TODO())
	update := bson.D{
		{"$set", bson.D{{"nick", user.Nick}}},
	}
	_, err := repository.getCollection().UpdateByID(context.TODO(), user.ID, update)
	if err != nil {
		return "0", err
	}*/
	return "0", nil
}

// Delete a user in database
func (repository Users) DeleteUser(user *models.User) (bool, error) {
	defer repository.db.Disconnect(context.TODO())
	_, err := repository.getCollection().DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return false, err
	}
	return true, nil
}

// FindAllUsers return a user slice
func (repository Users) FindAllUsers() ([]*models.User, error) {
	defer repository.db.Disconnect(context.TODO())
	var users []*models.User
	result, err := repository.getCollection().Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	return users, result.Decode(&users)
}
