package model

<<<<<<< HEAD
import "time"

// Node - cgs
type Node struct {
	UUID        string    `json:"uuid"`
	ServerUUID  string    `json:"server_uuid"`
	BmcMacAddr  string    `json:"bmc_mac_addr"`
	BmcIP       string    `json:"bmc_ip"`
	PXEMacAddr  string    `json:"pxe_mac_addr"`
	Status      string    `json:"status"`
	CPUCores    int       `json:"cpu_cores"`
	Memory      int       `json:"memory"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Active      int       `json:"active"`
	ForceOff    bool      `json:"force_off"`
}

// Nodes - cgs
type Nodes struct {
	Nodes []Node `json:"node"`
}

// NodeNum - cgs
type NodeNum struct {
	Number int `json:"number"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// Node : Contain infos of the node
type Node struct {
	UUID            string            `json:"uuid"`
	ServerUUID      string            `json:"server_uuid"`
	BmcMacAddr      string            `json:"bmc_mac_addr"`
	BmcIP           string            `json:"bmc_ip"`
	BmcIPSubnetMask string            `json:"bmc_ip_subnet_mask"`
	PXEMacAddr      string            `json:"pxe_mac_addr"`
	Status          string            `json:"status"`
	CPUCores        int               `json:"cpu_cores"`
	Memory          int               `json:"memory"`
	Description     string            `json:"description"`
	RackNumber      int               `json:"rack_number"`
	CreatedAt       time.Time         `json:"created_at"`
	Active          int               `json:"active"`
	ForceOff        bool              `json:"force_off"`
	Errors          []errors.HccError `json:"errors"`
}

// NodeList : Contain list of nodes
type NodeList struct {
	Nodes  []Node            `json:"node_list"`
	Errors []errors.HccError `json:"errors"`
}

// NodeNum : Contain the number of nodes
type NodeNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

// PowerControlNode : Contain the result of node's power control
type PowerControlNode struct {
	Results []string          `json:"results"`
	Errors  []errors.HccError `json:"errors"`
}

// PowerStateNode : Contain the power state of the node
type PowerStateNode struct {
	Result string            `json:"result"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
