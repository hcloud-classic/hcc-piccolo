package model

import "hcc/piccolo/lib/errors"

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	ServerUUID     string `json:"server_uuid"`
	PublicIP       string `json:"public_ip"`
	PrivateIP      string `json:"private_ip"`
	PrivateGateway string `json:"private_gateway"`
	Errors []errors.HccError `json:"errors"`
}

// AdaptiveIPServerList : Contain list of AdaptiveIPServers
type AdaptiveIPServerList struct {
	AdaptiveIPServers []AdaptiveIPServer `json:"adaptiveip_server_list"`
	Errors []errors.HccError `json:"errors"`
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
	Number int `json:"number"`
	Errors []errors.HccError `json:"errors"`
}
