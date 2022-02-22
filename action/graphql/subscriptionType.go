package graphql

import (
	"hcc/piccolo/action/graphql/fields"

	"github.com/graphql-go/graphql"
)

var subscriptionTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			// piccolo
			"list_user":               &fields.ListUser,
			"resource_usage":          &fields.ResourceUsage,
			"server_log":              &fields.ServerLog,
			"server_alarm_list":       &fields.ServerAlarmList,
			"num_unread_server_alarm": &fields.NumUnreadServerAlarm,
			// violin
			"list_server": &fields.ListServer,
			"all_server":  &fields.AllServer,
			// harp
			"list_subnet": &fields.ListSubnet,
			"all_subnet":  &fields.AllSubnet,
			// flute
			"list_node": &fields.ListNode,
			// piano
			"telegraf": &fields.Telegraf,
			// tuba
			"all_task": &fields.AllTask,
		},
	})
