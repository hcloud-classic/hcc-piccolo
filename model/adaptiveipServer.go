package model

<<<<<<< HEAD
// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	AdaptiveIPUUID string `json:"adaptiveip_uuid"`
	ServerUUID     string `json:"server_uuid"`
	PublicIP       string `json:"public_ip"`
	PrivateIP      string `json:"private_ip"`
	PrivateGateway string `json:"private_gateway"`
}

// AdaptiveIPServers - ish
type AdaptiveIPServers struct {
	AdaptiveIP []Subnet `json:"adaptiveip"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	ServerUUID     string            `json:"server_uuid"`
	PublicIP       string            `json:"public_ip"`
	PrivateIP      string            `json:"private_ip"`
	PrivateGateway string            `json:"private_gateway"`
	CreatedAt      time.Time         `json:"created_at"`
	Errors         []errors.HccError `json:"errors"`
}

// AdaptiveIPServerList : Contain list of AdaptiveIPServers
type AdaptiveIPServerList struct {
	AdaptiveIPServers []AdaptiveIPServer `json:"adaptiveip_server_list"`
	Errors            []errors.HccError  `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
<<<<<<< HEAD
	Number int `json:"number"`
=======
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
