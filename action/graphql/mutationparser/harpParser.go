package mutationparser

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/model"
)

func pbSubnetToModelSubnet(subnet *rpcharp.Subnet) (*model.Subnet, error) {
	createdAt, err := ptypes.Timestamp(subnet.CreatedAt)
	if err != nil {
		return nil, err
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

	return modelSubnet, err
}

// CreateSubnet : Create a subnet
func CreateSubnet(args map[string]interface{}) (interface{}, error) {
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

	resCreateSubnet, err := client.RC.CreateSubnet(&rpcharp.ReqCreateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return nil, err
	}

	modelSubnet, err := pbSubnetToModelSubnet(resCreateSubnet.Subnet)
	if err != nil {
		return nil, err
	}

	return *modelSubnet, nil
}

// UpdateSubnet : Update infos of the subnet
func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
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
		return nil, err
	}

	modelSubnet, err := pbSubnetToModelSubnet(resUpdateSubnet.Subnet)
	if err != nil {
		return nil, err
	}

	return *modelSubnet, nil
}

// DeleteSubnet : Delete the subnet
func DeleteSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var subnet model.Subnet
	uuid, err := client.RC.DeleteSubnet(requestedUUID)
	if err != nil {
		return nil, err
	}
	subnet.UUID = uuid

	return subnet, nil
}

// CreateDHCPDConf : Create the configuration of the DHCP server
func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	nodeUUIDs, nodeUUIDsOk := args["nodeUUIDs"].(string)
	if !subnetUUIDOk || !nodeUUIDsOk {
		return nil, errors.New("need subnet_uuid and nodeUUIDs arguments")
	}

	result, err := client.RC.CreateDHCPDConfig(subnetUUID, nodeUUIDs)
	if err != nil {
		return nil, err
	}

	return result, nil
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
		return nil, err
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
		return nil, errors.New("need a server_uuid argument")
	}

	serverUUID, err := client.RC.DeleteAdaptiveIPServer(requestedUUID)
	if err != nil {
		return nil, err
	}

	return model.AdaptiveIPServer{ServerUUID: serverUUID}, nil
}

// CreateAdaptiveIPSetting : Create settings of the adaptiveIP
func CreateAdaptiveIPSetting(args map[string]interface{}) (interface{}, error) {
	extIfaceIPAddress, extIfaceIPAddressOk := args["ext_ifaceip_address"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gatewayAddress, gatewayAddressOk := args["gateway_address"].(string)
	startIPaddress, startIPaddressOk := args["start_ip_address"].(string)
	endIPaddress, endIPaddressOk := args["end_ip_address"].(string)

	var reqCreateAdaptiveIPSetting rpcharp.ReqCreateAdaptiveIPSetting
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
		return nil, err
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
