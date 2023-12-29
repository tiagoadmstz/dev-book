package repositories

import (
	"api/src/config"
	"api/src/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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
	filter := bson.D{{"_id", user.ID}}

	update := bson.D{
		{"$set", bson.D{{"nick", user.Nick}}},
	}

	result, err := repository.getCollection().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "0", err
	}

	if result.ModifiedCount == 0 {
		return "0", errors.New("no user updated")
	}

	return "1", nil
}

// Delete a user in database
func (repository Users) DeleteUser(r *http.Request) (bool, error) {
	defer repository.db.Disconnect(context.TODO())

	userId := r.URL.Query().Get("id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return false, err
	}

	filter := bson.D{{"_id", objectId}}
	result, err := repository.getCollection().DeleteOne(context.TODO(), filter)

	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, err
	}

	return true, err
}

func (repository Users) FindUserById(r *http.Request) (*models.User, error) {
	defer repository.db.Disconnect(context.TODO())

	parameters := mux.Vars(r)
	objectId, err := primitive.ObjectIDFromHex(parameters["id"])

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objectId}}
	var user models.User

	err = repository.getCollection().FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			//return nil, responses.NewHTTPError(http.StatusNotFound, "User not found")
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

// FindAllUsers return a users slice
func (repository Users) FindAllUsers(w http.ResponseWriter, r *http.Request) ([]*models.User, error) {
	defer repository.db.Disconnect(context.TODO())

	nameOrNick := r.URL.Query().Get("user")
	var users []*models.User
	var cursor *mongo.Cursor
	var err error

	if nameOrNick != "" {
		regex := primitive.Regex{Pattern: nameOrNick, Options: "i"}
		filter := bson.M{"$or": []bson.M{
			{"name": bson.M{"$regex": regex}},
			{"nick": bson.M{"$regex": regex}},
		}}
		cursor, err = repository.getCollection().Find(context.TODO(), filter)
	} else {
		cursor, err = repository.getCollection().Find(context.TODO(), bson.D{{}})
	}

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
