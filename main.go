package main

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/server"
	"hcc/piccolo/action/http"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := logger.Init()
	if err != nil {
		errors.SetErrLogger(logger.Logger)
		errors.NewHccError(errors.PiccoloInternalInitFail, "logger.Init(): "+err.Error()).Fatal()
	}
	errors.SetErrLogger(logger.Logger)

	config.Init()

	err = mysql.Init()
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "mysql.Init(): "+err.Error()).Fatal()
	}

	err = client.Init()
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "client.Init(): "+err.Error()).Fatal()
	}
}

func end() {
	client.End()
	logger.End()
}

func main() {
	// Catch the exit signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		end()
		fmt.Println("Exiting piccolo module...")
		os.Exit(0)
	}()

	go server.Init()
	http.Init()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
