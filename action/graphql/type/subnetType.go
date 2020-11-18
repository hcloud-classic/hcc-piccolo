<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

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
<<<<<<< HEAD
=======
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// SubnetListType : Graphql object type of SubnetList
var SubnetListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SubnetList",
		Fields: graphql.Fields{
			"subnet_list": &graphql.Field{
				Type: graphql.NewList(SubnetType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// CreateDHCPConfResultType : Graphql object type of CreateDHCPConfResult
var CreateDHCPConfResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateDHCPConfResult",
		Fields: graphql.Fields{
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
		},
	},
)
