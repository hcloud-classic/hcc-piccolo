package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// Quota : Contain infos of the quota
type Quota struct {
	GroupID             int64  `json:"group_id"`
	GroupName           string `json:"group_name"`
	LimitCPUCores       int    `json:"limit_cpu_cores"`
	LimitMemoryGB       int    `json:"limit_memory_gb"`
	LimitSubnetHostBits int    `json:"limit_subnet_host_bits"`
	LimitAdaptiveIPCnt  int    `json:"limit_adaptive_ip_cnt"`
	LimitSSDGB          int    `json:"limit_ssd_gb"`
	LimitHDDGB          int    `json:"limit_hdd_gb"`
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
	GroupID             int64                     `json:"group_id"`
	GroupName           string                    `json:"group_name"`
	TotalCPUCores       int                       `json:"total_cpu_cores"`
	TotalMemoryGB       int                       `json:"total_memory_gb"`
	Nodes               []Node                    `json:"nodes"`
	TotalSubnetHostBits int                       `json:"total_subnet_host_bits"`
	TotalAdaptiveIPNum  int                       `json:"total_adaptive_ip_num"`
	Volumes             []Volume                  `json:"volumes"`
	Errors              []errconv.PiccoloHccError `json:"errors"`
}
