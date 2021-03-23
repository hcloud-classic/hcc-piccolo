package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/iputil"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/pb"

	"google.golang.org/grpc"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// AllTask : Get task list with provided options (Just call ListTask())
func AllTask(args map[string]interface{}) (interface{}, error) {
	serverAddress, serverAddressOk := args["server_address"].(string)
	serverPort, serverPortOk := args["server_port"].(int)
	sortBy, sortByOk := args["sort_by"].(string)
	reverseSorting, reverseSortingOk := args["reverse_sorting"].(bool)

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

	reqGetTaskList := &pb.ReqGetTaskList{
		SortBy:         "",
		ReverseSorting: false,
	}
	if sortByOk {
		reqGetTaskList.SortBy = sortBy
		if reverseSortingOk {
			reqGetTaskList.ReverseSorting = reverseSorting
		}
	}
	resGetTaskList, err := client.GetTaskList(tubaClient, reqGetTaskList)
	if err != nil {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetTaskList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.TaskListResult{
		Result: string(resGetTaskList.Result),
		Errors: Errors,
	}, nil
}
