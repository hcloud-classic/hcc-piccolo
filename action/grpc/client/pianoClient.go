package client

import (
	"context"
	"google.golang.org/grpc"
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"strconv"
	"time"
)

var pianoConn *grpc.ClientConn

func initPiano() error {
	var err error

	addr := config.Piano.ServerAddress + ":" + strconv.FormatInt(config.Piano.ServerPort, 10)
	pianoConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.piano = rpcpiano.NewPianoClient(pianoConn)
	logger.Logger.Println("gRPC piano client ready")

	return nil
}

func closePiano() {
	_ = pianoConn.Close()
}

// Telegraph : Get the metric data from influxDB
func (rc *RPCClient) Telegraph(in *rpcpiano.ReqMetricInfo) (*rpcpiano.ResMonitoringData, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Piano.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resMonitoringData, err := rc.piano.Telegraph(ctx, in)
	if err != nil {
		return nil, err
	}

	return resMonitoringData, nil
}
