package data

import "hcc/piccolo/model"

// AdaptiveIPServerData : Data structure of adaptiveip_server
type AdaptiveIPServerData struct {
	Data struct {
		AdaptiveIPServer model.AdaptiveIPServer `json:"adaptiveip_server"`
	} `json:"data"`
}

// ListAdaptiveIPServerData : Data structure of list_adaptiveip_server
type ListAdaptiveIPServerData struct {
	Data struct {
		ListAdaptiveIPServer []model.AdaptiveIPServer `json:"list_adaptiveip_server"`
	} `json:"data"`
}

// AllAdaptiveIPServerData : Data structure of all_adaptiveip_server
type AllAdaptiveIPServerData struct {
	Data struct {
		AllAdaptiveIPServer []model.AdaptiveIPServer `json:"all_adaptiveip_server"`
	} `json:"data"`
}

// NumAdaptiveIPServerData : Data structure of num_adaptiveip_server
type NumAdaptiveIPServerData struct {
	Data struct {
		NumAdaptiveIPServer model.AdaptiveIPServerNum `json:"num_adaptiveip_server"`
	} `json:"data"`
}

// CreateAdaptiveIPServerData : Data structure of create_adaptiveip_server
type CreateAdaptiveIPServerData struct {
	Data struct {
		AdaptiveIPServer model.AdaptiveIPServer `json:"create_adaptiveip_server"`
	} `json:"data"`
}

// DeleteAdaptiveIPServerData : Data structure of delete_adaptiveip_server
type DeleteAdaptiveIPServerData struct {
	Data struct {
		AdaptiveIPServer model.AdaptiveIPServer `json:"delete_adaptiveip_server"`
	} `json:"data"`
}
