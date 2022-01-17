package graphqltype

import "github.com/graphql-go/graphql"

// PermissionKey : Graphql object type of ServerList
var PermissionKey = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PemermissionKey",
		Fields: graphql.Fields{
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"pemkey": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
