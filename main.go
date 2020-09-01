package main

import (
	"fmt"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/action/grpc/client"
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

	graphql.Init()
}
