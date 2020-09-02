package graphql

import (
	"github.com/graphql-go/graphql"
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/userTool"
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / server")
					return queryparser.Server(params.Args)
				},
			},
			"list_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerType),
				Description: "Get server list",
				Args: graphql.FieldConfigArgument{
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / list_server")
					return queryparser.ListServer(params.Args)
				},
			},
			"all_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / all_server")
					return queryparser.AllServer(params.Args)
				},
			},
			"num_server": &graphql.Field{
				Type:        graphqlType.ServerNumType,
				Description: "Get the number of server",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / num_server")
					return queryparser.NumServer()
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / server_node")
					return queryparser.ServerNode(params.Args)
				},
			},
			"list_server_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerNodeType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / list_server_node")
					return queryparser.ListServerNode(params.Args)
				},
			},
			"all_server_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerNodeType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / all_server_node")
					return queryparser.AllServerNode(params.Args)
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin / num_nodes_server")
					return queryparser.NumServerNode(params.Args)
				},
			},
			// vnc
			"control_vnc": &graphql.Field{
				Type:        graphql.String,
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: violin-novnc: control_vnc")
					return queryparser.ControlVnc(params.Args)
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / subnet")
					return queryparser.Subnet(params.Args)
				},
			},
			"list_subnet": &graphql.Field{
				Type:        graphql.NewList(graphqlType.SubnetType),
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
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / list_subnet")
					return queryparser.ListSubnet(params.Args)
				},
			},
			"all_subnet": &graphql.Field{
				Type:        graphql.NewList(graphqlType.SubnetType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / all_subnet")
					return queryparser.AllSubnet(params.Args)
				},
			},
			"num_subnet": &graphql.Field{
				Type:        graphqlType.SubnetNumType,
				Description: "Get the number of subnet",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / num_subnet")
					return queryparser.NumSubnet()
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / adaptiveip_setting")
					return queryparser.GetAdaptiveIPSetting()
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / adaptiveip_server")
					return queryparser.AdaptiveIPServer(params.Args)
				},
			},
			"list_adaptiveip_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPServerType),
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
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / list_adaptiveip_server")
					return queryparser.ListAdaptiveIPServer(params.Args)
				},
			},
			"all_adaptiveip_server": &graphql.Field{
				Type:        graphql.NewList(graphqlType.AdaptiveIPServerType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / all_adaptiveip_server")
					return queryparser.AllAdaptiveIPServer(params.Args)
				},
			},
			"num_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerNumType,
				Description: "Get the number of adaptiveip_server",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: harp / num_adaptiveip_server")
					return queryparser.NumAdaptiveIPServer()
				},
			},
			// flute
			"power_state_node": &graphql.Field{
				Type:        graphql.String,
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: power_state_node")
					return queryparser.PowerStateNode(params.Args)
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: flute / node")
					return queryparser.Node(params.Args)
				},
			},
			"list_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.NodeType),
				Description: "Get node list",
				Args: graphql.FieldConfigArgument{
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: flute / list_node")
					return queryparser.ListNode(params.Args)
				},
			},
			"all_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.NodeType),
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: flute / all_node")
					return queryparser.AllNode(params.Args)
				},
			},
			"num_node": &graphql.Field{
				Type:        graphqlType.NodeNumType,
				Description: "Get the number of node",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: flute / num_node")
					return queryparser.NumNode()
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
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: flute / node_detail")
					return queryparser.NodeDetail(params.Args)
				},
			},
			// piano
			"telegraf": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "Get all cpu usage data",
				Args: graphql.FieldConfigArgument{
					"metric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"subMetric": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"period": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aggregateType": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := userTool.ValidateToken(params.Args)
					if err != nil {
						return nil, err
					}
					logger.Logger.Println("Resolving: piano / telegraf")
					return queryparser.Telegraf(params.Args)
				},
			},
		},
	})
