package graphql

import (
<<<<<<< HEAD
	"github.com/graphql-go/graphql"
	"hcc/piccolo/action/graphql/queryParser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"
=======
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/sqlite/serveractions"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
<<<<<<< HEAD
=======
			// piccolo
			"login": &graphql.Field{
				Type:        graphqlType.Token,
				Description: "Execute login process for piccolo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: piccolo / login")
					return queryparser.Login(params.Args)
				},
			},
			"user": &graphql.Field{
				Type:        graphqlType.UserType,
				Description: "Get the user list from piccolo",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / user")
					return queryparser.User(params.Args)
				},
			},
			"list_user": &graphql.Field{
				Type:        graphqlType.UserListType,
				Description: "Get the user list from piccolo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.UserList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / list_user")
					return queryparser.UserList(params.Args)
				},
			},
			"num_user": &graphql.Field{
				Type:        graphqlType.UserNumType,
				Description: "Get the number of users",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.UserNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / num_user")
					return queryparser.NumUser()
				},
			},
			"check_token": &graphql.Field{
				Type:        graphqlType.IsValid,
				Description: "Check validation of the token for piccolo",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: piccolo / check_token")
					return queryparser.CheckToken(params.Args)
				},
			},
			"resource_usage": &graphql.Field{
				Type:        graphqlType.ResourceUsageType,
				Description: "Get resource usage",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ResourceUsage{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / resource_usage")
					return queryparser.ResourceUsage()
				},
			},
			"server_log": &graphql.Field{
				Type:        serveractions.ServerActionsType,
				Description: "Get the server's log",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return serveractions.ServerActions{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / server_log")
					return serveractions.ShowServerActions(params.Args)
				},
			},
			"num_server_log": &graphql.Field{
				Type:        serveractions.ServerActionsNumType,
				Description: "Get the number of server's log",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return serveractions.ServerActionsNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piccolo / num_server_log")
					return serveractions.ShowServerActionsNum(params.Args)
				},
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			// violin
			"server": &graphql.Field{
				Type:        graphqlType.ServerType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / server")
					return queryParser.Server(params.Args)
				},
			},
			"list_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerType),
				Description: "Get server list",
				Args: graphql.FieldConfigArgument{
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / server")
					return queryparser.Server(params.Args)
				},
			},
			"list_server": &graphql.Field{
				Type:        graphqlType.ServerListType,
				Description: "Get server list",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / list_server")
					return queryParser.ListServer(params.Args)
				},
			},
			"all_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / list_server")
					return queryparser.ListServer(params.Args)
				},
			},
			"all_server": &graphql.Field{
				Type:        graphqlType.ServerListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get all server list",
				Args: graphql.FieldConfigArgument{
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / all_server")
					return queryParser.AllServer(params.Args)
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / all_server")
					return queryparser.AllServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"num_server": &graphql.Field{
				Type:        graphqlType.ServerNumType,
<<<<<<< HEAD
				Description: "Get the number of server",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / num_server")
					return queryParser.NumServer()
=======
				Description: "Get the number of servers",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / num_server")
					return queryparser.NumServer()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"server_node": &graphql.Field{
				Type:        graphqlType.ServerNodeType,
				Description: "Get server_node by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / server_node")
					return queryParser.ServerNode(params.Args)
				},
			},
			"list_server_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerNodeType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / server_node")
					return queryparser.ServerNode(params.Args)
				},
			},
			"list_server_node": &graphql.Field{
				Type:        graphqlType.ServerNodeListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get server_node list",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / list_server_node")
					return queryParser.ListServerNode(params.Args)
				},
			},
			"all_server_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerNodeType),
				Description: "Get all server_node list",
				Args:        graphql.FieldConfigArgument{},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / all_server_node")
					return queryParser.AllServerNode()
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / list_server_node")
					return queryparser.ListServerNode(params.Args)
				},
			},
			"all_server_node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / all_server_node")
					return queryparser.AllServerNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"num_nodes_server": &graphql.Field{
				Type:        graphqlType.ServerNumType,
				Description: "Get the number of nodes of server",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / num_nodes_server")
					return queryParser.NumNodesServer(params.Args)
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNodeNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin / num_nodes_server")
					return queryparser.NumServerNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			// vnc
			"control_vnc": &graphql.Field{
<<<<<<< HEAD
				Type:        graphqlType.VncNodeType,
				Description: "Create vnc",
=======
				Type:        graphqlType.VncPortType,
				Description: "Control VNC",
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
					"target_ip": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"target_port": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"target_pass": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"websocket_port": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"action": &graphql.ArgumentConfig{
=======
					"action": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
					logger.Logger.Println("Resolving: violin-novnc: control_vnc")
					return queryParser.ControlVnc(params.Args)
=======
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.VncPort{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: violin-novnc: control_vnc")
					return queryparser.ControlVnc(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			// harp
			"subnet": &graphql.Field{
				Type:        graphqlType.SubnetType,
				Description: "Get subnet by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / subnet")
					return queryParser.Subnet(params.Args)
				},
			},
			"list_subnet": &graphql.Field{
				Type:        graphql.NewList(graphqlType.SubnetType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / subnet")
					return queryparser.Subnet(params.Args)
				},
			},
			"list_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get subnet list",
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
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / list_subnet")
					return queryParser.ListSubnet(params.Args)
				},
			},
			"all_subnet": &graphql.Field{
				Type:        graphql.NewList(graphqlType.SubnetType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / list_subnet")
					return queryparser.ListSubnet(params.Args)
				},
			},
			"all_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get all subnet list",
				Args: graphql.FieldConfigArgument{
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / all_subnet")
					return queryParser.AllSubnet(params.Args)
				},
			},
			"num_subnet": &graphql.Field{
				Type:        graphqlType.SubnetNumType,
				Description: "Get the number of subnet",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / num_subnet")
					return queryParser.NumSubnet()
				},
			},
			"adaptiveip": &graphql.Field{
				Type:        graphqlType.AdaptiveIPType,
				Description: "Get adaptiveip by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / all_subnet")
					return queryparser.AllSubnet(params.Args)
				},
			},
			"available_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
				Description: "Get available subnet list",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
					logger.Logger.Println("Resolving: harp / adaptiveip")
					return queryParser.AdaptiveIP(params.Args)
				},
			},
			"list_adaptiveip": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPType),
				Description: "Get adaptiveip list",
				Args: graphql.FieldConfigArgument{
					"network_address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"netmask": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gateway": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"start_ip_address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"end_ip_address": &graphql.ArgumentConfig{
=======
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / available_subnet")
					return queryparser.AvailableSubnetList()
				},
			},
			"num_subnet": &graphql.Field{
				Type:        graphqlType.SubnetNumType,
				Description: "Get the number of subnets",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
					logger.Logger.Println("Resolving: harp / list_adaptiveip")
					return queryParser.ListAdaptiveIP(params.Args)
				},
			},
			"all_adaptiveip": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPType),
				Description: "Get all adaptiveip list",
				Args: graphql.FieldConfigArgument{
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / all_adaptiveip")
					return queryParser.AllAdaptiveIP(params.Args)
				},
			},
			"num_adaptiveip": &graphql.Field{
				Type:        graphqlType.AdaptiveIPNumType,
				Description: "Get the number of adaptiveip",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / num_adaptiveip")
					return queryParser.NumAdaptiveIP()
=======
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / num_subnet")
					return queryparser.NumSubnet()
				},
			},
			"adaptiveip_available_ip_list": &graphql.Field{
				Type:        graphqlType.AdaptiveIPAvailableIPListType,
				Description: "Get available ip list for adaptive ip",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPAvailableIPList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / adaptiveip_available_ip_list")
					return queryparser.GetAdaptiveIPAvailableIPList()
				},
			},
			"adaptiveip_setting": &graphql.Field{
				Type:        graphqlType.AdaptiveIPSettingType,
				Description: "Get settings of adaptiveip",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPSetting{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / adaptiveip_setting")
					return queryparser.GetAdaptiveIPSetting()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerType,
				Description: "Get adaptiveip by uuid",
				Args: graphql.FieldConfigArgument{
<<<<<<< HEAD
					"adaptiveip_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_uuid": &graphql.ArgumentConfig{
=======
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
					logger.Logger.Println("Resolving: harp / adaptiveip_server")
					return queryParser.AdaptiveIPServer(params.Args)
				},
			},
			"list_adaptiveip_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPServerType),
=======
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / adaptiveip_server")
					return queryparser.AdaptiveIPServer(params.Args)
				},
			},
			"list_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get adaptiveip_server list",
				Args: graphql.FieldConfigArgument{
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
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / list_adaptiveip_server")
					return queryParser.ListAdaptiveIPServer(params.Args)
				},
			},
			"all_adaptiveip_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPServerType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / list_adaptiveip_server")
					return queryparser.ListAdaptiveIPServer(params.Args)
				},
			},
			"all_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				Description: "Get all adaptiveip_server list",
				Args: graphql.FieldConfigArgument{
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / all_adaptiveip_server")
					return queryParser.AllAdaptiveIPServer(params.Args)
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / all_adaptiveip_server")
					return queryparser.AllAdaptiveIPServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"num_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerNumType,
<<<<<<< HEAD
				Description: "Get the number of adaptiveip_server",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / num_adaptiveip_server")
					return queryParser.NumAdaptiveIPServer()
				},
			},
			// flute
