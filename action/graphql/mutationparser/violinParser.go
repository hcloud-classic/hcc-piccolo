package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

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
	} else {
		reqCreateServer.Server.CPU = 0
	}
	if memoryOk {
		reqCreateServer.Server.Memory = int32(memory)
	} else {
		reqCreateServer.Server.Memory = 0
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
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resCreateServer.Server, &resCreateServer.HccErrorStack)

	return *modelServer, nil
}

// UpdateServer : Update the infos of the server
func UpdateServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
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
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resUpdateServer.Server, &resUpdateServer.HccErrorStack)

	return *modelServer, nil
}

// DeleteServer : Delete the server
func DeleteServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var server model.Server
	resDeleteServer, err := client.RC.DeleteServer(requestedUUID)
	if err != nil {
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	server.UUID = resDeleteServer.Server.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteServer.HccErrorStack)
	server.Errors = *hccErrStack.ConvertReportForm()
	if len(server.Errors) != 0 && server.Errors[0].ErrCode == 0 {
		server.Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

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
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServerNode := pbtomodel.PbServerNodeToModelServerNode(resCreateServerNode.ServerNode, nil, nil, &resCreateServerNode.HccErrorStack)

	return *modelServerNode, nil
}

// DeleteServerNode : Delete a info of server's node
func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var serverNode model.ServerNode
	resDeleteServerNode, err := client.RC.DeleteServerNode(requestedUUID)
	if err != nil {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	serverNode.UUID = resDeleteServerNode.ServerNode.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteServerNode.HccErrorStack)
	serverNode.Errors = *hccErrStack.ConvertReportForm()
	if len(serverNode.Errors) != 0 && serverNode.Errors[0].ErrCode == 0 {
		serverNode.Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return serverNode, nil
}
