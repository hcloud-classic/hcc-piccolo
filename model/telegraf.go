package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Telegraf - cgs
type Telegraf struct {
	UUID   string                `json:"id"`
	Result string                `json:"result"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
