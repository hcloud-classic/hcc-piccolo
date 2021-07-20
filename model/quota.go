package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Quota : Contain infos of the quota
type Quota struct {
	GroupID            int64                     `json:"group_id"`
	GroupName          string                    `json:"group_name"`
	TotalCPUCores      int                       `json:"total_cpu_cores"`
	TotalMemoryGB      int                       `json:"total_memory_gb"`
	LimitSubnetCnt     int                       `json:"limit_subnet_cnt"`
	LimitAdaptiveIPCnt int                       `json:"limit_adaptive_ip_cnt"`
	LimitNodeCnt       int                       `json:"limit_node_cnt"`
	PoolName           string                    `json:"pool_name"`
	LimitSSDGB         int                       `json:"limit_ssd_gb"`
	LimitHDDGB         int                       `json:"limit_hdd_gb"`
	Errors             []errconv.PiccoloHccError `json:"errors"`
}

// QuotaList : Contain list of quotas
type QuotaList struct {
	Quotas   []Quota                   `json:"quota_list"`
	TotalNum int                       `json:"total_num"`
	Errors   []errconv.PiccoloHccError `json:"errors"`
}

// QuotaNum : Contain the number of quotas
type QuotaNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// QuotaDetail : Contain detail infos of the quota
type QuotaDetail struct {
	GroupID            int64                     `json:"group_id"`
	GroupName          string                    `json:"group_name"`
	TotalCPUCores      int                       `json:"total_cpu_cores"`
	TotalMemoryGB      int                       `json:"total_memory_gb"`
	Nodes              []Node                    `json:"nodes"`
	TotalSubnetNum     int                       `json:"total_subnet_num"`
	TotalAdaptiveIPNum int                       `json:"total_adaptive_ip_num"`
	Volumes            []Volume                  `json:"volumes"`
	Errors             []errconv.PiccoloHccError `json:"errors"`
}
