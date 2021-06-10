package http

import (
	"github.com/graphql-go/handler"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/action/http/subscription"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
)

func initGraphQLServer(isProduction bool) {
	var serveType string
	var listenPort int

	serveMux := http.NewServeMux()

	if isProduction {
		serveType = "Production"
		listenPort = int(config.GraphQL.ProductionListenPort)
	} else {
		serveType = "Dev Internal"
		listenPort = int(config.GraphQL.DevInternalListenPort)
	}

	logger.Logger.Println("Opening " + serveType + " GraphQL server on port " + strconv.Itoa(listenPort) + "...")

	var graphqlHandler = handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   !isProduction && !config.GraphQL.DevInternalUsePlayground,
		Playground: !isProduction && config.GraphQL.DevInternalUsePlayground,
	})
	serveMux.Handle("/graphql", graphqlHandler)
	serveMux.Handle("/subscriptions", subscription.NewSubscriptionHandler())

	server := http.Server{
		Addr:    ":" + strconv.Itoa(listenPort),
		Handler: serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare " + serveType + "GraphQL server!")
	}
}

// Init : Initialize GraphQL server
func Init() {
	go initGraphQLServer(false)
	initGraphQLServer(true)
}
