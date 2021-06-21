package client

import (
	"context"
	"strconv"
	"time"

	"google.golang.org/grpc"

	rpcnovnc "hcc/piccolo/action/grpc/pb/rpcviolin_novnc"
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

	RC.novnc = rpcnovnc.NewNovncClient(novncConn)
	logger.Logger.Println("gRPC novnc client ready")

	return nil
}

func closeNovnc() {
	_ = novncConn.Close()
}

// ControlVNC : Set VNC with provided options
func (rc *RPCClient) ControlVNC(reqData map[string]interface{}) (interface{}, error) {
	//req data mapping
	var req rpcnovnc.ReqControlVNC
	req.Vnc = &rpcnovnc.VNC{
		Token:      reqData["token"].(string),
		ServerUUID: reqData["server_uuid"].(string),
		Action:     reqData["action"].(string),
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.ViolinNoVnc.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	r, err := rc.novnc.ControlVNC(ctx, &req)
	if err != nil {
		return nil, err
	}
	logger.Logger.Println(r)

	return r, nil
}
