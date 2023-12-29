package controllers

import (
	"api/src/models"
	"api/src/responses"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Create new user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if _, err := services.InsertUser(w, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// Find all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.FindAllUsers(w, r)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusFound, users)
}

// Find user by id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	user, err := services.FindUserById(r)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusFound, user)
}

// Updating an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if _, err := services.UpdateUser(&user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusAccepted, user)
}

// Deleting an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	services.DeleteUser(r)
	responses.JSON(w, http.StatusNoContent, nil)
}
