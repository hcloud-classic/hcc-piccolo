package model

<<<<<<< HEAD
// Telegraf - cgs
type Telegraf struct {
	Metric    string   `json:"metric"`
	SubMetric string   `json:"subMetric"`
	UUID      string   `json:"id"`
	Series    []Series `json:"data"`
}

// Series - cgs
type Series struct {
	//Time  string `json:"x"`
	Time  int `json:"x"`
	Value int `json:"y"`
=======
import "hcc/piccolo/lib/errors"

// Telegraf - cgs
type Telegraf struct {
	UUID   string            `json:"id"`
	Result string            `json:"result"`
	Errors []errors.HccError `json:"errors"`
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}
