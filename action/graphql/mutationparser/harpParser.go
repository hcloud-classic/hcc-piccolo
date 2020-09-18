package mutationparser

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
	"time"
)

func pbSubnetToModelSubnet(subnet *rpcharp.Subnet, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Subnet {
	var createdAt time.Time
	if subnet.CreatedAt == nil {
		createdAt, _ = ptypes.Timestamp(&timestamp.Timestamp{
			Seconds: 0,
			Nanos:   0,
		})
	} else {
		var err error

		createdAt, err = ptypes.Timestamp(subnet.CreatedAt)
		if err != nil {
			return &model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}
		}
	}

	modelSubnet := &model.Subnet{
		UUID:           subnet.UUID,
		NetworkIP:      subnet.NetworkIP,
		Netmask:        subnet.Netmask,
		Gateway:        subnet.Gateway,
		NextServer:     subnet.NextServer,
		NameServer:     subnet.NameServer,
		DomainName:     subnet.DomainName,
		ServerUUID:     subnet.ServerUUID,
		LeaderNodeUUID: subnet.LeaderNodeUUID,
		OS:             subnet.OS,
		SubnetName:     subnet.SubnetName,
		CreatedAt:      createdAt,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelSubnet.Errors = *hccErrStack
	}

	return modelSubnet
}

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

	modelSubnet := pbSubnetToModelSubnet(resCreateSubnet.Subnet, &resCreateSubnet.HccErrorStack)

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

	modelSubnet := pbSubnetToModelSubnet(resUpdateSubnet.Subnet, &resUpdateSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// DeleteSubnet : Delete the subnet
func DeleteSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	var subnet model.Subnet
	resDeleteSubnet, err := client.RC.DeleteSubnet(requestedUUID)
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}
	subnet.UUID = resDeleteSubnet.UUID

	return subnet, nil
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

	return model.CreateDHCPConfResult{Result: resCreateDHCPDConfig.Result}, nil
}

// CreateAdaptiveIPServer : Create a adaptiveIP server
func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
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

	resAdaptiveIPServer := resCreateadAptiveIPServer.AdaptiveipServer
	adaptiveIPServer := model.AdaptiveIPServer{
		ServerUUID:     resAdaptiveIPServer.ServerUUID,
		PublicIP:       resAdaptiveIPServer.PublicIP,
		PrivateIP:      resAdaptiveIPServer.PrivateIP,
		PrivateGateway: resAdaptiveIPServer.PrivateGateway,
	}

	return adaptiveIPServer, nil
}

// DeleteAdaptiveIPServer : Delete the adaptiveIP server
func DeleteAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["server_uuid"].(string)
	if !requestedUUIDOk {
		return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resDeleteAdaptiveIPServer, err := client.RC.DeleteAdaptiveIPServer(requestedUUID)
	if err != nil {
		return model.AdaptiveIPServer{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return model.AdaptiveIPServer{ServerUUID: resDeleteAdaptiveIPServer.ServerUUID}, nil
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
	return model.AdaptiveIPSetting{
		ExtIfaceIPAddress: adaptiveipSetting.ExtIfaceIPAddress,
		Netmask:           adaptiveipSetting.Netmask,
		GatewayAddress:    adaptiveipSetting.GatewayAddress,
		StartIPAddress:    adaptiveipSetting.StartIPAddress,
		EndIPAddress:      adaptiveipSetting.EndIPAddress,
	}, nil
}
