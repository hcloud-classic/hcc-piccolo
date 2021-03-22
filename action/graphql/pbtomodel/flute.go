package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// PbNodeToModelNode : Change node of proto type to model
func PbNodeToModelNode(node *pb.Node, hccGrpcErrStack *pb.HccErrorStack) *model.Node {
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
			return &model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelNode := &model.Node{
		GroupID:         node.GroupID,
		UUID:            node.UUID,
		ServerUUID:      node.ServerUUID,
		BmcMacAddr:      node.BmcMacAddr,
		BmcIP:           node.BmcIP,
		BmcIPSubnetMask: node.BmcIPSubnetMask,
		PXEMacAddr:      node.PXEMacAddr,
		Status:          node.Status,
		CPUCores:        int(node.CPUCores),
		Memory:          int(node.Memory),
		NICSpeedMbps:    int(node.NicSpeedMbps),
		Description:     node.Description,
		RackNumber:      int(node.RackNumber),
		ChargeCPU:       int(node.ChargeCPU),
		ChargeMemory:    int(node.ChargeMemory),
		ChargeNIC:       int(node.ChargeNIC),
		Active:          int(node.Active),
		CreatedAt:       createdAt,
		ForceOff:        node.ForceOff,
		Errors:          nil,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNode.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelNode.Errors) != 0 && modelNode.Errors[0].ErrCode == 0 {
			modelNode.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelNode
}

// PbNodeDetailToModelNodeDetail : Change nodeDetail of proto type to model
func PbNodeDetailToModelNodeDetail(nodeDetail *pb.NodeDetail, hccGrpcErrStack *pb.HccErrorStack) *model.NodeDetail {
	modelNodeDetail := &model.NodeDetail{
		NodeUUID:      nodeDetail.NodeUUID,
		CPUModel:      nodeDetail.CPUModel,
		CPUProcessors: int(nodeDetail.CPUProcessors),
		CPUThreads:    int(nodeDetail.CPUThreads),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNodeDetail.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelNodeDetail.Errors) != 0 && modelNodeDetail.Errors[0].ErrCode == 0 {
			modelNodeDetail.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelNodeDetail
}
