package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

func checkVncArgsAll(args map[string]interface{}) bool {
	_, serverUUIDOk := args["server_uuid"].(string)
	_, actionOk := args["action"].(string)

	return serverUUIDOk && actionOk
}

// ControlVnc : Set VNC with provided options
func ControlVnc(args map[string]interface{}) (interface{}, error) {
	if !checkVncArgsAll(args) {
		return model.VncPort{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "check needed arguments (server_uuid, action)")}, nil
	}

	serverUUID, _ := args["server_uuid"].(string)
	action, _ := args["action"].(string)

	vnc := pb.VNC{
		ServerUUID: serverUUID,
		Action:     action,
	}

	resControlVNC, err := client.RC.ControlVNC(&pb.ReqControlVNC{
		Vnc: &vnc,
	})
	if err != nil {
		return nil, err
	}

	hccErrStack := errconv.GrpcStackToHcc(resControlVNC.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.VncPort{Port: resControlVNC.Port, Errors: Errors}, nil
}
