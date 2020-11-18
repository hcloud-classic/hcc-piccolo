<<<<<<< HEAD
package graphqlType

import "github.com/graphql-go/graphql"

=======
package graphqltype

import "github.com/graphql-go/graphql"

// VolumeNum : GraphQL type of VolumeNum
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
var VolumeNum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VolumeNum",
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
