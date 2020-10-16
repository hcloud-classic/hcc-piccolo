package pbtomodel

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

// PbSubnetToModelSubnet : Change subnet of proto type to model
func PbSubnetToModelSubnet(subnet *rpcharp.Subnet, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Subnet {
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
			return &model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
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
		modelSubnet.Errors = *hccErrStack.ConvertReportForm()
		if len(modelSubnet.Errors) != 0 && modelSubnet.Errors[0].ErrCode == 0 {
			modelSubnet.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelSubnet
}
