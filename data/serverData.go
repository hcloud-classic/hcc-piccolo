package data

import "hcc/piccolo/model"

// Query

// ServerData : Data structure of server
type ServerData struct {
	Data struct {
		Server model.Server `json:"server"`
	} `json:"data"`
}

// ListServerData : Data structure of list_server
type ListServerData struct {
	Data struct {
		ListServer []model.Server `json:"list_server"`
	} `json:"data"`
}

// AllServerData : Data structure of all_server
type AllServerData struct {
	Data struct {
		AllServer []model.Server `json:"all_server"`
	} `json:"data"`
}

// NumServerData : Data structure of num_server
type NumServerData struct {
	Data struct {
		NumServer model.ServerNum `json:"num_server"`
	} `json:"data"`
}

// CreateServerData : Data structure of create_server
type CreateServerData struct {
	Data struct {
		Server model.Server `json:"create_server"`
	} `json:"data"`
}

// UpdateServerData : Data structure of update_server
type UpdateServerData struct {
	Data struct {
		Server model.Server `json:"update_server"`
	} `json:"data"`
}

// DeleteServerData : Data structure of delete_server
type DeleteServerData struct {
	Data struct {
		Server model.Server `json:"delete_server"`
	} `json:"data"`
}

// CreateServerNodeData : Data structure of create_server_node
type CreateServerNodeData struct {
	Data struct {
		Server model.ServerNode `json:"create_server_node"`
	} `json:"data"`
}

// DeleteServerNodeData : Data structure of delete_server_node
type DeleteServerNodeData struct {
	Data struct {
		Server model.ServerNode `json:"delete_server_node"`
	} `json:"data"`
}
