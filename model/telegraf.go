package model

import "hcc/piccolo/lib/errors"

// Telegraf - cgs
type Telegraf struct {
	Metric    string            `json:"metric"`
	SubMetric string            `json:"subMetric"`
	UUID      string            `json:"id"`
	Series    [][]float64       `json:"data"`
	Errors    []errors.HccError `json:"errors"`
}
