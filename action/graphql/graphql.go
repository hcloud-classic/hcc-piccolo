package graphql

import (
	"github.com/graphql-go/graphql"
)

// SchemaExported : GraphQL schema definition used for publisher
var SchemaExported graphql.Schema

// Schema : GraphQL schema definition
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:        queryTypes,
		Mutation:     mutationTypes,
		Subscription: subscriptionTypes,
	},
)
