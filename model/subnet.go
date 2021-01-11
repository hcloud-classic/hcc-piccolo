package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// Subnet : Contain infos of the subnet
type Subnet struct {
	UUID           string                `json:"uuid"`
	NetworkIP      string                `json:"network_ip"`
	Netmask        string                `json:"netmask"`
	Gateway        string                `json:"gateway"`
	NextServer     string                `json:"next_server"`
	NameServer     string                `json:"name_server"`
	DomainName     string                `json:"domain_name"`
	ServerUUID     string                `json:"server_uuid"`
	LeaderNodeUUID string                `json:"leader_node_uuid"`
	OS             string                `json:"os"`
	SubnetName     string                `json:"subnet_name"`
	CreatedAt      time.Time             `json:"created_at"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}

// SubnetList : Contain list of subnets
type SubnetList struct {
	Subnets []Subnet              `json:"subnet_list"`
	Errors  []errconv.PiccoloHccError `json:"errors"`
}

// SubnetNum : Contain the number of subnets
type SubnetNum struct {
	Number int                   `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// CreateDHCPConfResult : Contain result of creating DHCP configuration
type CreateDHCPConfResult struct {
	Result string                `json:"result"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
