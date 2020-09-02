package queryparser

import (
	"errors"
	"hcc/piccolo/action/grpc/client"
	rpcnovnc "hcc/piccolo/action/grpc/pb/rpcviolin_novnc"
)

func checkVncArgsAll(args map[string]interface{}) bool {
	_, serverUUIDOk := args["server_uuid"].(string)
	_, actionOk := args["action"].(string)

	return serverUUIDOk && actionOk
}

// ControlVnc : Set VNC with provided options
func ControlVnc(args map[string]interface{}) (interface{}, error) {
	if !checkVncArgsAll(args) {
		return nil, errors.New("check needed arguments (server_uuid, action)")
	}

	serverUUID, _ := args["server_uuid"].(string)
	action, _ := args["action"].(string)

	vnc := rpcnovnc.VNC{
		ServerUUID: serverUUID,
		Action:     action,
	}

	resControlVNC, err := client.RC.ControlVNC(&rpcnovnc.ReqControlVNC{
		Vnc: &vnc,
	})
	if err != nil {
		return nil, err
	}

	return resControlVNC.Port, nil
}
