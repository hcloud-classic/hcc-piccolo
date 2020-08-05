package data

import "hcc/piccolo/model"

type SubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"subnet"`
	} `json:"data"`
}

type ListSubnetData struct {
	Data struct {
		ListSubnet []model.Subnet `json:"list_subnet"`
	} `json:"data"`
}

type AllSubnetData struct {
	Data struct {
		AllSubnet []model.Subnet `json:"all_subnet"`
	} `json:"data"`
}

type NumSubnetData struct {
	Data struct {
		NumSubnet model.SubnetNum `json:"num_subnet"`
	} `json:"data"`
}

type CreateSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"create_subnet"`
	} `json:"data"`
}

type UpdateSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"update_subnet"`
	} `json:"data"`
}

type DeleteSubnetData struct {
	Data struct {
		Subnet model.Subnet `json:"delete_subnet"`
	} `json:"data"`
}

type CreateDHCPDConfData struct {
	Data struct {
		Result string `json:"create_dhcpd_conf"`
	} `json:"data"`
}