=======
				Description: "Get the number of AdaptiveIP Servers",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: harp / num_adaptiveip_server")
					return queryparser.NumAdaptiveIPServer()
				},
			},
			// flute
			"power_state_node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.PowerStateNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: power_state_node")
					return queryparser.PowerStateNode(params.Args)
				},
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			"node": &graphql.Field{
				Type:        graphqlType.NodeType,
				Description: "Get a node by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / node")
					return queryParser.Node(params.Args)
				},
			},
			"list_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.NodeType),
				Description: "Get node list",
				Args: graphql.FieldConfigArgument{
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: flute / node")
					return queryparser.Node(params.Args)
				},
			},
			"list_node": &graphql.Field{
				Type:        graphqlType.NodeListType,
				Description: "Get node list",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
=======
					"rack_number": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					"active": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / list_node")
					return queryParser.ListNode(params.Args)
				},
			},
			"all_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.NodeType),
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: flute / list_node")
					return queryparser.ListNode(params.Args)
				},
			},
			"all_node": &graphql.Field{
				Type:        graphqlType.NodeListType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / all_node")
					return queryParser.AllNode(params.Args)
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: flute / all_node")
					return queryparser.AllNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"num_node": &graphql.Field{
				Type:        graphqlType.NodeNumType,
<<<<<<< HEAD
				Description: "Get the number of node",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / num_node")
					return queryParser.NumNode()
=======
				Description: "Get the number of nodes",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: flute / num_node")
					return queryparser.NumNode()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			"detail_node": &graphql.Field{
				Type:        graphqlType.NodeDetailType,
				Description: "Get a node_detail by uuid",
				Args: graphql.FieldConfigArgument{
					"node_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / node_detail")
					return queryParser.NodeDetail(params.Args)
=======
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: flute / node_detail")
					return queryparser.NodeDetail(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
			// piano
			"telegraf": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "Get all cpu usage data",
				Args: graphql.FieldConfigArgument{
<<<<<<< HEAD
=======
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					"metric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"subMetric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"period": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
					"aggregateType": &graphql.ArgumentConfig{
=======
					"aggregateFn": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
<<<<<<< HEAD
					"uuid": &graphql.ArgumentConfig{
=======
					"time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"groupBy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"orderBy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
					logger.Logger.Println("Resolving: piano / telegraf")
					return queryParser.Telegraf(params.Args)
=======
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Telegraf{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: piano / telegraf")
					telegraf, err := queryparser.Telegraf(params.Args)

					Publisher()

					return telegraf, err
				},
			},
			// volume_list
			"volume_list": &graphql.Field{
				Type:        graphqlType.VolumeListType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"user_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: cello / volume_list")
					return queryparser.GetVolumeList(params.Args)
				},
			},
			// pool_list
			"pool_list": &graphql.Field{
				Type:        graphqlType.PoolListType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"user_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"row": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"action": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Resolving: cello / pool_list")
					return queryparser.GetPoolList(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
			},
		},
	})
