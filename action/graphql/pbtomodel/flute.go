package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/pb"
)

// PbNodeToModelNode : Change node of proto type to model
func PbNodeToModelNode(node *pb.Node, hccGrpcErrStack *pb.HccErrorStack) *model.Node {
	var nicSpeed = "Unknown"

	switch node.NicSpeedMbps {
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
	}

	modelNode := &model.Node{
		UUID:             node.UUID,
		NodeName:         node.NodeName,
		GroupID:          node.GroupID,
		NodeNum:          int64(node.NodeNum),
		NodeIP:           node.NodeIP,
		ServerUUID:       node.ServerUUID,
		BmcMacAddr:       node.BmcMacAddr,
		BmcIP:            node.BmcIP,
		BmcIPSubnetMask:  node.BmcIPSubnetMask,
		PXEMacAddr:       node.PXEMacAddr,
		Status:           node.Status,
		CPUCores:         int(node.CPUCores),
		Memory:           int(node.Memory),
		NICSpeed:         nicSpeed,
		Description:      node.Description,
		RackNumber:       int(node.RackNumber),
		Active:           int(node.Active),
		CreatedAt:        node.CreatedAt.AsTime(),
		ForceOff:         node.ForceOff,
		IPMIUserID:       node.IpmiUserID,
		IPMIUserPassword: node.IpmiUserPasswordEncryptedBytes,
		Errors:           nil,
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
		NicDetailData:  nodeDetail.NicDetailData,
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
