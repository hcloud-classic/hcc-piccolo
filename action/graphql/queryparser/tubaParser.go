package queryparser

import (
	"google.golang.org/grpc"
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/iputil"
	"hcc/piccolo/model"

	"github.com/hcloud-classic/hcc_errors"
	"github.com/hcloud-classic/pb"
)

// ListTask : Get task list with provided options
func ListTask(args map[string]interface{}) (interface{}, error) {
	serverAddress, serverAddressOk := args["server_address"].(string)
	serverPort, serverPortOk := args["server_port"].(int)

	cmd, cmdOk := args["cmd"].(string)
	state, stateOk := args["state"].(string)
	pid, pidOk := args["pid"].(int)
	ppid, ppidOk := args["ppid"].(int)
	pgid, pgidOk := args["pgid"].(int)
	sid, sidOk := args["sid"].(int)
	priority, priorityOk := args["priority"].(int)
	nice, niceOk := args["nice"].(int)
	numThreads, numThreadsOk := args["num_threads"].(int)

	epmType, epmTypeOk := args["epm_type"].(string)
	epmSource, epmSourceOk := args["epm_source"].(int)
	epmTarget, epmTargetOk := args["epm_target"].(int)

	if !serverAddressOk || !serverPortOk {
		return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need server_address and server_port arguments")}, nil
	}

	ipCheck := iputil.CheckValidIP(serverAddress)
	if ipCheck == nil {
		return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"invalid server address")}, nil
	}

	if serverPort < 1 || serverPort > 65535 {
		return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"out of port number range")}, nil
	}

	var reqListTask pb.ReqGetTaskList
	var reqTask pb.Task
	reqListTask.Task = &reqTask

	if cmdOk {
		reqListTask.Task.CMD = cmd
	}
	if stateOk {
		reqListTask.Task.State = state
	}
	if pidOk {
		reqListTask.Task.PID = int64(pid)
	}
	if ppidOk {
		reqListTask.Task.PPID = int64(ppid)
	}
	if pgidOk {
		reqListTask.Task.PGID = int64(pgid)
	}
	if sidOk {
		reqListTask.Task.SID = int64(sid)
	}
	if priorityOk {
		reqListTask.Task.Priority = int64(priority)
	}
	if niceOk {
		reqListTask.Task.Nice = int64(nice)
	}
	if numThreadsOk {
		reqListTask.Task.NumThreads = int64(numThreads)
	}
	if epmTypeOk {
		reqListTask.Task.EPMType = epmType
	}
	if epmSourceOk {
		reqListTask.Task.EPMSource = int64(epmSource)
	}
	if epmTargetOk {
		reqListTask.Task.EPMTarget = int64(epmTarget)
	}

	var conn *grpc.ClientConn
	tubaClient, err := client.InitTuba(serverAddress, serverPort, conn)
	if err != nil {
		return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	defer client.CloseTuba(conn)

	resTaskProcess, err := client.GetTaskList(tubaClient, &reqListTask)
	if err != nil {
		return model.TaskList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var processList []model.Task
	for _, pProcess := range resTaskProcess.Tasks {
		modelProcess := pbtomodel.PbTaskToModelTask(pProcess)
		processList = append(processList, *modelProcess)
	}

	hccErrStack := errconv.GrpcStackToHcc(resTaskProcess.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.TaskList{
		Tasks:                processList,
		TotalTasks:           int(resTaskProcess.TotalTasks),
		TotalMemUsage:        resTaskProcess.TotalMemUsage,
		TotalMem:             resTaskProcess.TotalMem,
		TotalMemUsagePercent: resTaskProcess.TotalMemUsagePercent,
		TotalCPUUsage:        resTaskProcess.TotalCPUUsage,
		Errors:               Errors,
	}, nil
}

// AllTask : Get task list with provided options (Just call ListTask())
func AllTask(args map[string]interface{}) (interface{}, error) {
	return ListTask(args)
}
