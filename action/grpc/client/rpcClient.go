package client

import (
	"hcc/piccolo/action/grpc/pb/rpccello"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/action/grpc/pb/rpcharp"
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/action/grpc/pb/rpcviolin"
	rpcnovnc "hcc/piccolo/action/grpc/pb/rpcviolin_novnc"
)

// RPCClient : Struct type of gRPC clients
type RPCClient struct {
	flute  rpcflute.FluteClient
	harp   rpcharp.HarpClient
	violin rpcviolin.ViolinClient
	novnc  rpcnovnc.NovncClient
	piano  rpcpiano.PianoClient
	cello  rpccello.CelloClient
}

// RC : Exported variable pointed to RPCClient
var RC = &RPCClient{}

// Init : Initialize clients of gRPC
func Init() error {
	err := initFlute()
	if err != nil {
		return err
	}

	err = initHarp()
	if err != nil {
		return err
	}

	err = initViolin()
	if err != nil {
		return err
	}

	err = initNovnc()
	if err != nil {
		return err
	}

	err = initPiano()
	if err != nil {
		return err
	}
	err = initCello()
	if err != nil {
		return err
	}

	return nil
}

// End : Close connections of gRPC clients
func End() {
	closePiano()
	closeNovnc()
	closeViolin()
	closeHarp()
	closeFlute()
	closeCello()
}
