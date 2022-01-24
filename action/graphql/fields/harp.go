package fields

import (
	"hcc/piccolo/action/graphql/mutationparser"
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
	"innogrid.com/hcloud-classic/hcc_errors"
)

// Query
var Subnet = graphql.Field{
	Type:        graphqlType.SubnetType,
	Description: "Get subnet by uuid",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.Subnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / subnet: " + err.Error())
		}
		return data, err
	},
}

var ListSubnet = graphql.Field{
	Type:        graphqlType.SubnetListType,
	Description: "Get subnet list",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"network_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"netmask": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"gateway": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"next_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"domain_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"leader_node_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"os": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"subnet_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.ListSubnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / list_subnet : " + err.Error())
		}
		return data, err
	},
}

var AllSubnet = graphql.Field{
	Type:        graphqlType.SubnetListType,
	Description: "Get all subnet list",
	Args: graphql.FieldConfigArgument{
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.AllSubnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / all_subnet: " + err.Error())
		}
		return data, err
	},
}

var AvailableSubnet = graphql.Field{
	Type:        graphqlType.SubnetListType,
	Description: "Get available subnet list",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.AvailableSubnetList(params.Args)
		if err != nil {
			logger.Logger.Println("harp / available_subnet: " + err.Error())
		}
		return data, err
	},
}

var NumSubnet = graphql.Field{
	Type:        graphqlType.SubnetNumType,
	Description: "Get the number of subnets",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.SubnetNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NumSubnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / num_subnet: " + err.Error())
		}
		return data, err
	},
}

var VaildCheckSubnet = graphql.Field{
	Type:        graphqlType.SubnetValidType,
	Description: "Check if we can create the subnet",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"network_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"netmask": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"gateway": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"is_update": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ValidCheckSubnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / valid_check_subnet: " + err.Error())
		}
		return data, err
	},
}

var AdaptiveIpAvailableIpList = graphql.Field{
	Type:        graphqlType.AdaptiveIPAvailableIPListType,
	Description: "Get available ip list for adaptive ip",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPAvailableIPList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.GetAdaptiveIPAvailableIPList()
		if err != nil {
			logger.Logger.Println("harp / adaptiveip_available_ip_list: " + err.Error())
		}
		return data, err
	},
}

var AdaptiveIpSetting = graphql.Field{
	Type:        graphqlType.AdaptiveIPSettingType,
	Description: "Get settings of adaptiveip",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Only master can change this setting!")}, nil
		}
		data, err := queryparser.GetAdaptiveIPSetting()
		if err != nil {
			logger.Logger.Println("harp / adaptiveip_setting: " + err.Error())
		}
		return data, err
	},
}

var AdaptiveIPSerer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerType,
	Description: "Get adaptiveip by uuid",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.AdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var ListAdaptiveIpServer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerListType,
	Description: "Get adaptiveip_server list",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"public_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"private_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"private_gateway": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		if !isAdmin && !isMaster {
			params.Args["user_uuid"] = id
		}
		data, err := queryparser.ListAdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / list_adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var AllAdaptiveIpServer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerListType,
	Description: "Get all adaptiveip_server list",
	Args: graphql.FieldConfigArgument{
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		if !isAdmin && !isMaster {
			params.Args["user_uuid"] = id
		}
		data, err := queryparser.AllAdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / all_adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var NumAdaptiveIpServer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerNumType,
	Description: "Get the number of AdaptiveIP Servers",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServerNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NumAdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / num_adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var ListPortForwarding = graphql.Field{
	Type:        graphqlType.PortForwardingListType,
	Description: "Get port_forwarding list",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"protocol": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"external_port": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"internal_port": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PortForwardingList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ListPortForwarding(params.Args)
		if err != nil {
			logger.Logger.Println("harp / list_port_forwarding: " + err.Error())
		}
		return data, err
	},
}

// Mutation
var CreateSubnet = graphql.Field{
	Type:        graphqlType.SubnetType,
	Description: "Create new subnet",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"network_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"netmask": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"gateway": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"next_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"domain_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"subnet_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := mutationparser.CreateSubnet(params.Args)
		if err != nil {
			logger.Logger.Println("harp / create_subnet: " + err.Error())
		}
		return data, err
	},
}

var UpdateSubnet = graphql.Field{
	Type:        graphqlType.SubnetType,
	Description: "Update subnet",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"network_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"netmask": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"gateway": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"next_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name_server": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"domain_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"os": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"subnet_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.UpdateSubnet(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("harp / update_subnet: " + err.Error())
		}
		return data, err
	},
}

var DeleteSubnet = graphql.Field{
	Type:        graphqlType.SubnetType,
	Description: "Delete subnet by uuid",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.DeleteSubnet(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("harp / delete_subnet: " + err.Error())
		}
		return data, err
	},
}

var CreateDHCPConf = graphql.Field{
	Type:        graphqlType.CreateDHCPConfResultType,
	Description: "Create new dhcpd config",
	Args: graphql.FieldConfigArgument{
		"subnet_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.CreateDHCPConfResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateDHCPDConf(params.Args)
		if err != nil {
			logger.Logger.Println("harp / create_dhcpd_conf: " + err.Error())
		}
		return data, err
	},
}

var CreateAdaptiveIpSetting = graphql.Field{
	Type:        graphqlType.AdaptiveIPSettingType,
	Description: "Create settings of adaptiveip",
	Args: graphql.FieldConfigArgument{
		"ext_ifaceip_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"netmask": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"gateway_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"internal_start_ip_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"internal_end_ip_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"external_start_ip_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"external_end_ip_address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Only master can change this setting!")}, nil
		}
		data, err := mutationparser.CreateAdaptiveIPSetting(params.Args)
		if err != nil {
			logger.Logger.Println("harp / create_adaptiveip_setting: " + err.Error())
		}
		return data, err
	},
}

var CreateAdaptiveIpServer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerType,
	Description: "Create new adaptiveip_server",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"public_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateAdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / create_adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var DeleteAdaptiveIpServer = graphql.Field{
	Type:        graphqlType.AdaptiveIPServerType,
	Description: "Delete adaptiveip_server by server_uuid",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.DeleteAdaptiveIPServer(params.Args)
		if err != nil {
			logger.Logger.Println("harp / delete_adaptiveip_server: " + err.Error())
		}
		return data, err
	},
}

var CreatePortForwarding = graphql.Field{
	Type:        graphqlType.PortForwardingType,
	Description: "Create new AdaptiveIP Port Forwarding",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"protocol": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"external_port": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"internal_port": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreatePortForwarding(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("harp / create_port_forwarding: " + err.Error())
		}
		return data, err
	},
}

var DeletePortForwarding = graphql.Field{
	Type:        graphqlType.PortForwardingType,
	Description: "Delete AdaptiveIP Port Forwarding by server_uuid",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"external_port": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.DeletePortForwarding(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("harp / delete_port_forwarding: " + err.Error())
		}
		return data, err
	},
}
