package client

import (
	"context"
	"errors"
	"hcc/piccolo/lib/config"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"innogrid.com/hcloud-classic/pb"
)

// InitTuba : Initialize a gRPC connection to tuba module
func InitTuba(serverAddress string, serverPort int, conn *grpc.ClientConn) (pb.TubaClient, error) {
	var err error

	addr := serverAddress + ":" + strconv.Itoa(serverPort)
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	tubaClient := pb.NewTubaClient(conn)

	return tubaClient, nil
}

// CloseTuba : Close the gRPC connection of tuba module
func CloseTuba(conn *grpc.ClientConn) {
	if conn != nil {
		_ = conn.Close()
	}
}

// GetTaskList : Get the list of tasks
func GetTaskList(tubaClient pb.TubaClient) (*pb.ResGetTaskList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Tuba.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	if tubaClient == nil {
		return nil, errors.New("tubaClient is nil")
	}
	taskListResult, err := tubaClient.GetTaskList(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	return taskListResult, nil
}
