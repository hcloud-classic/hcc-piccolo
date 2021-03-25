package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/pb"

	"google.golang.org/grpc"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// AllTask : Get task list with provided options
func AllTask(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	sortBy, sortByOk := args["sort_by"].(string)
	reverseSorting, reverseSortingOk := args["reverse_sorting"].(bool)
	hideThreads, hideThreadsOk := args["hide_threads"].(bool)

	if !serverUUIDOk {
		return model.TaskListResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need a server_uuid argument")}, nil
	}

	resGetNodeList, err := client.RC.GetNodeList(&pb.ReqGetNodeList{
		Node: &pb.Node{
			ServerUUID: serverUUID,
			NodeNum:    1,
		},
	})
	if err != nil {
		return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	if len(resGetNodeList.Node) == 0 {
		return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "node not found")}, nil
	}

	var conn *grpc.ClientConn
	tubaClient, err := client.InitTuba(resGetNodeList.Node[0].NodeIP, int(config.Tuba.ServerPort), conn)
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
	if hideThreadsOk {
		reqGetTaskList.HideThreads = hideThreads
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
