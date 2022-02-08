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
			return model.TimpaniService{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isAdmin && !isMaster {
			// params.Args["group_id"] = int(groupID)
			logger.Logger.Println("timpani / Not administrator: " + err.Error())
		}
		data, err := queryparser.TimapniServiceController(params.Args)
		if err != nil {
			logger.Logger.Println("timpani / timpani service control: " + err.Error())
		}
		return data, err
	},
}

var TimpaniMasterSync = graphql.Field{
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
	Description: "Timpani MasterSync API",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"newpw": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			// return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
			retData := model.MasterSync{}
			retData.Data.Errors.Errcode = ""
			retData.Data.Errors.Errmsg = err.Error() + " Token Error"
			return retData, nil
		}
		if !isMaster {
			// params.Args["group_id"] = int(groupID)
			logger.Logger.Println("timpani / Not administrator: " + err.Error())
			retData := model.MasterSync{}
			retData.Data.Errors.Errcode = ""
			retData.Data.Errors.Errmsg = err.Error() + " Not Master"
			return retData, nil
		}
		data, err := queryparser.TimpaniMasterSync(params.Args)
		if err != nil {
			logger.Logger.Println("timpani / master pass sync: " + err.Error())
		}
		return data, err
	},
}

var TimpaniBackup = graphql.Field{
	Type:        graphqlType.TimpaniMasterSync,
	Description: "Timpani Volume Backup API",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"usetype": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"nodetype": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			// return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
			retData := model.MasterSync{}
			retData.Data.Errors.Errcode = ""
			retData.Data.Errors.Errmsg = err.Error() + " Token Error"
			return retData, nil
		}
		params.Args["group_id"] = int(groupID)
		data, err := queryparser.TimpaniBackup(params.Args)
		if err != nil {
			logger.Logger.Println("timpani / volume backup: " + err.Error())
		}
		return data, err
	},
}

var Restore = graphql.Field{
	Type:        graphqlType.RestoreType,
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
				RunStatus: "",
				RunUUID:   "",
				Errors: model.Error{
					ErrMsg:  "Token Check Failed",
					ErrCode: "1003",
				},
			}, nil
		}

		return queryparser.Restore(params.Args)
	},
}
