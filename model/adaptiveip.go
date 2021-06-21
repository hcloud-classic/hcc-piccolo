package model

import "time"

// AdaptiveIP : Structure of AdaptiveIP
type AdaptiveIP struct {
	UUID           string    `json:"uuid"`
	NetworkAddress string    `json:"network_address"`
	Netmask        string    `json:"netmask"`
	Gateway        string    `json:"gateway"`
	StartIPAddress string    `json:"start_ip_address"`
	EndIPAddress   string    `json:"end_ip_address"`
	CreatedAt      time.Time `json:"created_at"`
}

// AdaptiveIPs : Structure of AdaptiveIPs
type AdaptiveIPs struct {
	AdaptiveIP []Subnet `json:"adaptiveip"`
}

// AdaptiveIPNum : Structure of AdaptiveIPNum
type AdaptiveIPNum struct {
	Number int `json:"number"`
}

// AdaptiveIPSetting : Structure of AdaptiveIPSetting
type AdaptiveIPSetting struct {
	ExtIfaceIPAddress string `json:"ext_ifaceip_address"`
	Netmask           string `json:"netmask"`
	GatewayAddress    string `json:"gateway_address"`
	StartIPAddress    string `json:"start_ip_address"`
	EndIPAddress      string `json:"end_ip_address"`
}
