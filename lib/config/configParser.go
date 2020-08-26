package config

import (
	"github.com/Terry-Mao/goconf"
	"hcc/piccolo/lib/logger"
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

func parseViolin() {
	config.ViolinConfig = conf.Get("violin")
	if config.ViolinConfig == nil {
		logger.Logger.Panicln("no violin section")
	}

	Violin = violin{}
	Violin.ServerAddress, err = config.ViolinConfig.String("violin_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Violin.ServerPort, err = config.ViolinConfig.Int("violin_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Violin.RequestTimeoutMs, err = config.ViolinConfig.Int("violin_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseViolinNoVnc() {
	config.ViolinNoVncConfig = conf.Get("violinnovnc")
	if config.ViolinNoVncConfig == nil {
		logger.Logger.Panicln("no violinnovnc section")
	}

	ViolinNoVnc = violinNoVnc{}
	ViolinNoVnc.ServerAddress, err = config.ViolinNoVncConfig.String("violinnovnc_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	ViolinNoVnc.ServerPort, err = config.ViolinNoVncConfig.Int("violinnovnc_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	ViolinNoVnc.RequestTimeoutMs, err = config.ViolinNoVncConfig.Int("violinnovnc_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parsePiano() {
	config.PianoConfig = conf.Get("piano")
	if config.PianoConfig == nil {
		logger.Logger.Panicln("no piano section")
	}

	Piano = piano{}
	Piano.ServerAddress, err = config.PianoConfig.String("piano_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Piano.ServerPort, err = config.PianoConfig.Int("piano_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Piano.RequestTimeoutMs, err = config.PianoConfig.Int("piano_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

// Init : Parse config file and initialize config structure
func Init() {
	if err = conf.Parse(configLocation); err != nil {
		logger.Logger.Panicln(err)
	}

	parseHTTP()
	parseFlute()
	parseCello()
	parseHarp()
	parseViolin()
	parseViolinNoVnc()
	parsePiano()
}
