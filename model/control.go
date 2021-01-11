package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Control : Struct of Control
type Control struct {
	HccCommand string                `json:"action"`
	HccIPRange string                `json:"iprange"`
	ServerUUID string                `json:"server_uuid"`
	Errors     []errconv.PiccoloHccError `json:"errors"`
}
