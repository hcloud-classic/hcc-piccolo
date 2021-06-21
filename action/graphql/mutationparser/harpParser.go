package mutationparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/sqlite/serveractions"
	"hcc/piccolo/model"
)

// CreateSubnet : Create a subnet
func CreateSubnet(args map[string]interface{}) (interface{}, error) {
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	var subnet rpcharp.Subnet
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

	resCreateSubnet, err := client.RC.CreateSubnet(&rpcharp.ReqCreateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resCreateSubnet.Subnet, &resCreateSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// UpdateSubnet : Update infos of the subnet
func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

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

	var subnet rpcharp.Subnet
	subnet.UUID = requestedUUID
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

	resUpdateSubnet, err := client.RC.UpdateSubnet(&rpcharp.ReqUpdateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resUpdateSubnet.Subnet, &resUpdateSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// DeleteSubnet : Delete the subnet
func DeleteSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resDeleteSubnet, err := client.RC.DeleteSubnet(requestedUUID)
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resDeleteSubnet.Subnet, &resDeleteSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// CreateDHCPDConf : Create the configuration of the DHCP server
func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	nodeUUIDs, nodeUUIDsOk := args["nodeUUIDs"].(string)
	if !subnetUUIDOk || !nodeUUIDsOk {
		return model.CreateDHCPConfResult{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need subnet_uuid and nodeUUIDs arguments")}, nil
	}

	resCreateDHCPDConfig, err := client.RC.CreateDHCPDConfig(subnetUUID, nodeUUIDs)
	if err != nil {
		return model.CreateDHCPConfResult{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resCreateDHCPDConfig.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return model.CreateDHCPConfResult{Result: resCreateDHCPDConfig.Result, Errors: Errors}, nil
}

// CreateAdaptiveIPServer : Create a adaptiveIP server
func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	tokenString, _ := args["token"].(string)

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	publicIP, publicIPOk := args["public_ip"].(string)

	var reqCreateAdaptiveIPServer rpcharp.ReqCreateAdaptiveIPServer
	if serverUUIDOk {
		reqCreateAdaptiveIPServer.ServerUUID = serverUUID
	}
	if publicIPOk {
		reqCreateAdaptiveIPServer.PublicIP = publicIP
	}

	resCreateadAptiveIPServer, err := client.RC.CreateAdaptiveIPServer(&reqCreateAdaptiveIPServer)
	if err != nil {
		return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(&resCreateadAptiveIPServer.HccErrorStack)

	resAdaptiveIPServer := resCreateadAptiveIPServer.AdaptiveipServer
	adaptiveIPServer := model.AdaptiveIPServer{
		ServerUUID:     resAdaptiveIPServer.ServerUUID,
		PublicIP:       resAdaptiveIPServer.PublicIP,
		PrivateIP:      resAdaptiveIPServer.PrivateIP,
		PrivateGateway: resAdaptiveIPServer.PrivateGateway,
	}

	var success bool
	var errStr = ""

	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errors.ReturnHccEmptyErrorPiccolo()
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

	err = serveractions.WriteServerAction(
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
		return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resDeleteAdaptiveIPServer, err := client.RC.DeleteAdaptiveIPServer(requestedUUID)
	if err != nil {
		return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var success bool
	var errStr = ""

	hccErrStack := errconv.GrpcStackToHcc(&resDeleteAdaptiveIPServer.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 {
		if Errors[0].ErrCode == 0 {
			success = true
			Errors = errors.ReturnHccEmptyErrorPiccolo()
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

	err = serveractions.WriteServerAction(
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

	var reqCreateAdaptiveIPSetting rpcharp.ReqCreateAdaptiveIPSetting
	var reqAdaptiveipSetting rpcharp.AdaptiveIPSetting

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
		return model.AdaptiveIPSetting{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	adaptiveipSetting := resCreateAdaptiveIPSetting.AdaptiveipSetting

	hccErrStack := errconv.GrpcStackToHcc(&resCreateAdaptiveIPSetting.HccErrorStack)
	Errors := *hccErrStack.ConvertReportForm()
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errors.ReturnHccEmptyErrorPiccolo()
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
