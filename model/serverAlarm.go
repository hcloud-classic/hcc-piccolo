package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// ServerAlarm : Structure of ServerAlarm
type ServerAlarm struct {
	No                 int                       `json:"no"`
	UserID             string                    `json:"user_id"`
	UserName           string                    `json:"user_name"`
	ServerUUID         string                    `json:"server_uuid"`
	ServerName         string                    `json:"server_name"`
	Reason             string                    `json:"reason"`
	Detail             string                    `json:"detail"`
	Time               time.Time                 `json:"time"`
	Unread             int                       `json:"unread"`
	AutoScaleTriggered int                       `json:"auto_scale_triggered"`
	Errors             []errconv.PiccoloHccError `json:"errors"`
}

// ServerAlarms : Struct of ServerAlarms
type ServerAlarms struct {
	ServerAlarms []ServerAlarm             `json:"server_alarm_list"`
	TotalNum     int                       `json:"total_num"`
	Errors       []errconv.PiccoloHccError `json:"errors"`
}

// ServerAlarmsNum : Contain the number of ServerAlarms
type ServerAlarmsNum struct {
	Number int                       `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
