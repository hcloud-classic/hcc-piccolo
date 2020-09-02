// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcviolin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	rpcmsgType "hcc/piccolo/action/grpc/pb/rpcmsgType"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ViolinClient is the client API for Violin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViolinClient interface {
	// Server
	CreateServer(ctx context.Context, in *ReqCreateServer, opts ...grpc.CallOption) (*ResCreateServer, error)
	GetServer(ctx context.Context, in *ReqGetServer, opts ...grpc.CallOption) (*ResGetServer, error)
	GetServerList(ctx context.Context, in *ReqGetServerList, opts ...grpc.CallOption) (*ResGetServerList, error)
	GetServerNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetServerNum, error)
	UpdateServer(ctx context.Context, in *ReqUpdateServer, opts ...grpc.CallOption) (*ResUpdateServer, error)
	DeleteServer(ctx context.Context, in *ReqDeleteServer, opts ...grpc.CallOption) (*ResDeleteServer, error)
	// ServerNode
	CreateServerNode(ctx context.Context, in *ReqCreateServerNode, opts ...grpc.CallOption) (*ResCreateServerNode, error)
	GetServerNode(ctx context.Context, in *ReqGetServerNode, opts ...grpc.CallOption) (*ResGetServerNode, error)
	GetServerNodeList(ctx context.Context, in *ReqGetServerNodeList, opts ...grpc.CallOption) (*ResGetServerNodeList, error)
	GetServerNodeNum(ctx context.Context, in *ReqGetServerNodeNum, opts ...grpc.CallOption) (*ResGetServerNodeNum, error)
	DeleteServerNode(ctx context.Context, in *ReqDeleteServerNode, opts ...grpc.CallOption) (*ResDeleteServerNode, error)
	DeleteServerNodeByServerUUID(ctx context.Context, in *ReqDeleteServerNodeByServerUUID, opts ...grpc.CallOption) (*ResDeleteServerNodeByServerUUID, error)
}

type violinClient struct {
	cc grpc.ClientConnInterface
}

func NewViolinClient(cc grpc.ClientConnInterface) ViolinClient {
	return &violinClient{cc}
}

