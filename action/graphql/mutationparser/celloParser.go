package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/model"
	"strconv"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// VolumeHandle : oboe to cello
func VolumeHandle(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

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
	var reqVolumeHandle pb.ReqVolumeHandler
	var reqVolume pb.Volume
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
		reqVolumeHandle.Volume.Gateway_IP = gatewayIP
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
			err2 := dao.WriteServerAction(
				serverUUID,
				"cello / volume_handle (action: "+action+")",
				"Failed",
				err.Error(),
				tokenString)
			if err2 != nil {
				logger.Logger.Println("WriteServerAction(): " + err2.Error())
			}

			return model.Volume{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		modelVolume = pbtomodel.PbVolumeToModelVolume(resVolumeHandle.Volume, resVolumeHandle.HccErrorStack)

	} else {
		err2 := dao.WriteServerAction(
			serverUUID,
			"cello / volume_handle (action: "+action+")",
			"Failed",
			"None Action",
			tokenString)
		if err2 != nil {
			logger.Logger.Println("WriteServerAction(): " + err2.Error())
		}

		return model.Volume{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "None Action")}, nil
	}

	success := true
	errStr := ""

	if len(modelVolume.Errors) != 0 {
		success = false
	}

	var result string
	if success {
		result = "Success"
	} else {
		result = "Failed"
	}

	err := dao.WriteServerAction(
		serverUUID,
		"cello / volume_handle (action: "+action+")",
		result,
		errStr,
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	return *modelVolume, nil
}
