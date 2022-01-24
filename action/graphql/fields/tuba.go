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

var AllTask = graphql.Field{
	Type:        graphqlType.TaskListResultType,
	Description: "all_task",
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
			logger.Logger.Println("tuba / all_task: " + err.Error())
		}

		return data, err
	},
}
