package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Resource : Contain infos of resources
type Resource struct {
	CPU     int `json:"cpu"`
	Memory  int `json:"memory"`
	Storage int `json:"storage"`
	Node    int `json:"node"`
}

// ResourceUsage : Contain usage info of resources
type ResourceUsage struct {
	Total  Resource              `json:"total"`
	InUse  Resource              `json:"in_use"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
