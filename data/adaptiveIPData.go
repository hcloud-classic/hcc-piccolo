package data

import "hcc/piccolo/model"

// AdaptiveIPData : Data structure of adaptiveip
type AdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"adaptiveip"`
	} `json:"data"`
}

type AdaptiveIPAvailableIPListData struct {
	Data struct {
		AdaptiveIPAvailableIPList model.AdaptiveIPAvailableIPList `json:"adaptiveip_available_ip_list"`
	} `json:"data"`
}

// CreateAdaptiveIPData : Data structure of create_adaptiveip
type CreateAdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"create_adaptiveip"`
	} `json:"data"`
}
