package graphqltype

import "github.com/graphql-go/graphql"

// NodeNumType : Graphql object type of NodeNum
var NodeNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeNum",
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
