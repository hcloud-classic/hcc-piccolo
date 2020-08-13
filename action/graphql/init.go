package graphql

import (
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
)

// Init : Initialize GraphQL HTTP server
func Init() {
	http.Handle("/graphql", graphqlHandler)
	logger.Logger.Println("Opening GraphQL HTTP server on port " + strconv.Itoa(int(config.HTTP.Port)) + "...")
	err := http.ListenAndServe(":"+strconv.Itoa(int(config.HTTP.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare GraphQL HTTP server!")
		return
	}
}
