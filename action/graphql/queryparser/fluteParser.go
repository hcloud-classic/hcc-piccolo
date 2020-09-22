package queryparser

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

func pbNodeToModelNode(node *rpcflute.Node, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Node {
	var createdAt time.Time
	if node.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(node.CreatedAt)
		if err != nil {
			return &model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelNode := &model.Node{
		UUID:        node.UUID,
		ServerUUID:  node.ServerUUID,
		BmcMacAddr:  node.BmcMacAddr,
		BmcIP:       node.BmcIP,
		PXEMacAddr:  node.PXEMacAddr,
		Status:      node.Status,
		CPUCores:    int(node.CPUCores),
		Memory:      int(node.Memory),
		Description: node.Description,
		CreatedAt:   createdAt,
		Active:      int(node.Active),
		ForceOff:    node.ForceOff,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNode.Errors = *hccErrStack.ConvertReportForm()
	}

	return modelNode
}

func pbNodeDetailToModelNodeDetail(nodeDetail *rpcflute.NodeDetail, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.NodeDetail {
	modelNodeDetail := &model.NodeDetail{
		NodeUUID:      nodeDetail.NodeUUID,
		CPUModel:      nodeDetail.CPUModel,
		CPUProcessors: int(nodeDetail.CPUProcessors),
		CPUThreads:    int(nodeDetail.CPUThreads),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelNodeDetail.Errors = *hccErrStack.ConvertReportForm()
	}

	return modelNodeDetail
}

// PowerStateNode : Get power state of the node
func PowerStateNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.PowerStateNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resNodePowerState, err := client.RC.GetNodePowerState(uuid)
	if err != nil {
		return model.PowerStateNode{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resNodePowerState.HccErrorStack)

	return model.PowerStateNode{Result: resNodePowerState.Result, Errors: *hccErrStack.ConvertReportForm()}, nil
}

// Node : Get infos of the node
func Node(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetNode, err := client.RC.GetNode(uuid)
	if err != nil {
		return model.Node{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	modelNode := pbNodeToModelNode(resGetNode.Node, &resGetNode.HccErrorStack)

	return *modelNode, nil
}

// ListNode : Get node list with provided options
func ListNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(int)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListNode rpcflute.ReqGetNodeList
	var reqNode rpcflute.Node
	reqListNode.Node = &reqNode

	if uuidOk {
		reqListNode.Node.UUID = uuid
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
	if descriptionOk {
		reqListNode.Node.Description = description
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
		return model.NodeList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var nodeList []model.Node
	for _, pNode := range resGetNodeList.Node {
		modelNode := pbNodeToModelNode(pNode, nil)
		nodeList = append(nodeList, *modelNode)
	}

	hccErrStack := errconv.GrpcStackToHcc(&resGetNodeList.HccErrorStack)

	return model.NodeList{Nodes: nodeList, Errors: *hccErrStack.ConvertReportForm()}, nil
}

// AllNode : Get node list with provided options (Just call ListNode())
func AllNode(args map[string]interface{}) (interface{}, error) {
	return ListNode(args)
}

// NumNode : Get number of nodes
func NumNode() (interface{}, error) {
	resGetNodeNum, err := client.RC.GetNodeNum()
	if err != nil {
		return model.NodeNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelNodeNum model.NodeNum
	modelNodeNum.Number = int(resGetNodeNum.Num)

	hccErrStack := errconv.GrpcStackToHcc(&resGetNodeNum.HccErrorStack)
	modelNodeNum.Errors = *hccErrStack.ConvertReportForm()

	return modelNodeNum, nil
}

// NodeDetail : Get infos of the detail of the node
func NodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	if !nodeUUIDOk {
		return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a node_uuid argument")}, nil
	}

	resGetNodeDetail, err := client.RC.GetNodeDetail(nodeUUID)
	if err != nil {
		return model.NodeDetail{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	modelNodeDetail := pbNodeDetailToModelNodeDetail(resGetNodeDetail.NodeDetail, &resGetNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}
