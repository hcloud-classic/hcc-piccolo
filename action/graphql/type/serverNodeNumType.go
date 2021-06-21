package graphqltype

import "github.com/graphql-go/graphql"

// ServerNodeNumType : Graphql object type of ServerNodeNumType
var ServerNodeNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerNodeNumType",
		Fields: graphql.Fields{
			"number": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
