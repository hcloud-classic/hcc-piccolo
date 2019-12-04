package data

import "hcc/piccolo/model"

// Query

// ServerNodeData : Data structure of server_node
type ServerNodeData struct {
	Data struct {
		ServerNode model.ServerNode `json:"server_node"`
	} `json:"data"`
}

// ListServerNodeData : Data structure of list_server_node
type ListServerNodeData struct {
	Data struct {
		ListServerNode []model.ServerNode `json:"list_server_node"`
	} `json:"data"`
}

// AllServerNodeData : Data structure of all_server_node
type AllServerNodeData struct {
	Data struct {
		AllServerNode []model.ServerNode `json:"all_server_node"`
	} `json:"data"`
}

// NumNodesServerData : Data structure of num_nodes_server
type NumNodesServerData struct {
	Data struct {
		NumNodesServer model.ServerNodeNum `json:"num_nodes_server"`
	} `json:"data"`
}
