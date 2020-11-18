package config

import (
	"github.com/Terry-Mao/goconf"
<<<<<<< HEAD
	"hcc/piccolo/lib/logger"
=======
	"hcc/piccolo/lib/errors"
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
)

var conf = goconf.New()
var config = piccoloConfig{}
var err error

<<<<<<< HEAD
func parseHTTP() {
	config.HTTPConfig = conf.Get("http")
	if config.HTTPConfig == nil {
		logger.Logger.Panicln("no http section")
=======
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
	Mysql.ConnectionRetryCount, err = config.MysqlConfig.Int("connection_retry_count")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.ConnectionRetryIntervalMs, err = config.MysqlConfig.Int("connection_retry_interval_ms")
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	HTTP = http{}
	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
	}

	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	HTTP.UsePlayground, err = config.HTTPConfig.Bool("use_playground")
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no flute section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no flute section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}
}

func parseCello() {
	config.CelloConfig = conf.Get("cello")
	if config.CelloConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no cello section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no cello section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Cello = cello{}
	Cello.ServerAddress, err = config.CelloConfig.String("cello_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Cello.ServerPort, err = config.CelloConfig.Int("cello_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Cello.RequestTimeoutMs, err = config.CelloConfig.Int("cello_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}
}

func parseHarp() {
	config.HarpConfig = conf.Get("harp")
	if config.HarpConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no harp section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no harp section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Harp = harp{}
	Harp.ServerAddress, err = config.HarpConfig.String("harp_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Harp.ServerPort, err = config.HarpConfig.Int("harp_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Harp.RequestTimeoutMs, err = config.HarpConfig.Int("harp_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
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
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}
}

func parseViolin() {
	config.ViolinConfig = conf.Get("violin")
	if config.ViolinConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no violin section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no violin section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Violin = violin{}
	Violin.ServerAddress, err = config.ViolinConfig.String("violin_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Violin.ServerPort, err = config.ViolinConfig.Int("violin_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Violin.RequestTimeoutMs, err = config.ViolinConfig.Int("violin_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}
}

func parseViolinNoVnc() {
	config.ViolinNoVncConfig = conf.Get("violinnovnc")
	if config.ViolinNoVncConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no violinnovnc section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no violinnovnc section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	ViolinNoVnc = violinNoVnc{}
	ViolinNoVnc.ServerAddress, err = config.ViolinNoVncConfig.String("violinnovnc_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	ViolinNoVnc.ServerPort, err = config.ViolinNoVncConfig.Int("violinnovnc_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	ViolinNoVnc.RequestTimeoutMs, err = config.ViolinNoVncConfig.Int("violinnovnc_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}
}

func parsePiano() {
	config.PianoConfig = conf.Get("piano")
	if config.PianoConfig == nil {
<<<<<<< HEAD
		logger.Logger.Panicln("no piano section")
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, "no piano section").Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Piano = piano{}
	Piano.ServerAddress, err = config.PianoConfig.String("piano_server_address")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Piano.ServerPort, err = config.PianoConfig.Int("piano_server_port")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}

	Piano.RequestTimeoutMs, err = config.PianoConfig.Int("piano_request_timeout_ms")
	if err != nil {
<<<<<<< HEAD
		logger.Logger.Panicln(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		logger.Logger.Panicln(err)
	}

=======
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseTimpani() {
	config.TimpaniConfig = conf.Get("timpani")
	if config.TimpaniConfig == nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, "no timpani section").Fatal()
	}

	Timpani = timpani{}
	Timpani.ServerAddress, err = config.TimpaniConfig.String("timpani_server_address")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.ServerPort, err = config.TimpaniConfig.Int("timpani_server_port")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.RequestTimeoutMs, err = config.TimpaniConfig.Int("timpani_request_timeout_ms")
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.RequestRetry, err = config.TimpaniConfig.Int("timpani_request_retry")
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
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	parseHTTP()
	parseFlute()
	parseCello()
	parseHarp()
<<<<<<< HEAD
	parseViola()
	parseViolin()
	parseViolinNoVnc()
	parsePiano()
=======
	parseViolin()
	parseViolinNoVnc()
	parsePiano()
	parseUser()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
