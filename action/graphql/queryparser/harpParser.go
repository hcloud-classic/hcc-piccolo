package queryparser

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

// Subnet : Get infos of the subnet
func Subnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	subnet, err := client.RC.GetSubnet(uuid)
	if err != nil {
		return nil, err
	}

	modelSubnet, err := pbSubnetToModelSubnet(subnet)
	if err != nil {
		return nil, err
	}

	return *modelSubnet, nil
}

// ListSubnet : Get subnet list with provided options
func ListSubnet(args map[string]interface{}) (interface{}, error) {
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
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListSubnet rpcharp.ReqGetSubnetList
	var reqSubnet rpcharp.Subnet
	reqListSubnet.Subnet = &reqSubnet

	if networkIPOk {
		reqListSubnet.Subnet.NetworkIP = networkIP
	}
	if netmaskOk {
		reqListSubnet.Subnet.Netmask = netmask
	}
	if gatewayOk {
		reqListSubnet.Subnet.Gateway = gateway
	}
	if nextServerOk {
		reqListSubnet.Subnet.NextServer = nextServer
	}
	if nameServerOk {
		reqListSubnet.Subnet.NameServer = nameServer
	}
	if domainNameOk {
		reqListSubnet.Subnet.DomainName = domainName
	}
	if serverUUIDOk {
		reqListSubnet.Subnet.ServerUUID = serverUUID
	}
	if leaderNodeUUIDOk {
		reqListSubnet.Subnet.LeaderNodeUUID = leaderNodeUUID
	}
	if osOk {
		reqListSubnet.Subnet.OS = os
	}
	if subnetNameOk {
		reqListSubnet.Subnet.SubnetName = subnetName
	}
	if rowOk {
		reqListSubnet.Row = int64(row)
	}
	if pageOk {
		reqListSubnet.Page = int64(page)
	}

	resListSubnet, err := client.RC.GetSubnetList(&reqListSubnet)
	if err != nil {
		return nil, err
	}

	var subnetList []model.Subnet
	for _, pSubnet := range resListSubnet.Subnet {
		modelSubnet, err := pbSubnetToModelSubnet(pSubnet)
		if err != nil {
			return nil, err
		}
		subnetList = append(subnetList, *modelSubnet)
	}

	return subnetList, nil
}

// AllSubnet : Get subnet list with provided options (Just call ListSubnet())
func AllSubnet(args map[string]interface{}) (interface{}, error) {
	return ListSubnet(args)
}

// NumSubnet : Get number of subnets
func NumSubnet() (interface{}, error) {
	num, err := client.RC.GetSubnetNum()
	if err != nil {
		return nil, err
	}

	var modelSubnetNum model.SubnetNum
	modelSubnetNum.Number = num

	return modelSubnetNum, nil
}

// GetAdaptiveIPSetting : Get infos of the adaptiveIP settings
func GetAdaptiveIPSetting() (interface{}, error) {
	resGetAdaptiveIPSetting, err := client.RC.GetAdaptiveIPSetting()
	if err != nil {
		return nil, err
	}

	adaptiveipSetting := resGetAdaptiveIPSetting.AdaptiveipSetting
	return model.AdaptiveIPSetting{
		ExtIfaceIPAddress: adaptiveipSetting.ExtIfaceIPAddress,
		Netmask:           adaptiveipSetting.Netmask,
		GatewayAddress:    adaptiveipSetting.GatewayAddress,
		StartIPAddress:    adaptiveipSetting.StartIPAddress,
		EndIPAddress:      adaptiveipSetting.EndIPAddress,
	}, nil
}

// AdaptiveIPServer : Get infos of the adaptiveIP server
func AdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	adaptiveIPServer, err := client.RC.GetAdaptiveIPServer(serverUUID)
	if err != nil {
		return nil, err
	}

	return model.AdaptiveIPServer{
		ServerUUID:     adaptiveIPServer.ServerUUID,
		PublicIP:       adaptiveIPServer.PublicIP,
		PrivateIP:      adaptiveIPServer.PrivateIP,
		PrivateGateway: adaptiveIPServer.PrivateGateway,
	}, nil
}

// ListAdaptiveIPServer : Get adaptiveIP server list with provided options
func ListAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	publicIP, publicIPOk := args["public_ip"].(string)
	privateIP, privateIPOk := args["private_ip"].(string)
	privateGateway, privateGatewayOk := args["private_gateway"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqGetAdaptiveIPServerList rpcharp.ReqGetAdaptiveIPServerList
	if serverUUIDOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.ServerUUID = serverUUID
	}
	if publicIPOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.PublicIP = publicIP
	}
	if privateIPOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.PrivateIP = privateIP
	}
	if privateGatewayOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.PrivateGateway = privateGateway
	}
	if rowOk {
		reqGetAdaptiveIPServerList.Row = int64(row)
	}
	if pageOk {
		reqGetAdaptiveIPServerList.Page = int64(page)
	}
	reqGetAdaptiveIPServerList.AdaptiveipServer.ServerUUID = serverUUID

	resAdaptiveIPServerList, err := client.RC.GetAdaptiveIPServerList(&reqGetAdaptiveIPServerList)
	if err != nil {
		return nil, err
	}

	var adaptiveIPServerList []model.AdaptiveIPServer
	for _, adaptiveIPServer := range resAdaptiveIPServerList.AdaptiveipServer {
		adaptiveIPServerList = append(adaptiveIPServerList, model.AdaptiveIPServer{
			ServerUUID:     adaptiveIPServer.ServerUUID,
			PublicIP:       adaptiveIPServer.PublicIP,
			PrivateIP:      adaptiveIPServer.PrivateIP,
			PrivateGateway: adaptiveIPServer.PrivateGateway,
		})
	}

	return adaptiveIPServerList, nil
}

// AllAdaptiveIPServer : Get adaptiveIP server list with provided options (Just call ListAdaptiveIPServer())
func AllAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	return ListAdaptiveIPServer(args)
}

// NumAdaptiveIPServer : Get number of adaptiveIP servers
func NumAdaptiveIPServer() (interface{}, error) {
	num, err := client.RC.GetAdaptiveIPServerNum()
	if err != nil {
		return nil, err
	}

	return model.AdaptiveIPServerNum{Number: num}, nil
}
