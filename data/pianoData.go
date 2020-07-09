package data

import "hcc/piccolo/model"

// TelegrafData : Data structure of telegraf
type TelegrafData struct {
	Data struct {
		Telegraf model.Telegraf `json:"telegraf"`
	} `json:"data"`
}
