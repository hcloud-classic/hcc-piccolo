package pbtomodel

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

// PbNodeToModelNode : Change node of proto type to model
func PbNodeToModelNode(node *rpcflute.Node, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Node {
	var createdAt time.Time
	if node.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(node.CreatedAt)
		if err != nil {
			return &model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelNode := &model.Node{
		UUID:            node.UUID,
		ServerUUID:      node.ServerUUID,
		BmcMacAddr:      node.BmcMacAddr,
		BmcIP:           node.BmcIP,
		BmcIPSubnetMask: node.BmcIPSubnetMask,
		PXEMacAddr:      node.PXEMacAddr,
		Status:          node.Status,
		CPUCores:        int(node.CPUCores),
		Memory:          int(node.Memory),
		Description:     node.Description,
		RackNumber:      int(node.RackNumber),
		CreatedAt:       createdAt,
		Active:          int(node.Active),
		ForceOff:        node.ForceOff,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNode.Errors = *hccErrStack.ConvertReportForm()
	}

	return modelNode
}

// PbNodeDetailToModelNodeDetail : Change nodeDetail of proto type to model
func PbNodeDetailToModelNodeDetail(nodeDetail *rpcflute.NodeDetail, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.NodeDetail {
	modelNodeDetail := &model.NodeDetail{
		NodeUUID:      nodeDetail.NodeUUID,
		CPUModel:      nodeDetail.CPUModel,
		CPUProcessors: int(nodeDetail.CPUProcessors),
		CPUThreads:    int(nodeDetail.CPUThreads),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNodeDetail.Errors = *hccErrStack.ConvertReportForm()
	}

	return modelNodeDetail
}
