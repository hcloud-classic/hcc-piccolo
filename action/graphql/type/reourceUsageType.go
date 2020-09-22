package graphqltype

import "github.com/graphql-go/graphql"

// ResourceType : Graphql object type of Resource
var ResourceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Resource",
		Fields: graphql.Fields{
			"cpu": &graphql.Field{
				Type: graphql.Int,
			},
			"memory": &graphql.Field{
				Type: graphql.Int,
			},
			"storage": &graphql.Field{
				Type: graphql.Int,
			},
			"node": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// ResourceUsageType : Graphql object type of ResourceUsage
var ResourceUsageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResourceUsage",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type: ResourceType,
			},
			"in_use": &graphql.Field{
				Type: ResourceType,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
