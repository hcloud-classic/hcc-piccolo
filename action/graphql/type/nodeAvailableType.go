package graphqltype

import "github.com/graphql-go/graphql"

// NodeAvailableType : Graphql object type of NodeAvailable
var NodeAvailableType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeAvailable",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type: ResourceType,
			},
			"available": &graphql.Field{
				Type: ResourceType,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
