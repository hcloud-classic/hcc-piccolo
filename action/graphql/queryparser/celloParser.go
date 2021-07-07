package queryparser

import (
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
