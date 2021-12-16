package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// ServerAction : Structure of ServerAction
type ServerAction struct {
	Action string    `json:"action"`
	Result string    `json:"result"`
	ErrStr string    `json:"err_str"`
	UserID string    `json:"user_id"`
	Time   time.Time `json:"time"`
}

// ServerActions : Struct of ServerActions
type ServerActions struct {
	ServerActions []ServerAction            `json:"server_actions"`
	Number        int                       `json:"number"`
	Errors        []errconv.PiccoloHccError `json:"errors"`
}
