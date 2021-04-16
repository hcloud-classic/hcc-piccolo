package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// Server : Contain infos of the server
type Server struct {
	UUID       string                    `json:"uuid"`
	GroupID    int64                     `json:"group_id"`
	GroupName  string                    `json:"group_name"`
	SubnetUUID string                    `json:"subnet_uuid"`
	NicSpeed   string                    `json:"nic_speed"`
	ExternalIP string                    `json:"external_ip"`
	PXEBootIP  string                    `json:"pxe_boot_ip"`
	NodeList   []Node                    `json:"node_list"`
	OS         string                    `json:"os"`
	ServerName string                    `json:"server_name"`
	ServerDesc string                    `json:"server_desc"`
	CPU        int                       `json:"cpu"`
	Memory     int                       `json:"memory"`
	DiskSize   int                       `json:"disk_size"`
	Nodes      int                       `json:"nodes"`
	Status     string                    `json:"status"`
	UserUUID   string                    `json:"user_uuid"`
	CreatedAt  time.Time                 `json:"created_at"`
	Errors     []errconv.PiccoloHccError `json:"errors"`
}

// ServerList : Contain list of servers
type ServerList struct {
	Servers  []Server                  `json:"server_list"`
	TotalNum int                       `json:"total_num"`
	Errors   []errconv.PiccoloHccError `json:"errors"`
}

// ServerNum : Contain the number of servers
type ServerNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
