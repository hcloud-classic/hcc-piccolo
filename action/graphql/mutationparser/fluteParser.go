package mutationparser

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
			return &model.Node{Errors: errors.ReturnHccError(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
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
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack).ConvertReportForm()
		modelNode.Errors = *hccErrStack
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
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack).ConvertReportForm()
		modelNodeDetail.Errors = *hccErrStack
	}

	return modelNodeDetail
}

// OnNode : Turn on the node
func OnNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	resNodePowerControl, err := client.RC.OnNode(UUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	return resNodePowerControl.Result[0], nil
}

// OffNode : Turn off the node
func OffNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	var forceOff bool
	forceOff, forceOffOk := args["force_off"].(bool)
	if !forceOffOk {
		forceOff = false
	}

	resNodePowerControl, err := client.RC.OffNode(UUID, forceOff)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	return resNodePowerControl.Result[0], nil
}

// ForceRestartNode : Force restart the node
func ForceRestartNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	resNodePowerControl, err := client.RC.ForceRestartNode(UUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	return resNodePowerControl.Result[0], nil
}

// CreateNode : Create a node
func CreateNode(args map[string]interface{}) (interface{}, error) {
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
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
	if descriptionOk {
		reqCreateNode.Node.Description = description
	}
	if activeOk {
		reqCreateNode.Node.Active = int32(active)
	}

	resCreateNode, err := client.RC.CreateNode(&reqCreateNode)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelNode := pbNodeToModelNode(resCreateNode.Node, &resCreateNode.HccErrorStack)

	return *modelNode, nil
}

// UpdateNode : Update the infos of the node
func UpdateNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(int)

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
	if descriptionOk {
		reqUpdateNode.Node.Description = description
	}
	if activeOk {
		reqUpdateNode.Node.Active = int32(active)
	}

	resUpdateNode, err := client.RC.UpdateNode(&reqUpdateNode)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelNode := pbNodeToModelNode(resUpdateNode.Node, &resUpdateNode.HccErrorStack)

	return *modelNode, nil
}

// DeleteNode : Delete the node
func DeleteNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a uuid argument").New()
	}

	var node model.Node
	resDeleteNode, err := client.RC.DeleteNode(requestedUUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}
	node.UUID = resDeleteNode.UUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteNode.HccErrorStack).ConvertReportForm()
	node.Errors = *hccErrStack

	return node, nil
}

// CreateNodeDetail : Create detail infos of the node
func CreateNodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)
	cpuModel, cpuModelOk := args["cpu_model"].(string)
	cpuProcessors, cpuProcessorsOk := args["cpu_processors"].(int)
	cpuThreads, cpuThreadsOk := args["cpu_threads"].(int)

	var reqCreateNodeDetail rpcflute.ReqCreateNodeDetail
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
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}

	modelNodeDetail := pbNodeDetailToModelNodeDetail(resCreateNodeDetail.NodeDetail, &resCreateNodeDetail.HccErrorStack)

	return *modelNodeDetail, nil
}

// DeleteNodeDetail : Delete the node detail of the node
func DeleteNodeDetail(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["node_uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a node_uuid argument").New()
	}

	var nodeDetail model.NodeDetail
	resDeleteNodeDetail, err := client.RC.DeleteNodeDetail(requestedUUID)
	if err != nil {
		return nil, errors.NewHccError(errors.PiccoloGrpcRequestError, err.Error()).New()
	}
	nodeDetail.NodeUUID = resDeleteNodeDetail.NodeUUID

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteNodeDetail.HccErrorStack).ConvertReportForm()
	nodeDetail.Errors = *hccErrStack

	return nodeDetail, nil
}
