package main

import (
	"fmt"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/server"
	"hcc/piccolo/action/http"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/syscheck"
	"os"
	"os/signal"
	"syscall"

	"innogrid.com/hcloud-classic/hcc_errors"
)

func init() {
	err := logger.Init()
	if err != nil {
		hcc_errors.SetErrLogger(logger.Logger)
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "logger.Init(): "+err.Error()).Fatal()
	}
	hcc_errors.SetErrLogger(logger.Logger)

	config.Init()

	err = mysql.Init()
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "mysql.Init(): "+err.Error()).Fatal()
	}

	err = syscheck.IncreaseRLimitToMax()
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "syscheck.IncreaseRLimitToMax(): "+err.Error()).Fatal()
	}

	err = client.Init()
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "client.Init(): "+err.Error()).Fatal()
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
}
