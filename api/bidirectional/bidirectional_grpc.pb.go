// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bidirectional

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

// BidirectionalServiceClient is the client API for BidirectionalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidirectionalServiceClient interface {
	GetUser(ctx context.Context, opts ...grpc.CallOption) (BidirectionalService_GetUserClient, error)
}

type bidirectionalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBidirectionalServiceClient(cc grpc.ClientConnInterface) BidirectionalServiceClient {
	return &bidirectionalServiceClient{cc}
}

func (c *bidirectionalServiceClient) GetUser(ctx context.Context, opts ...grpc.CallOption) (BidirectionalService_GetUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &BidirectionalService_ServiceDesc.Streams[0], "/bidirectional.BidirectionalService/GetUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidirectionalServiceGetUserClient{stream}
	return x, nil
}

type BidirectionalService_GetUserClient interface {
	Send(*GetUserRequest) error
	Recv() (*GetUserResponse, error)
	grpc.ClientStream
}

type bidirectionalServiceGetUserClient struct {
	grpc.ClientStream
}

func (x *bidirectionalServiceGetUserClient) Send(m *GetUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidirectionalServiceGetUserClient) Recv() (*GetUserResponse, error) {
	m := new(GetUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalServiceServer is the server API for BidirectionalService service.
// All implementations must embed UnimplementedBidirectionalServiceServer
// for forward compatibility
type BidirectionalServiceServer interface {
	GetUser(BidirectionalService_GetUserServer) error
	mustEmbedUnimplementedBidirectionalServiceServer()
}

// UnimplementedBidirectionalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBidirectionalServiceServer struct {
}

func (UnimplementedBidirectionalServiceServer) GetUser(BidirectionalService_GetUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedBidirectionalServiceServer) mustEmbedUnimplementedBidirectionalServiceServer() {}

// UnsafeBidirectionalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidirectionalServiceServer will
// result in compilation errors.
type UnsafeBidirectionalServiceServer interface {
	mustEmbedUnimplementedBidirectionalServiceServer()
}

func RegisterBidirectionalServiceServer(s grpc.ServiceRegistrar, srv BidirectionalServiceServer) {
	s.RegisterService(&BidirectionalService_ServiceDesc, srv)
}

func _BidirectionalService_GetUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidirectionalServiceServer).GetUser(&bidirectionalServiceGetUserServer{stream})
}

type BidirectionalService_GetUserServer interface {
	Send(*GetUserResponse) error
	Recv() (*GetUserRequest, error)
	grpc.ServerStream
}

type bidirectionalServiceGetUserServer struct {
	grpc.ServerStream
}

func (x *bidirectionalServiceGetUserServer) Send(m *GetUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidirectionalServiceGetUserServer) Recv() (*GetUserRequest, error) {
	m := new(GetUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalService_ServiceDesc is the grpc.ServiceDesc for BidirectionalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidirectionalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bidirectional.BidirectionalService",
	HandlerType: (*BidirectionalServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUser",
			Handler:       _BidirectionalService_GetUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/bidirectional/bidirectional.proto",
}