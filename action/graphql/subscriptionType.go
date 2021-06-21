package graphql

import (
	"github.com/graphql-go/graphql"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
)

var subscriptionTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			// piano
			"telegraf": &graphql.Field{
				Type:        graphqlType.TelegrafType,
				Description: "telegraf subscription",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					err := usertool.ValidateToken(params.Args)
					if err != nil {
						return model.Telegraf{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
					}
					logger.Logger.Println("Subscription: piano / telegraf")
					return model.Telegraf{}, nil
				},
			},
		},
	})
