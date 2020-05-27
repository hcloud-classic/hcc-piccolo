package config

import (
	"hcc/piccolo/lib/logger"

	"github.com/Terry-Mao/goconf"
)

var conf = goconf.New()
var config = piccoloConfig{}
var err error

func parseHTTP() {
	config.HTTPConfig = conf.Get("http")
	if config.HTTPConfig == nil {
		logger.Logger.Panicln("no http section")
	}

	HTTP = http{}
	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
		logger.Logger.Panicln("no flute section")
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		logger.Logger.Panicln(err)
	}

	parseHTTP()

}

func parseCello() {
	config.CelloConfig = conf.Get("cello")
	if config.CelloConfig == nil {
		logger.Logger.Panicln("no cello section")
	}

	Cello = cello{}
	Cello.ServerAddress, err = config.CelloConfig.String("cello_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Cello.ServerPort, err = config.CelloConfig.Int("cello_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Cello.RequestTimeoutMs, err = config.CelloConfig.Int("cello_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseHarp() {
	config.HarpConfig = conf.Get("harp")
	if config.HarpConfig == nil {
		logger.Logger.Panicln("no harp section")
	}

	Harp = harp{}
	Harp.ServerAddress, err = config.HarpConfig.String("harp_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Harp.ServerPort, err = config.HarpConfig.Int("harp_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Harp.RequestTimeoutMs, err = config.HarpConfig.Int("harp_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseViola() {
	config.ViolaConfig = conf.Get("viola")
	if config.ViolaConfig == nil {
		logger.Logger.Panicln("no viola section")
	}

	Viola = viola{}
	Viola.ServerAddress, err = config.ViolaConfig.String("viola_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Viola.ServerPort, err = config.ViolaConfig.Int("viola_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Viola.RequestTimeoutMs, err = config.ViolaConfig.Int("viola_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}
