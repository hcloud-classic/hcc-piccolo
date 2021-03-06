package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// PowerStateNode : Get power state of the node
func PowerStateNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.PowerStateNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resNodePowerState, err := client.RC.GetNodePowerState(uuid)
	if err != nil {
		return model.PowerStateNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resNodePowerState.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.PowerStateNode{Result: resNodePowerState.Result, Errors: Errors}, nil
}

// Node : Get infos of the node
func Node(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetNode, err := client.RC.GetNode(uuid)
	if err != nil {
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}
	modelNode := pbtomodel.PbNodeToModelNode(resGetNode.Node, resGetNode.HccErrorStack)

	// group_name
	if modelNode.GroupID != -1 {
		group, err := dao.ReadGroup(int(modelNode.GroupID))
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
		}
		modelNode.GroupName = group.Name
	} else {
		modelNode.GroupName = "Not allocated"
	}

	return *modelNode, nil
}

// ListNode : Get node list with provided options
func ListNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
	nodeName, nodeNameOk := args["node_name"].(string)
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
	nicSpeedMbps, nicSpeedMbpsOk := args["nic_speed_mbps"].(int)
	description, descriptionOk := args["description"].(string)
	rackNumber, rackNumberOk := args["rack_number"].(int)
	active, activeOk := args["active"].(int)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListNode pb.ReqGetNodeList
	var reqNode pb.Node
	reqListNode.Node = &reqNode

	if uuidOk {
		reqListNode.Node.UUID = uuid
	}
	if nodeNameOk {
		reqListNode.Node.NodeName = nodeName
	}
	if groupIDOk {
		reqListNode.Node.GroupID = int64(groupID)
	}
	if nodeNumOk {
		reqListNode.Node.NodeNum = int32(nodeNum)
	}
	if nodeIPOk {
		reqListNode.Node.NodeIP = nodeIP
	}
	if serverUUIDOk {
		reqListNode.Node.ServerUUID = serverUUID
	}
	if bmcMacAddrOk {
		reqListNode.Node.BmcMacAddr = bmcMacAddr
	}
	if bmcIPOk {
		reqListNode.Node.BmcIP = bmcIP
	}
	if pxeMacAddrOk {
		reqListNode.Node.PXEMacAddr = pxeMacAddr
	}
	if statusOk {
		reqListNode.Node.Status = status
	}
	if cpuCoresOk {
		reqListNode.Node.CPUCores = int32(cpuCores)
	}
	if memoryOk {
		reqListNode.Node.Memory = int32(memory)
	}
	if nicSpeedMbpsOk {
		reqListNode.Node.NicSpeedMbps = int32(nicSpeedMbps)
	}
	if descriptionOk {
		reqListNode.Node.Description = description
	}
	if rackNumberOk {
		reqListNode.Node.RackNumber = int32(rackNumber)
	}
	if activeOk {
		reqListNode.Node.Active = int32(active)
	}
	if rowOk {
		reqListNode.Row = int64(row)
	}
	if pageOk {
		reqListNode.Page = int64(page)
	}

	resGetNodeList, err := client.RC.GetNodeList(&reqListNode)
	if err != nil {
		return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var numNode int
	if rowOk && pageOk {
		reqListNode.Row = 0
		reqListNode.Page = 0
		resGetNodeList2, err := client.RC.GetNodeList(&reqListNode)
		if err != nil {
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		numNode = len(resGetNodeList2.Node)
	} else {
		numNode = len(resGetNodeList.Node)
	}

	var nodeList []model.Node
	for _, pNode := range resGetNodeList.Node {
		modelNode := pbtomodel.PbNodeToModelNode(pNode, nil)

		// group_name
		if modelNode.GroupID != -1 {
			group, err := dao.ReadGroup(int(modelNode.GroupID))
			if err != nil {
				return model.NodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
			}
			modelNode.GroupName = group.Name
		} else {
			modelNode.GroupName = "Not allocated"
		}

		nodeList = append(nodeList, *modelNode)
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetNodeList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.NodeList{Nodes: nodeList, TotalNum: numNode, Errors: Errors}, nil
}

// AllNode : Get node list with provided options (Just call ListNode())
func AllNode(args map[string]interface{}) (interface{}, error) {
	return ListNode(args)
}

// NumNode : Get number of nodes
func NumNode(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetNodeNum pb.ReqGetNodeNum
	if groupIDOk {
		reqGetNodeNum.GroupID = int64(groupID)
	}
	resGetNodeNum, err := client.RC.GetNodeNum(&reqGetNodeNum)
	if err != nil {
		return model.NodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelNodeNum model.NodeNum
	modelNodeNum.Number = int(resGetNodeNum.Num)

	hccErrStack := errconv.GrpcStackToHcc(resGetNodeNum.HccErrorStack)
	modelNodeNum.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(modelNodeNum.Errors) != 0 && modelNodeNum.Errors[0].ErrCode == 0 {
		modelNodeNum.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelNodeNum, nil
}

// NodeDetail : Get infos of the detail of the node
func NodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	if !nodeUUIDOk {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a node_uuid argument")}, nil
	}

	resGetNodeDetail, err := client.RC.GetNodeDetail(nodeUUID)
	if err != nil {
		return model.NodeDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	modelNodeDetail := pbtomodel.PbNodeDetailToModelNodeDetail(resGetNodeDetail.NodeDetail, resGetNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}
