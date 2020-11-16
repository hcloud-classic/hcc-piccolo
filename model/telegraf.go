package model

import "hcc/piccolo/lib/errors"

// Telegraf - cgs
type Telegraf struct {
	UUID   string            `json:"id"`
	Result []byte            `json:"result"`
	Errors []errors.HccError `json:"errors"`
}
