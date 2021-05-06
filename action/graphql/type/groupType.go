package graphqltype

import "github.com/graphql-go/graphql"

// GroupType : Graphql object type of group
var GroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Group",
		Fields: graphql.Fields{
			"group_id": &graphql.Field{
				Type: graphql.Int,
			},
			"group_name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// GroupListType : Graphql object type of GroupList
var GroupListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupList",
		Fields: graphql.Fields{
			"group_list": &graphql.Field{
				Type: graphql.NewList(GroupType),
			},
			"total_num": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
