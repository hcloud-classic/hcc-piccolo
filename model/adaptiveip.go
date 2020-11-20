package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// AdaptiveIP : Structure of AdaptiveIP
type AdaptiveIP struct {
	UUID           string                    `json:"uuid"`
	NetworkAddress string                    `json:"network_address"`
	Netmask        string                    `json:"netmask"`
	Gateway        string                    `json:"gateway"`
	StartIPAddress string                    `json:"start_ip_address"`
	EndIPAddress   string                    `json:"end_ip_address"`
	CreatedAt      time.Time                 `json:"created_at"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPNum : Structure of AdaptiveIPNum
type AdaptiveIPNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPSetting : Structure of AdaptiveIPSetting
type AdaptiveIPSetting struct {
	ExtIfaceIPAddress string                    `json:"ext_ifaceip_address"`
	Netmask           string                    `json:"netmask"`
	GatewayAddress    string                    `json:"gateway_address"`
	StartIPAddress    string                    `json:"start_ip_address"`
	EndIPAddress      string                    `json:"end_ip_address"`
	Errors            []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPAvailableIPList : Struct of AdaptiveIPAvailableIPList
type AdaptiveIPAvailableIPList struct {
	AvailableIPList []string                  `json:"available_ip_list"`
	Errors          []errconv.PiccoloHccError `json:"errors"`
}
