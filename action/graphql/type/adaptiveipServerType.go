<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import "github.com/graphql-go/graphql"

// AdaptiveIPServerType : Graphql object type of AdaptiveIPServer
var AdaptiveIPServerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPServer",
		Fields: graphql.Fields{
<<<<<<< HEAD
			"adaptiveip_uuid": &graphql.Field{
				Type: graphql.String,
			},
=======
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
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
<<<<<<< HEAD
=======
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
		},
	},
)
