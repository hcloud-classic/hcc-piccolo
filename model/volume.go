package model

<<<<<<< HEAD
import "time"

// Volume - cgs
type Volume struct {
	UUID       string    `json:"uuid"`
	Size       int       `json:"size"`
	Filesystem string    `json:"filesystem"` //os
	ServerUUID string    `json:"server_uuid"`
	UseType    string    `json:"use_type"` //
	UserUUID   string    `json:"user_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	NetworkIP  string    `json:"network_ip"`
}

// Volumes - cgs
type Volumes struct {
	Volumes []Volume `json:"volume"`
=======
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// VolumeNum - cgs
type VolumeNum struct {
<<<<<<< HEAD
	Number int `json:"number"`
=======
	Number int               `json:"number"`
	Errors []errors.HccError `json:"errors"`
}

type VolumeList struct {
	Volumes []Volume          `json:"volume_list"`
	Errors  []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
