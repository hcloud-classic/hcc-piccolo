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
func (rc *RPCClient) OnNode(nodeUUID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	node := pb.Node{
		UUID: nodeUUID,
	}
	nodes = append(nodes, &node)

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &rpcflute.ReqNodePowerControl{
		Node:       nodes,
		PowerState: rpcflute.PowerState_ON,
	})
	if err != nil {
		return "", err
	}

	return resNodePowerControl.Result[0], nil
}

// OffNode : Turn off selected node
func (rc *RPCClient) OffNode(nodeUUID string, forceOff bool) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	node := pb.Node{
		UUID: nodeUUID,
	}
	nodes = append(nodes, &node)

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
		return "", err
	}

	return resNodePowerControl.Result[0], nil
}

// ForceRestartNode : Force restart selected node
func (rc *RPCClient) ForceRestartNode(nodeUUID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	var nodes []*pb.Node
	node := pb.Node{
		UUID: nodeUUID,
	}
	nodes = append(nodes, &node)

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &rpcflute.ReqNodePowerControl{
		Node:       nodes,
		PowerState: rpcflute.PowerState_FORCE_RESTART,
	})
	if err != nil {
		return "", err
	}

	return resNodePowerControl.Result[0], nil
}

// GetNodePowerState : Get power state of selected node
func (rc *RPCClient) GetNodePowerState(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	resNodePowerState, err := rc.flute.GetNodePowerState(ctx, &rpcflute.ReqNodePowerState{
		UUID: uuid,
	})
	if err != nil {
		return "", err
	}

	return resNodePowerState.Result, nil
}

// CreateNodeDetail: Create a nodeDetail
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
func (rc *RPCClient) GetNodeDetail(uuid string) (*rpcflute.NodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	nodeDetail, err := rc.flute.GetNodeDetail(ctx, &rpcflute.ReqGetNodeDetail{NodeUUID: uuid})
	if err != nil {
		return nil, err
	}

	return nodeDetail.NodeDetail, nil
}

// DeleteNodeDetail : Delete of the nodeDetail
func (rc *RPCClient) DeleteNodeDetail(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNodeDetail, err := rc.flute.DeleteNodeDetail(ctx, &rpcflute.ReqDeleteNodeDetail{NodeUUID: uuid})
	if err != nil {
		return "", err
	}

	return resDeleteNodeDetail.NodeUUID, nil
}

// CreateNode: Create a node
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
func (rc *RPCClient) GetNode(uuid string) (*rpcflute.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	node, err := rc.flute.GetNode(ctx, &rpcflute.ReqGetNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return node.Node, nil
}

// GetNodeList : Get the list of nodes by server UUID.
func (rc *RPCClient) GetNodeList(in *rpcflute.ReqGetNodeList) (*rpcflute.ResGetNodeList, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	nodeList, err := rc.flute.GetNodeList(ctx, in)
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

// GetNodeNum : Get the number of nodes
func (rc *RPCClient) GetNodeNum() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	num, err := rc.flute.GetNodeNum(ctx, &rpcflute.Empty{})
	if err != nil {
		return 0, err
	}

	return int(num.Num), nil
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
func (rc *RPCClient) DeleteNode(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNode, err := rc.flute.DeleteNode(ctx, &rpcflute.ReqDeleteNode{UUID: uuid})
	if err != nil {
		return "", err
	}

	return resDeleteNode.UUID, nil
}