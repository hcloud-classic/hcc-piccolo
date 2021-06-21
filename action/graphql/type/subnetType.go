package graphqltype

import "github.com/graphql-go/graphql"

// SubnetType : Graphql object type of Subnet
var SubnetType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subnet",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"network_ip": &graphql.Field{
				Type: graphql.String,
			},
			"netmask": &graphql.Field{
				Type: graphql.String,
			},
			"gateway": &graphql.Field{
				Type: graphql.String,
			},
			"next_server": &graphql.Field{
				Type: graphql.String,
			},
			"name_server": &graphql.Field{
				Type: graphql.String,
			},
			"domain_name": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"leader_node_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"os": &graphql.Field{
				Type: graphql.String,
			},
			"subnet_name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
