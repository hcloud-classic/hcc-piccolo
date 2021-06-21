package model

import "hcc/piccolo/lib/errors"

// Control : Struct of Control
type Control struct {
	HccCommand string            `json:"action"`
	HccIPRange string            `json:"iprange"`
	ServerUUID string            `json:"server_uuid"`
	Errors     []errors.HccError `json:"errors"`
}
