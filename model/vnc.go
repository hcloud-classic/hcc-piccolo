package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// VncPort : Model struct of vnc port
type VncPort struct {
	Port   string                `json:"port"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