func (c *violinClient) CreateServer(ctx context.Context, in *ReqCreateServer, opts ...grpc.CallOption) (*ResCreateServer, error) {
	out := new(ResCreateServer)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServer(ctx context.Context, in *ReqGetServer, opts ...grpc.CallOption) (*ResGetServer, error) {
	out := new(ResGetServer)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServerList(ctx context.Context, in *ReqGetServerList, opts ...grpc.CallOption) (*ResGetServerList, error) {
	out := new(ResGetServerList)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServerNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetServerNum, error) {
	out := new(ResGetServerNum)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServerNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) UpdateServer(ctx context.Context, in *ReqUpdateServer, opts ...grpc.CallOption) (*ResUpdateServer, error) {
	out := new(ResUpdateServer)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/UpdateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) DeleteServer(ctx context.Context, in *ReqDeleteServer, opts ...grpc.CallOption) (*ResDeleteServer, error) {
	out := new(ResDeleteServer)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/DeleteServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) CreateServerNode(ctx context.Context, in *ReqCreateServerNode, opts ...grpc.CallOption) (*ResCreateServerNode, error) {
	out := new(ResCreateServerNode)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/CreateServerNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServerNode(ctx context.Context, in *ReqGetServerNode, opts ...grpc.CallOption) (*ResGetServerNode, error) {
	out := new(ResGetServerNode)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServerNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServerNodeList(ctx context.Context, in *ReqGetServerNodeList, opts ...grpc.CallOption) (*ResGetServerNodeList, error) {
	out := new(ResGetServerNodeList)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServerNodeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) GetServerNodeNum(ctx context.Context, in *ReqGetServerNodeNum, opts ...grpc.CallOption) (*ResGetServerNodeNum, error) {
	out := new(ResGetServerNodeNum)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/GetServerNodeNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) DeleteServerNode(ctx context.Context, in *ReqDeleteServerNode, opts ...grpc.CallOption) (*ResDeleteServerNode, error) {
	out := new(ResDeleteServerNode)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/DeleteServerNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *violinClient) DeleteServerNodeByServerUUID(ctx context.Context, in *ReqDeleteServerNodeByServerUUID, opts ...grpc.CallOption) (*ResDeleteServerNodeByServerUUID, error) {
	out := new(ResDeleteServerNodeByServerUUID)
	err := c.cc.Invoke(ctx, "/RpcViolin.Violin/DeleteServerNodeByServerUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViolinServer is the server API for Violin service.
// All implementations must embed UnimplementedViolinServer
// for forward compatibility
type ViolinServer interface {
	// Server
	CreateServer(context.Context, *ReqCreateServer) (*ResCreateServer, error)
	GetServer(context.Context, *ReqGetServer) (*ResGetServer, error)
	GetServerList(context.Context, *ReqGetServerList) (*ResGetServerList, error)
	GetServerNum(context.Context, *rpcmsgType.Empty) (*ResGetServerNum, error)
	UpdateServer(context.Context, *ReqUpdateServer) (*ResUpdateServer, error)
	DeleteServer(context.Context, *ReqDeleteServer) (*ResDeleteServer, error)
	// ServerNode
	CreateServerNode(context.Context, *ReqCreateServerNode) (*ResCreateServerNode, error)
	GetServerNode(context.Context, *ReqGetServerNode) (*ResGetServerNode, error)
	GetServerNodeList(context.Context, *ReqGetServerNodeList) (*ResGetServerNodeList, error)
	GetServerNodeNum(context.Context, *ReqGetServerNodeNum) (*ResGetServerNodeNum, error)
	DeleteServerNode(context.Context, *ReqDeleteServerNode) (*ResDeleteServerNode, error)
	DeleteServerNodeByServerUUID(context.Context, *ReqDeleteServerNodeByServerUUID) (*ResDeleteServerNodeByServerUUID, error)
	mustEmbedUnimplementedViolinServer()
}

// UnimplementedViolinServer must be embedded to have forward compatible implementations.
type UnimplementedViolinServer struct {
}

func (*UnimplementedViolinServer) CreateServer(context.Context, *ReqCreateServer) (*ResCreateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (*UnimplementedViolinServer) GetServer(context.Context, *ReqGetServer) (*ResGetServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (*UnimplementedViolinServer) GetServerList(context.Context, *ReqGetServerList) (*ResGetServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (*UnimplementedViolinServer) GetServerNum(context.Context, *rpcmsgType.Empty) (*ResGetServerNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerNum not implemented")
}
func (*UnimplementedViolinServer) UpdateServer(context.Context, *ReqUpdateServer) (*ResUpdateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (*UnimplementedViolinServer) DeleteServer(context.Context, *ReqDeleteServer) (*ResDeleteServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (*UnimplementedViolinServer) CreateServerNode(context.Context, *ReqCreateServerNode) (*ResCreateServerNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServerNode not implemented")
}
func (*UnimplementedViolinServer) GetServerNode(context.Context, *ReqGetServerNode) (*ResGetServerNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerNode not implemented")
}
func (*UnimplementedViolinServer) GetServerNodeList(context.Context, *ReqGetServerNodeList) (*ResGetServerNodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerNodeList not implemented")
}
func (*UnimplementedViolinServer) GetServerNodeNum(context.Context, *ReqGetServerNodeNum) (*ResGetServerNodeNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerNodeNum not implemented")
}
func (*UnimplementedViolinServer) DeleteServerNode(context.Context, *ReqDeleteServerNode) (*ResDeleteServerNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerNode not implemented")
}
func (*UnimplementedViolinServer) DeleteServerNodeByServerUUID(context.Context, *ReqDeleteServerNodeByServerUUID) (*ResDeleteServerNodeByServerUUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerNodeByServerUUID not implemented")
}
func (*UnimplementedViolinServer) mustEmbedUnimplementedViolinServer() {}

func RegisterViolinServer(s *grpc.Server, srv ViolinServer) {
	s.RegisterService(&_Violin_serviceDesc, srv)
}

func _Violin_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).CreateServer(ctx, req.(*ReqCreateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServer(ctx, req.(*ReqGetServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetServerList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServerList(ctx, req.(*ReqGetServerList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServerNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcmsgType.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServerNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServerNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServerNum(ctx, req.(*rpcmsgType.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/UpdateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).UpdateServer(ctx, req.(*ReqUpdateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/DeleteServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).DeleteServer(ctx, req.(*ReqDeleteServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_CreateServerNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateServerNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).CreateServerNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/CreateServerNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).CreateServerNode(ctx, req.(*ReqCreateServerNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServerNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetServerNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServerNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServerNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServerNode(ctx, req.(*ReqGetServerNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServerNodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetServerNodeList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServerNodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServerNodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServerNodeList(ctx, req.(*ReqGetServerNodeList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_GetServerNodeNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetServerNodeNum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).GetServerNodeNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/GetServerNodeNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).GetServerNodeNum(ctx, req.(*ReqGetServerNodeNum))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_DeleteServerNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteServerNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).DeleteServerNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/DeleteServerNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).DeleteServerNode(ctx, req.(*ReqDeleteServerNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Violin_DeleteServerNodeByServerUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteServerNodeByServerUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolinServer).DeleteServerNodeByServerUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcViolin.Violin/DeleteServerNodeByServerUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolinServer).DeleteServerNodeByServerUUID(ctx, req.(*ReqDeleteServerNodeByServerUUID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Violin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RpcViolin.Violin",
	HandlerType: (*ViolinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServer",
			Handler:    _Violin_CreateServer_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _Violin_GetServer_Handler,
		},
		{
			MethodName: "GetServerList",
			Handler:    _Violin_GetServerList_Handler,
		},
		{
			MethodName: "GetServerNum",
			Handler:    _Violin_GetServerNum_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _Violin_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _Violin_DeleteServer_Handler,
		},
		{
			MethodName: "CreateServerNode",
			Handler:    _Violin_CreateServerNode_Handler,
		},
		{
			MethodName: "GetServerNode",
			Handler:    _Violin_GetServerNode_Handler,
		},
		{
			MethodName: "GetServerNodeList",
			Handler:    _Violin_GetServerNodeList_Handler,
		},
		{
			MethodName: "GetServerNodeNum",
			Handler:    _Violin_GetServerNodeNum_Handler,
		},
		{
			MethodName: "DeleteServerNode",
			Handler:    _Violin_DeleteServerNode_Handler,
		},
		{
			MethodName: "DeleteServerNodeByServerUUID",
			Handler:    _Violin_DeleteServerNodeByServerUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "violin.proto",
}
