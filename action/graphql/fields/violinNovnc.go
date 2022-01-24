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

var CotrolVNC = graphql.Field{
	Type:        graphqlType.VncPortType,
	Description: "Control VNC",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"action": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.VncPort{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ControlVnc(params.Args)
		if err != nil {
			logger.Logger.Println("violin-novnc / control_vnc: " + err.Error())
		}
		return data, err
	},
}
