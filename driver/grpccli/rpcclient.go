package grpccli

import (
	rpcnovnc "hcc/piccolo/action/grpc/rpcviolin_novnc"
)

type RpcClient struct {
	novnc rpcnovnc.NovncClient
}

var RC = &RpcClient{}

func InitGRPCClient() {
	go initNovnc()
}

func ClenGRPCClient() {
	cleanNovnc()
}
