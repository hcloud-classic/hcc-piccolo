package fields

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

var TimpaniServiceMgmt = graphql.Field{
	Type:        graphqlType.TimapniService,
	Description: "Timpani Agent Controller",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"target": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"action": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isAdmin && !isMaster {
			// params.Args["group_id"] = int(groupID)
			logger.Logger.Println("timpani / Not administrator: " + err.Error())
		}
		data, err := queryparser.TimapniServiceController(params.Args)
		if err != nil {
			logger.Logger.Println("timpani / available pool_list: " + err.Error())
		}
		return data, err
	},
}

var TimapniMasterSync = graphql.Field{
	Type: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "timpani",
			Fields: graphql.Fields{
				"isvaild": &graphql.Field{
					Type: graphql.Boolean,
				},
				"errors": &graphql.Field{
					Type: graphql.NewList(graphqlType.Errors),
				},
			},
		},
	),
	Description: "Timpani Agent Controller",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"newpw": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isAdmin && !isMaster {
			// params.Args["group_id"] = int(groupID)
			logger.Logger.Println("timpani / Not administrator: " + err.Error())
		}
		data, err := queryparser.TimapniServiceController(params.Args)
		if err != nil {
			logger.Logger.Println("timpani / available pool_list: " + err.Error())
		}
		return data, err
	},
}

var Restore = graphql.Field{
	Type:        graphqlType.SubnetType,
	Description: "Restore server",
	Args: graphql.FieldConfigArgument{
		"snapname": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"usetype": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nodetype": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"isboot": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Restore{
				RestoreInfo: model.RestoreInfo{
					RunStatus: "",
					RunUUID:   "",
					Errors: model.Error{
						ErrMsg:  "Token Check Failed",
						ErrCode: "1003",
					},
				},
			}, nil
		}

		return queryparser.Restore(params.Args)
	},
}
