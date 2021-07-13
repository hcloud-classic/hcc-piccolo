package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// User : Contain infos of the user
type User struct {
	ID             string                    `json:"id"`
	GroupID        int64                     `json:"group_id"`
	Authentication string                    `json:"authentication"`
	Name           string                    `json:"name"`
	GroupName      string                    `json:"group_name"`
	Email          string                    `json:"email"`
	LoginAt        time.Time                 `json:"login_at"`
	CreatedAt      time.Time                 `json:"created_at"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}

// UserList : Contain list of users
type UserList struct {
	Users    []User                    `json:"user_list"`
	TotalNum int                       `json:"total_num"`
	Errors   []errconv.PiccoloHccError `json:"errors"`
}

// UserNum : Contain the number of users
type UserNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// Token : Contain the user token and authentication
type Token struct {
	Token  string                    `json:"token"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// IsValid : Contain the validation of the token
type IsValid struct {
	IsValid        bool                      `json:"isvalid"`
	UserID         string                    `json:"user_id"`
	GroupID        int64                     `json:"group_id"`
	GroupName      string                    `json:"group_name"`
	Authentication string                    `json:"authentication"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}
