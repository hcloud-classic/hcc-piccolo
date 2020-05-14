package main

import (
	"hcc/piccolo/action/graphql"
	hccGatewayEnd "hcc/piccolo/end"
	hccGatewayInit "hcc/piccolo/init"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"net/http"
	"strconv"
)

func init() {
	err := hccGatewayInit.MainInit()
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		hccGatewayEnd.MainEnd()
	}()

	http.Handle("/graphql", graphql.GraphqlHandler)
	logger.Logger.Println("Opening server on port " + strconv.Itoa(int(config.HTTP.Port)) + "...")
	err := http.ListenAndServe(":"+strconv.Itoa(int(config.HTTP.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare http server!")
		return
	}
}
