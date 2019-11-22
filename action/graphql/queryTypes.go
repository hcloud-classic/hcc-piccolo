package graphql

import (
	"github.com/graphql-go/graphql"
	"hcc/piccolo/action/graphql/queryParser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// violin
			"server": &graphql.Field{
				Type:        graphqlType.ServerType,
				Description: "Get server by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / list_server")
					return queryParser.ListServer(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / all_server")
					return queryParser.AllServer(params.Args)
				},
			},
			"num_server": &graphql.Field{
				Type:        graphqlType.ServerNumType,
				Description: "Get the number of server",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / num_server")
					return queryParser.NumServer()
				},
			},
			"server_node": &graphql.Field{
				Type:        graphqlType.ServerNodeType,
				Description: "Get server_node by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: violin / server_node")
					return queryParser.ServerNode(params.Args)
				},
			},
			"list_server_node": &graphql.Field{
				Type:        graphql.NewList(graphqlType.ServerNodeType),
				Description: "Get server_node list",
				Args: graphql.FieldConfigArgument{
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / subnet")
					return queryParser.Subnet(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / list_subnet")
					return queryParser.ListSubnet(params.Args)
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
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
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
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
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
				},
			},
			"adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerType,
				Description: "Get adaptiveip by uuid",
				Args: graphql.FieldConfigArgument{
					"adaptiveip_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / adaptiveip_server")
					return queryParser.AdaptiveIPServer(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / list_adaptiveip_server")
					return queryParser.ListAdaptiveIPServer(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / all_adaptiveip_server")
					return queryParser.AllAdaptiveIPServer(params.Args)
				},
			},
			"num_adaptiveip_server": &graphql.Field{
				Type:        graphqlType.AdaptiveIPServerNumType,
				Description: "Get the number of adaptiveip_server",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: harp / num_adaptiveip_server")
					return queryParser.NumAdaptiveIPServer()
				},
			},
			// flute
			"node": &graphql.Field{
				Type:        graphqlType.NodeType,
				Description: "Get a node by uuid",
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / list_node")
					return queryParser.ListNode(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / all_node")
					return queryParser.AllNode(params.Args)
				},
			},
			"num_node": &graphql.Field{
				Type:        graphqlType.NodeNumType,
				Description: "Get the number of node",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / num_node")
					return queryParser.NumNode()
				},
			},
			"detail_node": &graphql.Field{
				Type:        graphqlType.NodeDetailType,
				Description: "Get a node_detail by uuid",
				Args: graphql.FieldConfigArgument{
					"node_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: flute / node_detail")
					return queryParser.NodeDetail(params.Args)
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					logger.Logger.Println("Resolving: piano / telegraf")
					return queryParser.Telegraf(params.Args)
				},
			},
		},
	})
