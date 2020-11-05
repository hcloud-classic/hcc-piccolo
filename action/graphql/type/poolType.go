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
			"Capacity": &graphql.Field{
				Type: graphql.String,
			},
			"Health": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"AvailableSize": &graphql.Field{
				Type: graphql.String,
			},
			"Used": &graphql.Field{
				Type: graphql.String,
			},
			"Action": &graphql.Field{
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
