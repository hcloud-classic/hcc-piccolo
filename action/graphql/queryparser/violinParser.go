package queryparser

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcflute"
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
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelServer.Errors = *hccErrStack
	}

	return modelServer
}

func pbServerNodeToModelServerNode(serverNode *rpcviolin.ServerNode, node *rpcflute.Node,
	nodeDetail *rpcflute.NodeDetail, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.ServerNode {
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
		UUID:          serverNode.UUID,
		ServerUUID:    serverNode.ServerUUID,
		NodeUUID:      serverNode.NodeUUID,
		CPUModel:      nodeDetail.CPUModel,
		CPUProcessors: int(nodeDetail.CPUProcessors),
		CPUCores:      int(node.CPUCores),
		CPUThreads:    int(nodeDetail.CPUThreads),
		Memory:        int(node.Memory),
		CreatedAt:     createdAt,
		Errors:        *errors.NewHccErrorStack(),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelServerNode.Errors = *hccErrStack
	}

	return modelServerNode
}

// Server : Get infos of the server
func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetServer, err := client.RC.GetServer(uuid)
	if err != nil {
		return model.Server{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return *pbServerToModelServer(resGetServer.Server, &resGetServer.HccErrorStack), nil
}

// ListServer : Get server list with provided options
func ListServer(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
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

	if uuidOk {
		reqListServer.Server.UUID = uuid
	}
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
		return model.ServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverList []model.Server
	for _, pServer := range resListServer.Server {
		modelServer := pbServerToModelServer(pServer, nil)
		serverList = append(serverList, *modelServer)
	}

	hccErrStack := errconv.GrpcStackToHcc(&resListServer.HccErrorStack)

	return model.ServerList{Servers: serverList, Errors: *hccErrStack}, nil
}

// AllServer : Get server list with provided options (Just call ListServer())
func AllServer(args map[string]interface{}) (interface{}, error) {
	return ListServer(args)
}

// NumServer : Get number of servers
func NumServer() (interface{}, error) {
	resGetServerNum, err := client.RC.GetServerNum()
	if err != nil {
		return nil, err
	}

	var modelServerNum model.ServerNum
	modelServerNum.Number = int(resGetServerNum.Num)
	hccErrStack := errconv.GrpcStackToHcc(&resGetServerNum.HccErrorStack)
	modelServerNum.Errors = *hccErrStack

	return modelServerNum, nil
}

// ServerNode : Get infos of the serverNode
func ServerNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetServerNode, err := client.RC.GetServerNode(uuid)
	if err != nil {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	resGetNode, err := client.RC.GetNode(resGetServerNode.ServerNode.NodeUUID)
	if err != nil {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	resGetNodeDetail, err := client.RC.GetNodeDetail(resGetServerNode.ServerNode.NodeUUID)
	if err != nil {
		return model.ServerNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return *pbServerNodeToModelServerNode(resGetServerNode.ServerNode, resGetNode.Node, resGetNodeDetail.NodeDetail,
		&resGetServerNode.HccErrorStack), nil
}

// ListServerNode : Get serverNode list with provided options
func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	var reqListServerNode rpcviolin.ReqGetServerNodeList
	reqListServerNode.ServerUUID = serverUUID

	resListServerNode, err := client.RC.GetServerNodeList(&reqListServerNode)
	if err != nil {
		return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverNodeList []model.ServerNode
	for _, pServerNode := range resListServerNode.ServerNode {
		resGetNode, err := client.RC.GetNode(pServerNode.NodeUUID)
		if err != nil {
			return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		resGetNodeDetail, err := client.RC.GetNodeDetail(pServerNode.NodeUUID)
		if err != nil {
			return model.ServerNodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		modelServerNode := pbServerNodeToModelServerNode(pServerNode, resGetNode.Node, resGetNodeDetail.NodeDetail, nil)
		serverNodeList = append(serverNodeList, *modelServerNode)
	}

	hccErrStack := errconv.GrpcStackToHcc(&resListServerNode.HccErrorStack)

	return model.ServerNodeList{ServerNodes: serverNodeList, Errors: *hccErrStack}, nil
}

// AllServerNode : Get serverNode list with provided options (Just call ListServerNode())
func AllServerNode(args map[string]interface{}) (interface{}, error) {
	return ListServerNode(args)
}

// NumServerNode : Get number of serverNodes
func NumServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resGetServerNodeNum, err := client.RC.GetServerNodeNum(serverUUID)
	if err != nil {
		return model.ServerNodeNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelServerNodeNum model.ServerNodeNum
	modelServerNodeNum.Number = int(resGetServerNodeNum.Num)
	hccErrStack := errconv.GrpcStackToHcc(&resGetServerNodeNum.HccErrorStack)
	modelServerNodeNum.Errors = *hccErrStack

	return modelServerNodeNum, nil
}
