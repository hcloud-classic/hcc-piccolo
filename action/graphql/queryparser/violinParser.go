package queryparser

import (
	"github.com/golang/protobuf/ptypes"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

func pbServerToModelServer(server *rpcviolin.Server) *model.Server {
	createdAt, err := ptypes.Timestamp(server.CreatedAt)
	if err != nil {
		return &model.Server{Errors: errors.ReturnHccError(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
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
		Errors:     *errors.NewHccErrorStack(),
	}

	return modelServer
}

func pbServerNodeToModelServerNode(serverNode *rpcviolin.ServerNode) *model.ServerNode {
	createdAt, err := ptypes.Timestamp(serverNode.CreatedAt)
	if err != nil {
		return &model.ServerNode{Errors: errors.ReturnHccError(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
	}

	modelServerNode := &model.ServerNode{
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CreatedAt:  createdAt,
		Errors:     *errors.NewHccErrorStack(),
	}

	return modelServerNode
}

// Server : Get infos of the server
func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Server{Errors: errors.ReturnHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	server, err := client.RC.GetServer(uuid)
	if err != nil {
		return model.Server{Errors: errors.ReturnHccError(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return *pbServerToModelServer(server), nil
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
		return model.ServerList{Errors: errors.ReturnHccError(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverList []model.Server
	for _, pServer := range resListServer.Server {
		modelServer := pbServerToModelServer(pServer)
		serverList = append(serverList, *modelServer)
	}

	return model.ServerList{Servers: serverList}, nil
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
		return model.Server{Errors: errors.ReturnHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	serverNode, err := client.RC.GetServerNode(uuid)
	if err != nil {
		return model.Server{Errors: errors.ReturnHccError(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return *pbServerNodeToModelServerNode(serverNode), nil
}

// ListServerNode : Get serverNode list with provided options
func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeList{Errors: errors.ReturnHccError(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	var reqListServerNode rpcviolin.ReqGetServerNodeList
	reqListServerNode.ServerUUID = serverUUID

	resListServerNode, err := client.RC.GetServerNodeList(&reqListServerNode)
	if err != nil {
		return model.ServerNodeList{Errors: errors.ReturnHccError(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverNodeList []model.ServerNode
	for _, pServerNode := range resListServerNode.ServerNode {
		modelServerNode := pbServerNodeToModelServerNode(pServerNode)
		serverNodeList = append(serverNodeList, *modelServerNode)
	}

	return model.ServerNodeList{ServerNodes: serverNodeList}, nil
}

// AllServerNode : Get serverNode list with provided options (Just call ListServerNode())
func AllServerNode(args map[string]interface{}) (interface{}, error) {
	return ListServerNode(args)
}

// NumServerNode : Get number of serverNodes
func NumServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeNum{Errors: errors.ReturnHccError(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	num, err := client.RC.GetServerNodeNum(serverUUID)
	if err != nil {
		return model.ServerNodeNum{Errors: errors.ReturnHccError(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelServerNodeNum model.ServerNodeNum
	modelServerNodeNum.Number = num

	return modelServerNodeNum, nil
}
