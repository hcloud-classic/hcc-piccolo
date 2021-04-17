package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/timpani"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/pb"

	"innogrid.com/hcloud-classic/hcc_errors"
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
	groupID, groupIDOk := args["group_id"].(int)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	nicModel, _ := args["nic_model"].(string)
	nicSpeedMbps, nicSpeedMbpsOk := args["nic_speed_mbps"].(int)
	bmcNICModel, _ := args["bmc_model"].(string)
	bmcNICSpeedMbps, _ := args["bmc_nic_speed_mbps"].(int)
	description, descriptionOk := args["description"].(string)
	chargeCPU, chargeCPUOk := args["charge_cpu"].(int)
	chargeMemory, chargeMemoryOk := args["charge_memory"].(int)
	chargeNIC, chargeNICOk := args["charge_nic"].(int)

	var reqCreateNode pb.ReqCreateNode
	var reqNode pb.Node
	reqCreateNode.Node = &reqNode

	if !groupIDOk || !bmcIPOk || !nicSpeedMbpsOk || !descriptionOk || !chargeCPUOk || !chargeMemoryOk || !chargeNICOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need group_id and bmc_ip, nic_speed_mbps, description, charge_cpu, charge_memory, charge_nic arguments")}, nil
	}

	reqCreateNode.Node.GroupID = int64(groupID)
	reqCreateNode.Node.BmcIP = bmcIP
	reqCreateNode.Node.NicModel = nicModel
	reqCreateNode.Node.NicSpeedMbps = int32(nicSpeedMbps)
	reqCreateNode.Node.BmcNicModel = bmcNICModel
	reqCreateNode.Node.BmcNicSpeedMbps = int32(bmcNICSpeedMbps)
	reqCreateNode.Node.Description = description
	reqCreateNode.Node.ChargeCPU = int32(chargeCPU)
	reqCreateNode.Node.ChargeMemory = int32(chargeMemory)
	reqCreateNode.Node.ChargeNIC = int32(chargeNIC)

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

	groupID, groupIDOk := args["group_id"].(int)
	nodeNum, nodeNumOk := args["node_num"].(int)
	nodeIP, nodeIPOk := args["node_ip"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	nicModel, nicModelOk := args["nic_model"].(string)
	nicSpeedMbps, nicSpeedMbpsOk := args["nic_speed_mbps"].(int)
	bmcNICModel, bmcNICModelOk := args["bmc_nic_model"].(string)
	bmcNICSpeedMbps, bmcNICSpeedMbpsOk := args["bmc_nic_speed_mbps"].(int)
	description, descriptionOk := args["description"].(string)
	rackNumber, rackNumberOk := args["rack_number"].(int)
	chargeCPU, chargeCPUOk := args["charge_cpu"].(int)
	chargeMemory, chargeMemoryOk := args["charge_memory"].(int)
	chargeNIC, chargeNICOk := args["charge_nic"].(int)
	active, activeOk := args["active"].(int)

	var reqUpdateNode pb.ReqUpdateNode
	var reqNode pb.Node
	reqUpdateNode.Node = &reqNode

	reqUpdateNode.Node.UUID = requestedUUID
	if groupIDOk {
		reqUpdateNode.Node.GroupID = int64(groupID)
	}
	if nodeNumOk {
		reqUpdateNode.Node.NodeNum = int32(nodeNum)
	}
	if nodeIPOk {
		reqUpdateNode.Node.NodeIP = nodeIP
	}
	if serverUUIDOk {
		reqUpdateNode.Node.ServerUUID = serverUUID
	}
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
	if nicModelOk {
		reqUpdateNode.Node.NicModel = nicModel
	}
	if nicSpeedMbpsOk {
		reqUpdateNode.Node.NicSpeedMbps = int32(nicSpeedMbps)
	}
	if bmcNICModelOk {
		reqUpdateNode.Node.BmcNicModel = bmcNICModel
	}
	if bmcNICSpeedMbpsOk {
		reqUpdateNode.Node.BmcNicSpeedMbps = int32(bmcNICSpeedMbps)
	}
	if descriptionOk {
		reqUpdateNode.Node.Description = description
	}
	if rackNumberOk {
		reqUpdateNode.Node.RackNumber = int32(rackNumber)
	}
	if chargeCPUOk {
		reqUpdateNode.Node.ChargeCPU = int32(chargeCPU)
	}
	if chargeMemoryOk {
		reqUpdateNode.Node.ChargeMemory = int32(chargeMemory)
	}
	if chargeNICOk {
		reqUpdateNode.Node.ChargeNIC = int32(chargeNIC)
	}
	if activeOk {
		reqUpdateNode.Node.Active = int32(active)
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
	nodeDetailData, nodeDetailDataOk := args["node_detail_data"].(string)

	var reqCreateNodeDetail pb.ReqCreateNodeDetail
	var nodeDetail pb.NodeDetail
	reqCreateNodeDetail.NodeDetail = &nodeDetail

	if nodeUUIDOk {
		reqCreateNodeDetail.NodeDetail.NodeUUID = nodeUUID
	}
	if nodeDetailDataOk {
		reqCreateNodeDetail.NodeDetail.NodeDetailData = nodeDetailData
	}

	resCreateNodeDetail, err := client.RC.CreateNodeDetail(&reqCreateNodeDetail)
	if err != nil {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resCreateNodeDetail.NodeDetail, resCreateNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}

// UpdateNodeDetail : Update detail infos of the node
func UpdateNodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)
	nodeDetailData, nodeDetailDataOk := args["node_detail_data"].(string)

	var reqUpdateNodeDetail pb.ReqUpdateNodeDetail
	var nodeDetail pb.NodeDetail
	reqUpdateNodeDetail.NodeDetail = &nodeDetail

	if nodeUUIDOk {
		reqUpdateNodeDetail.NodeDetail.NodeUUID = nodeUUID
	}
	if nodeDetailDataOk {
		reqUpdateNodeDetail.NodeDetail.NodeDetailData = nodeDetailData
	}

	resUpdateNodeDetail, err := client.RC.UpdateNodeDetail(&reqUpdateNodeDetail)
	if err != nil {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resUpdateNodeDetail.NodeDetail, resUpdateNodeDetail.HccErrorStack)

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
