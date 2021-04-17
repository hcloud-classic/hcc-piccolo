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

func getNICSpeed(nicSpeedMbps int32) string {
	var nicSpeed string

	switch nicSpeedMbps {
	case 10:
		nicSpeed = "10Mbps"
	case 100:
		nicSpeed = "100Mbps"
	case 1000:
		nicSpeed = "1Gbps"
	case 2500:
		nicSpeed = "2.5Gbps"
	case 5000:
		nicSpeed = "5Gbps"
	case 10000:
		nicSpeed = "10Gbps"
	case 20000:
		nicSpeed = "20Gbps"
	case 40000:
		nicSpeed = "40Gbps"
	default:
		nicSpeed = "Unknown"
	}

	return nicSpeed
}

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
		UUID:            node.UUID,
		NodeName:        node.NodeName,
		GroupID:         node.GroupID,
		NodeNum:         int64(node.NodeNum),
		NodeIP:          node.NodeIP,
		ServerUUID:      node.ServerUUID,
		BmcMacAddr:      node.BmcMacAddr,
		BmcIP:           node.BmcIP,
		BmcIPSubnetMask: node.BmcIPSubnetMask,
		PXEMacAddr:      node.PXEMacAddr,
		Status:          node.Status,
		CPUCores:        int(node.CPUCores),
		Memory:          int(node.Memory),
		NICModel:        node.NicModel,
		NICSpeed:        getNICSpeed(node.NicSpeedMbps),
		BMCNICModel:     node.BmcNicModel,
		BMCNICSpeed:     getNICSpeed(node.BmcNicSpeedMbps),
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
		NodeUUID:       nodeDetail.NodeUUID,
		NodeDetailData: nodeDetail.NodeDetailData,
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
