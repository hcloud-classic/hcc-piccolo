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

// CreateSubnet : Create a subnet
func CreateSubnet(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	var subnet pb.Subnet
	if groupIDOk {
		subnet.GroupID = int64(groupID)
	}
	if networkIPOk {
		subnet.NetworkIP = networkIP
	}
	if netmaskOk {
		subnet.Netmask = netmask
	}
	if gatewayOk {
		subnet.Gateway = gateway
	}
	if nextServerOk {
		subnet.NextServer = nextServer
	}
	if nameServerOk {
		subnet.NameServer = nameServer
	}
	if domainNameOk {
		subnet.DomainName = domainName
	}
	if osOk {
		subnet.OS = os
	}
	if subnetNameOk {
		subnet.SubnetName = subnetName
	}

	resCreateSubnet, err := client.RC.CreateSubnet(&pb.ReqCreateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resCreateSubnet.Subnet, resCreateSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// UpdateSubnet : Update infos of the subnet
func UpdateSubnet(args map[string]interface{}, isMaster bool) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	if !isMaster {
		groupID, _ := args["group_id"].(int)
		subnet, err := queryparser.Subnet(args)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		if int(subnet.(model.Subnet).GroupID) != groupID {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other subnet if you are not a master")}, nil
		}
	}

	groupID, groupIDOk := args["group_id"].(int)
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	serverUUID, serverUUIDOk := args["sever_uuid"].(string)
	leaderNodeUUID, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	var subnet pb.Subnet
	subnet.UUID = requestedUUID
	if groupIDOk {
		subnet.GroupID = int64(groupID)
	}
	if networkIPOk {
		subnet.NetworkIP = networkIP
	}
	if netmaskOk {
		subnet.Netmask = netmask
	}
	if gatewayOk {
		subnet.Gateway = gateway
	}
	if nextServerOk {
		subnet.NextServer = nextServer
	}
	if nameServerOk {
		subnet.NameServer = nameServer
	}
	if domainNameOk {
		subnet.DomainName = domainName
	}
	if serverUUIDOk {
		subnet.ServerUUID = serverUUID
	}
	if leaderNodeUUIDOk {
		subnet.LeaderNodeUUID = leaderNodeUUID
	}
	if osOk {
		subnet.OS = os
	}
	if subnetNameOk {
		subnet.SubnetName = subnetName
	}

	resUpdateSubnet, err := client.RC.UpdateSubnet(&pb.ReqUpdateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resUpdateSubnet.Subnet, resUpdateSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// DeleteSubnet : Delete the subnet
func DeleteSubnet(args map[string]interface{}, isMaster bool) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	if !isMaster {
		groupID, _ := args["group_id"].(int)
		subnet, err := queryparser.Subnet(args)
		if err != nil {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}

		if int(subnet.(model.Subnet).GroupID) != groupID {
			return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't delete the other subnet if you are not a master")}, nil
		}
	}

	resDeleteSubnet, err := client.RC.DeleteSubnet(requestedUUID)
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resDeleteSubnet.Subnet, resDeleteSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// CreateDHCPDConf : Create the configuration of the DHCP server
func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	if !subnetUUIDOk {
		return model.CreateDHCPConfResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a subnet_uuid argument")}, nil
	}

	resCreateDHCPDConfig, err := client.RC.CreateDHCPDConfig(subnetUUID)
	if err != nil {
		return model.CreateDHCPConfResult{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resCreateDHCPDConfig.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.CreateDHCPConfResult{Result: resCreateDHCPDConfig.Result, Errors: Errors}, nil
}

// CreateAdaptiveIPServer : Create a adaptiveIP server
func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	groupID, groupIDOk := args["group_id"].(int)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	publicIP, publicIPOk := args["public_ip"].(string)

	var reqCreateAdaptiveIPServer pb.ReqCreateAdaptiveIPServer
	if groupIDOk {
		reqCreateAdaptiveIPServer.GroupID = int64(groupID)
	}
	if serverUUIDOk {
		reqCreateAdaptiveIPServer.ServerUUID = serverUUID
	}
	if publicIPOk {
		reqCreateAdaptiveIPServer.PublicIP = publicIP
	}

	resCreateadAptiveIPServer, err := client.RC.CreateAdaptiveIPServer(&reqCreateAdaptiveIPServer)
	if err != nil {
		return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resCreateadAptiveIPServer.HccErrorStack)

	resAdaptiveIPServer := resCreateadAptiveIPServer.AdaptiveipServer
	adaptiveIPServer := model.AdaptiveIPServer{
		ServerUUID:     resAdaptiveIPServer.ServerUUID,
		GroupID:        resAdaptiveIPServer.GroupID,
		PublicIP:       resAdaptiveIPServer.PublicIP,
		PrivateIP:      resAdaptiveIPServer.PrivateIP,
		PrivateGateway: resAdaptiveIPServer.PrivateGateway,
	}

	var success bool
	var errStr = ""

	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errconv.ReturnHccEmptyErrorPiccolo()
		} else {
			success = false
			errStr = Errors[0].ErrText
		}
	} else {
		success = true
	}

	var result string
	if success {
		result = "Success"
	} else {
		result = "Failed"
	}

	err = dao.WriteServerAction(
		serverUUID,
		"harp / create_adaptiveip_server",
		result,
		errStr,
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	adaptiveIPServer.Errors = Errors

	return adaptiveIPServer, nil
}

// DeleteAdaptiveIPServer : Delete the adaptiveIP server
func DeleteAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	requestedUUID, requestedUUIDOk := args["server_uuid"].(string)
	if !requestedUUIDOk {
		return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resDeleteAdaptiveIPServer, err := client.RC.DeleteAdaptiveIPServer(requestedUUID)
	if err != nil {
		return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var success bool
	var errStr = ""

	hccErrStack := errconv.GrpcStackToHcc(resDeleteAdaptiveIPServer.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errconv.ReturnHccEmptyErrorPiccolo()
		} else {
			success = false
			errStr = Errors[0].ErrText
		}
	} else {
		success = true
	}

	var result string
	if success {
		result = "Success"
	} else {
		result = "Failed"
	}

	err = dao.WriteServerAction(
		requestedUUID,
		"harp / delete_adaptiveip_server",
		result,
		errStr,
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	return model.AdaptiveIPServer{ServerUUID: resDeleteAdaptiveIPServer.ServerUUID, Errors: Errors}, nil
}

// CreateAdaptiveIPSetting : Create settings of the adaptiveIP
func CreateAdaptiveIPSetting(args map[string]interface{}) (interface{}, error) {
	extIfaceIPAddress, extIfaceIPAddressOk := args["ext_ifaceip_address"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gatewayAddress, gatewayAddressOk := args["gateway_address"].(string)
	startIPaddress, startIPaddressOk := args["start_ip_address"].(string)
	endIPaddress, endIPaddressOk := args["end_ip_address"].(string)

	var reqCreateAdaptiveIPSetting pb.ReqCreateAdaptiveIPSetting
	var reqAdaptiveipSetting pb.AdaptiveIPSetting

	reqCreateAdaptiveIPSetting.AdaptiveipSetting = &reqAdaptiveipSetting

	if extIfaceIPAddressOk {
		reqCreateAdaptiveIPSetting.AdaptiveipSetting.ExtIfaceIPAddress = extIfaceIPAddress
	}
	if netmaskOk {
		reqCreateAdaptiveIPSetting.AdaptiveipSetting.Netmask = netmask
	}
	if gatewayAddressOk {
		reqCreateAdaptiveIPSetting.AdaptiveipSetting.GatewayAddress = gatewayAddress
	}
	if startIPaddressOk {
		reqCreateAdaptiveIPSetting.AdaptiveipSetting.StartIPAddress = startIPaddress
	}
	if endIPaddressOk {
		reqCreateAdaptiveIPSetting.AdaptiveipSetting.EndIPAddress = endIPaddress
	}

	resCreateAdaptiveIPSetting, err := client.RC.CreateAdaptiveIPSetting(&reqCreateAdaptiveIPSetting)
	if err != nil {
		return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	adaptiveipSetting := resCreateAdaptiveIPSetting.AdaptiveipSetting

	hccErrStack := errconv.GrpcStackToHcc(resCreateAdaptiveIPSetting.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.AdaptiveIPSetting{
		ExtIfaceIPAddress: adaptiveipSetting.ExtIfaceIPAddress,
		Netmask:           adaptiveipSetting.Netmask,
		GatewayAddress:    adaptiveipSetting.GatewayAddress,
		StartIPAddress:    adaptiveipSetting.StartIPAddress,
		EndIPAddress:      adaptiveipSetting.EndIPAddress,
		Errors:            Errors,
	}, nil
}

// CreatePortForwarding : Create the AdaptiveIP Port Forwarding
func CreatePortForwarding(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	forwardTCP, forwardTCPOk := args["forwarding_tcp"].(bool)
	forwardUDP, forwardUDPOk := args["forwarding_udp"].(bool)
	externalPort, externalPortOk := args["external_port"].(int)
	internalPort, internalPortOk := args["internal_port"].(int)
	description, descriptionOk := args["description"].(string)

	if !serverUUIDOk || (!forwardTCPOk && !forwardUDPOk) || !externalPortOk || !internalPortOk || !descriptionOk {
		return model.PortForwarding{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need ServerUUID and ForwardTCP/ForwardUDP,"+
				"ExternalPort, InternalPort, Description arguments")}, nil
	}

	reqCreatePortForwarding := &pb.ReqCreatePortForwarding{
		PortForwarding: &pb.PortForwarding{
			ServerUUID:   serverUUID,
			ForwardTCP:   forwardTCP,
			ForwardUDP:   forwardUDP,
			ExternalPort: int64(externalPort),
			InternalPort: int64(internalPort),
			Description:  description,
		},
	}

	resCreatePortForwarding, err := client.RC.CreatePortForwarding(reqCreatePortForwarding)
	if err != nil {
		return model.PortForwarding{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resCreatePortForwarding.HccErrorStack)

	resPortForwarding := resCreatePortForwarding.PortForwarding

	var protocol string
	if resPortForwarding.ForwardTCP && resPortForwarding.ForwardUDP {
		protocol = "all"
	} else if forwardTCPOk {
		protocol = "tcp"
	} else if forwardUDPOk {
		protocol = "udp"
	}

	portForwarding := model.PortForwarding{
		ServerUUID:   resPortForwarding.ServerUUID,
		Protocol:     protocol,
		ExternalPort: resPortForwarding.ExternalPort,
		InternalPort: resPortForwarding.InternalPort,
		Description:  resPortForwarding.Description,
	}

	var success bool
	var errStr = ""

	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errconv.ReturnHccEmptyErrorPiccolo()
		} else {
			success = false
			errStr = Errors[0].ErrText
		}
	} else {
		success = true
	}

	var result string
	if success {
		result = "Success"
	} else {
		result = "Failed"
	}

	err = dao.WriteServerAction(
		serverUUID,
		"harp / create_port_forwarding",
		result,
		errStr,
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	portForwarding.Errors = Errors

	return portForwarding, nil
}

// DeletePortForwarding : Delete the AdaptiveIP Port Forwarding
func DeletePortForwarding(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	externalPort, externalPortOk := args["external_port"].(int)

	if !serverUUIDOk || !externalPortOk {
		return model.PortForwarding{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need server_uuid and external_port arguments")}, nil
	}

	resDeletePortForwarding, err := client.RC.DeletePortForwarding(&pb.ReqDeletePortForwarding{
		PortForwarding: &pb.PortForwarding{
			ServerUUID:   serverUUID,
			ExternalPort: int64(externalPort),
		},
	})
	if err != nil {
		return model.PortForwarding{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var success bool
	var errStr = ""

	hccErrStack := errconv.GrpcStackToHcc(resDeletePortForwarding.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errconv.ReturnHccEmptyErrorPiccolo()
		} else {
			success = false
			errStr = Errors[0].ErrText
		}
	} else {
		success = true
	}

	var result string
	if success {
		result = "Success"
	} else {
		result = "Failed"
	}

	err = dao.WriteServerAction(
		serverUUID,
		"harp / delete_port_forwarding",
		result,
		errStr,
		tokenString)
	if err != nil {
		logger.Logger.Println("WriteServerAction(): " + err.Error())
	}

	return model.PortForwarding{ServerUUID: resDeletePortForwarding.ServerUUID, Errors: Errors}, nil
}
