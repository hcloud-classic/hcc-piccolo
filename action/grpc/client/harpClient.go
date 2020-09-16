package client

import (
	"context"
	"google.golang.org/grpc"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"strconv"
	"time"
)

var harpConn *grpc.ClientConn

func initHarp() error {
	var err error

	addr := config.Harp.ServerAddress + ":" + strconv.FormatInt(config.Harp.ServerPort, 10)
	harpConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.harp = rpcharp.NewHarpClient(harpConn)
	logger.Logger.Println("gRPC harp client ready")

	return nil
}

func closeHarp() {
	_ = harpConn.Close()
}

// CreateSubnet : Create a subnet
func (rc *RPCClient) CreateSubnet(in *rpcharp.ReqCreateSubnet) (*rpcharp.ResCreateSubnet, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateSubnet, err := rc.harp.CreateSubnet(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateSubnet, nil
}

// GetSubnet : Get infos of the subnet
func (rc *RPCClient) GetSubnet(uuid string) (*rpcharp.ResGetSubnet, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetSubnet, err := rc.harp.GetSubnet(ctx, &rpcharp.ReqGetSubnet{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resGetSubnet, nil
}

// GetSubnetList : Get list of the subnet
func (rc *RPCClient) GetSubnetList(in *rpcharp.ReqGetSubnetList) (*rpcharp.ResGetSubnetList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	subnetList, err := rc.harp.GetSubnetList(ctx, in)
	if err != nil {
		return nil, err
	}

	return subnetList, nil
}

// GetSubnetNum : Get the number of subnets
func (rc *RPCClient) GetSubnetNum() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	num, err := rc.harp.GetSubnetNum(ctx, &rpcharp.Empty{})
	if err != nil {
		return 0, err
	}

	return int(num.Num), nil
}

// UpdateSubnet : Update infos of the subnet
func (rc *RPCClient) UpdateSubnet(in *rpcharp.ReqUpdateSubnet) (*rpcharp.ResUpdateSubnet, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resUpdateSubnet, err := rc.harp.UpdateSubnet(ctx, in)
	if err != nil {
		return nil, err
	}

	return resUpdateSubnet, nil
}

// DeleteSubnet : Delete of the subnet
func (rc *RPCClient) DeleteSubnet(uuid string) (*rpcharp.ResDeleteSubnet, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteSubnet, err := rc.harp.DeleteSubnet(ctx, &rpcharp.ReqDeleteSubnet{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resDeleteSubnet, nil
}

// CreateAdaptiveIPSetting : Create settings of AdaptiveIP
func (rc *RPCClient) CreateAdaptiveIPSetting(in *rpcharp.ReqCreateAdaptiveIPSetting) (*rpcharp.ResCreateAdaptiveIPSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateAdaptiveIPSetting, err := rc.harp.CreateAdaptiveIPSetting(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateAdaptiveIPSetting, nil
}

// GetAdaptiveIPAvailableIPList : Get available IP list of AdaptiveIP
func (rc *RPCClient) GetAdaptiveIPAvailableIPList() (*rpcharp.ResGetAdaptiveIPAvailableIPList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetAdaptiveIPAvailableIPList, err := rc.harp.GetAdaptiveIPAvailableIPList(ctx, &rpcharp.Empty{})
	if err != nil {
		return nil, err
	}

	return resGetAdaptiveIPAvailableIPList, nil
}

// GetAdaptiveIPSetting : Get settings of AdaptiveIP
func (rc *RPCClient) GetAdaptiveIPSetting() (*rpcharp.ResGetAdaptiveIPSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resAdaptiveIPSetting, err := rc.harp.GetAdaptiveIPSetting(ctx, &rpcharp.Empty{})
	if err != nil {
		return nil, err
	}

	return resAdaptiveIPSetting, nil
}

// CreateAdaptiveIPServer : Create AdaptiveIP server
func (rc *RPCClient) CreateAdaptiveIPServer(in *rpcharp.ReqCreateAdaptiveIPServer) (*rpcharp.ResCreateAdaptiveIPServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateAdaptiveIPServer, err := rc.harp.CreateAdaptiveIPServer(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateAdaptiveIPServer, nil
}

// GetAdaptiveIPServer : Get infos of the adaptiveIP server
func (rc *RPCClient) GetAdaptiveIPServer(serverUUID string) (*rpcharp.ResGetAdaptiveIPServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetAdaptiveIPServer, err := rc.harp.GetAdaptiveIPServer(ctx, &rpcharp.ReqGetAdaptiveIPServer{ServerUUID: serverUUID})
	if err != nil {
		return nil, err
	}

	return resGetAdaptiveIPServer, nil
}

// GetAdaptiveIPServerList : Get list of the adaptiveIP server
func (rc *RPCClient) GetAdaptiveIPServerList(in *rpcharp.ReqGetAdaptiveIPServerList) (*rpcharp.ResGetAdaptiveIPServerList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	adaptiveIPServerList, err := rc.harp.GetAdaptiveIPServerList(ctx, in)
	if err != nil {
		return nil, err
	}

	return adaptiveIPServerList, nil
}

// GetAdaptiveIPServerNum : Get the number of adaptiveIP server
func (rc *RPCClient) GetAdaptiveIPServerNum() (*rpcharp.ResGetAdaptiveIPServerNum, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetAdaptiveIPServerNum, err := rc.harp.GetAdaptiveIPServerNum(ctx, &rpcharp.Empty{})
	if err != nil {
		return nil, err
	}

	return resGetAdaptiveIPServerNum, nil
}

// DeleteAdaptiveIPServer : Delete of the adaptiveIP server
func (rc *RPCClient) DeleteAdaptiveIPServer(serverUUID string) (*rpcharp.ResDeleteAdaptiveIPServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteAdaptiveIPServer, err := rc.harp.DeleteAdaptiveIPServer(ctx, &rpcharp.ReqDeleteAdaptiveIPServer{ServerUUID: serverUUID})
	if err != nil {
		return nil, err
	}

	return resDeleteAdaptiveIPServer, nil
}

// CreateDHCPDConfig : Do dhcpd config file creation works
func (rc *RPCClient) CreateDHCPDConfig(subnetUUID string, nodeUUIDs string) (*rpcharp.ResCreateDHPCDConf, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Harp.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateDHPCDConf, err := rc.harp.CreateDHPCDConf(ctx, &rpcharp.ReqCreateDHPCDConf{
		SubnetUUID: subnetUUID,
		NodeUUIDs:  nodeUUIDs,
	})
	if err != nil {
		return nil, err
	}

	return resCreateDHPCDConf, nil
}
