package model

<<<<<<< HEAD
// Vnc : Model strucy of vnc
type Vnc struct {
	ServerUUID     string `json:"server_uuid"`
	TargetIP       string `json:"target_ip"`
	TargetPort     string `json:"target_port"`
	WebSocket      string `json:"websocket_port"`
	TargetPass     string `json:"target_pass"`
	Info           string `json:"vnc_info"`
	ActionClassify string `json:"action"`
=======
import "hcc/piccolo/lib/errors"

// VncPort : Model struct of vnc port
type VncPort struct {
	Port   string            `json:"port"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
