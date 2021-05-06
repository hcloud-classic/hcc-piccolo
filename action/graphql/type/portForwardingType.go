package graphqltype

import "github.com/graphql-go/graphql"

// PortForwardingType : Graphql object type of PortForwarding
var PortForwardingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PortForwarding",
		Fields: graphql.Fields{
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"protocol": &graphql.Field{
				Type: graphql.String,
			},
			"external_port": &graphql.Field{
				Type: graphql.Int,
			},
			"internal_port": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// PortForwardingListType : Graphql object type of PortForwardingList
var PortForwardingListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PortForwardingList",
		Fields: graphql.Fields{
			"port_forwarding_list": &graphql.Field{
				Type: graphql.NewList(PortForwardingType),
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
