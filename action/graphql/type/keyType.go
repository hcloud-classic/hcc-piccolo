package graphqltype

import "github.com/graphql-go/graphql"

// PermissionKey : Graphql object type of ServerList
var PermissionKey = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PemermissionKey",
		Fields: graphql.Fields{
			"service_name": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"action": &graphql.Field{
				Type: graphql.String,
			},
			"error": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
