package mutationparser

import (
	"errors"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
)

func checkVncArgsAll(args map[string]interface{}) bool {
	_, serverUUIDOk := args["server_uuid"].(string)
	_, targetIPOk := args["target_ip"].(string)
	_, targetPortOk := args["target_port"].(string)
	_, targetPassOk := args["target_pass"].(string)
	_, actionOk := args["action"].(string)

	return serverUUIDOk && targetIPOk && targetPortOk && targetPassOk && actionOk
}

// CreateVnc : Create infos of the VNC
func CreateVnc(args map[string]interface{}) (interface{}, error) {
	if !checkVncArgsAll(args) {
		return nil, errors.New("check needed arguments (server_uuid, target_ip, target_port, target_pass, action)")
	}

	serverUUID, _ := args["server_uuid"].(string)
	targetIP, _ := args["target_ip"].(string)
	targetPort, _ := args["target_port"].(string)
	targetPass, _ := args["target_pass"].(string)
	action, _ := args["action"].(string)

	var createVncData data.CreateVncData
	query := "mutation _ { create_vnc(server_uuid: \"" + serverUUID + "\",target_ip:\"" + targetIP + "\", target_port:\"" +
		targetPort + "\",target_pass:\"" + targetPass + "\", action:\"" + action + "\") {" +
		"server_uuid target_ip target_port target_pass websocket_port action } }"

	return http.DoHTTPRequest("violin-novnc", true, "CreateVncData", createVncData, query)
}

// ControlVnc : Set VNC with provided options
func ControlVnc(args map[string]interface{}) (interface{}, error) {
	if !checkVncArgsAll(args) {
		return nil, errors.New("check needed arguments (server_uuid, target_ip, target_port, target_pass, action)")
	}

	return client.RC.ControlVNC(args)
}
