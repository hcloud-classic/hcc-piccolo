package graphqltype

import "github.com/graphql-go/graphql"

// ServerNodeType : Graphql object type of ServerNode
var ServerNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerNode",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"node_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// ServerNodeListType : Graphql object type of ServerNodeList
var ServerNodeListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerNodeList",
		Fields: graphql.Fields{
			"server_node_list": &graphql.Field{
				Type: graphql.NewList(ServerNodeType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
