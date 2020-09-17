package mutationparser

import (
	"github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

func pbServerToModelServer(server *rpcviolin.Server, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Server {
	var createdAt time.Time
	if server.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(server.CreatedAt)
		if err != nil {
			return &model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelServer := &model.Server{
		UUID:       server.UUID,
		SubnetUUID: server.SubnetUUID,
		OS:         server.OS,
		ServerName: server.ServerName,
		ServerDesc: server.ServerDesc,
		CPU:        int(server.CPU),
		Memory:     int(server.Memory),
		DiskSize:   int(server.DiskSize),
		Status:     server.Status,
		UserUUID:   server.UserUUID,
		CreatedAt:  createdAt,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack).ConvertReportForm()
		modelServer.Errors = *hccErrStack
	}

	return modelServer
}

func pbServerNodeToModelServerNode(serverNode *rpcviolin.ServerNode, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.ServerNode {
	var createdAt time.Time
	if serverNode.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(serverNode.CreatedAt)
		if err != nil {
			return &model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelServerNode := &model.ServerNode{
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CreatedAt:  createdAt,
		Errors:     *errors.NewHccErrorStack(),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack).ConvertReportForm()
		modelServerNode.Errors = *hccErrStack
	}

	return modelServerNode
}

// CreateServer : Create a server
func CreateServer(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOK := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	nrNode, nrNodeOk := args["nr_node"].(int)

	var reqCreateServer rpcviolin.ReqCreateServer
	var reqServer rpcviolin.Server
	reqCreateServer.Server = &reqServer

	if subnetUUIDOk {
		reqCreateServer.Server.SubnetUUID = subnetUUID
	}
	if osOK {
		reqCreateServer.Server.OS = os
	}
	if serverNameOk {
		reqCreateServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqCreateServer.Server.ServerDesc = serverDesc
	}
	if cpuOk {
		reqCreateServer.Server.CPU = int32(cpu)
	}
	if memoryOk {
		reqCreateServer.Server.Memory = int32(memory)
	}
	if diskSizeOk {
		reqCreateServer.Server.DiskSize = int32(diskSize)
	}
	if userUUIDOk {
		reqCreateServer.Server.UserUUID = userUUID
	}
	if nrNodeOk {
		reqCreateServer.NrNode = int32(nrNode)
	}

	resCreateServer, err := client.RC.CreateServer(&reqCreateServer)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelServer := pbServerToModelServer(resCreateServer.Server, &resCreateServer.HccErrorStack)

	return *modelServer, nil
}

// UpdateServer : Update the infos of the server
func UpdateServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOK := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)

	var reqUpdateServer rpcviolin.ReqUpdateServer
	var reqServer rpcviolin.Server
	reqUpdateServer.Server = &reqServer

	reqUpdateServer.Server.UUID = requestedUUID
	if subnetUUIDOk {
		reqUpdateServer.Server.SubnetUUID = subnetUUID
	}
	if osOK {
		reqUpdateServer.Server.OS = os
	}
	if serverNameOk {
		reqUpdateServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqUpdateServer.Server.ServerDesc = serverDesc
	}
	if cpuOk {
		reqUpdateServer.Server.CPU = int32(cpu)
	}
	if memoryOk {
		reqUpdateServer.Server.Memory = int32(memory)
	}
	if diskSizeOk {
		reqUpdateServer.Server.DiskSize = int32(diskSize)
	}
	if statusOk {
		reqUpdateServer.Server.Status = status
	}
	if userUUIDOk {
		reqUpdateServer.Server.UserUUID = userUUID
	}

	resUpdateServer, err := client.RC.UpdateServer(&reqUpdateServer)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelServer := pbServerToModelServer(resUpdateServer.Server, &resUpdateServer.HccErrorStack)

	return *modelServer, nil
}

// DeleteServer : Delete the server
func DeleteServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	var server model.Server
	resDeleteServer, err := client.RC.DeleteServer(requestedUUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}
	server.UUID = resDeleteServer.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteServer.HccErrorStack).ConvertReportForm()
	server.Errors = *hccErrStack

	return server, nil
}

// CreateServerNode : Create a info of server's node
func CreateServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	var reqCreateServerNode rpcviolin.ReqCreateServerNode
	if serverUUIDOk {
		reqCreateServerNode.ServerNode.ServerUUID = serverUUID
	}
	if nodeUUIDOk {
		reqCreateServerNode.ServerNode.NodeUUID = nodeUUID
	}

	resCreateServerNode, err := client.RC.CreateServerNode(&reqCreateServerNode)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelServerNode := pbServerNodeToModelServerNode(resCreateServerNode.ServerNode, &resCreateServerNode.HccErrorStack)

	return *modelServerNode, nil
}

// DeleteServerNode : Delete a info of server's node
func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	var serverNode model.ServerNode
	resDeleteServerNode, err := client.RC.DeleteServerNode(requestedUUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}
	serverNode.UUID = resDeleteServerNode.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteServerNode.HccErrorStack).ConvertReportForm()
	serverNode.Errors = *hccErrStack

	return serverNode, nil
}
