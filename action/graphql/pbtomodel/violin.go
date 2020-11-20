package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hcloud-classic/hcc_errors"
	"github.com/hcloud-classic/pb"
)

// PbServerToModelServer : Change server of proto type to model
func PbServerToModelServer(server *pb.Server, hccGrpcErrStack *[]*pb.HccError) *model.Server {
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
			return &model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}
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
		modelServer.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelServer.Errors) != 0 && modelServer.Errors[0].ErrCode == 0 {
			modelServer.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelServer
}

// PbServerNodeToModelServerNode : Change serverNode of proto type to model
func PbServerNodeToModelServerNode(serverNode *pb.ServerNode, node *pb.Node,
	nodeDetail *pb.NodeDetail, hccGrpcErrStack *[]*pb.HccError) *model.ServerNode {
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

	modelServerNode := &model.ServerNode{
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CPUModel:   nodeDetail.CPUModel,
		CreatedAt:  createdAt,
		Errors:     errconv.HccErrorToPiccoloHccErr(*hcc_errors.NewHccErrorStack()),
	}

	if node != nil {
		modelServerNode.RackNumber = int(node.RackNumber)
	}

	if node != nil && nodeDetail != nil {
		modelServerNode.CPUProcessors = int(nodeDetail.CPUProcessors)
		modelServerNode.CPUCores = int(node.CPUCores)
		modelServerNode.CPUThreads = int(nodeDetail.CPUThreads)
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