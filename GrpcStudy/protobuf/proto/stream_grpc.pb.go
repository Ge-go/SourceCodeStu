// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: stream.proto

package proto

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

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	// 关键字stream指定启用流特性，参数部分是接收客户端参数的流，返回值是返回给客户端的流。
	//重新生成代码可以看到接口中新增加的Channel方法的定义
	GetStream(ctx context.Context, in *StreamRequestData, opts ...grpc.CallOption) (StreamService_GetStreamClient, error)
	PostStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_PostStreamClient, error)
	AllStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_AllStreamClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) GetStream(ctx context.Context, in *StreamRequestData, opts ...grpc.CallOption) (StreamService_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/StreamService/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_GetStreamClient interface {
	Recv() (*StreamResponsetData, error)
	grpc.ClientStream
}

type streamServiceGetStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceGetStreamClient) Recv() (*StreamResponsetData, error) {
	m := new(StreamResponsetData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) PostStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_PostStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/StreamService/PostStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServicePostStreamClient{stream}
	return x, nil
}

type StreamService_PostStreamClient interface {
	Send(*StreamRequestData) error
	CloseAndRecv() (*StreamResponsetData, error)
	grpc.ClientStream
}

type streamServicePostStreamClient struct {
	grpc.ClientStream
}

func (x *streamServicePostStreamClient) Send(m *StreamRequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServicePostStreamClient) CloseAndRecv() (*StreamResponsetData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponsetData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/StreamService/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceAllStreamClient{stream}
	return x, nil
}

type StreamService_AllStreamClient interface {
	Send(*StreamRequestData) error
	Recv() (*StreamResponsetData, error)
	grpc.ClientStream
}

type streamServiceAllStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceAllStreamClient) Send(m *StreamRequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceAllStreamClient) Recv() (*StreamResponsetData, error) {
	m := new(StreamResponsetData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	// 关键字stream指定启用流特性，参数部分是接收客户端参数的流，返回值是返回给客户端的流。
	//重新生成代码可以看到接口中新增加的Channel方法的定义
	GetStream(*StreamRequestData, StreamService_GetStreamServer) error
	PostStream(StreamService_PostStreamServer) error
	AllStream(StreamService_AllStreamServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) GetStream(*StreamRequestData, StreamService_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedStreamServiceServer) PostStream(StreamService_PostStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PostStream not implemented")
}
func (UnimplementedStreamServiceServer) AllStream(StreamService_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequestData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).GetStream(m, &streamServiceGetStreamServer{stream})
}

type StreamService_GetStreamServer interface {
	Send(*StreamResponsetData) error
	grpc.ServerStream
}

type streamServiceGetStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceGetStreamServer) Send(m *StreamResponsetData) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_PostStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).PostStream(&streamServicePostStreamServer{stream})
}

type StreamService_PostStreamServer interface {
	SendAndClose(*StreamResponsetData) error
	Recv() (*StreamRequestData, error)
	grpc.ServerStream
}

type streamServicePostStreamServer struct {
	grpc.ServerStream
}

func (x *streamServicePostStreamServer) SendAndClose(m *StreamResponsetData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServicePostStreamServer) Recv() (*StreamRequestData, error) {
	m := new(StreamRequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).AllStream(&streamServiceAllStreamServer{stream})
}

type StreamService_AllStreamServer interface {
	Send(*StreamResponsetData) error
	Recv() (*StreamRequestData, error)
	grpc.ServerStream
}

type streamServiceAllStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceAllStreamServer) Send(m *StreamResponsetData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceAllStreamServer) Recv() (*StreamRequestData, error) {
	m := new(StreamRequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _StreamService_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostStream",
			Handler:       _StreamService_PostStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _StreamService_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
