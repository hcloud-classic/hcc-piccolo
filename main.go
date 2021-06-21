package main

import (
	"fmt"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/syscheck"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := syscheck.CheckRoot()
	if err != nil {
		log.Fatalf("syscheck.CheckRoot(): %v", err.Error())
	}

	err = logger.Init()
	if err != nil {
		log.Fatalf("logger.Init(): %v", err.Error())
	}

	config.Init()

	err = mysql.Init()
	if err != nil {
		logger.Logger.Fatalf("mysql.Init(): %v", err.Error())
	}

	err = client.Init()
	if err != nil {
		logger.Logger.Fatalf("client.Init(): %v", err.Error())
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
