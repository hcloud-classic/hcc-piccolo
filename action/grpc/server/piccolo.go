package server

import (
	"context"
	"errors"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"innogrid.com/hcloud-classic/hcc_errors"

	"innogrid.com/hcloud-classic/pb"
)

type piccoloServer struct {
	pb.UnimplementedPiccoloServer
}

func (s *piccoloServer) WriteServerAction(_ context.Context, in *pb.ReqWriteServerAction) (*pb.ResWriteServerAction, error) {
	logger.Logger.Println("Request received: WriteServerAction()")

	if in.GetServerUUID() == "" {
		return nil, errors.New("ServerUUID is empty")
	}

	if in.GetServerAction() == nil {
		return nil, errors.New("ServerAction is empty")
	}

	action := in.ServerAction.GetAction()
	result := in.ServerAction.GetResult()
	errStr := in.ServerAction.GetErrStr()
	token := in.ServerAction.GetToken()

	err := dao.WriteServerAction(in.GetServerUUID(), action, result, errStr, token)
	if err != nil {
		return nil, err
	}

	return &pb.ResWriteServerAction{Result: "Success"}, nil
}

func (s *piccoloServer) GetGroupList(_ context.Context, _ *pb.Empty) (*pb.ResGetGroupList, error) {
	// logger.Logger.Println("Request received: GetGroupList()")

	groupList, errCode, errStr := dao.ReadGroupList()
	if errCode != 0 {
		errStack := hcc_errors.NewHccErrorStack(hcc_errors.NewHccError(errCode, errStr))
		return &pb.ResGetGroupList{Group: []*pb.Group{}, HccErrorStack: errconv.HccStackToGrpc(errStack)}, nil
	}

	return groupList, nil
}
