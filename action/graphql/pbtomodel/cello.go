package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"innogrid.com/hcloud-classic/pb"
)

// PbVolumeToModelVolume : Change volume of proto type to model
func PbVolumeToModelVolume(volume *pb.Volume, hccGrpcErrStack *pb.HccErrorStack) *model.Volume {
	if volume.CreatedAt == nil {
		volume.CreatedAt = timestamppb.New(time.Time{})
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
		GatewayIP:  volume.Gateway_IP,
		LunNum:     int(volume.Lun),
		Pool:       volume.Pool,
		State:      volume.State,
		GroupID:    int(volume.GroupID),
		CreatedAt:  volume.CreatedAt.AsTime(),
	}
	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelVolume.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelVolume.Errors) != 0 && modelVolume.Errors[0].ErrCode == 0 {
			modelVolume.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelVolume
}

// PbPoolToModelPool : Change volume of proto type to model
func PbPoolToModelPool(pool *pb.Pool, hccGrpcErrStack *pb.HccErrorStack) *model.Pool {
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
		modelPool.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelPool.Errors) != 0 && modelPool.Errors[0].ErrCode == 0 {
			modelPool.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	}

	return modelPool
}

// // PbVolumeListToModelVolumeList : Change volume of proto type to model
// func PbVolumeListToModelVolumeList(pbVolumeList []*pb.Volume, hccGrpcErrStack *[]*rpcmsgType.HccError) []*model.Volume {

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
// 			modelPool.Errors = errconv.ReturnHccEmptyErrorPiccolo()
// 		}
// 	}

// 	return modelPool
// }
