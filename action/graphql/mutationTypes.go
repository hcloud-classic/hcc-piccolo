package graphql

import (
	"hcc/piccolo/action/graphql/mutationparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/userTool"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// piccolo
		"signup": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Execute user sign up process for piccolo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: piccolo / signup")
				return mutationparser.SignUp(params.Args)
			},
		},
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: violin / create_server")
				return mutationparser.CreateServer(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: violin / update_server")
				return mutationparser.UpdateServer(params.Args)
			},
		},
		"delete_server": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: violin / delete_server")
				return mutationparser.DeleteServer(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: violin / create_server_node")
				return mutationparser.CreateServerNode(params.Args)
			},
		},
		"delete_server_node": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: violin / delete server_node")
				return mutationparser.DeleteServerNode(params.Args)
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				return mutationparser.CreateSubnet(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: harp / update_subnet")
				return mutationparser.UpdateSubnet(params.Args)
			},
		},
		"delete_subnet": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: harp / delete_subnet")
				return mutationparser.DeleteSubnet(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: harp / create_dhcpd_conf")
				return mutationparser.CreateDHCPDConf(params.Args)
			},
		},
		"create_adaptiveip_setting": &graphql.Field{
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
				"start_ip_address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"end_ip_address": &graphql.ArgumentConfig{
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
				logger.Logger.Println("Resolving: harp / create_adaptiveip_setting")
				return mutationparser.CreateAdaptiveIPSetting(params.Args)
			},
		},
		"create_adaptiveip_server": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: harp / create_adaptiveip_server")
				return mutationparser.CreateAdaptiveIPServer(params.Args)
			},
		},
		"delete_adaptiveip_server": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: harp / delete_adaptiveip_server")
				return mutationparser.DeleteAdaptiveIPServer(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / on_node")
				return mutationparser.OnNode(params.Args)
			},
		},
		"off_node": &graphql.Field{
			Type:        graphql.String,
			Description: "Off node",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"force_off": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
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
				logger.Logger.Println("Resolving: flute / off_node")
				return mutationparser.OffNode(params.Args)
			},
		},
		"force_restart_node": &graphql.Field{
			Type:        graphql.String,
			Description: "Force restart node",
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
				logger.Logger.Println("Resolving: flute / force_restart_node")
				return mutationparser.ForceRestartNode(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / create_node")
				return mutationparser.CreateNode(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / update_node")
				return mutationparser.UpdateNode(params.Args)
			},
		},
		"delete_node": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / delete_node")
				return mutationparser.DeleteNode(params.Args)
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / create_node_detail")
				return mutationparser.CreateNodeDetail(params.Args)
			},
		},
		"delete_node_detail": &graphql.Field{
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
				err := userTool.ValidateToken(params.Args)
				if err != nil {
					return nil, err
				}
				logger.Logger.Println("Resolving: flute / delete_node_detail")
				return mutationparser.DeleteNodeDetail(params.Args)
			},
		},
	},
})
