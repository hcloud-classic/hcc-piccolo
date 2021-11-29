package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// PoolHandler : Handler of zfs pool
func PoolHandler(args map[string]interface{}) (interface{}, error) {
	Action, ActionOk := args["action"].(string)
	GroupID, GroupIDOk := args["group_id"].(int)

	var reqPoolHandler pb.ReqPoolHandler
	var reqPool pb.Pool
	var reqGroup pb.Group
	reqPoolHandler.Pool = &reqPool
	reqPoolHandler.Group = &reqGroup

	if ActionOk {
		reqPoolHandler.Pool.Action = Action
	} else {
		reqPoolHandler.Pool.Action = "read"
	}
	if GroupIDOk {
		reqPoolHandler.Group.Id = (int64)(GroupID)
	} else {
		reqPoolHandler.Group.Id = 1
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
	Action, ActionOk := args["action"].(string)
	GroupID, GroupIDOk := args["group_id"].(int)

	var modelPoolList []model.Pool
	var reqGetPoolList pb.ReqGetPoolList
	var reqPool pb.Pool
	var reqGroup pb.Group
	reqGetPoolList.Pool = &reqPool
	reqGetPoolList.Group = &reqGroup

	if ActionOk {
		reqGetPoolList.Pool.Action = Action
	} else {
		reqGetPoolList.Pool.Action = "read"
	}
	if GroupIDOk {
		reqGetPoolList.Group.Id = (int64)(GroupID)
	} else {
		reqGetPoolList.Group.Id = 1
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
