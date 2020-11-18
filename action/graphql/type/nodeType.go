<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import "github.com/graphql-go/graphql"

// NodeType : Graphql object type of Node
var NodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Node",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"bmc_mac_addr": &graphql.Field{
				Type: graphql.String,
			},
			"bmc_ip": &graphql.Field{
				Type: graphql.String,
			},
<<<<<<< HEAD
=======
			"bmc_ip_subnet_mask": &graphql.Field{
				Type: graphql.String,
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			"pxe_mac_addr": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_cores": &graphql.Field{
				Type: graphql.Int,
			},
			"memory": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
<<<<<<< HEAD
=======
			"rack_number": &graphql.Field{
				Type: graphql.Int,
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			"active": &graphql.Field{
				Type: graphql.Int,
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

// NodeListType : Graphql object type of NodeList
var NodeListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeList",
		Fields: graphql.Fields{
			"node_list": &graphql.Field{
				Type: graphql.NewList(NodeType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// PowerControlNodeType : Graphql object type of PowerControlNodeType
var PowerControlNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PowerControlNodeType",
		Fields: graphql.Fields{
			"results": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// PowerStateNodeType : Graphql object type of PowerStateNode
var PowerStateNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PowerStateNode",
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
