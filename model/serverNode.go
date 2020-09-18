package model

import (
	"hcc/piccolo/lib/errors"
	"time"
)

// ServerNode : Contain infos of the server's node
type ServerNode struct {
	UUID          string            `json:"uuid"`
	ServerUUID    string            `json:"server_uuid"`
	NodeUUID      string            `json:"node_uuid"`
	CPUModel      string            `json:"cpu_model"`
	CPUProcessors int               `json:"cpu_processors"`
	CPUCores      int               `json:"cpu_cores"`
	CPUThreads    int               `json:"cpu_threads"`
	Memory        int               `json:"memory"`
	CreatedAt     time.Time         `json:"created_at"`
	Errors        []errors.HccError `json:"errors"`
}

// ServerNodeList : Contain list of server's nodes
type ServerNodeList struct {
	ServerNodes []ServerNode      `json:"server_node_list"`
	Errors      []errors.HccError `json:"errors"`
}

// ServerNodeNum : Contain the number of server's nodes
type ServerNodeNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}
