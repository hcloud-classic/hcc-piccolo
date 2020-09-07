package model

import "hcc/piccolo/lib/errors"

// Telegraf - cgs
type Telegraf struct {
	Metric    string   `json:"metric"`
	SubMetric string   `json:"subMetric"`
	UUID      string   `json:"id"`
	Series    []Series `json:"data"`
	Errors []errors.HccError `json:"errors"`
}

// Series - cgs
type Series struct {
	//Time  string `json:"x"`
	Time  int `json:"x"`
	Value int `json:"y"`
}
