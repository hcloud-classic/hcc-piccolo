package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

// OnNode : Turn on the node
func OnNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resNodePowerControl, err := client.RC.OnNode(UUID)
	if err != nil {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resNodePowerControl.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Result: resNodePowerControl.Result[0], Errors: Errors}, nil
}

// OffNode : Turn off the node
func OffNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var forceOff bool
	forceOff, forceOffOk := args["force_off"].(bool)
	if !forceOffOk {
		forceOff = false
	}

	resNodePowerControl, err := client.RC.OffNode(UUID, forceOff)
	if err != nil {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resNodePowerControl.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Result: resNodePowerControl.Result[0], Errors: Errors}, nil
}

// ForceRestartNode : Force restart the node
func ForceRestartNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resNodePowerControl, err := client.RC.ForceRestartNode(UUID)
	if err != nil {
		return model.PowerControlNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resNodePowerControl.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Result: resNodePowerControl.Result[0], Errors: Errors}, nil
}

// CreateNode : Create a node
func CreateNode(args map[string]interface{}) (interface{}, error) {
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	rackNumber, rackNumberOk := args["rack_number"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(int)

	var reqCreateNode rpcflute.ReqCreateNode
	var reqNode rpcflute.Node
	reqCreateNode.Node = &reqNode

	if bmcMacAddrOk {
		reqCreateNode.Node.BmcMacAddr = bmcMacAddr
	}
	if bmcIPOk {
		reqCreateNode.Node.BmcIP = bmcIP
	}
	if pxeMacAddrOk {
		reqCreateNode.Node.PXEMacAddr = pxeMacAddr
	}
	if statusOk {
		reqCreateNode.Node.Status = status
	}
	if cpuCoresOk {
		reqCreateNode.Node.CPUCores = int32(cpuCores)
	}
	if memoryOk {
		reqCreateNode.Node.Memory = int32(memory)
	}
	if rackNumberOk {
		reqCreateNode.Node.RackNumber = int32(rackNumber)
	}
	if descriptionOk {
		reqCreateNode.Node.Description = description
	}
	if activeOk {
		reqCreateNode.Node.Active = int32(active)
	}

	resCreateNode, err := client.RC.CreateNode(&reqCreateNode)
	if err != nil {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resCreateNode.Node, &resCreateNode.HccErrorStack)

	return *modelNode, nil
}

// UpdateNode : Update the infos of the node
func UpdateNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	rackNumber, rackNumberOk := args["rack_number"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(int)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)

	var reqUpdateNode rpcflute.ReqUpdateNode
	var reqNode rpcflute.Node
	reqUpdateNode.Node = &reqNode

	reqUpdateNode.Node.UUID = requestedUUID
	if bmcMacAddrOk {
		reqUpdateNode.Node.BmcMacAddr = bmcMacAddr
	}
	if bmcIPOk {
		reqUpdateNode.Node.BmcIP = bmcIP
	}
	if pxeMacAddrOk {
		reqUpdateNode.Node.PXEMacAddr = pxeMacAddr
	}
	if statusOk {
		reqUpdateNode.Node.Status = status
	}
	if cpuCoresOk {
		reqUpdateNode.Node.CPUCores = int32(cpuCores)
	}
	if memoryOk {
		reqUpdateNode.Node.Memory = int32(memory)
	}
	if rackNumberOk {
		reqUpdateNode.Node.RackNumber = int32(rackNumber)
	}
	if descriptionOk {
		reqUpdateNode.Node.Description = description
	}
	if activeOk {
		reqUpdateNode.Node.Active = int32(active)
	}
	if serverUUIDOk {
		reqUpdateNode.Node.ServerUUID = serverUUID
	}

	resUpdateNode, err := client.RC.UpdateNode(&reqUpdateNode)
	if err != nil {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resUpdateNode.Node, &resUpdateNode.HccErrorStack)

	return *modelNode, nil
}

// DeleteNode : Delete the node
func DeleteNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var node model.Node
	resDeleteNode, err := client.RC.DeleteNode(requestedUUID)
	if err != nil {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	node.UUID = resDeleteNode.Node.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteNode.HccErrorStack)
	node.Errors = *hccErrStack.ConvertReportForm()
	if len(node.Errors) != 0 && node.Errors[0].ErrCode == 0 {
		node.Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return node, nil
}

// CreateNodeDetail : Create detail infos of the node
func CreateNodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)
	cpuModel, cpuModelOk := args["cpu_model"].(string)
	cpuProcessors, cpuProcessorsOk := args["cpu_processors"].(int)
	cpuThreads, cpuThreadsOk := args["cpu_threads"].(int)

	var reqCreateNodeDetail rpcflute.ReqCreateNodeDetail
	var nodeDetail rpcflute.NodeDetail
	reqCreateNodeDetail.NodeDetail = &nodeDetail

	if nodeUUIDOk {
		reqCreateNodeDetail.NodeDetail.NodeUUID = nodeUUID
	}
	if cpuModelOk {
		reqCreateNodeDetail.NodeDetail.CPUModel = cpuModel
	}
	if cpuProcessorsOk {
		reqCreateNodeDetail.NodeDetail.CPUProcessors = int32(cpuProcessors)
	}
	if cpuThreadsOk {
		reqCreateNodeDetail.NodeDetail.CPUThreads = int32(cpuThreads)
	}

	resCreateNodeDetail, err := client.RC.CreateNodeDetail(&reqCreateNodeDetail)
	if err != nil {
		return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resCreateNodeDetail.NodeDetail, &resCreateNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}

// DeleteNodeDetail : Delete the node detail of the node
func DeleteNodeDetail(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["node_uuid"].(string)
	if !requestedUUIDOk {
		return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a node_uuid argument")}, nil
	}

	var nodeDetail model.NodeDetail
	resDeleteNodeDetail, err := client.RC.DeleteNodeDetail(requestedUUID)
	if err != nil {
		return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	nodeDetail.NodeUUID = resDeleteNodeDetail.NodeDetail.NodeUUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteNodeDetail.HccErrorStack)
	nodeDetail.Errors = *hccErrStack.ConvertReportForm()
	if len(nodeDetail.Errors) != 0 && nodeDetail.Errors[0].ErrCode == 0 {
		nodeDetail.Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return nodeDetail, nil
}
