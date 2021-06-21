package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
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
		if len(modelVolume.Errors) != 0 && modelVolume.Errors[0].ErrCode == 0 {
			modelVolume.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelVolume
}

// PbPoolToModelPool : Change volume of proto type to model
func PbPoolToModelPool(pool *rpccello.Pool, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Pool {

	modelPool := &model.Pool{
		UUID:          pool.UUID,
		Size:          pool.Size,
		Free:          pool.Free,
		Capacity:      pool.Capacity,
		Health:        pool.Health,
		Name:          pool.Name,
		AvailableSize: pool.AvailableSize,
		Action:        pool.Action,
		Used:          pool.Used,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelPool.Errors = *hccErrStack.ConvertReportForm()
		if len(modelPool.Errors) != 0 && modelPool.Errors[0].ErrCode == 0 {
			modelPool.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelPool
}

// // PbVolumeListToModelVolumeList : Change volume of proto type to model
// func PbVolumeListToModelVolumeList(pbVolumeList []*rpccello.Volume, hccGrpcErrStack *[]*rpcmsgType.HccError) []*model.Volume {

// 	var modelVolumeList []*model.Volume

// 	for _, args := range pbVolumeList {
// 		strSize,_:=strconv.Atoi(args.Size)
// 		tempPbVol := model.Volume{
// 			UUID:       args.UUID,
// 			Size:       strSize,
// 			Filesystem: args.Filesystem,
// 			ServerUUID: args.ServerUUID,
// 			UseType:    args.UseType,
// 			UserUUID:   args.UserUUID,
// 			Pool:       args.Pool,
// 			LunNum:        int(args.Lun),
// 			CreatedAt: time.Unix(args.CreatedAt.Seconds, int64(args.CreatedAt.Nanos)).UTC(),
// 		}
// 		modelVolumeList = append(modelVolumeList, &tempPbVol)
// 	}
// 	if hccGrpcErrStack != nil {
// 		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
// 		modelVolumeList.Errors = *hccErrStack.ConvertReportForm()
// 		if len(modelPool.Errors) != 0 && modelPool.Errors[0].ErrCode == 0 {
// 			modelPool.Errors = errors.ReturnHccEmptyErrorPiccolo()
// 		}
// 	}

// 	return modelPool
// }
