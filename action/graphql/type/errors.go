package graphqltype

import "github.com/graphql-go/graphql"

//// ErrorStack : Graphql object type of errors
//var ErrorStack = graphql.NewObject(
//	graphql.ObjectConfig{
//		Name: "ErrorStack",
//		Fields: graphql.Fields{
//			"errors": &graphql.Field{
//					Type: graphql.NewList(Errors),
//			},
//		},
//	},
//)

// Errors : Graphql object type of errors
var Errors = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Errors",
		Fields: graphql.Fields{
			"errcode": &graphql.Field{
				Type: graphql.Int,
			},
			"errtext": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
