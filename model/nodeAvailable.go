package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// NodeAvailable : Contain available info of nodes
type NodeAvailable struct {
	Total     Resource                  `json:"total"`
	Available Resource                  `json:"available"`
	Errors    []errconv.PiccoloHccError `json:"errors"`
}
