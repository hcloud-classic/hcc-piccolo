package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// NodeDetail : Contain detail infos of the node
type NodeDetail struct {
	NodeUUID       string                    `json:"node_uuid"`
	NodeDetailData string                    `json:"node_detail_data"`
	NicDetailData  string                    `json:"nic_detail_data"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}
