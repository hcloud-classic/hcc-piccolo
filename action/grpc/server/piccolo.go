package server

import (
	"context"
	"errors"
	"github.com/hcloud-classic/pb"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/sqlite/serveractions"
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

	err := serveractions.WriteServerAction(in.GetServerUUID(), action, result, errStr, token)
	if err != nil {
		return nil, err
	}

	return &pb.ResWriteServerAction{Result: "Success"}, nil
}
