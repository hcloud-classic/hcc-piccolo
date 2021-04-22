package graphql

import (
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
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
			"resource_usage": &graphql.Field{
				Type:        graphqlType.ResourceUsageType,
				Description: "Get resource usage",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err, _, isMaster, _, groupID := usertool.ValidateToken(params.Args, false)
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
					err, isAdmin, isMaster, id, groupID := usertool.ValidateToken(params.Args, false)
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
					err, _, isMaster, _, groupID := usertool.ValidateToken(params.Args, false)
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
					err, _, _, _, _ := usertool.ValidateToken(params.Args, false)
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
					err, _, _, _, _ := usertool.ValidateToken(params.Args, false)
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
