package config

import (
	"github.com/Terry-Mao/goconf"
	"hcc/piccolo/lib/errors"
)

var conf = goconf.New()
var config = piccoloConfig{}
var err error

func parseMysql() {
	config.MysqlConfig = conf.Get("mysql")
	if config.MysqlConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no mysql section").Fatal()
	}

	Mysql = mysql{}
	Mysql.ID, err = config.MysqlConfig.String("id")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Password, err = config.MysqlConfig.String("password")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Address, err = config.MysqlConfig.String("address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Port, err = config.MysqlConfig.Int("port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Database, err = config.MysqlConfig.String("database")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseGrpc() {
	config.GrpcConfig = conf.Get("grpc")
	if config.GrpcConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no grpc section").Fatal()
	}

	Grpc.Port, err = config.GrpcConfig.Int("port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseHTTP() {
	config.HTTPConfig = conf.Get("http")
	if config.HTTPConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no http section").Fatal()
	}

	HTTP = http{}
	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	HTTP.UsePlayground, err = config.HTTPConfig.Bool("use_playground")
}

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no flute section").Fatal()
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseCello() {
	config.CelloConfig = conf.Get("cello")
	if config.CelloConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no cello section").Fatal()
	}

	Cello = cello{}
	Cello.ServerAddress, err = config.CelloConfig.String("cello_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Cello.ServerPort, err = config.CelloConfig.Int("cello_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Cello.RequestTimeoutMs, err = config.CelloConfig.Int("cello_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseHarp() {
	config.HarpConfig = conf.Get("harp")
	if config.HarpConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no harp section").Fatal()
	}

	Harp = harp{}
	Harp.ServerAddress, err = config.HarpConfig.String("harp_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Harp.ServerPort, err = config.HarpConfig.Int("harp_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Harp.RequestTimeoutMs, err = config.HarpConfig.Int("harp_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseViolin() {
	config.ViolinConfig = conf.Get("violin")
	if config.ViolinConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no violin section").Fatal()
	}

	Violin = violin{}
	Violin.ServerAddress, err = config.ViolinConfig.String("violin_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Violin.ServerPort, err = config.ViolinConfig.Int("violin_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Violin.RequestTimeoutMs, err = config.ViolinConfig.Int("violin_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseViolinNoVnc() {
	config.ViolinNoVncConfig = conf.Get("violinnovnc")
	if config.ViolinNoVncConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no violinnovnc section").Fatal()
	}

	ViolinNoVnc = violinNoVnc{}
	ViolinNoVnc.ServerAddress, err = config.ViolinNoVncConfig.String("violinnovnc_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	ViolinNoVnc.ServerPort, err = config.ViolinNoVncConfig.Int("violinnovnc_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	ViolinNoVnc.RequestTimeoutMs, err = config.ViolinNoVncConfig.Int("violinnovnc_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parsePiano() {
	config.PianoConfig = conf.Get("piano")
	if config.PianoConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no piano section").Fatal()
	}

	Piano = piano{}
	Piano.ServerAddress, err = config.PianoConfig.String("piano_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Piano.ServerPort, err = config.PianoConfig.Int("piano_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Piano.RequestTimeoutMs, err = config.PianoConfig.Int("piano_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseUser() {
	config.UserConfig = conf.Get("user")
	if config.UserConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no user section").Fatal()
	}

	User = user{}
	User.TokenExpirationTimeMinutes, err = config.UserConfig.Int("token_expiration_time_minutes")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

// Init : Parse config file and initialize config structure
func Init() {
	if err = conf.Parse(configLocation); err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	parseMysql()
	parseGrpc()
	parseHTTP()
	parseFlute()
	parseCello()
	parseHarp()
	parseViolin()
	parseViolinNoVnc()
	parsePiano()
	parseUser()
}
