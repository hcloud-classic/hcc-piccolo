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

// PbSubnetToModelSubnet : Change subnet of proto type to model
func PbSubnetToModelSubnet(subnet *pb.Subnet, hccGrpcErrStack *[]*pb.HccError) *model.Subnet {
	var createdAt time.Time
	if subnet.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(subnet.CreatedAt)
		if err != nil {
			return &model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelSubnet := &model.Subnet{
		UUID:           subnet.UUID,
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
		CreatedAt:      createdAt,
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
