package graphql

import (
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
	"github.com/hcloud-classic/hcc_errors"
)

var subscriptionTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			// piccolo
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
						logger.Logger.Println("piccolo / resource_usage (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// violin
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllServer(params.Args)
					if err != nil {
						logger.Logger.Println("violin / all_server (Subscription): " + err.Error())
					}
					return data, err
				},
			},
			// harp
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
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.AllSubnet(params.Args)
					if err != nil {
						logger.Logger.Println("harp / all_subnet (Subscription): " + err.Error())
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
					err := usertool.ValidateToken(params.Args)
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
			"list_task": &graphql.Field{
				Type:        graphqlType.TaskListType,
				Description: "list_task subscription",
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
					"show_children": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"show_threads": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"cmd": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"state": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"pid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ppid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"pgid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"sid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"priority": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"nice": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"num_threads": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"epm_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"epm_source": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"epm_target": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					data, err := queryparser.ListTask(params.Args)
					if err != nil {
						logger.Logger.Println("tuba / list_task (Subscription): " + err.Error())
					}

					return data, err
				},
			},
			"all_task": &graphql.Field{
				Type:        graphqlType.TaskListType,
				Description: "all_task subscription",
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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
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
