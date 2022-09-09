package models

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents an user
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Nick     string             `json:"nick,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Created  time.Time          `json:"created,omitempty"`
}

// Prepare will call validate and format methods
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name is required, it cannot be empty")
	}
	if user.Nick == "" {
		return errors.New("The nick is required, it cannot be empty")
	}
	if user.Email == "" {
		return errors.New("The email is required, it cannot be empty")
	}
	if user.Password == "" {
		return errors.New("The password is required, it cannot be empty")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
