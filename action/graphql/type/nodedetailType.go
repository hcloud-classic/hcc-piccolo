package graphqltype

import "github.com/graphql-go/graphql"

// NodeDetailType : Graphql object type of NodeDetail
var NodeDetailType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeDetail",
		Fields: graphql.Fields{
			"node_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"node_detail_data": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
