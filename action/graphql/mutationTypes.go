package graphql

import (
<<<<<<< HEAD
	"hcc/piccolo/action/graphql/mutationParser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/logger"
=======
	"hcc/piccolo/action/graphql/mutationparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
<<<<<<< HEAD
=======
		// piccolo
		"signup": &graphql.Field{
			Type:        graphqlType.UserType,
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
		"unregister": &graphql.Field{
			Type:        graphqlType.UserType,
			Description: "Execute user unregister process for piccolo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := usertool.ValidateTokenForAdmin(params.Args)
				if err != nil {
					return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: piccolo / unregister")
				return mutationparser.Unregister(params.Args)
			},
		},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / create_server")
				return mutationParser.CreateServer(params.Args)
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
				logger.Logger.Println("Resolving: violin / create_server")
				return mutationparser.CreateServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / update_server")
				return mutationParser.UpdateServer(params.Args)
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
				logger.Logger.Println("Resolving: violin / update_server")
				return mutationparser.UpdateServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_server": &graphql.Field{
			Type:        graphqlType.ServerType,
			Description: "Delete server by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / delete_server")
				return mutationParser.DeleteServer(params.Args)
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
				logger.Logger.Println("Resolving: violin / delete_server")
				return mutationparser.DeleteServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: violin / create_server_node")
				return mutationParser.CreateServerNode(params.Args)
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
				logger.Logger.Println("Resolving: violin / create_server_node")
				return mutationparser.CreateServerNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_server_node": &graphql.Field{
			Type:        graphqlType.ServerNodeType,
			Description: "Delete server_node by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
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
=======
				"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
				logger.Logger.Println("Resolving: violin-novnc: create_vnc")
				return mutationParser.CreateVnc(params.Args)
=======
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: violin / delete server_node")
				return mutationparser.DeleteServerNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"leader_node_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
=======
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				"os": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"subnet_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return mutationParser.CreateSubnet(params.Args)
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
				logger.Logger.Println("Resolving: harp / create_subnet")
				return mutationparser.CreateSubnet(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / update_subnet")
				return mutationParser.UpdateSubnet(params.Args)
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
				logger.Logger.Println("Resolving: harp / update_subnet")
				return mutationparser.UpdateSubnet(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_subnet": &graphql.Field{
			Type:        graphqlType.SubnetType,
			Description: "Delete subnet by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / delete_subnet")
				return mutationParser.DeleteSubnet(params.Args)
			},
		},
		"create_dhcpd_conf": &graphql.Field{
			Type:        graphql.String,
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
				logger.Logger.Println("Resolving: harp / delete_subnet")
				return mutationparser.DeleteSubnet(params.Args)
			},
		},
		"create_dhcpd_conf": &graphql.Field{
			Type:        graphqlType.CreateDHCPConfResultType,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			Description: "Create new dhcpd config",
			Args: graphql.FieldConfigArgument{
				"subnet_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"node_uuids": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
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
=======
				"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
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
=======
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.CreateDHCPConfResult{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					Type: graphql.String,
				},
				"netmask": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
				"gateway": &graphql.ArgumentConfig{
=======
				"gateway_address": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					Type: graphql.String,
				},
				"start_ip_address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"end_ip_address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
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
=======
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.AdaptiveIPSetting{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: harp / create_adaptiveip_setting")
				return mutationparser.CreateAdaptiveIPSetting(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"create_adaptiveip_server": &graphql.Field{
			Type:        graphqlType.AdaptiveIPServerType,
			Description: "Create new adaptiveip_server",
			Args: graphql.FieldConfigArgument{
<<<<<<< HEAD
				"adaptiveip_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
=======
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"public_ip": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / create_adaptiveip_server")
				return mutationParser.CreateAdaptiveIPServer(params.Args)
=======
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: harp / create_adaptiveip_server")
				return mutationparser.CreateAdaptiveIPServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_adaptiveip_server": &graphql.Field{
			Type:        graphqlType.AdaptiveIPServerType,
			Description: "Delete adaptiveip_server by server_uuid",
			Args: graphql.FieldConfigArgument{
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: harp / delete_adaptiveip_server")
				return mutationParser.DeleteAdaptiveIPServer(params.Args)
=======
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: harp / delete_adaptiveip_server")
				return mutationparser.DeleteAdaptiveIPServer(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		// flute
		"on_node": &graphql.Field{
<<<<<<< HEAD
			Type:        graphql.String,
			Description: "On node",
			Args: graphql.FieldConfigArgument{
				"mac": &graphql.ArgumentConfig{
=======
			Type:        graphqlType.PowerControlNodeType,
			Description: "On node",
			Args: graphql.FieldConfigArgument{
				"uuids": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"token": &graphql.ArgumentConfig{
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
<<<<<<< HEAD
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
=======
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: flute / on_node")
				return mutationparser.OnNode(params.Args)
			},
		},
		"off_node": &graphql.Field{
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
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: flute / off_node")
				return mutationparser.OffNode(params.Args)
			},
		},
		"force_restart_node": &graphql.Field{
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
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: flute / force_restart_node")
				return mutationparser.ForceRestartNode(params.Args)
			},
		},
		"create_node": &graphql.Field{
			Type:        graphqlType.NodeType,
			Description: "Create new node",
			Args: graphql.FieldConfigArgument{
				"bmc_ip": &graphql.ArgumentConfig{
					Type: graphql.String,
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
<<<<<<< HEAD
				"active": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / create_node")
				return mutationParser.CreateNode(params.Args)
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
				logger.Logger.Println("Resolving: flute / create_node")
				return mutationparser.CreateNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / update_node")
				return mutationParser.UpdateNode(params.Args)
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
				logger.Logger.Println("Resolving: flute / update_node")
				return mutationparser.UpdateNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_node": &graphql.Field{
			Type:        graphqlType.NodeType,
			Description: "Delete node by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / delete_node")
				return mutationParser.DeleteNode(params.Args)
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
				logger.Logger.Println("Resolving: flute / delete_node")
				return mutationparser.DeleteNode(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / create_node_detail")
				return mutationParser.CreateNodeDetail(params.Args)
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
				logger.Logger.Println("Resolving: flute / create_node_detail")
				return mutationparser.CreateNodeDetail(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
		"delete_node_detail": &graphql.Field{
			Type:        graphqlType.NodeDetailType,
			Description: "Delete node_detail by node_uuid",
			Args: graphql.FieldConfigArgument{
				"node_uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
<<<<<<< HEAD
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: flute / delete_node_detail")
				return mutationParser.DeleteNodeDetail(params.Args)
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
				logger.Logger.Println("Resolving: flute / delete_node_detail")
				return mutationparser.DeleteNodeDetail(params.Args)
			},
		},

		//Cello
		// volume DB
		"volume_handle": &graphql.Field{
			Type:        graphqlType.VolumeType,
			Description: "Create new volume",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"filesystem": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"network_ip": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"use_type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"user_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"gateway_ip": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"lun_num": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"pool": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"action": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				err := usertool.ValidateToken(params.Args)
				if err != nil {
					return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				logger.Logger.Println("Resolving: cello / createVolume")

				return mutationparser.VolumeHandle(params.Args)
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
	},
})
