package data

import "hcc/piccolo/model"

// Query

// ControlVncData : Data structure of control_vnc
type ControlVncData struct {
	Data struct {
		Vnc model.Vnc `json:"control_vnc"`
	} `json:"data"`
}

// CreateVncData : Data structure of create_vnc
type CreateVncData struct {
	Data struct {
		Vnc model.Vnc `json:"create_vnc"`
	} `json:"data"`
}
