package fields

import (
	"hcc/piccolo/action/graphql/mutationparser"
	"hcc/piccolo/action/graphql/queryparser"
	"hcc/piccolo/action/graphql/queryparserext"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
	"innogrid.com/hcloud-classic/hcc_errors"
)

var VolumeHandle = graphql.Field{
	Type:        graphqlType.VolumeType,
	Description: "Create new volume",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_id": &graphql.ArgumentConfig{
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
		_, isMaster, userID, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
			params.Args["user_uuid"] = string(userID)
		}
		data, err := mutationparser.VolumeHandle(params.Args)
		if err != nil {
			logger.Logger.Println("cello / createVolume: " + err.Error())
		}
		return data, err
	},
}

var MountHandle = graphql.Field{
	Type:        graphqlType.VolumeType,
	Description: "Volume mount handler",
	Args: graphql.FieldConfigArgument{
		"uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_id": &graphql.ArgumentConfig{
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
		_, isMaster, userID, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
			params.Args["user_uuid"] = string(userID)
		}
		data, err := mutationparser.MountHandler(params.Args)
		if err != nil {
			logger.Logger.Println("cello / mountHandle: " + err.Error())
		}
		return data, err
	},
}

var VolumeList = graphql.Field{
	Type:        graphqlType.VolumeListType,
	Description: "Get server by uuid",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
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
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparserext.GetVolumeList(params.Args)
		if err != nil {
			logger.Logger.Println("cello / volume_list: " + err.Error())
		}
		return data, err
	},
}

var PoolList = graphql.Field{
	Type:        graphqlType.PoolListType,
	Description: "Get disk pool list",
	Args: graphql.FieldConfigArgument{
		"user_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"action": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.GetPoolList(params.Args)
		if err != nil {
			logger.Logger.Println("cello / pool_list: " + err.Error())
		}
		return data, err
	},
}

var AvailablePoolList = graphql.Field{
	Type:        graphqlType.PoolListType,
	Description: "Get available pool list",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
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
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.AvailablePoolList(params.Args)
		if err != nil {
			logger.Logger.Println("cello / available pool_list: " + err.Error())
		}
		return data, err
	},
}
