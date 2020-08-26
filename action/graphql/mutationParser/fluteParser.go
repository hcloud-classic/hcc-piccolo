package mutationParser

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/model"
)

func pbNodeToModelNode(node *rpcflute.Node) (*model.Node, error) {
	createdAt, err := ptypes.Timestamp(node.CreatedAt)
	if err != nil {
		return nil, err
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

	return modelNode, nil
}

func pbNodeDetailToModelNodeDetail(nodeDetail *rpcflute.NodeDetail) *model.NodeDetail {
	modelNodeDetail := &model.NodeDetail{
		NodeUUID:      nodeDetail.NodeUUID,
		CPUModel:      nodeDetail.CPUModel,
		CPUProcessors: int(nodeDetail.CPUProcessors),
		CPUThreads:    int(nodeDetail.CPUThreads),
	}

	return modelNodeDetail
}

func OnNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.New("need a UUID argument")
	}

	return client.RC.OnNode(UUID)
}

func OffNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.New("need a UUID argument")
	}

	forceOff, _ := args["force_off"].(bool)

	return client.RC.OffNode(UUID, forceOff)
}

func ForceRestartNode(args map[string]interface{}) (interface{}, error) {
	UUID, UUIDOk := args["uuid"].(string)
	if !UUIDOk {
		return nil, errors.New("need a UUID argument")
	}

	return client.RC.ForceRestartNode(UUID)
}

func CreateNode(args map[string]interface{}) (interface{}, error) {
	bmcMacAddr, _ := args["bmc_mac_addr"].(string)
	bmcIP, _ := args["bmc_ip"].(string)
	pxeMacAddr, _ := args["pxe_mac_addr"].(string)
	status, _ := args["status"].(string)
	cpuCores, _ := args["cpu_cores"].(int)
	memory, _ := args["memory"].(int)
	description, _ := args["description"].(string)
	active, _ := args["active"].(int)

	var node rpcflute.Node
	node.BmcMacAddr = bmcMacAddr
	node.BmcIP = bmcIP
	node.PXEMacAddr = pxeMacAddr
	node.Status = status
	node.CPUCores = int32(cpuCores)
	node.Memory = int32(memory)
	node.Description = description
	node.Active = int32(active)

	resCreateNode, err := client.RC.CreateNode(&rpcflute.ReqCreateNode{
		Node: &node,
	})
	if err != nil {
		return nil, err
	}

	modelNode, err := pbNodeToModelNode(resCreateNode.Node)

	return modelNode, nil
}

func UpdateNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	bmcMacAddr, _ := args["bmc_mac_addr"].(string)
	bmcIP, _ := args["bmc_ip"].(string)
	pxeMacAddr, _ := args["pxe_mac_addr"].(string)
	status, _ := args["status"].(string)
	cpuCores, _ := args["cpu_cores"].(int)
	memory, _ := args["memory"].(int)
	description, _ := args["description"].(string)
	active, _ := args["active"].(int)

	var node rpcflute.Node
	node.UUID = requestedUUID
	node.BmcMacAddr = bmcMacAddr
	node.BmcIP = bmcIP
	node.PXEMacAddr = pxeMacAddr
	node.Status = status
	node.CPUCores = int32(cpuCores)
	node.Memory = int32(memory)
	node.Description = description
	node.Active = int32(active)

	resUpdateNode, err := client.RC.UpdateNode(&rpcflute.ReqUpdateNode{
		Node: &node,
	})
	if err != nil {
		return nil, err
	}

	modelNode, err := pbNodeToModelNode(resUpdateNode.Node)

	return modelNode, nil
}

func DeleteNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var node model.Node
	uuid, err := client.RC.DeleteNode(requestedUUID)
	if err != nil {
		return nil, err
	}
	node.UUID = uuid

	return node, nil
}

func CreateNodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, _ := args["node_uuid"].(string)
	cpuModel, _ := args["cpu_model"].(string)
	cpuProcessors, _ := args["cpu_processors"].(int)
	cpuThreads, _ := args["cpu_threads"].(int)

	var nodeDetail rpcflute.NodeDetail
	nodeDetail.NodeUUID = nodeUUID
	nodeDetail.CPUModel = cpuModel
	nodeDetail.CPUProcessors = int32(cpuProcessors)
	nodeDetail.CPUThreads = int32(cpuThreads)

	resCreateNodeDetail, err := client.RC.CreateNodeDetail(&rpcflute.ReqCreateNodeDetail{
		NodeDetail: &nodeDetail,
	})
	if err != nil {
		return nil, err
	}

	modelNodeDetail := pbNodeDetailToModelNodeDetail(resCreateNodeDetail.NodeDetail)

	return modelNodeDetail, nil
}

func DeleteNodeDetail(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["node_uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a node_uuid argument")
	}

	var nodeDetail model.NodeDetail
	nodeUUID, err := client.RC.DeleteNodeDetail(requestedUUID)
	if err != nil {
		return nil, err
	}
	nodeDetail.NodeUUID = nodeUUID

	return nodeDetail, nil
}
