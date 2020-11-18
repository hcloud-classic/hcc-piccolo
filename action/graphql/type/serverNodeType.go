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
			"cpu_model": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_processors": &graphql.Field{
				Type: graphql.Int,
			},
			"cpu_cores": &graphql.Field{
				Type: graphql.Int,
			},
			"cpu_threads": &graphql.Field{
				Type: graphql.Int,
			},
			"memory": &graphql.Field{
				Type: graphql.Int,
			},
			"rack_number": &graphql.Field{
				Type: graphql.Int,
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
