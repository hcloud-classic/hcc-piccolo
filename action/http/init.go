package http

import (
	"github.com/graphql-go/handler"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/action/websocket"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
)

// Init : Initialize HTTP server
func Init() {
	logger.Logger.Println("Opening HTTP server on port " + strconv.Itoa(int(config.HTTP.Port)) + "...")

	var graphqlHandler = handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   !config.HTTP.UsePlayground,
		Playground: config.HTTP.UsePlayground,
	})
	http.Handle("/graphql", graphqlHandler)
	logger.Logger.Println("Serving GraphQL requests on /graphql")

	graphql.SchemaExported = graphql.Schema

	http.HandleFunc("/subscriptions", websocket.Handler)
	logger.Logger.Println("Serving GraphQL's subscription websocket requests on /subscriptions")

	err := http.ListenAndServe(":"+strconv.Itoa(int(config.HTTP.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare HTTP server!")
		return
	}
}
