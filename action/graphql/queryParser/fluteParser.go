package queryParser

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

func PowerStateNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	return client.RC.GetNodePowerState(uuid)
}

func Node(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	return client.RC.GetSubnet(uuid)
}

func ListNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, _ := args["server_uuid"].(string)
	bmcMacAddr, _ := args["bmc_mac_addr"].(string)
	bmcIP, _ := args["bmc_ip"].(string)
	pxeMacAdr, _ := args["pxe_mac_addr"].(string)
	status, _ := args["status"].(string)
	cpuCores, _ := args["cpu_cores"].(int)
	memory, _ := args["memory"].(int)
	description, _ := args["description"].(string)
	active, _ := args["active"].(int)
	row, _ := args["row"].(int)
	page, _ := args["page"].(int)

	var reqListNode rpcflute.ReqGetNodeList
	reqListNode.Node.ServerUUID = serverUUID
	reqListNode.Node.BmcMacAddr = bmcMacAddr
	reqListNode.Node.BmcIP = bmcIP
	reqListNode.Node.PXEMacAddr = pxeMacAdr
	reqListNode.Node.Status = status
	reqListNode.Node.CPUCores = int32(cpuCores)
	reqListNode.Node.Memory = int32(memory)
	reqListNode.Node.Description = description
	reqListNode.Node.Active = int32(active)
	reqListNode.Row = int64(row)
	reqListNode.Page = int64(page)

	resListNode, err := client.RC.GetNodeList(&reqListNode)
	if err != nil {
		return nil, err
	}

	var nodeList []model.Node
	for _, pNode := range resListNode.Node {
		modelNode, err := pbNodeToModelNode(pNode)
		if err != nil {
			return nil, err
		}
		nodeList = append(nodeList, *modelNode)
	}

	return nodeList, nil
}

func AllNode(args map[string]interface{}) (interface{}, error) {
	return ListNode(args)
}

func NumNode() (interface{}, error) {
	num, err := client.RC.GetNodeNum()
	if err != nil {
		return nil, err
	}

	var modelNodeNum model.NodeNum
	modelNodeNum.Number = num

	return modelNodeNum, nil
}

func NodeDetail(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	return client.RC.GetNodeDetail(uuid)
}
