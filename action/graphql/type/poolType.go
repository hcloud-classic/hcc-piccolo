package graphqltype

import "github.com/graphql-go/graphql"

// PoolType : GraphQL type of PoolType
var PoolType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pool",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"size": &graphql.Field{
				Type: graphql.String,
			},
			"free": &graphql.Field{
				Type: graphql.String,
			},
			"capacity": &graphql.Field{
				Type: graphql.String,
			},
			"health": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"availableSize": &graphql.Field{
				Type: graphql.String,
			},
			"used": &graphql.Field{
				Type: graphql.String,
			},
			"action": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// PoolListType : Graphql object type of PoolListType
var PoolListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PoolList",
		Fields: graphql.Fields{
			"pool_list": &graphql.Field{
				Type: graphql.NewList(PoolType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
