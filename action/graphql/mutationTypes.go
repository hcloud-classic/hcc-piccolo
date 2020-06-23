package graphql

import (
	"hcc/piccolo/action/graphql/mutationParser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// violin
		"create_server": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Create new server",
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
				"user_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"nr_node": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / create_server")
				return mutationParser.CreateServer(params.Args)
			},
		},
		"update_server": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Update server",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / update_server")
				return mutationParser.UpdateServer(params.Args)
			},
		},
		"delete_server": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Delete server by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / delete_server")
				return mutationParser.DeleteServer(params.Args)
			},
		},
		"create_server_node": &graphql.Field{
			Type:        graphqlType.ServerNodeType,
			Description: "Create new server_node",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"node_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / create_server_node")
				return mutationParser.CreateServerNode(params.Args)
			},
		},
		"delete_server_node": &graphql.Field{
			Type:        graphqlType.ServerNodeType,
			Description: "Delete server_node by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / delete server_node")
				return mutationParser.DeleteServerNode(params.Args)
			},
		},
		// vnc
		"create_vnc": &graphql.Field{
			Type:        graphqlType.VncNodeType,
			Description: "Create vnc",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
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
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin-novnc: create_vnc")
				return mutationParser.CreateVnc(params.Args)
			},
		},
		// harp
		"create_subnet": &graphql.Field{
			Type:        graphqlType.SubnetType,
			Description: "Create new subnet",
			Args: graphql.FieldConfigArgument{
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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return mutationParser.CreateSubnet(params.Args)
			},
		},
		"update_subnet": &graphql.Field{
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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / update_subnet")
				return mutationParser.UpdateSubnet(params.Args)
			},
		},
		"delete_subnet": &graphql.Field{
			Type:        graphqlType.SubnetType,
			Description: "Delete subnet by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / delete_subnet")
				return mutationParser.DeleteSubnet(params.Args)
			},
		},
		"create_dhcpd_conf": &graphql.Field{
			Type:        graphql.String,
			Description: "Create new dhcpd config",
			Args: graphql.FieldConfigArgument{
				"subnet_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"node_uuids": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / create_dhcpd_conf")
				return mutationParser.CreateDHCPDConf(params.Args)
			},
		},
		"create_adaptiveip": &graphql.Field{
			Type:        graphqlType.AdaptiveIPType,
			Description: "Create new adaptiveip",
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
				logger.Logger.Println("Resolving: harp / create_adaptiveip")
				return mutationParser.CreateAdaptiveIP(params.Args)
			},
		},
		"update_adaptiveip": &graphql.Field{
			Type:        graphqlType.AdaptiveIPType,
			Description: "Update adaptiveip",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
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
				logger.Logger.Println("Resolving: harp / update_adaptiveip")
				return mutationParser.UpdateAdaptiveIP(params.Args)
			},
		},
		"delete_adaptiveip": &graphql.Field{
			Type:        graphqlType.AdaptiveIPType,
			Description: "Delete adaptiveip by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / delete_subnet")
				return mutationParser.DeleteAdaptiveIP(params.Args)
			},
		},
		"create_adaptiveip_server": &graphql.Field{
			Type:        graphqlType.AdaptiveIPServerType,
			Description: "Create new adaptiveip_server",
			Args: graphql.FieldConfigArgument{
				"adaptiveip_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"public_ip": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / create_adaptiveip_server")
				return mutationParser.CreateAdaptiveIPServer(params.Args)
			},
		},
		"delete_adaptiveip_server": &graphql.Field{
			Type:        graphqlType.AdaptiveIPServerType,
			Description: "Delete adaptiveip_server by server_uuid",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / delete_adaptiveip_server")
				return mutationParser.DeleteAdaptiveIPServer(params.Args)
			},
		},
		// flute
		"on_node": &graphql.Field{
			Type:        graphql.String,
			Description: "On node",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / on_node")
				return mutationParser.OnNode(params.Args)
			},
		},
		"create_node": &graphql.Field{
			Type:        graphqlType.NodeType,
			Description: "Create new node",
			Args: graphql.FieldConfigArgument{
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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / create_node")
				return mutationParser.CreateNode(params.Args)
			},
		},
		"update_node": &graphql.Field{
			Type:        graphqlType.NodeType,
			Description: "Update node",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
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
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"active": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / update_node")
				return mutationParser.UpdateNode(params.Args)
			},
		},
		"delete_node": &graphql.Field{
			Type:        graphqlType.NodeType,
			Description: "Delete node by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / delete_node")
				return mutationParser.DeleteNode(params.Args)
			},
		},
		// node_detail DB
		"create_node_detail": &graphql.Field{
			Type:        graphqlType.NodeDetailType,
			Description: "Create new node_detail",
			Args: graphql.FieldConfigArgument{
				"node_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cpu_model": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cpu_processors": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"cpu_threads": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / create_node_detail")
				return mutationParser.CreateNodeDetail(params.Args)
			},
		},
		"delete_node_detail": &graphql.Field{
			Type:        graphqlType.NodeDetailType,
			Description: "Delete node_detail by node_uuid",
			Args: graphql.FieldConfigArgument{
				"node_uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / delete_node_detail")
				return mutationParser.DeleteNodeDetail(params.Args)
			},
		},
	},
})
