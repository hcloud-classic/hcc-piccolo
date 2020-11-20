package client

import (
	"context"
	"github.com/hcloud-classic/pb"
	"google.golang.org/grpc"
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

	RC.flute = pb.NewFluteClient(fluteConn)
	logger.Logger.Println("gRPC flute client ready")

	return nil
}

func closeFlute() {
	_ = fluteConn.Close()
}

// OnNode : Turn on selected node
func (rc *RPCClient) OnNode(nodeUUIDs []string) (*pb.ResNodePowerControl, error) {
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

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &pb.ReqNodePowerControl{
		Node:       nodes,
		PowerState: pb.PowerState_ON,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// OffNode : Turn off selected node
func (rc *RPCClient) OffNode(nodeUUIDs []string, forceOff bool) (*pb.ResNodePowerControl, error) {
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

	var powerState pb.PowerState
	if forceOff {
		powerState = pb.PowerState_FORCE_OFF
	} else {
		powerState = pb.PowerState_OFF
	}

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &pb.ReqNodePowerControl{
		Node:       nodes,
		PowerState: powerState,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// ForceRestartNode : Force restart selected node
func (rc *RPCClient) ForceRestartNode(nodeUUIDs []string) (*pb.ResNodePowerControl, error) {
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

	resNodePowerControl, err := rc.flute.NodePowerControl(ctx, &pb.ReqNodePowerControl{
		Node:       nodes,
		PowerState: pb.PowerState_FORCE_RESTART,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerControl, nil
}

// GetNodePowerState : Get power state of selected node
func (rc *RPCClient) GetNodePowerState(uuid string) (*pb.ResNodePowerState, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()

	resNodePowerState, err := rc.flute.GetNodePowerState(ctx, &pb.ReqNodePowerState{
		UUID: uuid,
	})
	if err != nil {
		return nil, err
	}

	return resNodePowerState, nil
}

// CreateNodeDetail : Create a nodeDetail
func (rc *RPCClient) CreateNodeDetail(in *pb.ReqCreateNodeDetail) (*pb.ResCreateNodeDetail, error) {
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
func (rc *RPCClient) GetNodeDetail(uuid string) (*pb.ResGetNodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNodeDetail, err := rc.flute.GetNodeDetail(ctx, &pb.ReqGetNodeDetail{NodeUUID: uuid})
	if err != nil {
		return nil, err
	}

	return resGetNodeDetail, nil
}

// DeleteNodeDetail : Delete of the nodeDetail
func (rc *RPCClient) DeleteNodeDetail(uuid string) (*pb.ResDeleteNodeDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNodeDetail, err := rc.flute.DeleteNodeDetail(ctx, &pb.ReqDeleteNodeDetail{NodeUUID: uuid})
	if err != nil {
		return nil, err
	}

	return resDeleteNodeDetail, nil
}

// CreateNode : Create a node
func (rc *RPCClient) CreateNode(in *pb.ReqCreateNode) (*pb.ResCreateNode, error) {
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
func (rc *RPCClient) GetNode(uuid string) (*pb.ResGetNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNode, err := rc.flute.GetNode(ctx, &pb.ReqGetNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resGetNode, nil
}

// GetNodeList : Get the list of nodes
func (rc *RPCClient) GetNodeList(in *pb.ReqGetNodeList) (*pb.ResGetNodeList, error) {
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
func (rc *RPCClient) GetNodeNum() (*pb.ResGetNodeNum, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resGetNodeNum, err := rc.flute.GetNodeNum(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	return resGetNodeNum, nil
}

// UpdateNode : Update infos of the node
func (rc *RPCClient) UpdateNode(in *pb.ReqUpdateNode) (*pb.ResUpdateNode, error) {
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
func (rc *RPCClient) DeleteNode(uuid string) (*pb.ResDeleteNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	resDeleteNode, err := rc.flute.DeleteNode(ctx, &pb.ReqDeleteNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return resDeleteNode, nil
}