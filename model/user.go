package model

import (
	"hcc/piccolo/lib/errors"
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
	Errors    string    `json:"errors"`
}

// Token : Contain the user token
type Token struct {
	Token  string               `json:"token"`
	Errors []errors.HccError `json:"errors"`
}

// IsValid : Contain the validation of the token
type IsValid struct {
	IsValid bool                 `json:"isvalid"`
	Errors []errors.HccError `json:"errors"`
}
