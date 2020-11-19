package graphqltype

import "github.com/graphql-go/graphql"

// VolumeNum : GraphQL type of VolumeNum
var VolumeNum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VolumeNum",
		Fields: graphql.Fields{
			"number": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
