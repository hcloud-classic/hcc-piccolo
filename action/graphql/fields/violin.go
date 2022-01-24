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

var Server = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Get server by uuid",
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
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.Server(params.Args)
		if err != nil {
			logger.Logger.Println("violin / server: " + err.Error())
		}
		return data, err
	},
}

var ListServer = graphql.Field{
	Type:        graphqlType.ServerListType,
	Description: "Get server list",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"subnet_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"os": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cpu": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"memory": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"disk_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_uuid": &graphql.ArgumentConfig{
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
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		if !isAdmin && !isMaster {
			params.Args["user_uuid"] = id
		}
		data, err := queryparser.ListServer(params.Args)
		if err != nil {
			logger.Logger.Println("violin / list_server: " + err.Error())
		}
		return data, err
	},
}

var AllServer = graphql.Field{
	Type:        graphqlType.ServerListType,
	Description: "Get all server list",
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
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		if !isAdmin && !isMaster {
			params.Args["user_uuid"] = id
		}
		data, err := queryparser.AllServer(params.Args)
		if err != nil {
			logger.Logger.Println("violin / all_server: " + err.Error())
		}
		return data, err
	},
}

var NumServer = graphql.Field{
	Type:        graphqlType.ServerNumType,
	Description: "Get the number of servers",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ServerNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NumServer(params.Args)
		if err != nil {
			logger.Logger.Println("violin / num_server: " + err.Error())
		}
		return data, err
	},
}

var ServerNode = graphql.Field{
	Type:        graphqlType.ServerNodeType,
	Description: "Get server_node by uuid",
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
			return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / server_node: " + err.Error())
		}
		return data, err
	},
}

var ListServerNode = graphql.Field{
	Type:        graphqlType.ServerNodeListType,
	Description: "Get server_node list",
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
			return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ListServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / list_server_node: " + err.Error())
		}
		return data, err
	},
}

var AllServerNode = graphql.Field{
	Type:        graphqlType.ServerNodeListType,
	Description: "Get all server_node list",
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
			return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.AllServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / all_server_node: " + err.Error())
		}
		return data, err
	},
}

var NumNodesServer = graphql.Field{
	Type:        graphqlType.ServerNumType,
	Description: "Get the number of nodes of server",
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
			return model.ServerNodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.NumServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / num_nodes_server: " + err.Error())
		}
		return data, err
	},
}

var CreateServer = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Create new server",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"subnet_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"os": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cpu": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"memory": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"disk_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nr_node": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := mutationparser.CreateServer(params.Args)
		if err != nil {
			logger.Logger.Println("violin / create_server: " + err.Error())
		}
		return data, err
	},
}

var UpdateServer = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Update server",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"server_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.UpdateServer(params.Args, isAdmin, isMaster, id)
		if err != nil {
			logger.Logger.Println("violin / update_server: " + err.Error())
		}
		return data, err
	},
}

var UpdateServerNodes = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Update nodes of the server",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"selected_nodes": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"trigger_alarm": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.UpdateServerNodes(params.Args, isAdmin, isMaster, id)
		if err != nil {
			logger.Logger.Println("violin / update_server_nodes: " + err.Error())
		}
		return data, err
	},
}

var ScaleUpServer = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Scale up the server",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.ScaleUpServer(params.Args, isAdmin, isMaster, id)
		if err != nil {
			logger.Logger.Println("violin / scale_up_server: " + err.Error())
		}
		return data, err
	},
}

var DeleteServer = graphql.Field{
	Type:        graphqlType.ServerType,
	Description: "Delete server by uuid",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := mutationparser.DeleteServer(params.Args, isAdmin, isMaster, id)
		if err != nil {
			logger.Logger.Println("violin / delete_server: " + err.Error())
		}
		return data, err
	},
}

var CreateServerNode = graphql.Field{
	Type:        graphqlType.ServerNodeType,
	Description: "Create new server_node",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
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
			return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / create_server_node: " + err.Error())
		}
		return data, err
	},
}

var DeleteServerNode = graphql.Field{
	Type:        graphqlType.ServerNodeType,
	Description: "Delete server_node by uuid",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.DeleteServerNode(params.Args)
		if err != nil {
			logger.Logger.Println("violin / delete server_node: " + err.Error())
		}
		return data, err
	},
}
