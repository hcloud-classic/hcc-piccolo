package graphql

import (
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"

	"github.com/graphql-go/graphql"
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// piccolo
			"login": &graphql.Field{
				Type:        graphqlType.Token,
				Description: "Execute login process for piccolo",
				Args: graphql.FieldConfigArgument{
					"group_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					data, err := queryparser.Login(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / login: " + err.Error())
					}
					return data, err
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
						return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.User(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / user: " + err.Error())
					}
					return data, err
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
						return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.UserList(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / list_user: " + err.Error())
					}
					return data, err
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
						return model.UserNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumUser()
					if err != nil {
						logger.Logger.Println("piccolo / num_user: " + err.Error())
					}
					return data, err
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
					data, err := queryparser.CheckToken(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / check_token: " + err.Error())
					}
					return data, err
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
						return model.ResourceUsage{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ResourceUsage()
					if err != nil {
						logger.Logger.Println("piccolo / resource_usage: " + err.Error())
					}
					return data, err
				},
			},
			"server_log": &graphql.Field{
				Type:        graphqlType.ServerActionsType,
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
						return model.ServerActions{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := dao.ShowServerActions(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / server_log: " + err.Error())
					}
					return data, err
				},
			},
			"num_server_log": &graphql.Field{
				Type:        graphqlType.ServerActionsNumType,
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
						return model.ServerActionsNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := dao.ShowServerActionsNum(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / num_server_log: " + err.Error())
					}
					return data, err
				},
			},
			// violin
			"server": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.Server(params.Args)
					if err != nil {
						logger.Logger.Println("violin / server: " + err.Error())
					}
					return data, err
				},
			},
			"list_server": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListServer(params.Args)
					if err != nil {
						logger.Logger.Println("violin / list_server: " + err.Error())
					}
					return data, err
				},
			},
			"all_server": &graphql.Field{
				Type:        graphqlType.ServerListType,
				Description: "Get all server list",
				Args: graphql.FieldConfigArgument{
					"group_id": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllServer(params.Args)
					if err != nil {
						logger.Logger.Println("violin / all_server: " + err.Error())
					}
					return data, err
				},
			},
			"num_server": &graphql.Field{
				Type:        graphqlType.ServerNumType,
				Description: "Get the number of servers",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumServer()
					if err != nil {
						logger.Logger.Println("violin / num_server: " + err.Error())
					}
					return data, err
				},
			},
			"server_node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ServerNode(params.Args)
					if err != nil {
						logger.Logger.Println("violin / server_node: " + err.Error())
					}
					return data, err
				},
			},
			"list_server_node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListServerNode(params.Args)
					if err != nil {
						logger.Logger.Println("violin / list_server_node: " + err.Error())
					}
					return data, err
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
						return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllServerNode(params.Args)
					if err != nil {
						logger.Logger.Println("violin / all_server_node: " + err.Error())
					}
					return data, err
				},
			},
			"num_nodes_server": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerNodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumServerNode(params.Args)
					if err != nil {
						logger.Logger.Println("violin / num_nodes_server: " + err.Error())
					}
					return data, err
				},
			},
			// vnc
			"control_vnc": &graphql.Field{
				Type:        graphqlType.VncPortType,
				Description: "Control VNC",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"action": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.VncPort{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ControlVnc(params.Args)
					if err != nil {
						logger.Logger.Println("violin-novnc / control_vnc: " + err.Error())
					}
					return data, err
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
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.Subnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / subnet: " + err.Error())
					}
					return data, err
				},
			},
			"list_subnet": &graphql.Field{
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
					"name": &graphql.ArgumentConfig{
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
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListSubnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / list_subnet : " + err.Error())
					}
					return data, err
				},
			},
			"all_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
				Description: "Get all subnet list",
				Args: graphql.FieldConfigArgument{
					"group_id": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllSubnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / all_subnet: " + err.Error())
					}
					return data, err
				},
			},
			"available_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
				Description: "Get available subnet list",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AvailableSubnetList()
					if err != nil {
						logger.Logger.Println("harp / available_subnet: " + err.Error())
					}
					return data, err
				},
			},
			"num_subnet": &graphql.Field{
				Type:        graphqlType.SubnetNumType,
				Description: "Get the number of subnets",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumSubnet()
					if err != nil {
						logger.Logger.Println("harp / num_subnet: " + err.Error())
					}
					return data, err
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
						return model.AdaptiveIPAvailableIPList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.GetAdaptiveIPAvailableIPList()
					if err != nil {
						logger.Logger.Println("harp / adaptiveip_available_ip_list: " + err.Error())
					}
					return data, err
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
						return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.GetAdaptiveIPSetting()
					if err != nil {
						logger.Logger.Println("harp / adaptiveip_setting: " + err.Error())
					}
					return data, err
				},
			},
			"adaptiveip_server": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AdaptiveIPServer(params.Args)
					if err != nil {
						logger.Logger.Println("harp / adaptiveip_server: " + err.Error())
					}
					return data, err
				},
			},
			"list_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerListType,
				Description: "Get adaptiveip_server list",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"group_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
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
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListAdaptiveIPServer(params.Args)
					if err != nil {
						logger.Logger.Println("harp / list_adaptiveip_server: " + err.Error())
					}
					return data, err
				},
			},
			"all_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerListType,
				Description: "Get all adaptiveip_server list",
				Args: graphql.FieldConfigArgument{
					"group_id": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllAdaptiveIPServer(params.Args)
					if err != nil {
						logger.Logger.Println("harp / all_adaptiveip_server: " + err.Error())
					}
					return data, err
				},
			},
			"num_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerNumType,
				Description: "Get the number of AdaptiveIP Servers",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.AdaptiveIPServerNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumAdaptiveIPServer()
					if err != nil {
						logger.Logger.Println("harp / num_adaptiveip_server: " + err.Error())
					}
					return data, err
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
						return model.PowerStateNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.PowerStateNode(params.Args)
					if err != nil {
						logger.Logger.Println("flute / power_state_node: " + err.Error())
					}
					return data, err
				},
			},
			"node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.Node(params.Args)
					if err != nil {
						logger.Logger.Println("flute / node: " + err.Error())
					}
					return data, err
				},
			},
			"list_node": &graphql.Field{
				Type:        graphqlType.NodeListType,
				Description: "Get node list",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"group_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
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
					"charge_cpu": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"charge_memory": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"charge_nic": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListNode(params.Args)
					if err != nil {
						logger.Logger.Println("flute / list_node: " + err.Error())
					}
					return data, err
				},
			},
			"all_node": &graphql.Field{
				Type:        graphqlType.NodeListType,
				Description: "Get all node list",
				Args: graphql.FieldConfigArgument{
					"group_id": &graphql.ArgumentConfig{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllNode(params.Args)
					if err != nil {
						logger.Logger.Println("flute / all_node: " + err.Error())
					}
					return data, err
				},
			},
			"num_node": &graphql.Field{
				Type:        graphqlType.NodeNumType,
				Description: "Get the number of nodes",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NumNode()
					if err != nil {
						logger.Logger.Println("flute / num_node: " + err.Error())
					}
					return data, err
				},
			},
			"detail_node": &graphql.Field{
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.NodeDetail(params.Args)
					if err != nil {
						logger.Logger.Println("flute / node_detail: " + err.Error())
					}
					return data, err
				},
			},
			// piano
			"telegraf": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "Get all cpu usage data",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"metric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"subMetric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"period": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aggregateFn": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
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
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Telegraf{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.Telegraf(params.Args)
					if err != nil {
						logger.Logger.Println("piano / telegraf: " + err.Error())
					}

					return data, err
				},
			},
			"billing_data": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "Get the billing data",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"group_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"billing_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"date_start": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"date_end": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.GetBillingData(params.Args)
					if err != nil {
						logger.Logger.Println("piano / billing_data: " + err.Error())
					}

					return data, err
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
						return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.GetVolumeList(params.Args)
					if err != nil {
						logger.Logger.Println("cello / volume_list: " + err.Error())
					}
					return data, err
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
						return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.GetPoolList(params.Args)
					if err != nil {
						logger.Logger.Println("cello / pool_list: " + err.Error())
					}
					return data, err
				},
			},
			// tuba
			"all_task": &graphql.Field{
				Type:        graphqlType.TaskListResultType,
				Description: "all_task",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_port": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"sort_by": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"reverse_sorting": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"hide_threads": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllTask(params.Args)
					if err != nil {
						logger.Logger.Println("tuba / all_task: " + err.Error())
					}

					return data, err
				},
			},
		},
	})
