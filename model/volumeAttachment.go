package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// VolumeAttachment - cgs
type VolumeAttachment struct {
	UUID       string                    `json:"uuid"`
	VolumeUUID string                    `json:"volume_uuid"`
	ServerUUID string                    `json:"server_uuid"`
	CreatedAt  time.Time                 `json:"created_at"`
	UpdatedAt  time.Time                 `json:"updated_at"`
	Errors     []errconv.PiccoloHccError `json:"errors"`
}
