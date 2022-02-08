package graphql

import (
	"hcc/piccolo/action/graphql/fields"

	"github.com/graphql-go/graphql"
)

var queryTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// piccolo
			"login":                   &fields.Login,
			"user":                    &fields.User,
			"list_user":               &fields.ListUser,
			"num_user":                &fields.NumUser,
			"all_group":               &fields.AllGroup,
			"check_token":             &fields.CheckToken,
			"node_available":          &fields.NodeAvailable,
			"server_log":              &fields.ServerLog,
			"server_alarm_list":       &fields.ServerAlarmList,
			"num_unread_server_alarm": &fields.NumUnreadServerAlarm,
			"quota":                   &fields.Quota,
			"list_quota":              &fields.ListQuota,
			"quota_detail":            &fields.QuotaDetail,
			// violin
			"server":           &fields.Server,
			"list_server":      &fields.ListServer,
			"all_server":       &fields.AllServer,
			"num_server":       &fields.NumServer,
			"server_node":      &fields.ServerNode,
			"list_server_node": &fields.ListServerNode,
			"all_server_node":  &fields.AllServerNode,
			"num_nodes_server": &fields.NumNodesServer,
			// vnc
			"control_vnc": &fields.CotrolVNC,
			// harp
			"subnet":                       &fields.Subnet,
			"list_subnet":                  &fields.ListSubnet,
			"all_subnet":                   &fields.AllSubnet,
			"available_subnet":             &fields.AvailableSubnet,
			"num_subnet":                   &fields.NumSubnet,
			"valid_check_subnet":           &fields.VaildCheckSubnet,
			"adaptiveip_available_ip_list": &fields.AdaptiveIpAvailableIpList,
			"adaptiveip_setting":           &fields.AdaptiveIpSetting,
			"adaptiveip_server":            &fields.AdaptiveIPSerer,
			"list_adaptiveip_server":       &fields.ListAdaptiveIpServer,
			"all_adaptiveip_server":        &fields.AllAdaptiveIpServer,
			"num_adaptiveip_server":        &fields.NumAdaptiveIpServer,
			"list_port_forwarding":         &fields.ListPortForwarding,
			// flute
			"power_state_node":         &fields.PowerStateNode,
			"node":                     &fields.Node,
			"list_node":                &fields.ListNode,
			"all_node":                 &fields.AllNode,
			"all_server_prepared_node": &fields.AllServerPreparedNode,
			"num_node":                 &fields.NumNode,
			"detail_node":              &fields.DetailNode,
			// piano
			"telegraf":       &fields.Telegraf,
			"billing_data":   &fields.BillingData,
			"billing_detail": &fields.BillingDetail,
			// cello
			"volume_list":         &fields.VolumeList,
			"pool_list":           &fields.PoolList,
			"available_pool_list": &fields.AvailablePoolList,

			// tuba
			"all_task": &fields.AllTask,
			// viola
			"get_pemkey": &fields.GetPemKey,

			// Timpani
			"mastersync":            &fields.TimapniMasterSync,
			"timpani_volume_backup": &fields.TimpaniBackup,
			// "timpani_volume_backup_scheduler": &fields.TimpaniBackup,
			"restore": &fields.Restore,
		},
	})
