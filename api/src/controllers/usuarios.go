package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create new user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.CreateUserRepository(db)
	repository.Insert(&user)
}

// Find all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}

// Find user by id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user by id"))
}

// Updating an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

// Deleting an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
