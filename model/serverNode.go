package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// ServerNode : Contain infos of the server's node
type ServerNode struct {
	UUID          string                    `json:"uuid"`
	ServerUUID    string                    `json:"server_uuid"`
	NodeName      string                    `json:"node_name"`
	NodeUUID      string                    `json:"node_uuid"`
	CPUModel      string                    `json:"cpu_model"`
	CPUProcessors int                       `json:"cpu_processors"`
	CPUCores      int                       `json:"cpu_cores"`
	CPUThreads    int                       `json:"cpu_threads"`
	Memory        int                       `json:"memory"`
	RackNumber    int                       `json:"rack_number"`
	CreatedAt     time.Time                 `json:"created_at"`
	Errors        []errconv.PiccoloHccError `json:"errors"`
}

// ServerNodeList : Contain list of server's nodes
type ServerNodeList struct {
	ServerNodes []ServerNode              `json:"server_node_list"`
	Errors      []errconv.PiccoloHccError `json:"errors"`
}

// ServerNodeNum : Contain the number of server's nodes
type ServerNodeNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
