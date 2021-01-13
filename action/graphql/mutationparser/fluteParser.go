package mutationparser

import (
	"github.com/hcloud-classic/pb"
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/timpani"
	"hcc/piccolo/model"

	"github.com/hcloud-classic/hcc_errors"
)

// OnNode : Turn on the node
func OnNode(args map[string]interface{}) (interface{}, error) {
	UUIDs, UUIDsOk := args["uuids"].([]interface{})
	if !UUIDsOk {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuids argument")}, nil
	}

	var strUUIDs []string
	for _, UUID := range UUIDs {
		strUUIDs = append(strUUIDs, UUID.(string))
	}

	resNodePowerControl, err := client.RC.OnNode(strUUIDs)
	if err != nil {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resNodePowerControl.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Results: resNodePowerControl.Result, Errors: Errors}, nil
}

// OffNode : Turn off the node
func OffNode(args map[string]interface{}) (interface{}, error) {
	UUIDs, UUIDsOk := args["uuids"].([]interface{})
	if !UUIDsOk {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var forceOff bool
	forceOff, forceOffOk := args["force_off"].(bool)
	if !forceOffOk {
		forceOff = false
	}

	var strUUIDs []string
	for _, UUID := range UUIDs {
		strUUIDs = append(strUUIDs, UUID.(string))
	}

	resNodePowerControl, err := client.RC.OffNode(strUUIDs, forceOff)
	if err != nil {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resNodePowerControl.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Results: resNodePowerControl.Result, Errors: Errors}, nil
}

// ForceRestartNode : Force restart the node
func ForceRestartNode(args map[string]interface{}) (interface{}, error) {
	UUIDs, UUIDsOk := args["uuids"].([]interface{})
	if !UUIDsOk {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var strUUIDs []string
	for _, UUID := range UUIDs {
		_, err := timpani.NormalRebootNotification(UUID.(string))
		if err != nil {
			return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		strUUIDs = append(strUUIDs, UUID.(string))
	}

	resNodePowerControl, err := client.RC.ForceRestartNode(strUUIDs)
	if err != nil {
		return model.PowerControlNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resNodePowerControl.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerControlNode{Results: resNodePowerControl.Result, Errors: Errors}, nil
}

// CreateNode : Create a node
func CreateNode(args map[string]interface{}) (interface{}, error) {
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	description, descriptionOk := args["description"].(string)

	var reqCreateNode pb.ReqCreateNode
	var reqNode pb.Node
	reqCreateNode.Node = &reqNode

	if !bmcIPOk || !descriptionOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need bmc_ip and description arguments")}, nil
	}

	reqCreateNode.Node.BmcIP = bmcIP
	reqCreateNode.Node.Description = description

	resCreateNode, err := client.RC.CreateNode(&reqCreateNode)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resCreateNode.Node, resCreateNode.HccErrorStack)

	return *modelNode, nil
}

// UpdateNode : Update the infos of the node
func UpdateNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
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

	var reqUpdateNode pb.ReqUpdateNode
	var reqNode pb.Node
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
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resUpdateNode.Node, resUpdateNode.HccErrorStack)

	return *modelNode, nil
}

// DeleteNode : Delete the node
func DeleteNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resDeleteNode, err := client.RC.DeleteNode(requestedUUID)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resDeleteNode.Node, resDeleteNode.HccErrorStack)

	return *modelNode, nil
}

// CreateNodeDetail : Create detail infos of the node
func CreateNodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)
	cpuModel, cpuModelOk := args["cpu_model"].(string)
	cpuProcessors, cpuProcessorsOk := args["cpu_processors"].(int)
	cpuThreads, cpuThreadsOk := args["cpu_threads"].(int)

	var reqCreateNodeDetail pb.ReqCreateNodeDetail
	var nodeDetail pb.NodeDetail
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
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resCreateNodeDetail.NodeDetail, resCreateNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}

// DeleteNodeDetail : Delete the node detail of the node
func DeleteNodeDetail(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["node_uuid"].(string)
	if !requestedUUIDOk {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a node_uuid argument")}, nil
	}

	resDeleteNodeDetail, err := client.RC.DeleteNodeDetail(requestedUUID)
	if err != nil {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resDeleteNodeDetail.NodeDetail, resDeleteNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}
