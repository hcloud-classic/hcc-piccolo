package client

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"

	"innogrid.com/hcloud-classic/pb"
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

func pingNovnc() bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(config.ViolinNoVnc.ServerAddress,
		strconv.FormatInt(config.ViolinNoVnc.ServerPort, 10)),
		time.Duration(config.Grpc.ClientPingTimeoutMs)*time.Millisecond)
	if err != nil {
		return false
	}
	if conn != nil {
		defer func() {
			_ = conn.Close()
		}()
		return true
	}

	return false
}

func checkNovnc() {
	ticker := time.NewTicker(time.Duration(config.Grpc.ClientPingIntervalMs) * time.Millisecond)
	go func() {
		connOk := true
		for range ticker.C {
			pingOk := pingNovnc()
			if pingOk {
				if !connOk {
					logger.Logger.Println("checkNovnc(): Ping Ok! Resetting connection...")
					closeNovnc()
					err := initNovnc()
					if err != nil {
						logger.Logger.Println("checkNovnc(): " + err.Error())
						continue
					}
					connOk = true
				}
			} else {
				if connOk {
					logger.Logger.Println("checkNovnc(): Novnc module seems dead. Pinging...")
				}
				connOk = false
			}
		}
	}()
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
