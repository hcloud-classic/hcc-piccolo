package queryparser

import (
	"encoding/json"
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// Server : Get infos of the server
func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetServer, err := client.RC.GetServer(uuid)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := *pbtomodel.PbServerToModelServer(resGetServer.Server, resGetServer.HccErrorStack)

	// group_name
	group, err := dao.ReadGroup(int(modelServer.GroupID))
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}
	modelServer.GroupName = group.Name

	// external_ip
	resGetAdaptiveIPServer, err := client.RC.GetAdaptiveIPServer(uuid)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	if resGetAdaptiveIPServer.AdaptiveipServer != nil {
		modelServer.ExternalIP = resGetAdaptiveIPServer.AdaptiveipServer.PublicIP
	}

	// Get nodes
	var nodeList []model.Node
	resGetNodeList, err := client.RC.GetNodeList(&pb.ReqGetNodeList{
		Node: &pb.Node{
			ServerUUID: uuid,
		},
	})
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	if resGetNodeList.Node != nil {
		for _, node := range resGetNodeList.Node {
			modelNode := *pbtomodel.PbNodeToModelNode(node, nil)
			resGetNodeDetail, _ := client.RC.GetNodeDetail(node.UUID)
			if resGetNodeDetail.NodeDetail == nil {
				continue
			}
			// cpu_model of the node
			nodeDetailData := resGetNodeDetail.NodeDetail.GetNodeDetailData()
			var nodeDetailJSON model.NodeDetailData
			err := json.Unmarshal([]byte(nodeDetailData), &nodeDetailJSON)
			if err != nil {
				continue
			}
			modelNode.CPUModel = nodeDetailJSON.CPUs[0].Model
			nodeList = append(nodeList, modelNode)
		}
	}

	// node_list
	modelServer.NodeList = nodeList

	// nic_speed, pxe_boot_ip
	if len(nodeList) != 0 {
		modelServer.PXEBootIP = nodeList[0].NodeIP
		modelServer.NicSpeed = nodeList[0].NICSpeed
	}

	// TODO : Need to get mount path from cello

	queryArgs := make(map[string]interface{})
	queryArgs["server_uuid"] = modelServer.UUID
	numServerNode, _ := NumServerNode(queryArgs)
	modelServerNodeNum := numServerNode.(model.ServerNodeNum)
	modelServer.Nodes = modelServerNodeNum.Number

	return modelServer, nil
}

