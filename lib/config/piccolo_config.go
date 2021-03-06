package config

import "github.com/Terry-Mao/goconf"

var configLocation = "/etc/hcc/piccolo/piccolo.conf"

type piccoloConfig struct {
	MysqlConfig       *goconf.Section
	GrpcConfig        *goconf.Section
	GraphQLConfig     *goconf.Section
	FluteConfig       *goconf.Section
	CelloConfig       *goconf.Section
	HarpConfig        *goconf.Section
	ViolinConfig      *goconf.Section
	ViolinNoVncConfig *goconf.Section
	PianoConfig       *goconf.Section
	TubaConfig        *goconf.Section
	TimpaniConfig     *goconf.Section
	UserConfig        *goconf.Section
}
