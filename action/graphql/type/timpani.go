package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// TimapniService : GraphQL type of TimapniService
var TimapniService = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "timpani",
		Fields: graphql.Fields{
			"target": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// TimpaniMasterSync : GraphQL type of TimpaniMasterSync
var TimpaniMasterSync = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "timpani",
		Fields: graphql.Fields{
			"target": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
