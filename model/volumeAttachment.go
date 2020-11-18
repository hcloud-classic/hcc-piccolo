package model

<<<<<<< HEAD
import "time"

// VolumeAttachment - cgs
type VolumeAttachment struct {
	UUID       string    `json:"uuid"`
	VolumeUUID string    `json:"volume_uuid"`
	ServerUUID string    `json:"server_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// VolumeAttachments - cgs
type VolumeAttachments struct {
	VolumeAttachments []VolumeAttachment `json:"volumeAttachment"`
=======
import (
	"hcc/piccolo/lib/errors"
	"time"
)

// VolumeAttachment - cgs
type VolumeAttachment struct {
	UUID       string            `json:"uuid"`
	VolumeUUID string            `json:"volume_uuid"`
	ServerUUID string            `json:"server_uuid"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Errors     []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
