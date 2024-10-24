// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: grpc/proto.proto

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

const (
	SendShareService_SendShare_FullMethodName = "/proto.SendShareService/SendShare"
)

// SendShareServiceClient is the client API for SendShareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendShareServiceClient interface {
	SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledgement, error)
}

type sendShareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendShareServiceClient(cc grpc.ClientConnInterface) SendShareServiceClient {
	return &sendShareServiceClient{cc}
}

func (c *sendShareServiceClient) SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, SendShareService_SendShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendShareServiceServer is the server API for SendShareService service.
// All implementations must embed UnimplementedSendShareServiceServer
// for forward compatibility
type SendShareServiceServer interface {
	SendShare(context.Context, *Share) (*Acknowledgement, error)
	mustEmbedUnimplementedSendShareServiceServer()
}

// UnimplementedSendShareServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendShareServiceServer struct {
}

func (UnimplementedSendShareServiceServer) SendShare(context.Context, *Share) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShare not implemented")
}
func (UnimplementedSendShareServiceServer) mustEmbedUnimplementedSendShareServiceServer() {}

// UnsafeSendShareServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendShareServiceServer will
// result in compilation errors.
type UnsafeSendShareServiceServer interface {
	mustEmbedUnimplementedSendShareServiceServer()
}

func RegisterSendShareServiceServer(s grpc.ServiceRegistrar, srv SendShareServiceServer) {
	s.RegisterService(&SendShareService_ServiceDesc, srv)
}

func _SendShareService_SendShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Share)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendShareServiceServer).SendShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SendShareService_SendShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendShareServiceServer).SendShare(ctx, req.(*Share))
	}
	return interceptor(ctx, in, info, handler)
}

// SendShareService_ServiceDesc is the grpc.ServiceDesc for SendShareService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendShareService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendShareService",
	HandlerType: (*SendShareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendShare",
			Handler:    _SendShareService_SendShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}

const (
	SendAggregatedShareService_SendAggregatedShare_FullMethodName = "/proto.SendAggregatedShareService/SendAggregatedShare"
)

// SendAggregatedShareServiceClient is the client API for SendAggregatedShareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendAggregatedShareServiceClient interface {
	SendAggregatedShare(ctx context.Context, in *AggregatedShare, opts ...grpc.CallOption) (*Acknowledgement, error)
}

type sendAggregatedShareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendAggregatedShareServiceClient(cc grpc.ClientConnInterface) SendAggregatedShareServiceClient {
	return &sendAggregatedShareServiceClient{cc}
}

func (c *sendAggregatedShareServiceClient) SendAggregatedShare(ctx context.Context, in *AggregatedShare, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, SendAggregatedShareService_SendAggregatedShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendAggregatedShareServiceServer is the server API for SendAggregatedShareService service.
// All implementations must embed UnimplementedSendAggregatedShareServiceServer
// for forward compatibility
type SendAggregatedShareServiceServer interface {
	SendAggregatedShare(context.Context, *AggregatedShare) (*Acknowledgement, error)
	mustEmbedUnimplementedSendAggregatedShareServiceServer()
}

// UnimplementedSendAggregatedShareServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendAggregatedShareServiceServer struct {
}

func (UnimplementedSendAggregatedShareServiceServer) SendAggregatedShare(context.Context, *AggregatedShare) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAggregatedShare not implemented")
}
func (UnimplementedSendAggregatedShareServiceServer) mustEmbedUnimplementedSendAggregatedShareServiceServer() {
}

// UnsafeSendAggregatedShareServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendAggregatedShareServiceServer will
// result in compilation errors.
type UnsafeSendAggregatedShareServiceServer interface {
	mustEmbedUnimplementedSendAggregatedShareServiceServer()
}

func RegisterSendAggregatedShareServiceServer(s grpc.ServiceRegistrar, srv SendAggregatedShareServiceServer) {
	s.RegisterService(&SendAggregatedShareService_ServiceDesc, srv)
}

func _SendAggregatedShareService_SendAggregatedShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregatedShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendAggregatedShareServiceServer).SendAggregatedShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SendAggregatedShareService_SendAggregatedShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendAggregatedShareServiceServer).SendAggregatedShare(ctx, req.(*AggregatedShare))
	}
	return interceptor(ctx, in, info, handler)
}

// SendAggregatedShareService_ServiceDesc is the grpc.ServiceDesc for SendAggregatedShareService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendAggregatedShareService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendAggregatedShareService",
	HandlerType: (*SendAggregatedShareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendAggregatedShare",
			Handler:    _SendAggregatedShareService_SendAggregatedShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
