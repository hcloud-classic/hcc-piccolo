package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Pool :
type Pool struct {
	UUID          string                `json:"uuid"`
	Size          string                `json:"size"`
	Free          string                `json:"free"`
	Capacity      string                `json:"capacity"`
	Health        string                `json:"health"`
	Name          string                `json:"name"`
	AvailableSize string                `json:"availablesize"`
	Used          string                `json:"used"`
	Action        string                `json:"action"`
	Errors        []errconv.PiccoloHccError `json:"errors"`
}

// PoolList :
type PoolList struct {
	Pools  []Pool                `json:"pool_list"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
