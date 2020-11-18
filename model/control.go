package model

<<<<<<< HEAD
// Control : Struct of Control
type Control struct {
	HccCommand string `json:"action"`
	HccIPRange string `json:"iprange"`
	ServerUUID string `json:"server_uuid"`
}

// Controls : Array struct of Control
type Controls struct {
	Controls Control `json:"control"`
=======
import "hcc/piccolo/lib/errors"

// Control : Struct of Control
type Control struct {
	HccCommand string            `json:"action"`
	HccIPRange string            `json:"iprange"`
	ServerUUID string            `json:"server_uuid"`
	Errors     []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
