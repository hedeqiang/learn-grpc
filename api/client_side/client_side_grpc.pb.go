// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package client_side

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientSideClient is the client API for ClientSide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientSideClient interface {
	GetUser(ctx context.Context, opts ...grpc.CallOption) (ClientSide_GetUserClient, error)
}

type clientSideClient struct {
	cc grpc.ClientConnInterface
}

func NewClientSideClient(cc grpc.ClientConnInterface) ClientSideClient {
	return &clientSideClient{cc}
}

func (c *clientSideClient) GetUser(ctx context.Context, opts ...grpc.CallOption) (ClientSide_GetUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientSide_ServiceDesc.Streams[0], "/client_side.ClientSide/GetUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientSideGetUserClient{stream}
	return x, nil
}

type ClientSide_GetUserClient interface {
	Send(*GetUserRequest) error
	CloseAndRecv() (*GetUserResponse, error)
	grpc.ClientStream
}

type clientSideGetUserClient struct {
	grpc.ClientStream
}

func (x *clientSideGetUserClient) Send(m *GetUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientSideGetUserClient) CloseAndRecv() (*GetUserResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientSideServer is the server API for ClientSide service.
// All implementations must embed UnimplementedClientSideServer
// for forward compatibility
type ClientSideServer interface {
	GetUser(ClientSide_GetUserServer) error
	mustEmbedUnimplementedClientSideServer()
}

// UnimplementedClientSideServer must be embedded to have forward compatible implementations.
type UnimplementedClientSideServer struct {
}

func (UnimplementedClientSideServer) GetUser(ClientSide_GetUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedClientSideServer) mustEmbedUnimplementedClientSideServer() {}

// UnsafeClientSideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientSideServer will
// result in compilation errors.
type UnsafeClientSideServer interface {
	mustEmbedUnimplementedClientSideServer()
}

func RegisterClientSideServer(s grpc.ServiceRegistrar, srv ClientSideServer) {
	s.RegisterService(&ClientSide_ServiceDesc, srv)
}

func _ClientSide_GetUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientSideServer).GetUser(&clientSideGetUserServer{stream})
}

type ClientSide_GetUserServer interface {
	SendAndClose(*GetUserResponse) error
	Recv() (*GetUserRequest, error)
	grpc.ServerStream
}

type clientSideGetUserServer struct {
	grpc.ServerStream
}

func (x *clientSideGetUserServer) SendAndClose(m *GetUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientSideGetUserServer) Recv() (*GetUserRequest, error) {
	m := new(GetUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientSide_ServiceDesc is the grpc.ServiceDesc for ClientSide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientSide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client_side.ClientSide",
	HandlerType: (*ClientSideServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUser",
			Handler:       _ClientSide_GetUser_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/client_side/client_side.proto",
}
