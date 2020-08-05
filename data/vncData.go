package data

import "hcc/piccolo/model"

type ControlVncData struct {
	Data struct {
		Vnc model.Vnc `json:"control_vnc"`
	} `json:"data"`
}
type CreateVncData struct {
	Data struct {
		Vnc model.Vnc `json:"create_vnc"`
	} `json:"data"`
}
