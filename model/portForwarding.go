package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// PortForwarding : Contain port forwarding info of the AdaptiveIP
type PortForwarding struct {
	ServerUUID   string                    `json:"server_uuid"`
	Protocol     string                    `json:"protocol"`
	ExternalPort int64                     `json:"external_port"`
	InternalPort int64                     `json:"internal_port"`
	Description  string                    `json:"description"`
	Errors       []errconv.PiccoloHccError `json:"errors"`
}

// PortForwardingList : Contain list of PortForwarding
type PortForwardingList struct {
	PortForwardings []PortForwarding          `json:"port_forwarding_list"`
	TotalNum        int                       `json:"total_num"`
	Errors          []errconv.PiccoloHccError `json:"errors"`
}
