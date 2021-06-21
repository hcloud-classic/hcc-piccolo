package client

import (
	"context"
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

var celloConn *grpc.ClientConn

func initCello() error {
	var err error

	addr := config.Cello.ServerAddress + ":" + strconv.FormatInt(config.Cello.ServerPort, 10)
	celloConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.cello = rpccello.NewCelloClient(celloConn)
	logger.Logger.Println("gRPC violin client ready")

	return nil
}

func closeCello() {
	_ = celloConn.Close()
}

// VolumeHandler : VolumeHandler
func (rc *RPCClient) VolumeHandler(in *rpccello.ReqVolumeHandler) (*rpccello.ResVolumeHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Cello.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resVolumeHandle, err := rc.cello.VolumeHandler(ctx, in)
	if err != nil {
		return nil, err
	}

	return resVolumeHandle, nil
}

// PoolHandler : PoolHandler
func (rc *RPCClient) PoolHandler(in *rpccello.ReqPoolHandler) (*rpccello.ResPoolHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Cello.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resPoolhandler, err := rc.cello.PoolHandler(ctx, in)
	if err != nil {
		return nil, err
	}

	return resPoolhandler, nil
}

// GetPoolList : GetPoolList
func (rc *RPCClient) GetPoolList(in *rpccello.ReqGetPoolList) (*rpccello.ResGetPoolList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Cello.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resPoolList, err := rc.cello.GetPoolList(ctx, in)
	if err != nil {
		return nil, err
	}

	return resPoolList, nil
}

// GetVolumeList : GetVolumeList
func (rc *RPCClient) GetVolumeList(in *rpccello.ReqGetVolumeList) (*rpccello.ResGetVolumeList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Cello.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetVolumeList, err := rc.cello.GetVolumeList(ctx, in)
	if err != nil {
		return nil, err
	}

	return resGetVolumeList, nil
}
