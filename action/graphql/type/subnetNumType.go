package graphqltype

import "github.com/graphql-go/graphql"

// SubnetNumType : Graphql object type of SubnetNum
var SubnetNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SubnetNum",
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
