package http

import (
	"context"
	"github.com/functionalfoundry/graphqlws"
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
	"time"
)

// Init : Initialize GraphQL server
func Init() {
	logger.Logger.Println("Opening GraphQL server on port " + strconv.Itoa(int(config.GraphQL.Port)) + "...")

	var graphqlHandler = handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   !config.GraphQL.UsePlayground,
		Playground: config.GraphQL.UsePlayground,
	})
	http.Handle("/graphql", graphqlHandler)
	logger.Logger.Println("Serving GraphQL requests on /graphql")

	subscriptionManager := graphqlws.NewSubscriptionManager(&graphql.Schema)
	graphqlWSHandler := graphqlws.NewHandler(graphqlws.HandlerConfig{
		SubscriptionManager: subscriptionManager,
	})
	http.Handle("/subscriptions", graphqlWSHandler)
	logger.Logger.Println("Serving GraphQL's subscription websocket requests on /subscriptions")

	go func() {
		for true {
			// This assumes you have access to the above subscription manager
			subscriptions := subscriptionManager.Subscriptions()

			for conn := range subscriptions {
				for _, subscription := range subscriptions[conn] {
					// Prepare an execution context for running the query
					ctx := context.Background()

					// Re-execute the subscription query
					params := graphqlgo.Params{
						Schema:         graphql.Schema, // The GraphQL schema
						RequestString:  subscription.Query,
						VariableValues: subscription.Variables,
						OperationName:  subscription.OperationName,
						Context:        ctx,
					}
					result := graphqlgo.Do(params)

					logger.Logger.Println(result)

					// Send query results back to the subscriber at any point
					data := graphqlws.DataMessagePayload{
						// Data can be anything (interface{})
						Data: result.Data,
						// Errors is optional ([]error)
						Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
					}
					subscription.SendData(&data)
				}
			}
			time.Sleep(time.Millisecond * time.Duration(config.GraphQL.SubscriptionInterval))
		}
	}()

	err := http.ListenAndServe(":"+strconv.Itoa(int(config.GraphQL.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare GraphQL server!")
		return
	}
}
