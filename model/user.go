package model

import (
	"hcc/piccolo/lib/errors"
	"time"
)

// User : Contain infos of a user
type User struct {
	UUID      string            `json:"uuid"`
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	LoginAt   time.Time         `json:"login_at"`
	CreatedAt time.Time         `json:"created_at"`
	Errors    []errors.HccError `json:"errors"`
}

// UserList : Contain list of users
type UserList struct {
	Users  []User            `json:"user_list"`
	Errors []errors.HccError `json:"errors"`
}

// Token : Contain the user token
type Token struct {
	Token  string            `json:"token"`
	Errors []errors.HccError `json:"errors"`
}

// IsValid : Contain the validation of the token
type IsValid struct {
	IsValid bool              `json:"isvalid"`
	Errors  []errors.HccError `json:"errors"`
}
