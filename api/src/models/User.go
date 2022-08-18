package models

import (
	"time"
)

// Represents an user
type User struct {
	ID       string    `json:"_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}
