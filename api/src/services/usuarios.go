package services

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"net/http"
)

func InsertUser(w http.ResponseWriter, user *models.User) (bool, error) {
	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return false, err
	}

	db, err := database.Connect()
	if err != nil {
		return false, err
	}

	repository := repositories.CreateUserRepository(db)
	repository.InsertUser(user)
	return true, nil
}

func UpdateUser(user *models.User) (string, error) {
	repository, err := getUserRepository()
	if err != nil {
		return "0", err
	}
	return repository.UpdateUser(user)
}

func DeleteUser(r *http.Request) (bool, error) {
	repository, err := getUserRepository()
	if err != nil {
		return false, err
	}
	return repository.DeleteUser(r)
}

func FindUserById(r *http.Request) (*models.User, error) {
	repository, err := getUserRepository()
	if err != nil {
		return nil, err
	}
	return repository.FindUserById(r)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) ([]*models.User, error) {
	repository, err := getUserRepository()
	if err != nil {
		return nil, err
	}
	return repository.FindAllUsers(w, r)
}

func getUserRepository() (*repositories.Users, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return repositories.CreateUserRepository(db), nil
}
