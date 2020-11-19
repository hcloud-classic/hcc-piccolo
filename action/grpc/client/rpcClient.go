package client

import (
	"github.com/hcloud-classic/pb"
)

// RPCClient : Struct type of gRPC clients
type RPCClient struct {
	flute  pb.FluteClient
	harp   pb.HarpClient
	violin pb.ViolinClient
	novnc  pb.NovncClient
	piano  pb.PianoClient
	cello  pb.CelloClient
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
