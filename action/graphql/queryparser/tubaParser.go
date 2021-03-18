package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/iputil"
	"hcc/piccolo/model"

	"google.golang.org/grpc"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// AllTask : Get task list with provided options (Just call ListTask())
func AllTask(args map[string]interface{}) (interface{}, error) {
	serverAddress, serverAddressOk := args["server_address"].(string)
	serverPort, serverPortOk := args["server_port"].(int)

	if !serverAddressOk || !serverPortOk {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need server_address and server_port arguments")}, nil
	}

	ipCheck := iputil.CheckValidIP(serverAddress)
	if ipCheck == nil {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"invalid server address")}, nil
	}

	if serverPort < 1 || serverPort > 65535 {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"out of port number range")}, nil
	}

	var conn *grpc.ClientConn
	tubaClient, err := client.InitTuba(serverAddress, serverPort, conn)
	if err != nil {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	defer client.CloseTuba(conn)

	resTaskListResult, err := client.GetTaskList(tubaClient)
	if err != nil {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resTaskListResult.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.TaskListResult{
		Result: resTaskListResult.Result,
		Errors: Errors,
	}, nil
}
