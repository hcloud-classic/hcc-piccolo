package pbtomodel

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

// PbServerToModelServer : Change server of proto type to model
func PbServerToModelServer(server *rpcviolin.Server, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Server {
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
		modelServer.Errors = *hccErrStack.ConvertReportForm()
		if len(modelServer.Errors) != 0 && modelServer.Errors[0].ErrCode == 0 {
			modelServer.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelServer
}

// PbServerNodeToModelServerNode : Change serverNode of proto type to model
func PbServerNodeToModelServerNode(serverNode *rpcviolin.ServerNode, node *rpcflute.Node,
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
		UUID:       serverNode.UUID,
		ServerUUID: serverNode.ServerUUID,
		NodeUUID:   serverNode.NodeUUID,
		CPUModel:   nodeDetail.CPUModel,
		CreatedAt:  createdAt,
		Errors:     *errors.NewHccErrorStack(),
	}

	if node != nil && nodeDetail != nil {
		modelServerNode.CPUProcessors = int(nodeDetail.CPUProcessors)
		modelServerNode.CPUCores = int(node.CPUCores)
		modelServerNode.CPUThreads = int(nodeDetail.CPUThreads)
		modelServerNode.Memory = int(node.Memory)
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelServerNode.Errors = *hccErrStack.ConvertReportForm()
		if len(modelServerNode.Errors) != 0 && modelServerNode.Errors[0].ErrCode == 0 {
			modelServerNode.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelServerNode
}
