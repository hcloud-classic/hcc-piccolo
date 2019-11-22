package config

import "github.com/Terry-Mao/goconf"

var configLocation = "/etc/hcc/piccolo/piccolo.conf"

type piccoloConfig struct {
	HTTPConfig   *goconf.Section
	FluteConfig  *goconf.Section
	CelloConfig  *goconf.Section
	HarpConfig   *goconf.Section
	ViolaConfig  *goconf.Section
	ViolinConfig *goconf.Section
	PianoConfig *goconf.Section
}
