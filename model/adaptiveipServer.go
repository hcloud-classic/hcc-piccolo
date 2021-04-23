package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	ServerUUID     string                    `json:"server_uuid"`
	GroupID        int64                     `json:"group_id"`
	GroupName      string                    `json:"group_name"`
	PublicIP       string                    `json:"public_ip"`
	PrivateIP      string                    `json:"private_ip"`
	PrivateGateway string                    `json:"private_gateway"`
	CreatedAt      time.Time                 `json:"created_at"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPServerList : Contain list of AdaptiveIPServers
type AdaptiveIPServerList struct {
	AdaptiveIPServers []AdaptiveIPServer        `json:"adaptiveip_server_list"`
	TotalNum          int                       `json:"total_num"`
	Errors            []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
