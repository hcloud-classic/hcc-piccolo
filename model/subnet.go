package model

<<<<<<< HEAD
import "time"

// Subnet - cgs
type Subnet struct {
	UUID           string    `json:"uuid"`
	NetworkIP      string    `json:"network_ip"`
	Netmask        string    `json:"netmask"`
	Gateway        string    `json:"gateway"`
	NextServer     string    `json:"next_server"`
	NameServer     string    `json:"name_server"`
	DomainName     string    `json:"domain_name"`
	ServerUUID     string    `json:"server_uuid"`
	LeaderNodeUUID string    `json:"leader_node_uuid"`
	OS             string    `json:"os"`
	SubnetName     string    `json:"subnet_name"`
	CreatedAt      time.Time `json:"created_at"`
}

// Subnets - cgs
type Subnets struct {
	Subnets []Subnet `json:"subnet"`
}

// SubnetNum - cgs
type SubnetNum struct {
	Number int `json:"number"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// Subnet : Contain infos of the subnet
type Subnet struct {
	UUID           string            `json:"uuid"`
	NetworkIP      string            `json:"network_ip"`
	Netmask        string            `json:"netmask"`
	Gateway        string            `json:"gateway"`
	NextServer     string            `json:"next_server"`
	NameServer     string            `json:"name_server"`
	DomainName     string            `json:"domain_name"`
	ServerUUID     string            `json:"server_uuid"`
	LeaderNodeUUID string            `json:"leader_node_uuid"`
	OS             string            `json:"os"`
	SubnetName     string            `json:"subnet_name"`
	CreatedAt      time.Time         `json:"created_at"`
	Errors         []errors.HccError `json:"errors"`
}

// SubnetList : Contain list of subnets
type SubnetList struct {
	Subnets []Subnet          `json:"subnet_list"`
	Errors  []errors.HccError `json:"errors"`
}

// SubnetNum : Contain the number of subnets
type SubnetNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

// CreateDHCPConfResult : Contain result of creating DHCP configuration
type CreateDHCPConfResult struct {
	Result string            `json:"result"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
