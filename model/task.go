package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Task : Struct of task
type Task struct {
	CMD        string `json:"cmd"`
	State      string `json:"state"`
	PID        int    `json:"pid"`
	PPID       int    `json:"ppid"`
	PGID       int    `json:"pgid"`
	SID        int    `json:"sid"`
	Priority   int    `json:"priority"`
	Nice       int    `json:"nice"`
	NumThreads int    `json:"num_threads"`
	StartTime  string `json:"start_time"`
	Children   []Task `json:"children"`
	Threads    []Task `json:"threads"`
	CPUUsage   string `json:"cpu_usage"`
	MemUsage   string `json:"mem_usage"`
	EPMType    string `json:"epm_type"`
	EPMSource  int    `json:"epm_source"`
	EPMTarget  int    `json:"epm_target"`
}

// TaskList : Array struct of tasks
type TaskList struct {
	Tasks                []Task                    `json:"task_list"`
	TotalTasks           int                       `json:"total_tasks"`
	TotalMemUsage        string                    `json:"total_mem_usage"`
	TotalMem             string                    `json:"total_mem"`
	TotalMemUsagePercent string                    `json:"total_mem_usage_percent"`
	TotalCPUUsage        string                    `json:"total_cpu_usage"`
	Errors               []errconv.PiccoloHccError `json:"errors"`
}
