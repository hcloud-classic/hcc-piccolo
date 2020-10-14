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

// CreateVolume : Create a server
func (rc *RPCClient) CreateVolume(in *rpccello.ReqVolumeHandler) (*rpccello.ResVolumeHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Cello.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateVolume, err := rc.cello.VolumeHandler(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateVolume, nil
}
