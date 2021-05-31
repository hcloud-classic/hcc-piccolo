package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// BillingType : GraphQL type of Billing
var BillingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "billing",
		Fields: graphql.Fields{
			"result": &graphql.Field{
				Type: graphql.String,
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
