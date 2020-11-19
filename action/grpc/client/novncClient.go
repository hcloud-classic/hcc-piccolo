package client

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/hcloud-classic/pb"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
)

var novncConn *grpc.ClientConn

func initNovnc() error {
	var err error

	addr := config.ViolinNoVnc.ServerAddress + ":" + strconv.FormatInt(config.ViolinNoVnc.ServerPort, 10)
	novncConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.novnc = pb.NewNovncClient(novncConn)
	logger.Logger.Println("gRPC novnc client ready")

	return nil
}

func closeNovnc() {
	_ = novncConn.Close()
}

// ControlVNC : Set VNC with provided options
func (rc *RPCClient) ControlVNC(in *pb.ReqControlVNC) (*pb.ResControlVNC, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.ViolinNoVnc.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	fmt.Println(in)
	resControlVNC, err := rc.novnc.ControlVNC(ctx, in)
	if err != nil {
		return nil, err
	}

	return resControlVNC, nil
}
