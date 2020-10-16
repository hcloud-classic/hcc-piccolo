package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// SchemaExported : GraphQL schema definition used for publisher
var SchemaExported graphql.Schema

// Schema : GraphQL schema definition
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryTypes,
		Mutation: mutationTypes,
		Subscription: subscriptionTypes,
	},
)

// Handler : HTTP handler for GraphQL
// Handler config options
//
// Schema : GraphQL schema definition variable name
//
// Pretty : Show sorted json code in GraphiQL
//
// GraphiQL : Show GraphQL GUI request form in web browser
var Handler = handler.New(&handler.Config{
	Schema:   &Schema,
	Pretty:   true,
	GraphiQL: false,
	Playground: true,
})
