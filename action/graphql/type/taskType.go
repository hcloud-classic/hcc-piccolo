package graphqltype

import "github.com/graphql-go/graphql"

// TaskListResultType : Graphql object type of TaskListResultType
var TaskListResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TaskListResultType",
		Fields: graphql.Fields{
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
