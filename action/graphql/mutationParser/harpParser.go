package mutationParser

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

func CreateSubnet(args map[string]interface{}) (interface{}, error) {
	networkIP, _ := args["network_ip"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	nextServer, _ := args["next_server"].(string)
	nameServer, _ := args["name_server"].(string)
	domainName, _ := args["domain_name"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	leaderNodeUUID, _ := args["leader_node_uuid"].(string)
	os, _ := args["os"].(string)
	subnetName, _ := args["subnet_name"].(string)

	var subnet rpcharp.Subnet
	subnet.NetworkIP = networkIP
	subnet.Netmask = netmask
	subnet.Gateway = gateway
	subnet.NextServer = nextServer
	subnet.NameServer = nameServer
	subnet.DomainName = domainName
	subnet.ServerUUID = serverUUID
	subnet.LeaderNodeUUID = leaderNodeUUID
	subnet.OS = os
	subnet.SubnetName = subnetName

	resCreateSubnet, err := client.RC.CreateSubnet(&rpcharp.ReqCreateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return nil, err
	}

	modelSubnet, err := pbSubnetToModelSubnet(resCreateSubnet.Subnet)
	
	return modelSubnet, nil
}

func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	networkIP, _ := args["network_ip"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	nextServer, _ := args["next_server"].(string)
	nameServer, _ := args["name_server"].(string)
	domainName, _ := args["domain_name"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	leaderNodeUUID, _ := args["leader_node_uuid"].(string)
	os, _ := args["os"].(string)
	subnetName, _ := args["subnet_name"].(string)

	var subnet rpcharp.Subnet
	subnet.UUID = requestedUUID
	subnet.NetworkIP = networkIP
	subnet.Netmask = netmask
	subnet.Gateway = gateway
	subnet.NextServer = nextServer
	subnet.NameServer = nameServer
	subnet.DomainName = domainName
	subnet.ServerUUID = serverUUID
	subnet.LeaderNodeUUID = leaderNodeUUID
	subnet.OS = os
	subnet.SubnetName = subnetName

	resUpdateSubnet, err := client.RC.UpdateSubnet(&rpcharp.ReqUpdateSubnet{
		Subnet: &subnet,
	})
	if err != nil {
		return nil, err
	}

	modelSubnet, err := pbSubnetToModelSubnet(resUpdateSubnet.Subnet)

	return modelSubnet, nil
}

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

func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	node_uuids, node_uuidsOk := args["node_uuids"].(string)
	if !subnetUUIDOk || !node_uuidsOk {
		return nil, errors.New("need subnet_uuid and node_uuids arguments")
	}

	result, err := client.RC.CreateDHCPDConfig(subnetUUID, node_uuids)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, _ := args["server_uuid"].(string)
	publicIP, _ := args["public_ip"].(string)

	resCreateadAptiveIPServer, err := client.RC.CreateAdaptiveIPServer(&rpcharp.ReqCreateAdaptiveIPServer{
		ServerUUID: serverUUID,
		PublicIP:   publicIP,
	})
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

func CreateAdaptiveIPSetting(args map[string]interface{}) (interface{}, error) {
	extIfaceIPAddress, _ := args["ext_ifaceip_address"].(string)
	netmask, _ := args["netmask"].(string)
	gatewayAddress, _ := args["gateway_address"].(string)
	startIPaddressOk, _ := args["start_ip_address"].(string)
	endIPaddressOk, _ := args["end_ip_address"].(string)

	reqCreateAdaptiveIPSetting := &rpcharp.ReqCreateAdaptiveIPSetting {
		AdaptiveipSetting: &rpcharp.AdaptiveIPSetting{
			ExtIfaceIPAddress: extIfaceIPAddress,
			Netmask:           netmask,
			GatewayAddress:    gatewayAddress,
			StartIPAddress:    startIPaddressOk,
			EndIPAddress:      endIPaddressOk,
		},
	}

	resCreateAdaptiveIPSetting, err := client.RC.CreateAdaptiveIPSetting(reqCreateAdaptiveIPSetting)
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
