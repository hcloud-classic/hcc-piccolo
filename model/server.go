package model

<<<<<<< HEAD
import "time"

// Server - cgs
type Server struct {
	UUID       string    `json:"uuid"`
	SubnetUUID string    `json:"subnet_uuid"`
	OS         string    `json:"os"`
	ServerName string    `json:"server_name"`
	ServerDesc string    `json:"server_desc"`
	CPU        int       `json:"cpu"`
	Memory     int       `json:"memory"`
	DiskSize   int       `json:"disk_size"`
	Status     string    `json:"status"`
	UserUUID   string    `json:"user_uuid"`
	CreatedAt  time.Time `json:"created_at"`
}

// Servers - cgs
type Servers struct {
	Server []Server `json:"server"`
}

// ServerNum - cgs
type ServerNum struct {
	Number int `json:"number"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// Server : Contain infos of the server
type Server struct {
	UUID       string            `json:"uuid"`
	SubnetUUID string            `json:"subnet_uuid"`
	OS         string            `json:"os"`
	ServerName string            `json:"server_name"`
	ServerDesc string            `json:"server_desc"`
	CPU        int               `json:"cpu"`
	Memory     int               `json:"memory"`
	DiskSize   int               `json:"disk_size"`
	Status     string            `json:"status"`
	UserUUID   string            `json:"user_uuid"`
	CreatedAt  time.Time         `json:"created_at"`
	Errors     []errors.HccError `json:"errors"`
}

// ServerList : Contain list of servers
type ServerList struct {
	Servers []Server          `json:"server_list"`
	Errors  []errors.HccError `json:"errors"`
}

// ServerNum : Contain the number of servers
type ServerNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
