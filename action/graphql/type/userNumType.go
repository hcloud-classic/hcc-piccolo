package graphqltype

import "github.com/graphql-go/graphql"

// UserNumType : Graphql object type of UserNumType
var UserNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserNumType",
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
