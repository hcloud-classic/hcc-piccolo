package model

import "hcc/piccolo/action/grpc/errconv"

// Group : Contain infos of the group
type Group struct {
	ID   int64  `json:"group_id"`
	Name string `json:"group_name"`
}

// GroupList : Contain list of groups
type GroupList struct {
	Groups   []Group                   `json:"group_list"`
	TotalNum int                       `json:"total_num"`
	Errors   []errconv.PiccoloHccError `json:"errors"`
}
