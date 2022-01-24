package graphql

import (
	"hcc/piccolo/action/graphql/fields"

	"github.com/graphql-go/graphql"
)

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// piccolo
		"signup":              &fields.SignUp,
		"unregister":          &fields.Unregister,
		"update_user":         &fields.UpdateUser,
		"create_group":        &fields.CreateUser,
		"update_group":        &fields.UpdateGroup,
		"delete_group":        &fields.DeleteGroup,
		"create_quota":        &fields.CreateGroup,
		"update_quota":        &fields.UpdateQuota,
		"delete_quota":        &fields.DeleteQuota,
		"delete_server_alarm": &fields.DeleteServerAlarm,
		// violin
		"create_server":       &fields.CreateServer,
		"update_server":       &fields.UpdateServer,
		"update_server_nodes": &fields.UpdateServerNodes,
		"scale_up_server":     &fields.ScaleUpServer,
		"delete_server":       &fields.DeleteServer,
		"create_server_node":  &fields.CreateServerNode,
		"delete_server_node":  &fields.DeleteServerNode,
		// harp
		"create_subnet":             &fields.CreateSubnet,
		"update_subnet":             &fields.UpdateSubnet,
		"delete_subnet":             &fields.DeleteSubnet,
		"create_dhcpd_conf":         &fields.CreateDHCPConf,
		"create_adaptiveip_setting": &fields.CreateAdaptiveIpSetting,
		"create_adaptiveip_server":  &fields.CreateAdaptiveIpServer,
		"delete_adaptiveip_server":  &fields.DeleteAdaptiveIpServer,
		"create_port_forwarding":    &fields.CreatePortForwarding,
		"delete_port_forwarding":    &fields.DeletePortForwarding,
		// flute
		"on_node":            &fields.OnNode,
		"off_node":           &fields.OffNode,
		"force_restart_node": &fields.ForceRestartNode,
		"create_node":        &fields.CreateNode,
		"update_node":        &fields.UpdateNode,
		"delete_node":        &fields.DeleteNode,
		// node_detail DB
		"create_node_detail": &fields.CreateNodeDetail,
		"update_node_detail": &fields.UpdateNodeDetail,
		"delete_node_detail": &fields.DeleteNodeDetail,
		// Cello
		"volume_handle": &fields.VolumeHandle,
		"mount_handle":  &fields.MountHandle,

		// Violin to Viola
		"create_pemkey": &fields.CreatePemKey,
	},
})
