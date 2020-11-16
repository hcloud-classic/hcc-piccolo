package data

import "hcc/piccolo/model"

// Query

// AdaptiveIPData : Data structure of adaptiveip
type AdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"adaptiveip"`
	} `json:"data"`
}

// ListAdaptiveIPData : Data structure of list_adaptiveip
type ListAdaptiveIPData struct {
	Data struct {
		ListAdaptiveIP []model.AdaptiveIP `json:"list_adaptiveip"`
	} `json:"data"`
}

// AllAdaptiveIPData : Data structure of all_adaptiveip
type AllAdaptiveIPData struct {
	Data struct {
		AllAdaptiveIP []model.AdaptiveIP `json:"all_adaptiveip"`
	} `json:"data"`
}

// NumAdaptiveIPData : Data structure of num_adaptiveip
type NumAdaptiveIPData struct {
	Data struct {
		NumAdaptiveIP model.AdaptiveIPNum `json:"num_adaptiveip"`
	} `json:"data"`
}

// CreateAdaptiveIPData : Data structure of create_adaptiveip
type CreateAdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"create_adaptiveip"`
	} `json:"data"`
}

// UpdateAdaptiveIPData : Data structure of update_adaptiveip
type UpdateAdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"update_adaptiveip"`
	} `json:"data"`
}

// DeleteAdaptiveIPData : Data structure of delete_adaptiveip
type DeleteAdaptiveIPData struct {
	Data struct {
		AdaptiveIP model.AdaptiveIP `json:"delete_adaptiveip"`
	} `json:"data"`
}
