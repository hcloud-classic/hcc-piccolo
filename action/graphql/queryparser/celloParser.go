package queryparser

import (
	"fmt"
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// TODO: Need to handle group_id - ish

// PoolHandler : Handler of zfs pool
func PoolHandler(args map[string]interface{}) (interface{}, error) {

	// UUID, UUIDOk := args["uuid"].(string)
	// Size, SizeOk := args["size"].(string)
	// Free, FreeOk := args["free"].(string)
	// Capacity, CapacityOk := args["capacity"].(string)
	// Health, HealthOk := args["health"].(string)
	// Name, NameOk := args["name"].(string)
	Action, ActionOk := args["action"].(string)

	var reqPoolHandler pb.ReqPoolHandler
	var reqPool pb.Pool
	reqPoolHandler.Pool = &reqPool

	if ActionOk {
		reqPoolHandler.Pool.Action = Action
	}

	resPoolHandler, err := client.RC.PoolHandler(&reqPoolHandler)
	if err != nil {
		return model.Pool{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelPool := pbtomodel.PbPoolToModelPool(resPoolHandler.Pool, resPoolHandler.HccErrorStack)
	return *modelPool, nil
}

// GetPoolList : pool list
func GetPoolList(args map[string]interface{}) (interface{}, error) {
	// TODO: Need to handle group_id - ish

	// UUID, UUIDOk := args["uuid"].(string)
	// Size, SizeOk := args["size"].(string)
	// Free, FreeOk := args["free"].(string)
	// Capacity, CapacityOk := args["capacity"].(string)
	// Health, HealthOk := args["health"].(string)
	// Name, NameOk := args["name"].(string)
	Action, ActionOk := args["action"].(string)

	var modelPoolList []model.Pool
	var reqGetPoolList pb.ReqGetPoolList
	var reqPool pb.Pool
	reqGetPoolList.Pool = &reqPool

	if ActionOk {
		reqGetPoolList.Pool.Action = Action
	} else {
		reqGetPoolList.Pool.Action = "read"
	}

	resPoolList, err := client.RC.GetPoolList(&reqGetPoolList)
	if err != nil {
		return model.PoolList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	for _, args := range resPoolList.Pool {
		tempPool := pbtomodel.PbPoolToModelPool(args, resPoolList.HccErrorStack)
		modelPoolList = append(modelPoolList, *tempPool)
	}

	hccErrStack := errconv.GrpcStackToHcc(resPoolList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)

	return model.PoolList{Pools: modelPoolList, Errors: Errors}, nil
}

func GetVolumeList(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	// Size, SizeOk := args["size"].(string)
	// Free, FreeOk := args["free"].(string)
	// Capacity, CapacityOk := args["capacity"].(string)
	// Health, HealthOk := args["health"].(string)
	// Name, NameOk := args["name"].(string)
	action, actionOk := args["action"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqVolumeListHandler pb.ReqGetVolumeList
	var reqVolumeList pb.Volume
	var modelVolumeList []model.Volume
	reqVolumeListHandler.Volume = &reqVolumeList

	if actionOk {
		reqVolumeListHandler.Volume.Action = action
	} else {
		reqVolumeListHandler.Volume.Action = "read_list"
	}
	if rowOk {
		reqVolumeListHandler.Row = int64(row)
	} else {
		reqVolumeListHandler.Row = int64(10)
	}
	if pageOk {
		reqVolumeListHandler.Page = int64(page)
	} else {
		reqVolumeListHandler.Page = int64(1)
	}

	if serverUUIDOk {
		reqVolumeListHandler.Volume.ServerUUID = serverUUID
	}
	if userUUIDOk {
		reqVolumeListHandler.Volume.UserUUID = userUUID
	}

	resGetVolumeList, err := client.RC.GetVolumeList(&reqVolumeListHandler)
	if err != nil {
		return model.VolumeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var numVolume int
	if rowOk && pageOk {
		reqVolumeListHandler.Row = 0
		reqVolumeListHandler.Page = 0
		resGetVolumeList2, err := client.RC.GetVolumeList(&reqVolumeListHandler)
		if err != nil {
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		numVolume = len(resGetVolumeList2.Volume)
	} else {
		numVolume = len(resGetVolumeList.Volume)
	}

	fmt.Println(resGetVolumeList.Volume)
	for _, args := range resGetVolumeList.Volume {
		tempVol := pbtomodel.PbVolumeToModelVolume(args, resGetVolumeList.HccErrorStack)
		modelVolumeList = append(modelVolumeList, *tempVol)
	}
	fmt.Println("modelVolumeList", modelVolumeList)

	hccErrStack := errconv.GrpcStackToHcc(resGetVolumeList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.VolumeList{Volumes: modelVolumeList, TotalNum: numVolume, Errors: Errors}, nil
}
