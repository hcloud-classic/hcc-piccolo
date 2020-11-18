package graphqltype

import "github.com/graphql-go/graphql"

// Errors : Graphql object type of errors
var Errors = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Errors",
		Fields: graphql.Fields{
			"errcode": &graphql.Field{
				Type: graphql.Int,
			},
			"errtext": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
