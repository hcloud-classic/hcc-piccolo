package graphql

import (
	"hcc/piccolo/action/graphql/mutationparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// piccolo
		"signup": &graphql.Field{
			Type:        graphqlType.UserType,
			Description: "Execute user sign up process for piccolo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"authentication": &graphql.ArgumentConfig{
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				if !isMaster {
					params.Args["group_id"] = int(groupID)
				}
				data, err := mutationparser.SignUp(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / signup: " + err.Error())
				}
				return data, err
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
				isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.Unregister(params.Args, isAdmin, isMaster, id, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / unregister: " + err.Error())
				}
				return data, err
			},
		},
		"update_user": &graphql.Field{
			Type:        graphqlType.UserType,
			Description: "Update user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authentication": &graphql.ArgumentConfig{
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
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.UpdateUser(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / update_user: " + err.Error())
				}
				return data, err
			},
		},
		"create_group": &graphql.Field{
			Type:        graphqlType.GroupType,
			Description: "Get the group info",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"group_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.CreateGroup(params.Args, isMaster)
				if err != nil {
					logger.Logger.Println("piccolo / create_group: " + err.Error())
				}
				return data, err
			},
		},
		"update_group": &graphql.Field{
			Type:        graphqlType.GroupType,
			Description: "Update the group info",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"group_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.UpdateGroup(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / update_group: " + err.Error())
				}
				return data, err
			},
		},
		"delete_group": &graphql.Field{
			Type:        graphqlType.GroupType,
			Description: "Delete the group info",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.DeleteGroup(params.Args, isMaster)
				if err != nil {
					logger.Logger.Println("piccolo / delete_group: " + err.Error())
				}
				return data, err
			},
		},
		"create_quota": &graphql.Field{
			Type:        graphqlType.QuotaType,
			Description: "Create quota",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"pool_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"ssd_size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"hdd_size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"subnet_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"adaptive_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"node_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				if !isMaster {
					params.Args["group_id"] = int(groupID)
				}
				data, err := dao.CreateQuota(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / create_quota: " + err.Error())
				}
				return data, err
			},
		},
		"update_quota": &graphql.Field{
			Type:        graphqlType.QuotaType,
			Description: "Update quota",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"pool_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"ssd_size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"hdd_size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"subnet_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"adaptive_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"node_cnt": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				if !isMaster {
					params.Args["group_id"] = int(groupID)
				}
				data, err := dao.UpdateQuota(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / update_quota: " + err.Error())
				}
				return data, err
			},
		},
		"delete_quota": &graphql.Field{
			Type:        graphqlType.QuotaType,
			Description: "Delete quota",
			Args: graphql.FieldConfigArgument{
				"group_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				if !isMaster {
					params.Args["group_id"] = int(groupID)
				}
				data, err := dao.DeleteQuota(params.Args, isAdmin, isMaster, int(groupID))
				if err != nil {
					logger.Logger.Println("piccolo / delete_quota: " + err.Error())
				}
				return data, err
			},
		},
		"delete_server_alarm": &graphql.Field{
			Type:        graphqlType.ServerAlarmType,
			Description: "Delete the server alarm",
			Args: graphql.FieldConfigArgument{
				"no": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"token": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, _, id, _, err := usertool.ValidateToken(params.Args, true)
				if err != nil {
					return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := dao.DeleteServerAlarm(params.Args, id)
				if err != nil {
					logger.Logger.Println("piccolo / delete_group: " + err.Error())
				}
				return data, err
			},
		},
		// violin
		"create_server": &graphql.Field{
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
		},
		"update_server": &graphql.Field{
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
		},
		"update_server_nodes": &graphql.Field{
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
		},
		// harp
		"create_subnet": &graphql.Field{
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
				if !isMaster {
					params.Args["group_id"] = int(groupID)
				}
				data, err := mutationparser.CreateSubnet(params.Args)
				if err != nil {
					logger.Logger.Println("harp / create_subnet: " + err.Error())
				}
				return data, err
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
		},
		"create_dhcpd_conf": &graphql.Field{
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
		},
		"create_port_forwarding": &graphql.Field{
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
		},
		"delete_port_forwarding": &graphql.Field{
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
		},
		// flute
		"on_node": &graphql.Field{
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
		},
		"create_node": &graphql.Field{
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
		},
		"update_node": &graphql.Field{
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
		},
		// node_detail DB
		"create_node_detail": &graphql.Field{
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
		},
		"update_node_detail": &graphql.Field{
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
				_, _, _, _, err := usertool.ValidateToken(params.Args, false)
				if err != nil {
					return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
				}
				data, err := mutationparser.VolumeHandle(params.Args)
				if err != nil {
					logger.Logger.Println("cello / createVolume: " + err.Error())
				}
				return data, err
			},
		},
	},
})
