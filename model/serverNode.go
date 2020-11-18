package model

<<<<<<< HEAD
import "time"

// ServerNode - cgs
type ServerNode struct {
	UUID       string    `json:"uuid"`
	ServerUUID string    `json:"server_uuid"`
	NodeUUID   string    `json:"node_uuid"`
	CreatedAt  time.Time `json:"created_at"`
}

// ServerNodes - cgs
type ServerNodes struct {
	Server []Server `json:"server_node"`
}

// ServerNodeNum - ish
type ServerNodeNum struct {
	Number int `json:"number"`
=======
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
	RackNumber    int               `json:"rack_number"`
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
