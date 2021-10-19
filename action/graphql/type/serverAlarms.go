package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// ServerAlarmType : Graphql object type of ServerAlarm
var ServerAlarmType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerAlarm",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.String,
			},
			"user_name": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"server_name": &graphql.Field{
				Type: graphql.String,
			},
			"reason": &graphql.Field{
				Type: graphql.String,
			},
			"detail": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// ServerAlarmsType : Graphql object type of ServerAlarms
var ServerAlarmsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ServerAlarms",
		Fields: graphql.Fields{
			"server_alarm_list": &graphql.Field{
				Type: graphql.NewList(ServerAlarmType),
			},
			"total_num": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
