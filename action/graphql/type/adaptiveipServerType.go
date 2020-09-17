package graphqltype

import "github.com/graphql-go/graphql"

// AdaptiveIPServerType : Graphql object type of AdaptiveIPServer
var AdaptiveIPServerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPServer",
		Fields: graphql.Fields{
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"public_ip": &graphql.Field{
				Type: graphql.String,
			},
			"private_ip": &graphql.Field{
				Type: graphql.String,
			},
			"private_gateway": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// AdaptiveIPServerListType : Graphql object type of AdaptiveIPServerList
var AdaptiveIPServerListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPServerList",
		Fields: graphql.Fields{
			"adaptiveip_server_list": &graphql.Field{
				Type: graphql.NewList(AdaptiveIPServerType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
