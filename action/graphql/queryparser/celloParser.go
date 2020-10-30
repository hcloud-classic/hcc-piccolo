package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

// PoolHandler : Handler of zfs pool
func PoolHandler(args map[string]interface{}) (interface{}, error) {

	// UUID, UUIDOk := args["uuid"].(string)
	// Size, SizeOk := args["size"].(string)
	// Free, FreeOk := args["free"].(string)
	// Capacity, CapacityOk := args["capacity"].(string)
	// Health, HealthOk := args["health"].(string)
	// Name, NameOk := args["name"].(string)
	Action, ActionOk := args["action"].(string)

	var reqPoolHandler rpccello.ReqPoolHandler
	var reqPool rpccello.Pool
	reqPoolHandler.Pool = &reqPool

	if ActionOk {
		reqPoolHandler.Pool.Action = Action
	}

	resPoolHandler, err := client.RC.PoolHandler(&reqPoolHandler)
	if err != nil {
		return model.Pool{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelPool := pbtomodel.PbPoolToModelPool(resPoolHandler.Pool, &resPoolHandler.HccErrorStack)
	return *modelPool, nil
}
