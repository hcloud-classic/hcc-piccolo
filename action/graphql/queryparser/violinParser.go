package queryparser

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

// Server : Get infos of the server
func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	server, err := client.RC.GetServer(uuid)
	if err != nil {
		return nil, err
	}
	modelServer, err := pbServerToModelServer(server)
	if err != nil {
		return nil, err
	}

	return *modelServer, nil
}

// ListServer : Get server list with provided options
func ListServer(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOk := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListServer rpcviolin.ReqGetServerList
	var reqServer rpcviolin.Server
	reqListServer.Server = &reqServer

	if subnetUUIDOk {
		reqListServer.Server.SubnetUUID = subnetUUID
	}
	if osOk {
		reqListServer.Server.OS = os
	}
	if serverNameOk {
		reqListServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqListServer.Server.ServerDesc = serverDesc
	}
	if cpuOk {
		reqListServer.Server.CPU = int32(cpu)
	}
	if memoryOk {
		reqListServer.Server.Memory = int32(memory)
	}
	if diskSizeOk {
		reqListServer.Server.DiskSize = int32(diskSize)
	}
	if statusOk {
		reqListServer.Server.Status = status
	}
	if userUUIDOk {
		reqListServer.Server.UserUUID = userUUID
	}
	if rowOk {
		reqListServer.Row = int64(row)
	}
	if pageOk {
		reqListServer.Page = int64(page)
	}

	resListServer, err := client.RC.GetServerList(&reqListServer)
	if err != nil {
		return nil, err
	}

	var serverList []model.Server
	for _, pServer := range resListServer.Server {
		modelServer, err := pbServerToModelServer(pServer)
		if err != nil {
			return nil, err
		}
		serverList = append(serverList, *modelServer)
	}

	return serverList, nil
}

// AllServer : Get server list with provided options (Just call ListServer())
func AllServer(args map[string]interface{}) (interface{}, error) {
	return ListServer(args)
}

// NumServer : Get number of servers
func NumServer() (interface{}, error) {
	num, err := client.RC.GetServerNum()
	if err != nil {
		return nil, err
	}

	var modelServerNum model.ServerNum
	modelServerNum.Number = num

	return modelServerNum, nil
}

// ServerNode : Get infos of the serverNode
func ServerNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	serverNode, err := client.RC.GetServerNode(uuid)
	if err != nil {
		return nil, err
	}

	modelServerNode, err := pbServerNodeToModelServerNode(serverNode)
	if err != nil {
		return nil, err
	}

	return *modelServerNode, nil
}

// ListServerNode : Get serverNode list with provided options
func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	var reqListServerNode rpcviolin.ReqGetServerNodeList
	reqListServerNode.ServerUUID = serverUUID

	resListServerNode, err := client.RC.GetServerNodeList(&reqListServerNode)
	if err != nil {
		return nil, err
	}

	var serverNodeList []model.ServerNode
	for _, pServerNode := range resListServerNode.ServerNode {
		modelServerNode, err := pbServerNodeToModelServerNode(pServerNode)
		if err != nil {
			return nil, err
		}
		serverNodeList = append(serverNodeList, *modelServerNode)
	}

	return serverNodeList, nil
}

// AllServerNode : Get serverNode list with provided options (Just call ListServerNode())
func AllServerNode(args map[string]interface{}) (interface{}, error) {
	return ListServerNode(args)
}

// NumServerNode : Get number of serverNodes
func NumServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	num, err := client.RC.GetServerNodeNum(serverUUID)
	if err != nil {
		return nil, err
	}

	var modelServerNodeNum model.ServerNodeNum
	modelServerNodeNum.Number = num

	return modelServerNodeNum, nil
}
