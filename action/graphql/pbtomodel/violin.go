package pbtomodel

import (
	"encoding/json"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// PbServerToModelServer : Change server of proto type to model
func PbServerToModelServer(server *pb.Server, hccGrpcErrStack *pb.HccErrorStack) *model.Server {
	modelServer := &model.Server{
		UUID:       server.UUID,
		GroupID:    server.GroupID,
		SubnetUUID: server.SubnetUUID,
		OS:         server.OS,
		ServerName: server.ServerName,
		ServerDesc: server.ServerDesc,
		CPU:        int(server.CPU),
		Memory:     int(server.Memory),
		DiskSize:   int(server.DiskSize),
		Status:     server.Status,
		UserUUID:   server.UserUUID,
		CreatedAt:  server.CreatedAt.AsTime(),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelServer.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelServer.Errors) != 0 && modelServer.Errors[0].ErrCode == 0 {
			modelServer.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelServer
}

// PbServerNodeToModelServerNode : Change serverNode of proto type to model
func PbServerNodeToModelServerNode(serverNode *pb.ServerNode, node *pb.Node,
	nodeDetail *pb.NodeDetail, hccGrpcErrStack *pb.HccErrorStack) *model.ServerNode {
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
			return &model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	var nodeDetailJSON model.NodeDetailData
	err := json.Unmarshal([]byte(nodeDetail.NodeDetailData), &nodeDetailJSON)
	if err != nil {
		return &model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}
	}
	modelServerNode := &model.ServerNode{
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CPUModel:   nodeDetailJSON.CPUs[0].Model,
		CreatedAt:  createdAt,
		Errors:     errconv.HccErrorToPiccoloHccErr(*hcc_errors.NewHccErrorStack()),
	}

	if node != nil {
		modelServerNode.NodeName = node.NodeName
		modelServerNode.RackNumber = int(node.RackNumber)
	}

	if node != nil && nodeDetail != nil {
		modelServerNode.CPUProcessors = len(nodeDetailJSON.CPUs)
		modelServerNode.CPUCores = int(node.CPUCores)
		var cpuThreads = 0
		for _, cpu := range nodeDetailJSON.CPUs {
			cpuThreads += cpu.Threads
		}
		modelServerNode.CPUThreads = cpuThreads
		modelServerNode.Memory = int(node.Memory)
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelServerNode.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelServerNode.Errors) != 0 && modelServerNode.Errors[0].ErrCode == 0 {
			modelServerNode.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelServerNode
}
