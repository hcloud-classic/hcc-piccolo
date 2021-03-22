package model

import (
	"hcc/piccolo/action/grpc/errconv"
)

// BillingData : Contain infos of the billingData
type BillingData struct {
	BillingType string                    `json:"billing_type"`
	GroupID     []int32                   `json:"group_id"`
	Result      string                    `json:"result"`
	Errors      []errconv.PiccoloHccError `json:"errors"`
}
