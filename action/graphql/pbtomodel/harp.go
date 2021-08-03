package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/pb"
)

// PbSubnetToModelSubnet : Change subnet of proto type to model
func PbSubnetToModelSubnet(subnet *pb.Subnet, hccGrpcErrStack *pb.HccErrorStack) *model.Subnet {
	modelSubnet := &model.Subnet{
		UUID:           subnet.UUID,
		GroupID:        subnet.GroupID,
		NetworkIP:      subnet.NetworkIP,
		Netmask:        subnet.Netmask,
		Gateway:        subnet.Gateway,
		NextServer:     subnet.NextServer,
		NameServer:     subnet.NameServer,
		DomainName:     subnet.DomainName,
		ServerUUID:     subnet.ServerUUID,
		LeaderNodeUUID: subnet.LeaderNodeUUID,
		OS:             subnet.OS,
		SubnetName:     subnet.SubnetName,
		CreatedAt:      subnet.CreatedAt.AsTime(),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelSubnet.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelSubnet.Errors) != 0 && modelSubnet.Errors[0].ErrCode == 0 {
			modelSubnet.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelSubnet
}
