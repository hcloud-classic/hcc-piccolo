package graphqltype

import "github.com/graphql-go/graphql"

// AdaptiveIPServerNumType : Graphql object type of AdaptiveIPServerNum
var AdaptiveIPServerNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPServerNum",
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
