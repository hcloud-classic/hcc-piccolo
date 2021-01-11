package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// Server : Contain infos of the server
type Server struct {
	UUID       string                `json:"uuid"`
	SubnetUUID string                `json:"subnet_uuid"`
	OS         string                `json:"os"`
	ServerName string                `json:"server_name"`
	ServerDesc string                `json:"server_desc"`
	CPU        int                   `json:"cpu"`
	Memory     int                   `json:"memory"`
	DiskSize   int                   `json:"disk_size"`
	Status     string                `json:"status"`
	UserUUID   string                `json:"user_uuid"`
	CreatedAt  time.Time             `json:"created_at"`
	Errors     []errconv.PiccoloHccError `json:"errors"`
}

// ServerList : Contain list of servers
type ServerList struct {
	Servers []Server              `json:"server_list"`
	Errors  []errconv.PiccoloHccError `json:"errors"`
}

// ServerNum : Contain the number of servers
type ServerNum struct {
	Number int                   `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
