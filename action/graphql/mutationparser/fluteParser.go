package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/graphql/queryparser"
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
func CreateNode(args map[string]interface{}, isMaster bool) (interface{}, error) {
	if !isMaster {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	nodeName, nodeNameOk := args["node_name"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	nicSpeedMbps, nicSpeedMbpsOk := args["nic_speed_mbps"].(int)
	description, descriptionOk := args["description"].(string)

	ipmiUserID, ipmiUserIDOk := args["ipmi_user_id"].(string)
	ipmiUserPassword, ipmiUserPasswordOk := args["ipmi_user_password"].(string)

	nicDetailData, nicDetailDataOk := args["nic_detail_data"].(string)

	var reqCreateNode pb.ReqCreateNode
	var reqNode pb.Node
	reqCreateNode.Node = &reqNode

	if !nodeNameOk || !bmcIPOk || !nicSpeedMbpsOk || !descriptionOk ||
		!ipmiUserIDOk || !ipmiUserPasswordOk ||
		!nicDetailDataOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need node_name and bmc_ip, nic_speed_mbps, description, nic_detail_data arguments")}, nil
	}

	reqCreateNode.Node.NodeName = nodeName
	reqCreateNode.Node.BmcIP = bmcIP
	reqCreateNode.Node.NicSpeedMbps = int32(nicSpeedMbps)
	reqCreateNode.Node.Description = description

	reqCreateNode.Node.IpmiUserID = ipmiUserID
	reqCreateNode.Node.IpmiUserPassword = ipmiUserPassword

	reqCreateNode.NicDetailData = nicDetailData

	resCreateNode, err := client.RC.CreateNode(&reqCreateNode)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resCreateNode.Node, resCreateNode.HccErrorStack)

	return *modelNode, nil
}

// UpdateNode : Update the infos of the node
func UpdateNode(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	if !isMaster {
		groupID, _ := args["group_id"].(int)
		node, err := queryparser.Node(args)
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		if int(node.(model.Node).GroupID) != groupID {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group's node if you are not a master")}, nil
		}
	}

	nodeName, nodeNameOk := args["node_name"].(string)
	nodeNum, nodeNumOk := args["node_num"].(int)
	nodeIP, nodeIPOk := args["node_ip"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	nicSpeedMbps, nicSpeedMbpsOk := args["nic_speed_mbps"].(int)
	description, descriptionOk := args["description"].(string)
	rackNumber, rackNumberOk := args["rack_number"].(int)
	active, activeOk := args["active"].(int)

	ipmiUserID, ipmiUserIDOk := args["ipmi_user_id"].(string)
	ipmiUserPassword, ipmiUserPasswordOk := args["ipmi_user_password"].(string)

	var reqUpdateNode pb.ReqUpdateNode
	var reqNode pb.Node
	reqUpdateNode.Node = &reqNode

	reqUpdateNode.Node.UUID = requestedUUID
	if nodeNameOk {
		reqUpdateNode.Node.NodeName = nodeName
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
	if nicSpeedMbpsOk {
		reqUpdateNode.Node.NicSpeedMbps = int32(nicSpeedMbps)
	}
	if descriptionOk {
		reqUpdateNode.Node.Description = description
	}
	if rackNumberOk {
		reqUpdateNode.Node.RackNumber = int32(rackNumber)
	}
	if activeOk {
		reqUpdateNode.Node.Active = int32(active)
	}

	if ipmiUserIDOk {
		reqUpdateNode.Node.IpmiUserID = ipmiUserID
	}
	if ipmiUserPasswordOk {
		reqUpdateNode.Node.IpmiUserPassword = ipmiUserPassword
	}

	resUpdateNode, err := client.RC.UpdateNode(&reqUpdateNode)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resUpdateNode.Node, resUpdateNode.HccErrorStack)

	return *modelNode, nil
}

// DeleteNode : Delete the node
func DeleteNode(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	if !isMaster {
		groupID, _ := args["group_id"].(int)
		node, err := queryparser.Node(args)
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		if int(node.(model.Node).GroupID) != groupID {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't delete the other group's node if you are not a master")}, nil
		}
	}

	resDeleteNode, err := client.RC.DeleteNode(requestedUUID)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelNode := pbtomodel.PbNodeToModelNode(resDeleteNode.Node, resDeleteNode.HccErrorStack)

	return *modelNode, nil
}

// CreateNodeDetail : Create detail infos of the node
func CreateNodeDetail(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

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
func UpdateNodeDetail(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

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
func DeleteNodeDetail(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

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
