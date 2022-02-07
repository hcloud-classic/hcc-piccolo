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
