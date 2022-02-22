package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// Volume : Contain infos of the volume
type Volume struct {
	UUID       string                    `json:"uuid"`
	Size       int                       `json:"size"`
	Filesystem string                    `json:"filesystem"` // os
	ServerUUID string                    `json:"server_uuid"`
	UseType    string                    `json:"use_type"` //
	UserUUID   string                    `json:"user_uuid"`
	CreatedAt  time.Time                 `json:"created_at"`
	NetworkIP  string                    `json:"network_ip"`
	GatewayIP  string                    `json:"gateway_ip"`
	Errors     []errconv.PiccoloHccError `json:"errors"`
	LunNum     int                       `json:"lun_num"`
	Pool       string                    `json:"pool"`
	GroupID    int                       `json:"group_id"`
	State      string                    `json:"state"`
}

// VolumeList : Contain list of volumes
type VolumeList struct {
	Volumes  []Volume                  `json:"volume_list"`
	TotalNum int                       `json:"total_num"`
	Errors   []errconv.PiccoloHccError `json:"errors"`
}

// VolumeNum : Contain the number of volumes
type VolumeNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
