<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import "github.com/graphql-go/graphql"

// AdaptiveIPServerNumType : Graphql object type of AdaptiveIPServerNum
var AdaptiveIPServerNumType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPServerNum",
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
