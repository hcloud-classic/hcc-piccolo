package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/model"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
)

// Subnet : Get infos of the subnet
func Subnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a uuid argument")}, nil
	}

	resGetSubnet, err := client.RC.GetSubnet(uuid)
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelSubnet := pbtomodel.PbSubnetToModelSubnet(resGetSubnet.Subnet, resGetSubnet.HccErrorStack)

	// group_name
	group, err := dao.ReadGroup(int(modelSubnet.GroupID))
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}
	modelSubnet.GroupName = group.Name

	// Get Leader Node
	if len(modelSubnet.LeaderNodeUUID) != 0 {
		resGetNode, _ := client.RC.GetNode(modelSubnet.LeaderNodeUUID)
		if resGetNode != nil && resGetNode.Node != nil {
			// pxe_boot_ip
			modelSubnet.PXEBootIP = resGetNode.Node.NodeIP
			// leader_node_name
			modelSubnet.LeaderNodeName = resGetNode.Node.NodeName
		}
	}

	// Get Server
	if len(modelSubnet.ServerUUID) != 0 {
		resGetServer, _ := client.RC.GetServer(modelSubnet.ServerUUID)
		if resGetServer != nil && resGetServer.Server != nil {
			// server_name
			modelSubnet.ServerName = resGetServer.Server.ServerName
		}
	}

	return *modelSubnet, nil
}

