package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"strconv"
)

// CreateVolume : oboe to cello
func CreateVolume(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	fileSystem, fileSystemOK := args["filesystem"].(string)
	diskSize, diskSizeOk := args["size"].(int)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	useType, useTypeOk := args["use_type"].(string)
	networkIP, networkIPOk := args["network_ip"].(string)
	gatewayIP, gatewayIPOk := args["gateway_ip"].(string)
	lunNum, lunNumOk := args["lun_num"].(string)
	pool, poolOk := args["pool"].(string)

	var reqCreateVolume rpccello.ReqVolumeHandler
	var reqVolume rpccello.Volume
	reqCreateVolume.Volume = &reqVolume

	if serverUUIDOk {
		reqCreateVolume.Volume.ServerUUID = serverUUID
	}

	if fileSystemOK {
		reqCreateVolume.Volume.Filesystem = fileSystem
	}
	strSize := strconv.Itoa(diskSize)
	if diskSizeOk {
		reqCreateVolume.Volume.Size = strSize
	}

	if userUUIDOk {
		reqCreateVolume.Volume.UserUUID = userUUID
	}
	if useTypeOk {
		reqCreateVolume.Volume.UseType = useType
	}
	if networkIPOk {
		reqCreateVolume.Volume.Network_IP = networkIP
	}
	if gatewayIPOk {
		reqCreateVolume.Volume.GatewayIp = gatewayIP
	}
	if lunNumOk {
		convInt, _ := strconv.Atoi(lunNum)
		reqCreateVolume.Volume.Lun = int64(convInt)
	}
	if poolOk {
		reqCreateVolume.Volume.Pool = pool
	}
	reqCreateVolume.Volume.Action = "create"
	resCreateVolume, err := client.RC.CreateVolume(&reqCreateVolume)
	if err != nil {
		return model.Volume{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelVolume := pbtomodel.PbVolumeToModelVolume(resCreateVolume.Volume, &resCreateVolume.HccErrorStack)
	return *modelVolume, nil
}
