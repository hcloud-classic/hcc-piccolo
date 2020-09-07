package model

import "hcc/piccolo/lib/errors"

// VncPort : Model struct of vnc port
type VncPort struct {
	Port     string `json:"port"`
	Errors []errors.HccError `json:"errors"`
}
