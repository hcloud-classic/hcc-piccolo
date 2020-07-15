package data

import "hcc/piccolo/model"

type ServerData struct {
	Data struct {
		Server model.Server `json:"server"`
	} `json:"data"`
}

type ListServerData struct {
	Data struct {
		ListServer []model.Server `json:"list_server"`
	} `json:"data"`
}

type AllServerData struct {
	Data struct {
		AllServer []model.Server `json:"all_server"`
	} `json:"data"`
}

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

type UpdateServerData struct {
	Data struct {
		Server model.Server `json:"update_server"`
	} `json:"data"`
}

type DeleteServerData struct {
	Data struct {
		Server model.Server `json:"delete_server"`
	} `json:"data"`
}

type CreateServerNodeData struct {
	Data struct {
		Server model.ServerNode `json:"create_server_node"`
	} `json:"data"`
}

type DeleteServerNodeData struct {
	Data struct {
		Server model.ServerNode `json:"delete_server_node"`
	} `json:"data"`
}
