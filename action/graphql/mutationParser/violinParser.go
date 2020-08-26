package mutationParser

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/model"
)

func pbServerToModelServer(server *rpcviolin.Server) (*model.Server, error) {
	createdAt, err := ptypes.Timestamp(server.CreatedAt)
	if err != nil {
		return nil, err
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

	return modelServer, nil
}

func pbServerNodeToModelServerNode(serverNode *rpcviolin.ServerNode) (*model.ServerNode, error) {
	createdAt, err := ptypes.Timestamp(serverNode.CreatedAt)
	if err != nil {
		return nil, err
	}

	modelServerNode := &model.ServerNode{
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CreatedAt:  createdAt,
	}

	return modelServerNode, nil
}

func CreateServer(args map[string]interface{}) (interface{}, error) {
	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	userUUID, _ := args["user_uuid"].(string)
	nrNode, _ := args["nr_node"].(int)

	var reqCreateServer rpcviolin.ReqCreateServer

	reqCreateServer.Server = &rpcviolin.Server{
		SubnetUUID: subnetUUID,
		OS:         os,
		ServerName: serverName,
		ServerDesc: serverDesc,
		CPU:        int32(cpu),
		Memory:     int32(memory),
		DiskSize:   int32(diskSize),
		UserUUID:   userUUID,
	}
	reqCreateServer.NrNode = int32(nrNode)

	resCreateServer, err := client.RC.CreateServer(&reqCreateServer)
	if err != nil {
		return nil, err
	}

	modelServer, err := pbServerToModelServer(resCreateServer.Server)

	return modelServer, nil
}

func UpdateServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	status, _ := args["status"].(string)
	userUUID, _ := args["user_uuid"].(string)

	var server rpcviolin.Server
	server.UUID = requestedUUID
	server.SubnetUUID = subnetUUID
	server.OS = os
	server.ServerName = serverName
	server.ServerDesc = serverDesc
	server.CPU = int32(cpu)
	server.Memory = int32(memory)
	server.DiskSize = int32(diskSize)
	server.Status = status
	server.UUID = userUUID

	resUpdateServer, err := client.RC.UpdateServer(&rpcviolin.ReqUpdateServer{
		Server: &server,
	})
	if err != nil {
		return nil, err
	}

	modelServer, err := pbServerToModelServer(resUpdateServer.Server)

	return modelServer, nil
}

func DeleteServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var server model.Server
	uuid, err := client.RC.DeleteServer(requestedUUID)
	if err != nil {
		return nil, err
	}
	server.UUID = uuid

	return server, nil
}

func CreateServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, _ := args["server_uuid"].(string)
	nodeUUID, _ := args["node_uuid"].(string)

	var reqCreateServerNode rpcviolin.ReqCreateServerNode

	reqCreateServerNode.ServerNode = &rpcviolin.ServerNode{
		ServerUUID: serverUUID,
		NodeUUID:   nodeUUID,
	}

	resCreateServerNode, err := client.RC.CreateServerNode(&reqCreateServerNode)
	if err != nil {
		return nil, err
	}

	modelServerNode, err := pbServerNodeToModelServerNode(resCreateServerNode.ServerNode)

	return modelServerNode, nil
}

func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var serverNode model.ServerNode
	uuid, err := client.RC.DeleteServerNode(requestedUUID)
	if err != nil {
		return nil, err
	}
	serverNode.UUID = uuid

	return serverNode, nil
}
