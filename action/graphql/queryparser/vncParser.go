package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	rpcnovnc "hcc/piccolo/action/grpc/pb/rpcviolin_novnc"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

func checkVncArgsAll(args map[string]interface{}) bool {
	_, serverUUIDOk := args["server_uuid"].(string)
	_, actionOk := args["action"].(string)

	return serverUUIDOk && actionOk
}

// ControlVnc : Set VNC with provided options
func ControlVnc(args map[string]interface{}) (interface{}, error) {
	if !checkVncArgsAll(args) {
		return model.VncPort{Errors: errors.ReturnHccError(errors.PiccoloGraphQLArgumentError, "check needed arguments (server_uuid, action)")}, nil
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

	return model.VncPort{Port: resControlVNC.Port}, nil
}
