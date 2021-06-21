package model

import "hcc/piccolo/lib/errors"

// NodeDetail : Contain detail infos of the node
type NodeDetail struct {
	NodeUUID      string            `json:"node_uuid"`
	CPUModel      string            `json:"cpu_model"`
	CPUProcessors int               `json:"cpu_processors"`
	CPUThreads    int               `json:"cpu_threads"`
	Errors        []errors.HccError `json:"errors"`
}