// ListSubnet : Get subnet list with provided options
func ListSubnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)
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
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqListSubnet pb.ReqGetSubnetList
	var reqSubnet pb.Subnet
	reqListSubnet.Subnet = &reqSubnet

	if uuidOk {
		reqListSubnet.Subnet.UUID = uuid
	}
	if groupIDOk {
		reqListSubnet.Subnet.GroupID = int64(groupID)
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
		return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var numSubnet int
	if rowOk && pageOk {
		reqListSubnet.Row = 0
		reqListSubnet.Page = 0
		resListSubnet2, err := client.RC.GetSubnetList(&reqListSubnet)
		if err != nil {
			return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		numSubnet = len(resListSubnet2.Subnet)
	} else {
		numSubnet = len(resListSubnet.Subnet)
	}

	var subnetList []model.Subnet
	for _, pSubnet := range resListSubnet.Subnet {
		modelSubnet := pbtomodel.PbSubnetToModelSubnet(pSubnet, nil)

		// group_name
		group, err := dao.ReadGroup(int(modelSubnet.GroupID))
		if err != nil {
			return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
		}
		modelSubnet.GroupName = group.Name

		// Get Leader Node
		if len(modelSubnet.LeaderNodeUUID) != 0 {
			resGetNode, _ := client.RC.GetNode(modelSubnet.LeaderNodeUUID)
			if resGetNode != nil && resGetNode.Node != nil {
				// pxe_boot_ip
				modelSubnet.PXEBootIP = resGetNode.Node.NodeIP
				// leader_node_name
				modelSubnet.LeaderNodeName = resGetNode.Node.NodeName
			}
		}

		// Get Server
		if len(modelSubnet.ServerUUID) != 0 {
			resGetServer, _ := client.RC.GetServer(modelSubnet.ServerUUID)
			if resGetServer != nil && resGetServer.Server != nil {
				// server_name
				modelSubnet.ServerName = resGetServer.Server.ServerName
			}
		}

		subnetList = append(subnetList, *modelSubnet)
	}

	hccErrStack := errconv.GrpcStackToHcc(resListSubnet.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.SubnetList{Subnets: subnetList, TotalNum: numSubnet, Errors: Errors}, nil
}

// AllSubnet : Get subnet list with provided options (Just call ListSubnet())
func AllSubnet(args map[string]interface{}) (interface{}, error) {
	return ListSubnet(args)
}

// AvailableSubnetList : Get available subnet list
func AvailableSubnetList(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetAvailableSubnetList pb.ReqGetAvailableSubnetList
	if groupIDOk {
		reqGetAvailableSubnetList.GroupID = int64(groupID)
	}
	resListSubnet, err := client.RC.GetAvailableSubnetList(&reqGetAvailableSubnetList)
	if err != nil {
		return model.SubnetList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var subnetList []model.Subnet
	for _, pSubnet := range resListSubnet.Subnet {
		modelSubnet := pbtomodel.PbSubnetToModelSubnet(pSubnet, nil)
		subnetList = append(subnetList, *modelSubnet)
	}

	hccErrStack := errconv.GrpcStackToHcc(resListSubnet.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.SubnetList{Subnets: subnetList, Errors: Errors}, nil
}

// NumSubnet : Get number of subnets
func NumSubnet(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetSubnetNum pb.ReqGetSubnetNum
	if groupIDOk {
		reqGetSubnetNum.GroupID = int64(groupID)
	}
	resGetSubnetNum, err := client.RC.GetSubnetNum(&reqGetSubnetNum)
	if err != nil {
		return model.Subnet{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var modelSubnetNum model.SubnetNum
	modelSubnetNum.Number = int(resGetSubnetNum.Num)

	hccErrStack := errconv.GrpcStackToHcc(resGetSubnetNum.HccErrorStack)
	modelSubnetNum.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(modelSubnetNum.Errors) != 0 && modelSubnetNum.Errors[0].ErrCode == 0 {
		modelSubnetNum.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelSubnetNum, nil
}

// ValidCheckSubnet : Check if we can create the subnet
func ValidCheckSubnet(args map[string]interface{}) (interface{}, error) {
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	isUpdate, isUpdateOk := args["is_update"].(bool)

	var subnet pb.Subnet
	if networkIPOk {
		subnet.NetworkIP = networkIP
	}
	if netmaskOk {
		subnet.Netmask = netmask
	}
	if gatewayOk {
		subnet.Gateway = gateway
	}

	resValidCheckSubnet, err := client.RC.ValidCheckSubnet(&pb.ReqValidCheckSubnet{
		Subnet:   &subnet,
		IsUpdate: isUpdateOk && isUpdate,
	})
	if err != nil {
		return model.SubnetValid{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var valid = false
	var errMsg = ""
	switch resValidCheckSubnet.ErrorCode {
	case SubnetValid:
		valid = true
	case SubnetValidErrorArgumentError:
		errMsg = "필수 항목들을 모두 입력해 주십시오."
	case SubnetValidErrorStartIPNot1:
		errMsg = "시작 IP 주소가 x.x.x.1 의 형태가 되어야 합니다."
	case SubnetValidErrorInvalidNetworkAddress:
		errMsg = "올바르지 않은 네트워크 주소입니다."
	case SubnetValidErrorInvalidNetmask:
		errMsg = "올바르지 않은 넷마스크 주소 입니다."
	case SubnetValidErrorSubnetConflict:
		errMsg = "생성하려는 서브넷이 다른 기존의 서브넷과 충돌합니다."
	case SubnetValidErrorNotPrivate:
		errMsg = "사설 네트워크 주소를 입력해 주십시오."
	case SubnetValidErrorInvalidGatewayAddress:
		errMsg = "게이트웨이 주소가 올바르지 않습니다."
	case SubnetValidErrorGatewayNotInSubnet:
		errMsg = "게이트웨이 주소가 서브넷에 포함되지 않습니다."
	case SubnetValidErrorSubnetIsUsedByIface:
		errMsg = "해당 서브넷은 내부 네트워크 인터페이스에서 사용중입니다."
	}

	return model.SubnetValid{Valid: valid, ErrMsg: errMsg, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// GetAdaptiveIPAvailableIPList : Get available IP list of AdaptiveIP
func GetAdaptiveIPAvailableIPList() (interface{}, error) {
	resGetAdaptiveIPAvailableIPList, err := client.RC.GetAdaptiveIPAvailableIPList()
	if err != nil {
		return model.AdaptiveIPAvailableIPList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var adaptiveIPAvailableIPList []string

	for _, availableIP := range resGetAdaptiveIPAvailableIPList.AdaptiveipAvailableipList.AvailableIp {
		adaptiveIPAvailableIPList = append(adaptiveIPAvailableIPList, availableIP)
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetAdaptiveIPAvailableIPList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.AdaptiveIPAvailableIPList{
		AvailableIPList: adaptiveIPAvailableIPList,
		Errors:          Errors,
	}, nil
}

// GetAdaptiveIPSetting : Get infos of the adaptiveIP settings
func GetAdaptiveIPSetting() (interface{}, error) {
	resGetAdaptiveIPSetting, err := client.RC.GetAdaptiveIPSetting()
	if err != nil {
		return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	adaptiveipSetting := resGetAdaptiveIPSetting.AdaptiveipSetting
	hccErrStack := errconv.GrpcStackToHcc(resGetAdaptiveIPSetting.HccErrorStack)
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

// AdaptiveIPServer : Get infos of the adaptiveIP server
func AdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	resGetAdaptiveIPServer, err := client.RC.GetAdaptiveIPServer(serverUUID)
	if err != nil {
		return model.AdaptiveIPSetting{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetAdaptiveIPServer.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	var _createdAt time.Time
	if resGetAdaptiveIPServer.AdaptiveipServer.CreatedAt != nil {
		_createdAt, err = ptypes.Timestamp(resGetAdaptiveIPServer.AdaptiveipServer.CreatedAt)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}, nil
		}
	}

	modelAdaptiveIPServer := model.AdaptiveIPServer{
		ServerUUID:     resGetAdaptiveIPServer.AdaptiveipServer.ServerUUID,
		GroupID:        resGetAdaptiveIPServer.AdaptiveipServer.GroupID,
		PublicIP:       resGetAdaptiveIPServer.AdaptiveipServer.PublicIP,
		PrivateIP:      resGetAdaptiveIPServer.AdaptiveipServer.PrivateIP,
		PrivateGateway: resGetAdaptiveIPServer.AdaptiveipServer.PrivateGateway,
		CreatedAt:      _createdAt,
		Errors:         Errors,
	}

	// group_name
	group, err := dao.ReadGroup(int(modelAdaptiveIPServer.GroupID))
	if err != nil {
		return model.AdaptiveIPServer{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}
	modelAdaptiveIPServer.GroupName = group.Name

	return modelAdaptiveIPServer, nil
}

// ListAdaptiveIPServer : Get adaptiveIP server list with provided options
func ListAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	groupID, groupIDOk := args["group_id"].(int)
	publicIP, publicIPOk := args["public_ip"].(string)
	privateIP, privateIPOk := args["private_ip"].(string)
	privateGateway, privateGatewayOk := args["private_gateway"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	userUUID, userUUIDOk := args["user_uuid"].(string)

	var reqGetAdaptiveIPServerList pb.ReqGetAdaptiveIPServerList
	var reqAdaptiveIPServerList pb.AdaptiveIPServer
	reqGetAdaptiveIPServerList.AdaptiveipServer = &reqAdaptiveIPServerList

	if serverUUIDOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.ServerUUID = serverUUID
	}
	if groupIDOk {
		reqGetAdaptiveIPServerList.AdaptiveipServer.GroupID = int64(groupID)
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

	resAdaptiveIPServerList, err := client.RC.GetAdaptiveIPServerList(&reqGetAdaptiveIPServerList)
	if err != nil {
		return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var numAdaptiveIPServer int
	if rowOk && pageOk {
		reqGetAdaptiveIPServerList.Row = 0
		reqGetAdaptiveIPServerList.Page = 0
		resAdaptiveIPServerList2, err := client.RC.GetAdaptiveIPServerList(&reqGetAdaptiveIPServerList)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		numAdaptiveIPServer = len(resAdaptiveIPServerList2.AdaptiveipServer)
	} else {
		numAdaptiveIPServer = len(resAdaptiveIPServerList.AdaptiveipServer)
	}

	var adaptiveIPServerList []model.AdaptiveIPServer
	for _, adaptiveIPServer := range resAdaptiveIPServerList.AdaptiveipServer {
		if userUUIDOk {
			queryArgs := make(map[string]interface{})
			queryArgs["uuid"] = adaptiveIPServer.ServerUUID
			server, err := Server(queryArgs)
			if err != nil {
				logger.Logger.Println("ListAdaptiveIPServer(): Failed to get server information for server_uuid = " + adaptiveIPServer.ServerUUID)
				continue
			}
			if server.(model.Server).UserUUID != userUUID {
				continue
			}
		}

		_createdAt, err := ptypes.Timestamp(adaptiveIPServer.CreatedAt)
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTimestampConversionError, err.Error())}, nil
		}

		modelAdaptiveIPServer := model.AdaptiveIPServer{
			ServerUUID:     adaptiveIPServer.ServerUUID,
			GroupID:        adaptiveIPServer.GroupID,
			PublicIP:       adaptiveIPServer.PublicIP,
			PrivateIP:      adaptiveIPServer.PrivateIP,
			PrivateGateway: adaptiveIPServer.PrivateGateway,
			CreatedAt:      _createdAt,
		}

		// group_name
		group, err := dao.ReadGroup(int(modelAdaptiveIPServer.GroupID))
		if err != nil {
			return model.AdaptiveIPServerList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
		}
		modelAdaptiveIPServer.GroupName = group.Name

		adaptiveIPServerList = append(adaptiveIPServerList, modelAdaptiveIPServer)
	}

	hccErrStack := errconv.GrpcStackToHcc(resAdaptiveIPServerList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.AdaptiveIPServerList{AdaptiveIPServers: adaptiveIPServerList, TotalNum: numAdaptiveIPServer, Errors: Errors}, nil
}

// AllAdaptiveIPServer : Get adaptiveIP server list with provided options (Just call ListAdaptiveIPServer())
func AllAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	return ListAdaptiveIPServer(args)
}

// NumAdaptiveIPServer : Get number of adaptiveIP servers
func NumAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetAdaptiveIPServerNum pb.ReqGetAdaptiveIPServerNum
	if groupIDOk {
		reqGetAdaptiveIPServerNum.GroupID = int64(groupID)
	}
	resGetAdaptiveIPServerNum, err := client.RC.GetAdaptiveIPServerNum(&reqGetAdaptiveIPServerNum)
	if err != nil {
		return model.AdaptiveIPServerNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetAdaptiveIPServerNum.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.AdaptiveIPServerNum{Number: int(resGetAdaptiveIPServerNum.Num), Errors: Errors}, nil
}

// ListPortForwarding : Get AdaptiveIP Port Forwarding list with provided options
func ListPortForwarding(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	protocol, protocolOk := args["protocol"].(string)
	externalPort, externalPortOk := args["external_port"].(int)
	internalPort, internalPortOk := args["internal_port"].(int)
	description, descriptionOk := args["description"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	var reqGetPortForwardingList pb.ReqGetPortForwardingList
	var reqPortForwardingList pb.PortForwarding
	reqGetPortForwardingList.PortForwarding = &reqPortForwardingList

	if !serverUUIDOk {
		return model.PortForwardingList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"Need a server_uuid argument")}, nil
	}
	reqGetPortForwardingList.PortForwarding.ServerUUID = serverUUID

	if protocolOk {
		protocol = strings.ToLower(protocol)
		if protocol == "tcp" {
			reqGetPortForwardingList.PortForwarding.ForwardTCP = true
		} else if protocol == "udp" {
			reqGetPortForwardingList.PortForwarding.ForwardUDP = true
		} else if protocol == "all" {
			reqGetPortForwardingList.PortForwarding.ForwardTCP = true
			reqGetPortForwardingList.PortForwarding.ForwardUDP = true
		} else {
			reqGetPortForwardingList.PortForwarding.ForwardTCP = false
			reqGetPortForwardingList.PortForwarding.ForwardUDP = false
		}
	}
	if externalPortOk {
		reqGetPortForwardingList.PortForwarding.ExternalPort = int64(externalPort)
	}
	if internalPortOk {
		reqGetPortForwardingList.PortForwarding.InternalPort = int64(internalPort)
	}
	if descriptionOk {
		reqGetPortForwardingList.PortForwarding.Description = description
	}
	if rowOk {
		reqGetPortForwardingList.Row = int64(row)
	}
	if pageOk {
		reqGetPortForwardingList.Page = int64(page)
	}

	resGetPortForwardingList, err := client.RC.GetPortForwardingList(&reqGetPortForwardingList)
	if err != nil {
		return model.PortForwardingList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	var numPortForwarding int
	if rowOk && pageOk {
		reqGetPortForwardingList.Row = 0
		reqGetPortForwardingList.Page = 0
		resGetPortForwardingList2, err := client.RC.GetPortForwardingList(&reqGetPortForwardingList)
		if err != nil {
			return model.PortForwardingList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
		}
		numPortForwarding = len(resGetPortForwardingList2.PortForwarding)
	} else {
		numPortForwarding = len(resGetPortForwardingList.PortForwarding)
	}

	var portForwardingList []model.PortForwarding
	for _, portForwarding := range resGetPortForwardingList.PortForwarding {
		var protocol string

		if portForwarding.ForwardTCP && portForwarding.ForwardUDP {
			protocol = "all"
		} else if portForwarding.ForwardTCP {
			protocol = "tcp"
		} else if portForwarding.ForwardUDP {
			protocol = "udp"
		}

		modelPortForwarding := model.PortForwarding{
			ServerUUID:   portForwarding.ServerUUID,
			Protocol:     protocol,
			ExternalPort: portForwarding.ExternalPort,
			InternalPort: portForwarding.InternalPort,
			Description:  portForwarding.Description,
		}

		portForwardingList = append(portForwardingList, modelPortForwarding)
	}

	hccErrStack := errconv.GrpcStackToHcc(resGetPortForwardingList.HccErrorStack)
	Errors := errconv.HccErrorToPiccoloHccErr(*hccErrStack)
	if len(Errors) != 0 && Errors[0].ErrCode == 0 {
		Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return model.PortForwardingList{PortForwardings: portForwardingList, TotalNum: numPortForwarding, Errors: Errors}, nil
}
