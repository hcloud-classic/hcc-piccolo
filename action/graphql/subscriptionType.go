package graphql

import (
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
	"innogrid.com/hcloud-classic/hcc_errors"
)

var subscriptionTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			// piccolo
			"list_user": &graphql.Field{
				Type:        graphqlType.UserListType,
				Description: "Get the user list from piccolo",
				Args: graphql.FieldConfigArgument{
					"group_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"authentication": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"group_name": &graphql.ArgumentConfig{
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
					_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					if !isMaster {
						params.Args["group_id"] = int(groupID)
					}
					data, err := queryparser.UserList(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / list_user (Subscription): " + err.Error())
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
					_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.ResourceUsage{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					if !isMaster {
						params.Args["group_id"] = int(groupID)
					}
					data, err := queryparser.ResourceUsage(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / resource_usage (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			"server_alarm_list": &graphql.Field{
				Type:        graphqlType.ServerAlarmsType,
				Description: "Get the server's alarm list",
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"user_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"reason": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"detail": &graphql.ArgumentConfig{
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
					_, _, _, _, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.ServerAlarms{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := dao.ShowServerAlarms(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / server_alarm_list (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			"num_unread_server_alarm": &graphql.Field{
				Type:        graphqlType.ServerAlarmsNumType,
				Description: "Get the number of unread server's alarm",
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					_, _, _, _, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.ServerAlarmsNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := dao.ShowUnreadServerAlarmsNum(params.Args)
					if err != nil {
						logger.Logger.Println("piccolo / num_unread_server_alarm (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// violin
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
						logger.Logger.Println("violin / list_server (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			"all_server": &graphql.Field{
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
						logger.Logger.Println("violin / all_server (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// harp
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
					"subnet_name": &graphql.ArgumentConfig{
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
					_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					if !isMaster {
						params.Args["group_id"] = int(groupID)
					}
					data, err := queryparser.ListSubnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / list_subnet (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			"all_subnet": &graphql.Field{
				Type:        graphqlType.SubnetListType,
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
					_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					if !isMaster {
						params.Args["group_id"] = int(groupID)
					}
					data, err := queryparser.AllSubnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / all_subnet (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// flute
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
						logger.Logger.Println("flute / list_node (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// piano
			"telegraf": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "telegraf subscription",
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
					_, _, _, _, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.Telegraf{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.Telegraf(params.Args)
					if err != nil {
						logger.Logger.Println("piano / telegraf (Subscription): " + err.Error())
					}

					return data, err
				},
			},
			// tuba
			"all_task": &graphql.Field{
				Type:        graphqlType.TaskListResultType,
				Description: "all_task subscription",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"server_uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
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
					_, _, _, _, err := usertool.ValidateToken(params.Args, false)
					if err != nil {
						return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllTask(params.Args)
					if err != nil {
						logger.Logger.Println("tuba / all_task (Subscription): " + err.Error())
					}

					return data, err
				},
			},
		},
	})
