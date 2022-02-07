package model

import "hcc/piccolo/action/grpc/errconv"

type Service struct {
	Target string                    `json:"target"`
	Status string                    `json:"result"`
	Action string                    `json:"action"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
