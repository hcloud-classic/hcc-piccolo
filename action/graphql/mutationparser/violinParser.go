package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/graphql/queryparser"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// CreateServer : Create a server
func CreateServer(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	groupID, groupIDOk := args["group_id"].(int)
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOK := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	userUUID, userUUIDOk := args["user_uuid"].(string)
	nrNode, nrNodeOk := args["nr_node"].(int)

	var reqCreateServer pb.ReqCreateServer
	var reqServer pb.Server
	reqCreateServer.Server = &reqServer

	if groupIDOk {
		reqCreateServer.Server.GroupID = int64(groupID)
	}
	if subnetUUIDOk {
		reqCreateServer.Server.SubnetUUID = subnetUUID
	}
	if osOK {
		reqCreateServer.Server.OS = os
	}
	if serverNameOk {
		reqCreateServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqCreateServer.Server.ServerDesc = serverDesc
	}
	if cpuOk {
		reqCreateServer.Server.CPU = int32(cpu)
	} else {
		reqCreateServer.Server.CPU = 0
	}
	if memoryOk {
		reqCreateServer.Server.Memory = int32(memory)
	} else {
		reqCreateServer.Server.Memory = 0
	}
	if diskSizeOk {
		reqCreateServer.Server.DiskSize = int32(diskSize)
	}
	if userUUIDOk {
		reqCreateServer.Server.UserUUID = userUUID
	}
	if nrNodeOk {
		reqCreateServer.NrNode = int32(nrNode)
	}

	reqCreateServer.Token = tokenString

	resCreateServer, err := client.RC.CreateServer(&reqCreateServer)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resCreateServer.Server, resCreateServer.HccErrorStack)

	err = dao.WriteServerAction(
		resCreateServer.Server.UUID,
		"violin / create_server",
		"Success",
		"",
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	return *modelServer, nil
}

// UpdateServer : Update the infos of the server
func UpdateServer(args map[string]interface{}, isAdmin bool, isMaster bool, id string) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	status, statusOk := args["status"].(string)

	if !isMaster || !isAdmin {
		server, err := queryparser.Server(args)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		if server.(model.Server).Errors != nil && len(server.(model.Server).Errors)!= 0 {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(server.(model.Server).Errors[0].ErrCode, server.(model.Server).Errors[0].ErrText)}, nil
		}

		if !isMaster {
			groupID, _ := args["group_id"].(int)

			if int(server.(model.Server).GroupID) != groupID {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group's server if you are not a master")}, nil
			}
		}
		if !isMaster && !isAdmin {
			if server.(model.Server).UserUUID != id {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other server if you are not a master or the admin")}, nil
			}
		}
	}

	var reqUpdateServer pb.ReqUpdateServer
	var reqServer pb.Server
	reqUpdateServer.Server = &reqServer

	reqUpdateServer.Server.UUID = requestedUUID
	if serverNameOk {
		reqUpdateServer.Server.ServerName = serverName
	}
	if serverDescOk {
		reqUpdateServer.Server.ServerDesc = serverDesc
	}
	if statusOk {
		reqUpdateServer.Server.Status = status
	}

	resUpdateServer, err := client.RC.UpdateServer(&reqUpdateServer)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resUpdateServer.Server, resUpdateServer.HccErrorStack)

	return *modelServer, nil
}

// UpdateServerNodes : Update nodes of the server
func UpdateServerNodes(args map[string]interface{}, isAdmin bool, isMaster bool, id string) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["server_uuid"].(string)
	selectedNodes, selectedNodesOk := args["selected_nodes"].(string)

	if !requestedUUIDOk || !selectedNodesOk {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need server_uuid and selected_nodes argument")}, nil
	}

	if !isMaster || !isAdmin {
		server, err := queryparser.Server(args)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		if server.(model.Server).Errors != nil && len(server.(model.Server).Errors)!= 0 {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(server.(model.Server).Errors[0].ErrCode, server.(model.Server).Errors[0].ErrText)}, nil
		}

		if !isMaster {
			groupID, _ := args["group_id"].(int)

			if int(server.(model.Server).GroupID) != groupID {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group's server if you are not a master")}, nil
			}
		}
		if !isMaster && !isAdmin {
			if server.(model.Server).UserUUID != id {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other server if you are not a master or the admin")}, nil
			}
		}
	}

	var reqUpdateServerNodes pb.ReqUpdateServerNodes

	reqUpdateServerNodes.ServerUUID = requestedUUID
	reqUpdateServerNodes.SelectedNodes = selectedNodes
	resUpdateServerNodes, err := client.RC.UpdateServerNodes(&reqUpdateServerNodes)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resUpdateServerNodes.Server, resUpdateServerNodes.HccErrorStack)

	return *modelServer, nil
}

// DeleteServer : Delete the server
func DeleteServer(args map[string]interface{}, isAdmin bool, isMaster bool, id string) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	if !isMaster || !isAdmin {
		server, err := queryparser.Server(args)
		if err != nil {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		if server.(model.Server).Errors != nil && len(server.(model.Server).Errors)!= 0 {
			return model.Server{Errors: errconv.ReturnHccErrorPiccolo(server.(model.Server).Errors[0].ErrCode, server.(model.Server).Errors[0].ErrText)}, nil
		}

		if !isMaster {
			groupID, _ := args["group_id"].(int)

			if int(server.(model.Server).GroupID) != groupID {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't delete the other group's server if you are not a master")}, nil
			}
		}
		if !isMaster && !isAdmin {
			if server.(model.Server).UserUUID != id {
				return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't delete the other server if you are not a master or the admin")}, nil
			}
		}
	}

	resDeleteServer, err := client.RC.DeleteServer(requestedUUID)
	if err != nil {
		return model.Server{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServer := pbtomodel.PbServerToModelServer(resDeleteServer.Server, resDeleteServer.HccErrorStack)

	// *** We are using ARCHIVE engine for server_actions table ***
	//err = dao.DeleteServerAction(requestedUUID)
	//if err != nil {
	//	logger.Logger.Println(err)
	//}

	return *modelServer, nil
}

// CreateServerNode : Create a info of server's node
func CreateServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	var reqCreateServerNode pb.ReqCreateServerNode
	if serverUUIDOk {
		reqCreateServerNode.ServerNode.ServerUUID = serverUUID
	}
	if nodeUUIDOk {
		reqCreateServerNode.ServerNode.NodeUUID = nodeUUID
	}

	resCreateServerNode, err := client.RC.CreateServerNode(&reqCreateServerNode)
	if err != nil {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServerNode := pbtomodel.PbServerNodeToModelServerNode(resCreateServerNode.ServerNode, nil, nil, resCreateServerNode.HccErrorStack)

	return *modelServerNode, nil
}

// DeleteServerNode : Delete a info of server's node
func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resDeleteServerNode, err := client.RC.DeleteServerNode(requestedUUID)
	if err != nil {
		return model.ServerNode{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelServerNode := pbtomodel.PbServerNodeToModelServerNode(resDeleteServerNode.ServerNode, nil, nil, resDeleteServerNode.HccErrorStack)

	return *modelServerNode, nil
}
