package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// TaskListResult : Array struct of taskListResult
type TaskListResult struct {
	Result string                    `json:"result"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
