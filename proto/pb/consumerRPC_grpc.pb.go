// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: consumerRPC.proto

package pb

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

// ConsumerClient is the client API for Consumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerClient interface {
	DataStream(ctx context.Context, in *RequestDataStream, opts ...grpc.CallOption) (Consumer_DataStreamClient, error)
	GetLastPrice(ctx context.Context, in *RequestGetLastPrice, opts ...grpc.CallOption) (*ResponseGetLastPrice, error)
}

type consumerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerClient(cc grpc.ClientConnInterface) ConsumerClient {
	return &consumerClient{cc}
}

func (c *consumerClient) DataStream(ctx context.Context, in *RequestDataStream, opts ...grpc.CallOption) (Consumer_DataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Consumer_ServiceDesc.Streams[0], "/pb.Consumer/DataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Consumer_DataStreamClient interface {
	Recv() (*ResponseDataStream, error)
	grpc.ClientStream
}

type consumerDataStreamClient struct {
	grpc.ClientStream
}

func (x *consumerDataStreamClient) Recv() (*ResponseDataStream, error) {
	m := new(ResponseDataStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerClient) GetLastPrice(ctx context.Context, in *RequestGetLastPrice, opts ...grpc.CallOption) (*ResponseGetLastPrice, error) {
	out := new(ResponseGetLastPrice)
	err := c.cc.Invoke(ctx, "/pb.Consumer/GetLastPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServer is the server API for Consumer service.
// All implementations must embed UnimplementedConsumerServer
// for forward compatibility
type ConsumerServer interface {
	DataStream(*RequestDataStream, Consumer_DataStreamServer) error
	GetLastPrice(context.Context, *RequestGetLastPrice) (*ResponseGetLastPrice, error)
	mustEmbedUnimplementedConsumerServer()
}

// UnimplementedConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServer struct {
}

func (UnimplementedConsumerServer) DataStream(*RequestDataStream, Consumer_DataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DataStream not implemented")
}
func (UnimplementedConsumerServer) GetLastPrice(context.Context, *RequestGetLastPrice) (*ResponseGetLastPrice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastPrice not implemented")
}
func (UnimplementedConsumerServer) mustEmbedUnimplementedConsumerServer() {}

// UnsafeConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServer will
// result in compilation errors.
type UnsafeConsumerServer interface {
	mustEmbedUnimplementedConsumerServer()
}

func RegisterConsumerServer(s grpc.ServiceRegistrar, srv ConsumerServer) {
	s.RegisterService(&Consumer_ServiceDesc, srv)
}

func _Consumer_DataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestDataStream)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsumerServer).DataStream(m, &consumerDataStreamServer{stream})
}

type Consumer_DataStreamServer interface {
	Send(*ResponseDataStream) error
	grpc.ServerStream
}

type consumerDataStreamServer struct {
	grpc.ServerStream
}

func (x *consumerDataStreamServer) Send(m *ResponseDataStream) error {
	return x.ServerStream.SendMsg(m)
}

func _Consumer_GetLastPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetLastPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetLastPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Consumer/GetLastPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetLastPrice(ctx, req.(*RequestGetLastPrice))
	}
	return interceptor(ctx, in, info, handler)
}

// Consumer_ServiceDesc is the grpc.ServiceDesc for Consumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Consumer",
	HandlerType: (*ConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastPrice",
			Handler:    _Consumer_GetLastPrice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataStream",
			Handler:       _Consumer_DataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "consumerRPC.proto",
}
