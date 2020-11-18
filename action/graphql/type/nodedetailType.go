<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import "github.com/graphql-go/graphql"

// NodeDetailType : Graphql object type of NodeDetail
var NodeDetailType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeDetail",
		Fields: graphql.Fields{
			"node_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_model": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_processors": &graphql.Field{
				Type: graphql.Int,
			},
			"cpu_threads": &graphql.Field{
				Type: graphql.Int,
			},
<<<<<<< HEAD
=======
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
		},
	},
)
