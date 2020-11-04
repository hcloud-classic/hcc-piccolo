package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/sqlite/serveractions"
	"hcc/piccolo/model"
	"strconv"
)

// VolumeHandle : oboe to cello
func VolumeHandle(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	fileSystem, fileSystemOK := args["filesystem"].(string)
	diskSize, diskSizeOk := args["size"].(int)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	useType, useTypeOk := args["use_type"].(string)
	networkIP, networkIPOk := args["network_ip"].(string)
	gatewayIP, gatewayIPOk := args["gateway_ip"].(string)
	lunNum, lunNumOk := args["lun_num"].(string)
	pool, poolOk := args["pool"].(string)
	action, actionOk := args["action"].(string)
	var modelVolume *model.Volume
	var reqVolumeHandle rpccello.ReqVolumeHandler
	var reqVolume rpccello.Volume
	reqVolumeHandle.Volume = &reqVolume

	if UUIDOk {
		reqVolumeHandle.Volume.UUID = UUID
	}

	if serverUUIDOk {
		reqVolumeHandle.Volume.ServerUUID = serverUUID
	}

	if fileSystemOK {
		reqVolumeHandle.Volume.Filesystem = fileSystem
	}
	strSize := strconv.Itoa(diskSize)
	if diskSizeOk {
		reqVolumeHandle.Volume.Size = strSize
	}

	if userUUIDOk {
		reqVolumeHandle.Volume.UserUUID = userUUID
	}
	if useTypeOk {
		reqVolumeHandle.Volume.UseType = useType
	}
	if networkIPOk {
		reqVolumeHandle.Volume.Network_IP = networkIP
	}
	if gatewayIPOk {
		reqVolumeHandle.Volume.GatewayIp = gatewayIP
	}
	if lunNumOk {
		convInt, _ := strconv.Atoi(lunNum)
		reqVolumeHandle.Volume.Lun = int64(convInt)
	}
	if poolOk {
		reqVolumeHandle.Volume.Pool = pool
	}
	if actionOk {
		reqVolumeHandle.Volume.Action = action
	}

	if reqVolumeHandle.Volume.Action != "" {
		resVolumeHandle, err := client.RC.VolumeHandler(&reqVolumeHandle)
		if err != nil {
			return model.Volume{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		modelVolume = pbtomodel.PbVolumeToModelVolume(resVolumeHandle.Volume, &resVolumeHandle.HccErrorStack)

	} else {
		return model.Volume{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, "None Action")}, nil

	}

	return *modelVolume, nil
}
