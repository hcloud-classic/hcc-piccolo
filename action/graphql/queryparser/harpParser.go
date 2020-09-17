package queryparser

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	rpcmsgType "hcc/piccolo/action/grpc/pb/rpcmsgType"
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

// Subnet : Get infos of the subnet
func Subnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetSubnet, err := client.RC.GetSubnet(uuid)
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbSubnetToModelSubnet(resGetSubnet.Subnet, &resGetSubnet.HccErrorStack)

	return *modelSubnet, nil
}

// ListSubnet : Get subnet list with provided options
func ListSubnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
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

	if uuidOk {
		reqListSubnet.Subnet.UUID = uuid
	}
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
		return model.SubnetList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var subnetList []model.Subnet
	for _, pSubnet := range resListSubnet.Subnet {
		modelSubnet := pbSubnetToModelSubnet(pSubnet, nil)
		subnetList = append(subnetList, *modelSubnet)
	}

	hccErrStack := errconv.GrpcStackToHcc(&resListSubnet.HccErrorStack)

	return model.SubnetList{Subnets: subnetList, Errors: *hccErrStack}, nil
}

// AllSubnet : Get subnet list with provided options (Just call ListSubnet())
func AllSubnet(args map[string]interface{}) (interface{}, error) {
	return ListSubnet(args)
}

// NumSubnet : Get number of subnets
func NumSubnet() (interface{}, error) {
	num, err := client.RC.GetSubnetNum()
	if err != nil {
		return model.Subnet{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelSubnetNum model.SubnetNum
	modelSubnetNum.Number = num

	return modelSubnetNum, nil
}

// GetAdaptiveIPAvailableIPList : Get available IP list of AdaptiveIP
func GetAdaptiveIPAvailableIPList() (interface{}, error) {
	resGetAdaptiveIPAvailableIPList, err := client.RC.GetAdaptiveIPAvailableIPList()
	if err != nil {
		return model.AdaptiveIPAvailableIPList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var adaptiveIPAvailableIPList []string

	for _, availableIP := range resGetAdaptiveIPAvailableIPList.AdaptiveipAvailableipList.AvailableIp {
		adaptiveIPAvailableIPList = append(adaptiveIPAvailableIPList, availableIP)
	}

	return model.AdaptiveIPAvailableIPList{
		AvailableIPList: adaptiveIPAvailableIPList,
	}, nil
}

// GetAdaptiveIPSetting : Get infos of the adaptiveIP settings
func GetAdaptiveIPSetting() (interface{}, error) {
	resGetAdaptiveIPSetting, err := client.RC.GetAdaptiveIPSetting()
	if err != nil {
		return model.AdaptiveIPSetting{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
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
		return model.AdaptiveIPSetting{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resGetAdaptiveIPServer, err := client.RC.GetAdaptiveIPServer(serverUUID)
	if err != nil {
		return model.AdaptiveIPSetting{Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
	}

	return model.AdaptiveIPServer{
		ServerUUID:     resGetAdaptiveIPServer.AdaptiveipServer.ServerUUID,
		PublicIP:       resGetAdaptiveIPServer.AdaptiveipServer.PublicIP,
		PrivateIP:      resGetAdaptiveIPServer.AdaptiveipServer.PrivateIP,
		PrivateGateway: resGetAdaptiveIPServer.AdaptiveipServer.PrivateGateway,
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
		return model.AdaptiveIPServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var adaptiveIPServerList []model.AdaptiveIPServer
	for _, adaptiveIPServer := range resAdaptiveIPServerList.AdaptiveipServer {
		_createdAt, err := ptypes.Timestamp(adaptiveIPServer.CreatedAt)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTimestampConversionError, err.Error())}, nil
		}

		adaptiveIPServerList = append(adaptiveIPServerList, model.AdaptiveIPServer{
			ServerUUID:     adaptiveIPServer.ServerUUID,
			PublicIP:       adaptiveIPServer.PublicIP,
			PrivateIP:      adaptiveIPServer.PrivateIP,
			PrivateGateway: adaptiveIPServer.PrivateGateway,
			CreatedAt:      _createdAt,
		})
	}

	return model.AdaptiveIPServerList{AdaptiveIPServers: adaptiveIPServerList}, nil
}

// AllAdaptiveIPServer : Get adaptiveIP server list with provided options (Just call ListAdaptiveIPServer())
func AllAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	return ListAdaptiveIPServer(args)
}

// NumAdaptiveIPServer : Get number of adaptiveIP servers
func NumAdaptiveIPServer() (interface{}, error) {
	resGetAdaptiveIPServerNum, err := client.RC.GetAdaptiveIPServerNum()
	if err != nil {
		return model.AdaptiveIPServerNum{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	return model.AdaptiveIPServerNum{Number: int(resGetAdaptiveIPServerNum.Num)}, nil
}
