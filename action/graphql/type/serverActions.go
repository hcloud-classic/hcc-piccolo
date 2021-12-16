package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// ServerActionType : Graphql object type of ServerAction
var ServerActionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerAction",
		Fields: graphql.Fields{
			"action": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"err_str": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// ServerActionsType : Graphql object type of ServerActions
var ServerActionsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerActions",
		Fields: graphql.Fields{
			"server_actions": &graphql.Field{
				Type: graphql.NewList(ServerActionType),
			},
			"number": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
