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

var PowerStateNode = graphql.Field{
	Type:        graphqlType.PowerStateNodeType,
	Description: "Get the node's power state",
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
			return model.PowerStateNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.PowerStateNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / power_state_node: " + err.Error())
		}
		return data, err
	},
}

var Node = graphql.Field{
	Type:        graphqlType.NodeType,
	Description: "Get a node by uuid",
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
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.Node(params.Args)
		if err != nil {
			logger.Logger.Println("flute / node: " + err.Error())
		}
		return data, err
	},
}

var ListNode = graphql.Field{
	Type:        graphqlType.NodeListType,
	Description: "Get node list",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"node_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"node_num": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"node_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"bmc_mac_addr": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"bmc_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pxe_mac_addr": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cpu_cores": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"memory": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"nic_speed_mbps": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"rack_number": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"active": &graphql.ArgumentConfig{
			Type: graphql.Int,
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
			return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.ListNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / list_node: " + err.Error())
		}
		return data, err
	},
}

var AllNode = graphql.Field{
	Type:        graphqlType.NodeListType,
	Description: "Get all node list",
	Args: graphql.FieldConfigArgument{
		"active": &graphql.ArgumentConfig{
			Type: graphql.Int,
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
			return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.AllNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / all_node: " + err.Error())
		}
		return data, err
	},
}

var AllServerPreparedNode = graphql.Field{
	Type:        graphqlType.NodeListType,
	Description: "Get all node list prepared for server",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		params.Args["server_uuid"] = "---"
		data, err := queryparser.AllNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / all_server_prepared_node: " + err.Error())
		}
		return data, err
	},
}

var NumNode = graphql.Field{
	Type:        graphqlType.NodeNumType,
	Description: "Get the number of nodes",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NumNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / num_node: " + err.Error())
		}
		return data, err
	},
}

var DetailNode = graphql.Field{
	Type:        graphqlType.NodeDetailType,
	Description: "Get a node_detail by uuid",
	Args: graphql.FieldConfigArgument{
		"node_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.NodeDetail(params.Args)
		if err != nil {
			logger.Logger.Println("flute / node_detail: " + err.Error())
		}
		return data, err
	},
}

var OnNode = graphql.Field{
	Type:        graphqlType.PowerControlNodeType,
	Description: "On node",
	Args: graphql.FieldConfigArgument{
		"uuids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.OnNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / on_node: " + err.Error())
		}
		return data, err
	},
}

var OffNode = graphql.Field{
	Type:        graphqlType.PowerControlNodeType,
	Description: "Off node",
	Args: graphql.FieldConfigArgument{
		"uuids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"force_off": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.OffNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / off_node: " + err.Error())
		}
		return data, err
	},
}

var ForceRestartNode = graphql.Field{
	Type:        graphqlType.PowerControlNodeType,
	Description: "Force restart node",
	Args: graphql.FieldConfigArgument{
		"uuids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.ForceRestartNode(params.Args)
		if err != nil {
			logger.Logger.Println("flute / force_restart_node: " + err.Error())
		}
		return data, err
	},
}

var CreateNode = graphql.Field{
	Type:        graphqlType.NodeType,
	Description: "Create new node",
	Args: graphql.FieldConfigArgument{
		"node_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"bmc_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nic_speed_mbps": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nic_detail_data": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"ipmi_user_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"ipmi_user_password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateNode(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("flute / create_node: " + err.Error())
		}
		return data, err
	},
}

var UpdateNode = graphql.Field{
	Type:        graphqlType.NodeType,
	Description: "Update node",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"node_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"node_num": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"node_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"bmc_mac_addr": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"bmc_ip": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pxe_mac_addr": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cpu_cores": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"memory": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"nic_speed_mbps": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"active": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"ipmi_user_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"ipmi_user_password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.UpdateNode(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("flute / update_node: " + err.Error())
		}
		return data, err
	},
}

var DeleteNode = graphql.Field{
	Type:        graphqlType.NodeType,
	Description: "Delete node by uuid",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.DeleteNode(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("flute / delete_node: " + err.Error())
		}
		return data, err
	},
}

var CreateNodeDetail = graphql.Field{
	Type:        graphqlType.NodeDetailType,
	Description: "Create new node_detail",
	Args: graphql.FieldConfigArgument{
		"node_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"node_detail_data": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateNodeDetail(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("flute / create_node_detail: " + err.Error())
		}
		return data, err
	},
}

var UpdateNodeDetail = graphql.Field{
	Type:        graphqlType.NodeDetailType,
	Description: "Update node_detail by node_uuid",
	Args: graphql.FieldConfigArgument{
		"node_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"node_detail_data": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.UpdateNodeDetail(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("flute / update_node_detail: " + err.Error())
		}
		return data, err
	},
}

var DeleteNodeDetail = graphql.Field{
	Type:        graphqlType.NodeDetailType,
	Description: "Delete node_detail by node_uuid",
	Args: graphql.FieldConfigArgument{
		"node_uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.DeleteNodeDetail(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("flute / delete_node_detail: " + err.Error())
		}
		return data, err
	},
}
