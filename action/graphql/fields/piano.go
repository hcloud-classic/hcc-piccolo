package fields

import (
	"hcc/piccolo/action/graphql/queryparser"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
	"strconv"

	"github.com/graphql-go/graphql"
	"innogrid.com/hcloud-classic/hcc_errors"
)

var Telegraf = graphql.Field{
	Type:        graphqlType.TelegrafType,
	Description: "Get all cpu usage data",
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
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Telegraf{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.Telegraf(params.Args)
		if err != nil {
			logger.Logger.Println("piano / telegraf: " + err.Error())
		}

		return data, err
	},
}

var BillingData = graphql.Field{
	Type:        graphqlType.BillingType,
	Description: "Get the billing data",
	Args: graphql.FieldConfigArgument{
		"group_ids": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"billing_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"date_start": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"date_end": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
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
		isAdmin, isMaster, _, loginGroupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_ids"] = strconv.Itoa(int(loginGroupID))
		}
		data, err := queryparser.GetBillingData(params.Args, isAdmin, isMaster, loginGroupID)
		if err != nil {
			logger.Logger.Println("piano / billing_data: " + err.Error())
		}

		return data, err
	},
}

var BillingDetail = graphql.Field{
	Type:        graphqlType.BillingType,
	Description: "Get the billingDetail data",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"billing_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"date_start": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, loginGroupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_ids"] = strconv.Itoa(int(loginGroupID))
		}
		data, err := queryparser.GetBillingDetail(params.Args, isAdmin, isMaster, loginGroupID)
		if err != nil {
			logger.Logger.Println("piano / billing_detail: " + err.Error())
		}

		return data, err
	},
}
