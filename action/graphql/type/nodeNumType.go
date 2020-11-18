<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import "github.com/graphql-go/graphql"

// NodeNumType : Graphql object type of NodeNum
var NodeNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeNum",
		Fields: graphql.Fields{
			"number": &graphql.Field{
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
