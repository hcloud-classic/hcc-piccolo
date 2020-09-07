package model

import (
	"hcc/piccolo/lib/errors"
	"time"
)

// ServerNode - cgs
type ServerNode struct {
	UUID       string    `json:"uuid"`
	ServerUUID string    `json:"server_uuid"`
	NodeUUID   string    `json:"node_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	Errors []errors.HccError `json:"errors"`
}

// ServerNodeList : Contain list of serverNodes
type ServerNodeList struct {
	ServerNodes []ServerNode `json:"server_node_list"`
	Errors []errors.HccError `json:"errors"`
}

// ServerNodeNum - ish
type ServerNodeNum struct {
	Number int `json:"number"`
	Errors []errors.HccError `json:"errors"`
}
