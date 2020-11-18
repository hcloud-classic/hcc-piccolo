package model

<<<<<<< HEAD
import "time"

// AdaptiveIP - ish
type AdaptiveIP struct {
	UUID           string    `json:"uuid"`
	NetworkAddress string    `json:"network_address"`
	Netmask        string    `json:"netmask"`
	Gateway        string    `json:"gateway"`
	StartIPAddress string    `json:"start_ip_address"`
	EndIPAddress   string    `json:"end_ip_address"`
	CreatedAt      time.Time `json:"created_at"`
}

// AdaptiveIPs - ish
type AdaptiveIPs struct {
	AdaptiveIP []Subnet `json:"adaptiveip"`
}

// AdaptiveIPNum - ish
type AdaptiveIPNum struct {
	Number int `json:"number"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// AdaptiveIP : Structure of AdaptiveIP
type AdaptiveIP struct {
	UUID           string            `json:"uuid"`
	NetworkAddress string            `json:"network_address"`
	Netmask        string            `json:"netmask"`
	Gateway        string            `json:"gateway"`
	StartIPAddress string            `json:"start_ip_address"`
	EndIPAddress   string            `json:"end_ip_address"`
	CreatedAt      time.Time         `json:"created_at"`
	Errors         []errors.HccError `json:"errors"`
}

// AdaptiveIPNum : Structure of AdaptiveIPNum
type AdaptiveIPNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

// AdaptiveIPSetting : Structure of AdaptiveIPSetting
type AdaptiveIPSetting struct {
	ExtIfaceIPAddress string            `json:"ext_ifaceip_address"`
	Netmask           string            `json:"netmask"`
	GatewayAddress    string            `json:"gateway_address"`
	StartIPAddress    string            `json:"start_ip_address"`
	EndIPAddress      string            `json:"end_ip_address"`
	Errors            []errors.HccError `json:"errors"`
}

// AdaptiveIPAvailableIPList : Struct of AdaptiveIPAvailableIPList
type AdaptiveIPAvailableIPList struct {
	AvailableIPList []string          `json:"available_ip_list"`
	Errors          []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
