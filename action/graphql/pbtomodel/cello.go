package pbtomodel

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"strconv"
	"strings"
	"time"
)

// PbVolumeToModelVolume : Change volume of proto type to model
func PbVolumeToModelVolume(volume *rpccello.Volume, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Volume {
	var createdAt time.Time
	if volume.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(volume.CreatedAt)
		if err != nil {
			return &model.Volume{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}
	convSize := strings.Split(volume.Size, "G")
	onlySize, _ := strconv.Atoi(convSize[0])
	modelVolume := &model.Volume{
		UUID:       volume.UUID,
		Size:       onlySize,
		Filesystem: volume.Filesystem,
		ServerUUID: volume.ServerUUID,
		UseType:    volume.UseType,
		UserUUID:   volume.UserUUID,
		NetworkIP:  volume.Network_IP,
		GatewayIP:  volume.GatewayIp,
		LunNum:     int(volume.Lun),
		Pool:       volume.Pool,
		CreatedAt:  createdAt,
	}
	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelVolume.Errors = *hccErrStack.ConvertReportForm()
		if modelVolume.Errors[0].ErrCode == 0 {
			modelVolume.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelVolume
}
