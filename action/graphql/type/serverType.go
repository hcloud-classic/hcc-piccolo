package graphqltype

import "github.com/graphql-go/graphql"

// ServerType : Graphql object type of Server
var ServerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Server",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"subnet_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"os": &graphql.Field{
				Type: graphql.String,
			},
			"server_name": &graphql.Field{
				Type: graphql.String,
			},
			"server_desc": &graphql.Field{
				Type: graphql.String,
			},
			"cpu": &graphql.Field{
				Type: graphql.Int,
			},
			"memory": &graphql.Field{
				Type: graphql.Int,
			},
			"disk_size": &graphql.Field{
				Type: graphql.Int,
			},
			"nodes": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"user_uuid": &graphql.Field{
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

// ServerListType : Graphql object type of ServerList
var ServerListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerList",
		Fields: graphql.Fields{
			"server_list": &graphql.Field{
				Type: graphql.NewList(ServerType),
			},
			"total_num": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
