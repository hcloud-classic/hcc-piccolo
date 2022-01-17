package fields

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

var CreatePemKey = graphql.Field{
	Type:        graphqlType.PermissionKey,
	Description: "Create new Permission Key",
	Args: graphql.FieldConfigArgument{

		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}

		data, err := mutationparser.CreatePermissionKey(params.Args)
		if err != nil {
			logger.Logger.Println("violin / CreatePemKey: " + err.Error())
		}
		return data, err
	},
}

var GetPemKey = graphql.Field{
	Type:        graphqlType.PermissionKey,
	Description: "Get new Permission Key",
	Args: graphql.FieldConfigArgument{

		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}

		data, err := mutationparser.GetPermissionKey(params.Args)
		if err != nil {
			logger.Logger.Println("violin / GetPemKey: " + err.Error())
		}
		return data, err
	},
}
