package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// Node : Contain infos of the node
type Node struct {
	UUID            string                    `json:"uuid"`
	NodeName        string                    `json:"node_name"`
	GroupID         int64                     `json:"group_id"`
	GroupName       string                    `json:"group_name"`
	NodeNum         int64                     `json:"node_num"`
	NodeIP          string                    `json:"node_ip"`
	ServerUUID      string                    `json:"server_uuid"`
	BmcMacAddr      string                    `json:"bmc_mac_addr"`
	BmcIP           string                    `json:"bmc_ip"`
	BmcIPSubnetMask string                    `json:"bmc_ip_subnet_mask"`
	PXEMacAddr      string                    `json:"pxe_mac_addr"`
	Status          string                    `json:"status"`
	CPUModel        string                    `json:"cpu_model"`
	CPUCores        int                       `json:"cpu_cores"`
	Memory          int                       `json:"memory"`
	NICSpeed        string                    `json:"nic_speed"`
	Description     string                    `json:"description"`
	RackNumber      int                       `json:"rack_number"`
	ChargeCPU       int                       `json:"charge_cpu"`
	ChargeMemory    int                       `json:"charge_memory"`
	ChargeNIC       int                       `json:"charge_nic"`
	Active          int                       `json:"active"`
	CreatedAt       time.Time                 `json:"created_at"`
	ForceOff        bool                      `json:"force_off"`
	Errors          []errconv.PiccoloHccError `json:"errors"`
}

// NodeList : Contain list of nodes
type NodeList struct {
	Nodes  []Node                    `json:"node_list"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// NodeNum : Contain the number of nodes
type NodeNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// PowerControlNode : Contain the result of node's power control
type PowerControlNode struct {
	Results []string                  `json:"results"`
	Errors  []errconv.PiccoloHccError `json:"errors"`
}

// PowerStateNode : Contain the power state of the node
type PowerStateNode struct {
	Result string                    `json:"result"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
