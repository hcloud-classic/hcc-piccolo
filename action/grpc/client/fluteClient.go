package client

import (
	"context"
	"google.golang.org/grpc"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	pb "hcc/piccolo/action/grpc/pb/rpcviolin"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"strconv"
	"time"
)

var fluteConn *grpc.ClientConn

func initFlute() error {
	var err error

	addr := config.Flute.ServerAddress + ":" + strconv.FormatInt(config.Flute.ServerPort, 10)
	fluteConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.flute = rpcflute.NewFluteClient(fluteConn)
	logger.Logger.Println("gRPC flute client ready")

	return nil
}

func closeFlute() {
	_ = fluteConn.Close()
}

// OnNode : Turn on selected node
func (rc *RPCClient) OnNode(nodeUUIDs []string) (*rpcflute.ResNodePowerControl, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	for i := range nodeUUIDs {
		node := pb.Node{
			UUID: nodeUUIDs[i],
		}
		nodes = append(nodes, &node)
	}

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &rpcflute.ReqNodePowerControl{
		Node:       nodes,
		PowerState: rpcflute.PowerState_ON,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// OffNode : Turn off selected node
func (rc *RPCClient) OffNode(nodeUUIDs []string, forceOff bool) (*rpcflute.ResNodePowerControl, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	for i := range nodeUUIDs {
		node := pb.Node{
			UUID: nodeUUIDs[i],
		}
		nodes = append(nodes, &node)
	}

	var powerState rpcflute.PowerState
	if forceOff {
		powerState = rpcflute.PowerState_FORCE_OFF
	} else {
		powerState = rpcflute.PowerState_OFF
	}

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &rpcflute.ReqNodePowerControl{
		Node:       nodes,
		PowerState: powerState,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// ForceRestartNode : Force restart selected node
func (rc *RPCClient) ForceRestartNode(nodeUUIDs []string) (*rpcflute.ResNodePowerControl, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	for i := range nodeUUIDs {
		node := pb.Node{
			UUID: nodeUUIDs[i],
		}
		nodes = append(nodes, &node)
	}

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &rpcflute.ReqNodePowerControl{
		Node:       nodes,
		PowerState: rpcflute.PowerState_FORCE_RESTART,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// GetNodePowerState : Get power state of selected node
func (rc *RPCClient) GetNodePowerState(uuid string) (*rpcflute.ResNodePowerState, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	resNodePowerState, err := rc.flute.GetNodePowerState(ctx, &rpcflute.ReqNodePowerState{
		UUID: uuid,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerState, nil
}

// CreateNodeDetail : Create a nodeDetail
func (rc *RPCClient) CreateNodeDetail(in *rpcflute.ReqCreateNodeDetail) (*rpcflute.ResCreateNodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateNodeDetail, err := rc.flute.CreateNodeDetail(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateNodeDetail, nil
}

// GetNodeDetail : Get infos of the nodeDetail
func (rc *RPCClient) GetNodeDetail(uuid string) (*rpcflute.ResGetNodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNodeDetail, err := rc.flute.GetNodeDetail(ctx, &rpcflute.ReqGetNodeDetail{NodeUUID: uuid})
	if err != nil {
		return nil, err
	}

	return resGetNodeDetail, nil
}

// DeleteNodeDetail : Delete of the nodeDetail
func (rc *RPCClient) DeleteNodeDetail(uuid string) (*rpcflute.ResDeleteNodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNodeDetail, err := rc.flute.DeleteNodeDetail(ctx, &rpcflute.ReqDeleteNodeDetail{NodeUUID: uuid})
	if err != nil {
		return nil, err
	}

	return resDeleteNodeDetail, nil
}

// CreateNode : Create a node
func (rc *RPCClient) CreateNode(in *rpcflute.ReqCreateNode) (*rpcflute.ResCreateNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resCreateNode, err := rc.flute.CreateNode(ctx, in)
	if err != nil {
		return nil, err
	}

	return resCreateNode, nil
}

// GetNode : Get infos of the node
func (rc *RPCClient) GetNode(uuid string) (*rpcflute.ResGetNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNode, err := rc.flute.GetNode(ctx, &rpcflute.ReqGetNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resGetNode, nil
}

// GetNodeList : Get the list of nodes
func (rc *RPCClient) GetNodeList(in *rpcflute.ReqGetNodeList) (*rpcflute.ResGetNodeList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNodeList, err := rc.flute.GetNodeList(ctx, in)
	if err != nil {
		return nil, err
	}

	return resGetNodeList, nil
}

// GetNodeNum : Get the number of nodes
func (rc *RPCClient) GetNodeNum() (*rpcflute.ResGetNodeNum, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNodeNum, err := rc.flute.GetNodeNum(ctx, &rpcflute.Empty{})
	if err != nil {
		return nil, err
	}

	return resGetNodeNum, nil
}

// UpdateNode : Update infos of the node
func (rc *RPCClient) UpdateNode(in *rpcflute.ReqUpdateNode) (*rpcflute.ResUpdateNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resUpdateNode, err := rc.flute.UpdateNode(ctx, in)
	if err != nil {
		return nil, err
	}

	return resUpdateNode, nil
}

// DeleteNode : Delete of the node
func (rc *RPCClient) DeleteNode(uuid string) (*rpcflute.ResDeleteNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNode, err := rc.flute.DeleteNode(ctx, &rpcflute.ReqDeleteNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resDeleteNode, nil
}
