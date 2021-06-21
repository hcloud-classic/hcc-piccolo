package model

import (
	"hcc/piccolo/lib/errors"
	"time"
)

// Volume - cgs
type Volume struct {
	UUID       string            `json:"uuid"`
	Size       int               `json:"size"`
	Filesystem string            `json:"filesystem"` //os
	ServerUUID string            `json:"server_uuid"`
	UseType    string            `json:"use_type"` //
	UserUUID   string            `json:"user_uuid"`
	CreatedAt  time.Time         `json:"created_at"`
	NetworkIP  string            `json:"network_ip"`
	GatewayIP  string            `json:"gateway_ip"`
	Errors     []errors.HccError `json:"errors"`
	LunNum     int               `json:"lun_num"`
	Pool       string            `json:"pool"`
}

// VolumeNum - cgs
type VolumeNum struct {
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

type VolumeList struct {
	Volumes []Volume          `json:"volume_list"`
	Errors  []errors.HccError `json:"errors"`
}
