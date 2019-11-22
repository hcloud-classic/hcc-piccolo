package data

import "hcc/piccolo/model"

// Query

// SubnetData : Data structure of subnet
type SubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"subnet"`
	} `json:"data"`
}

// ListSubnetData : Data structure of list_subnet
type ListSubnetData struct {
	Data struct {
		ListSubnet []model.Subnet `json:"list_subnet"`
	} `json:"data"`
}

// AllSubnetData : Data structure of all_subnet
type AllSubnetData struct {
	Data struct {
		AllSubnet []model.Subnet `json:"all_subnet"`
	} `json:"data"`
}

// NumSubnetData : Data structure of num_subnet
type NumSubnetData struct {
	Data struct {
		NumSubnet model.SubnetNum `json:"num_subnet"`
	} `json:"data"`
}

// CreateSubnetData : Data structure of create_subnet
type CreateSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"create_subnet"`
	} `json:"data"`
}

// UpdateSubnetData : Data structure of update_subnet
type UpdateSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"update_subnet"`
	} `json:"data"`
}

// DeleteSubnetData : Data structure of delete_subnet
type DeleteSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"delete_subnet"`
	} `json:"data"`
}

// CreateDHCPDConfData : Data structure of create_dhcpd_conf
type CreateDHCPDConfData struct {
	Data struct {
		Result string `json:"create_dhcpd_conf"`
	} `json:"data"`
}
