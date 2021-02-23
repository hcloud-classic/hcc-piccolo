package http

import (
	"context"
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func telegrafSubscriptionQueryTimeChange(query string, newTime string) string {
	newQuery := ""

	querySlice := strings.Split(query, ",")
	for i := range querySlice {
		if strings.Contains(querySlice[i], "time") {
			if strings.Contains(querySlice[i], "$time") {
				continue
			}
			querySlice[i] = strings.Replace(querySlice[i], " ", "", -1)
			querySlice[i] = strings.Replace(querySlice[i], "\t", "", -1)
			querySlice[i] = strings.Replace(querySlice[i], "\n", "", -1)
			s := strings.Split(querySlice[i], ":")
			if len(s) == 2 {
				querySlice[i] = "time: " + "\"" + newTime + "\""
			}
		}
		newQuery += querySlice[i] + ","
	}

	logger.Logger.Println("newQuery", newQuery)

	return newQuery
}

func telegrafSubscriptionGetNewTime(dataStr string) string {
	newTime := ""

	if strings.Contains(dataStr, "values") {
		dataStr = strings.TrimSpace(dataStr)
		dataStr = strings.Replace(dataStr, " ", "", -1)
		dataStr = strings.Replace(dataStr, "\t", "", -1)
		dataStr = strings.Replace(dataStr, "\n", "", -1)
		dataStr = strings.Replace(dataStr, "\\", "", -1)
		dataStr = strings.Replace(dataStr, "\"", "", -1)
		dataStr = strings.Replace(dataStr, "[", "", -1)
		dataStr = strings.Replace(dataStr, "]", "", -1)
		slices := strings.Split(dataStr, ":")
		length := len(slices)
		for i := range slices {
			if strings.Contains(slices[i], "values") && length >= i+2 {
				s := strings.Split(slices[i+1], ",")
				newTime = s[0]
				break
			}
		}
	}

	logger.Logger.Println("newTime", newTime)

	return newTime
}

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
		newTime := ""

		for true {
			// This assumes you have access to the above subscription manager
			subscriptions := subscriptionManager.Subscriptions()

			for conn := range subscriptions {
				for _, subscription := range subscriptions[conn] {
					query := strings.TrimSpace(subscription.Query)

					if strings.HasPrefix(query, "subscription") &&
						strings.Contains(query, "telegraf") {
						query = telegrafSubscriptionQueryTimeChange(query, newTime)
					}

					// Prepare an execution context for running the query
					ctx := context.Background()

					// Re-execute the subscription query
					params := graphqlgo.Params{
						Schema:         graphql.Schema, // The GraphQL schema
						RequestString:  query,
						VariableValues: subscription.Variables,
						OperationName:  subscription.OperationName,
						Context:        ctx,
					}
					result := graphqlgo.Do(params)

					if strings.HasPrefix(query, "subscription") &&
						strings.Contains(query, "telegraf") {
						dataStr := fmt.Sprintf("%v", result.Data)
						newTime = telegrafSubscriptionGetNewTime(dataStr)
					}

					// Send query results back to the subscriber at any point
					data := graphqlws.DataMessagePayload{
						// Data can be anything (interface{})
						Data: result.Data,
						// Errors is optional ([]error)
						Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
					}
					subscription.SendData(&data)
					logger.Logger.Println(data.Errors)
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
