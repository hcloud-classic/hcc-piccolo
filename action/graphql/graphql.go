package graphql

import (
	"github.com/graphql-go/graphql"
)

// Schema : GraphQL schema definition
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:        queryTypes,
		Mutation:     mutationTypes,
		Subscription: subscriptionTypes,
	},
)