// ListServer : Get server list with provided options
func ListServer(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
	groupID, groupIDOk := args["group_id"].(int)
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOk := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListServer pb.ReqGetServerList
	var reqServer pb.Server
	reqListServer.Server = &reqServer

	if uuidOk {
		reqListServer.Server.UUID = uuid
	}
	if groupIDOk {
		reqListServer.Server.GroupID = int64(groupID)
	}
	if subnetUUIDOk {
		reqListServer.Server.SubnetUUID = subnetUUID
	}
	if osOk {
		reqListServer.Server.OS = os
	}
	if serverNameOk {
		reqListServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqListServer.Server.ServerDesc = serverDesc
	}
	if cpuOk {
		reqListServer.Server.CPU = int32(cpu)
	}
	if memoryOk {
		reqListServer.Server.Memory = int32(memory)
	}
	if diskSizeOk {
		reqListServer.Server.DiskSize = int32(diskSize)
	}
	if statusOk {
		reqListServer.Server.Status = status
	}
	if userUUIDOk {
		reqListServer.Server.UserUUID = userUUID
	}
	if rowOk {
		reqListServer.Row = int64(row)
	}
	if pageOk {
		reqListServer.Page = int64(page)
	}

	resListServer, err := client.RC.GetServerList(&reqListServer)
	if err != nil {
		return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverList []model.Server
	for _, pServer := range resListServer.Server {
		modelServer := pbtomodel.PbServerToModelServer(pServer, nil)

		queryArgs := make(map[string]interface{})
		queryArgs["server_uuid"] = modelServer.UUID
		numServerNode, _ := NumServerNode(queryArgs)
		modelServerNodeNum := numServerNode.(model.ServerNodeNum)
		modelServer.Nodes = modelServerNodeNum.Number

		// group_name
		group, err := dao.ReadGroup(int(modelServer.GroupID))
		if err != nil {
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
		}
		modelServer.GroupName = group.Name

		// external_ip
		resGetAdaptiveIPServer, err := client.RC.GetAdaptiveIPServer(uuid)
		if err != nil {
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		if resGetAdaptiveIPServer.AdaptiveipServer != nil {
			modelServer.ExternalIP = resGetAdaptiveIPServer.AdaptiveipServer.PublicIP
		}

		// Get nodes
		var nodeList []model.Node
		resGetNodeList, err := client.RC.GetNodeList(&pb.ReqGetNodeList{
			Node: &pb.Node{
				ServerUUID: uuid,
			},
		})
		if err != nil {
			return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		if resGetNodeList.Node != nil {
			for _, node := range resGetNodeList.Node {
				modelNode := *pbtomodel.PbNodeToModelNode(node, nil)
				resGetNodeDetail, _ := client.RC.GetNodeDetail(node.UUID)
				if resGetNodeDetail.NodeDetail == nil {
					continue
				}
				// cpu_model of the node
				nodeDetailData := resGetNodeDetail.NodeDetail.GetNodeDetailData()
				var nodeDetailJSON model.NodeDetailData
				err := json.Unmarshal([]byte(nodeDetailData), &nodeDetailJSON)
				if err != nil {
					continue
				}
				modelNode.CPUModel = nodeDetailJSON.CPUs[0].Model
				nodeList = append(nodeList, modelNode)
			}
		}

		// node_list
		modelServer.NodeList = nodeList

		// nic_speed, pxe_boot_ip
		if len(nodeList) != 0 {
			modelServer.PXEBootIP = nodeList[0].NodeIP
			modelServer.NicSpeed = nodeList[0].NICSpeed
		}

		// TODO : Need to get mount path from cello

		serverList = append(serverList, *modelServer)
	}

	hccErrStack := errconv.GrpcStackToHcc(resListServer.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	numServer, err := NumServer(args)
	if err != nil {
		return model.ServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	modelServerNum := numServer.(model.ServerNum)
	if len(modelServerNum.Errors) != 0 {
		for _, numError := range modelServerNum.Errors {
			Errors = append(Errors, numError)
		}
		modelServerNum.Number = 0
	}

	return model.ServerList{Servers: serverList, TotalNum: modelServerNum.Number, Errors: Errors}, nil
}

// AllServer : Get server list with provided options (Just call ListServer())
func AllServer(args map[string]interface{}) (interface{}, error) {
	return ListServer(args)
}

// NumServer : Get number of servers
func NumServer(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetServerNum pb.ReqGetServerNum
	if groupIDOk {
		reqGetServerNum.GroupID = int64(groupID)
	}

	resGetServerNum, err := client.RC.GetServerNum(&reqGetServerNum)
	if err != nil {
		return nil, err
	}

	var modelServerNum model.ServerNum
	modelServerNum.Number = int(resGetServerNum.Num)

	hccErrStack := errconv.GrpcStackToHcc(resGetServerNum.HccErrorStack)
	modelServerNum.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(modelServerNum.Errors) != 0 && modelServerNum.Errors[0].ErrCode == 0 {
		modelServerNum.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelServerNum, nil
}

// ServerNode : Get infos of the serverNode
func ServerNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetServerNode, err := client.RC.GetServerNode(uuid)
	if err != nil {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	resGetNode, err := client.RC.GetNode(resGetServerNode.ServerNode.NodeUUID)
	if err != nil {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	resGetNodeDetail, err := client.RC.GetNodeDetail(resGetServerNode.ServerNode.NodeUUID)
	if err != nil {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return *pbtomodel.PbServerNodeToModelServerNode(resGetServerNode.ServerNode, resGetNode.Node, resGetNodeDetail.NodeDetail,
		resGetServerNode.HccErrorStack), nil
}

// ListServerNode : Get serverNode list with provided options
func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	var reqListServerNode pb.ReqGetServerNodeList
	reqListServerNode.ServerUUID = serverUUID

	resListServerNode, err := client.RC.GetServerNodeList(&reqListServerNode)
	if err != nil {
		return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var serverNodeList []model.ServerNode
	for _, pServerNode := range resListServerNode.ServerNode {
		resGetNode, err := client.RC.GetNode(pServerNode.NodeUUID)
		if err != nil {
			return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		resGetNodeDetail, err := client.RC.GetNodeDetail(pServerNode.NodeUUID)
		if err != nil {
			return model.ServerNodeList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		modelServerNode := pbtomodel.PbServerNodeToModelServerNode(pServerNode, resGetNode.Node, resGetNodeDetail.NodeDetail, nil)
		serverNodeList = append(serverNodeList, *modelServerNode)
	}

	hccErrStack := errconv.GrpcStackToHcc(resListServerNode.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.ServerNodeList{ServerNodes: serverNodeList, Errors: Errors}, nil
}

// AllServerNode : Get serverNode list with provided options (Just call ListServerNode())
func AllServerNode(args map[string]interface{}) (interface{}, error) {
	return ListServerNode(args)
}

// NumServerNode : Get number of serverNodes
func NumServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerNodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resGetServerNodeNum, err := client.RC.GetServerNodeNum(serverUUID)
	if err != nil {
		return model.ServerNodeNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelServerNodeNum model.ServerNodeNum
	modelServerNodeNum.Number = int(resGetServerNodeNum.Num)
	hccErrStack := errconv.GrpcStackToHcc(resGetServerNodeNum.HccErrorStack)
	modelServerNodeNum.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(modelServerNodeNum.Errors) != 0 && modelServerNodeNum.Errors[0].ErrCode == 0 {
		modelServerNodeNum.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelServerNodeNum, nil
}
