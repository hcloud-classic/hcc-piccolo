package client

import (
	"context"
	"google.golang.org/grpc"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"strconv"
	"time"
)

var violinConn *grpc.ClientConn

func initViolin() error {
	var err error

	addr := config.Violin.ServerAddress + ":" + strconv.FormatInt(config.Violin.ServerPort, 10)
	violinConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.violin = rpcviolin.NewViolinClient(violinConn)
	logger.Logger.Println("gRPC violin client ready")

	return nil
}

func closeViolin() {
	_ = violinConn.Close()
}

// CreateServer : Create a server
func (rc *RPCClient) CreateServer(in *rpcviolin.ReqCreateServer) (*rpcviolin.ResCreateServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateServer, err := rc.violin.CreateServer(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateServer, nil
}

// GetServer : Get infos of the server
func (rc *RPCClient) GetServer(uuid string) (*rpcviolin.Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	server, err := rc.violin.GetServer(ctx, &rpcviolin.ReqGetServer{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return server.Server, nil
}

// GetServerList : Get list of the server
func (rc *RPCClient) GetServerList(in *rpcviolin.ReqGetServerList) (*rpcviolin.ResGetServerList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	serverList, err := rc.violin.GetServerList(ctx, in)
	if err != nil {
		return nil, err
	}

	return serverList, nil
}

// GetServerNum : Get the number of servers
func (rc *RPCClient) GetServerNum() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	num, err := rc.violin.GetServerNum(ctx, &rpcviolin.Empty{})
	if err != nil {
		return 0, err
	}

	return int(num.Num), nil
}

// UpdateServer : Update infos of the server
func (rc *RPCClient) UpdateServer(in *rpcviolin.ReqUpdateServer) (*rpcviolin.ResUpdateServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resUpdateServer, err := rc.violin.UpdateServer(ctx, in)
	if err != nil {
		return nil, err
	}

	return resUpdateServer, nil
}

// DeleteServer : Delete of the server
func (rc *RPCClient) DeleteServer(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteServer, err := rc.violin.DeleteServer(ctx, &rpcviolin.ReqDeleteServer{UUID: uuid})
	if err != nil {
		return "", err
	}

	return resDeleteServer.UUID, nil
}

// CreateServerNode : Create a server node
func (rc *RPCClient) CreateServerNode(in *rpcviolin.ReqCreateServerNode) (*rpcviolin.ResCreateServerNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateServerNode, err := rc.violin.CreateServerNode(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateServerNode, nil
}

// GetServerNode : Get infos of the server
func (rc *RPCClient) GetServerNode(uuid string) (*rpcviolin.ServerNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	server, err := rc.violin.GetServerNode(ctx, &rpcviolin.ReqGetServerNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return server.ServerNode, nil
}

// GetServerNodeList : Get list of the server
func (rc *RPCClient) GetServerNodeList(in *rpcviolin.ReqGetServerNodeList) (*rpcviolin.ResGetServerNodeList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	serverList, err := rc.violin.GetServerNodeList(ctx, in)
	if err != nil {
		return nil, err
	}

	return serverList, nil
}

// GetServerNodeNum : Get the number of servers
func (rc *RPCClient) GetServerNodeNum(serverUUID string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	num, err := rc.violin.GetServerNodeNum(ctx, &rpcviolin.ReqGetServerNodeNum{ServerUUID: serverUUID})
	if err != nil {
		return 0, err
	}

	return int(num.Num), nil
}

// DeleteServerNode : Delete of the serverNode
func (rc *RPCClient) DeleteServerNode(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteServerNode, err := rc.violin.DeleteServerNode(ctx, &rpcviolin.ReqDeleteServerNode{UUID: uuid})
	if err != nil {
		return "", err
	}

	return resDeleteServerNode.UUID, nil
}

// DeleteServerNodeByServerUUID : Delete of the server
func (rc *RPCClient) DeleteServerNodeByServerUUID(serverUUID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Violin.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteServerNodeByServerUUID, err := rc.violin.DeleteServerNodeByServerUUID(ctx, &rpcviolin.ReqDeleteServerNodeByServerUUID{ServerUUID: serverUUID})
	if err != nil {
		return "", err
	}

	return resDeleteServerNodeByServerUUID.ServerUUID, nil
}
