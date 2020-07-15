package data

import "hcc/piccolo/model"

type ServerNodeData struct {
	Data struct {
		ServerNode model.ServerNode `json:"server_node"`
	} `json:"data"`
}

type ListServerNodeData struct {
	Data struct {
		ListServerNode []model.ServerNode `json:"list_server_node"`
	} `json:"data"`
}

type AllServerNodeData struct {
	Data struct {
		AllServerNode []model.ServerNode `json:"all_server_node"`
	} `json:"data"`
}

type NumNodesServerData struct {
	Data struct {
		NumNodesServer model.ServerNodeNum `json:"num_nodes_server"`
	} `json:"data"`
}
