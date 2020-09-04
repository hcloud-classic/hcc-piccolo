package model

import (
	"time"
)

// User : Contain infos of a user
type User struct {
	UUID      string    `json:"uuid"`
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	LoginAt   time.Time `json:"login_at"`
}

// Token : Contain the user token
type Token struct {
	Token string `json:"token"`
}

// IsValid : Contain the validation of the token
type IsValid struct {
	IsValid bool `json:"isvalid"`
}
