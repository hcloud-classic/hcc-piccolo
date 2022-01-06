package filed

import (
	"hcc/piccolo/action/graphql/mutationparser"
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
