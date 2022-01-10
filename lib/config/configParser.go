package config

import (
	"github.com/Terry-Mao/goconf"
	"hcc/piccolo/lib/logger"
	"innogrid.com/hcloud-classic/hcc_errors"
)

var conf = goconf.New()
var config = piccoloConfig{}
var err error

func parseRsakey() {
	config.RsakeyConfig = conf.Get("rsakey")
	if config.RsakeyConfig == nil {
		logger.Logger.Panicln("no rsakey section")
	}

	Rsakey.PrivateKeyFile, err = config.RsakeyConfig.String("private_key_file")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseMysql() {
	config.MysqlConfig = conf.Get("mysql")
	if config.MysqlConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no mysql section").Fatal()
	}

	Mysql = mysql{}
	Mysql.ID, err = config.MysqlConfig.String("id")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Address, err = config.MysqlConfig.String("address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Port, err = config.MysqlConfig.Int("port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.Database, err = config.MysqlConfig.String("database")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
	Mysql.ConnectionRetryCount, err = config.MysqlConfig.Int("connection_retry_count")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Mysql.ConnectionRetryIntervalMs, err = config.MysqlConfig.Int("connection_retry_interval_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

}

func parseGrpc() {
	config.GrpcConfig = conf.Get("grpc")
	if config.GrpcConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no grpc section").Fatal()
	}

	Grpc.Port, err = config.GrpcConfig.Int("port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Grpc.ClientPingIntervalMs, err = config.GrpcConfig.Int("client_ping_interval_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Grpc.ClientPingTimeoutMs, err = config.GrpcConfig.Int("client_ping_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseGraphQL() {
	config.GraphQLConfig = conf.Get("graphql")
	if config.GraphQLConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no graphql section").Fatal()
	}

	GraphQL = graphql{}
	GraphQL.ProductionListenPort, err = config.GraphQLConfig.Int("production_listen_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	GraphQL.DevInternalListenPort, err = config.GraphQLConfig.Int("dev_internal_listen_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	GraphQL.DevInternalUsePlayground, _ = config.GraphQLConfig.Bool("dev_internal_use_playground")

	GraphQL.SubscriptionInterval, err = config.GraphQLConfig.Int("subscription_interval_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseHorn() {
	config.HornConfig = conf.Get("horn")
	if config.HornConfig == nil {
		logger.Logger.Panicln("no horn section")
	}

	Horn = horn{}
	Horn.ServerAddress, err = config.HornConfig.String("horn_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Horn.ServerPort, err = config.HornConfig.Int("horn_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Horn.ConnectionTimeOutMs, err = config.HornConfig.Int("horn_connection_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Horn.ConnectionRetryCount, err = config.HornConfig.Int("horn_connection_retry_count")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Horn.RequestTimeoutMs, err = config.HornConfig.Int("horn_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no flute section").Fatal()
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseCello() {
	config.CelloConfig = conf.Get("cello")
	if config.CelloConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no cello section").Fatal()
	}

	Cello = cello{}
	Cello.ServerAddress, err = config.CelloConfig.String("cello_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Cello.ServerPort, err = config.CelloConfig.Int("cello_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Cello.RequestTimeoutMs, err = config.CelloConfig.Int("cello_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseHarp() {
	config.HarpConfig = conf.Get("harp")
	if config.HarpConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no harp section").Fatal()
	}

	Harp = harp{}
	Harp.ServerAddress, err = config.HarpConfig.String("harp_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Harp.ServerPort, err = config.HarpConfig.Int("harp_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Harp.RequestTimeoutMs, err = config.HarpConfig.Int("harp_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseViolin() {
	config.ViolinConfig = conf.Get("violin")
	if config.ViolinConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no violin section").Fatal()
	}

	Violin = violin{}
	Violin.ServerAddress, err = config.ViolinConfig.String("violin_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Violin.ServerPort, err = config.ViolinConfig.Int("violin_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Violin.RequestTimeoutMs, err = config.ViolinConfig.Int("violin_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseViolinNoVnc() {
	config.ViolinNoVncConfig = conf.Get("violinnovnc")
	if config.ViolinNoVncConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no violinnovnc section").Fatal()
	}

	ViolinNoVnc = violinNoVnc{}
	ViolinNoVnc.ServerAddress, err = config.ViolinNoVncConfig.String("violinnovnc_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	ViolinNoVnc.ServerPort, err = config.ViolinNoVncConfig.Int("violinnovnc_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	ViolinNoVnc.RequestTimeoutMs, err = config.ViolinNoVncConfig.Int("violinnovnc_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parsePiano() {
	config.PianoConfig = conf.Get("piano")
	if config.PianoConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no piano section").Fatal()
	}

	Piano = piano{}
	Piano.ServerAddress, err = config.PianoConfig.String("piano_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Piano.ServerPort, err = config.PianoConfig.Int("piano_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Piano.RequestTimeoutMs, err = config.PianoConfig.Int("piano_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseTuba() {
	config.TubaConfig = conf.Get("tuba")
	if config.TubaConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no tuba section").Fatal()
	}

	Tuba = tuba{}
	Tuba.ServerPort, err = config.TubaConfig.Int("tuba_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Tuba.RequestTimeoutMs, err = config.TubaConfig.Int("tuba_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseTimpani() {
	config.TimpaniConfig = conf.Get("timpani")
	if config.TimpaniConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no timpani section").Fatal()
	}

	Timpani = timpani{}
	Timpani.ServerAddress, err = config.TimpaniConfig.String("timpani_server_address")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.ServerPort, err = config.TimpaniConfig.Int("timpani_server_port")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.RequestTimeoutMs, err = config.TimpaniConfig.Int("timpani_request_timeout_ms")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	Timpani.RequestRetry, err = config.TimpaniConfig.Int("timpani_request_retry")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

func parseUser() {
	config.UserConfig = conf.Get("user")
	if config.UserConfig == nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, "no user section").Fatal()
	}

	User = user{}
	User.TokenExpirationTimeMinutes, err = config.UserConfig.Int("token_expiration_time_minutes")
	if err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}
}

// Init : Parse config file and initialize config structure
func Init() {
	if err = conf.Parse(configLocation); err != nil {
		hcc_errors.NewHccError(hcc_errors.PiccoloInternalInitFail, err.Error()).Fatal()
	}

	parseRsakey()
	parseMysql()
	parseGrpc()
	parseGraphQL()
	parseHorn()
	parseFlute()
	parseCello()
	parseHarp()
	parseViolin()
	parseViolinNoVnc()
	parsePiano()
	parseTuba()
	parseUser()
}
