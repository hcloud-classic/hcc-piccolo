package graphql

import (
	"github.com/graphql-go/graphql"
<<<<<<< HEAD
	"github.com/graphql-go/handler"
)

// Schema : GraphQL schema definition
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryTypes,
		Mutation: mutationTypes,
	},
)

// GraphqlHandler : GraphQL schema handler
//
// Handler config options
//
// Schema : GraphQL schema definition variable name
//
// Pretty : Show sorted json code in GraphiQL
//
// GraphiQL : Show GraphQL GUI request form in web browser
var GraphqlHandler = handler.New(&handler.Config{
	Schema:   &Schema,
	Pretty:   true,
	GraphiQL: true,
})
=======
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
