package main

import (
	"fmt"
	"hcc/piccolo/action/graphql"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/syscheck"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := syscheck.CheckRoot()
	if err != nil {
		panic(err)
	}

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	config.Parser()
}

func end(){
	logger.End()
}

func main() {
	// Catch the exit signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(){
		<- sigChan
		end()
		fmt.Println("Exiting piccolo module...")
		os.Exit(0)
	}()

	graphql.Init()
}
