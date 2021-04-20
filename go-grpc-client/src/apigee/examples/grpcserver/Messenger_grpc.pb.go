// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: src/apigee/examples/grpcserver/Messenger.proto

package grpcserver

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

// MessengerServiceClient is the client API for MessengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerServiceClient interface {
	GetGreeting(ctx context.Context, in *MessengerRequest, opts ...grpc.CallOption) (*MessengerResponse, error)
	GetPirateGreeting(ctx context.Context, in *MessengerRequest, opts ...grpc.CallOption) (*MessengerResponse, error)
}

type messengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerServiceClient(cc grpc.ClientConnInterface) MessengerServiceClient {
	return &messengerServiceClient{cc}
}

func (c *messengerServiceClient) GetGreeting(ctx context.Context, in *MessengerRequest, opts ...grpc.CallOption) (*MessengerResponse, error) {
	out := new(MessengerResponse)
	err := c.cc.Invoke(ctx, "/com.apigee.examples.grpc.server.grpcserver.MessengerService/getGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetPirateGreeting(ctx context.Context, in *MessengerRequest, opts ...grpc.CallOption) (*MessengerResponse, error) {
	out := new(MessengerResponse)
	err := c.cc.Invoke(ctx, "/com.apigee.examples.grpc.server.grpcserver.MessengerService/getPirateGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServiceServer is the server API for MessengerService service.
// All implementations must embed UnimplementedMessengerServiceServer
// for forward compatibility
type MessengerServiceServer interface {
	GetGreeting(context.Context, *MessengerRequest) (*MessengerResponse, error)
	GetPirateGreeting(context.Context, *MessengerRequest) (*MessengerResponse, error)
	mustEmbedUnimplementedMessengerServiceServer()
}

// UnimplementedMessengerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServiceServer struct {
}

func (UnimplementedMessengerServiceServer) GetGreeting(context.Context, *MessengerRequest) (*MessengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeting not implemented")
}
func (UnimplementedMessengerServiceServer) GetPirateGreeting(context.Context, *MessengerRequest) (*MessengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPirateGreeting not implemented")
}
func (UnimplementedMessengerServiceServer) mustEmbedUnimplementedMessengerServiceServer() {}

// UnsafeMessengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServiceServer will
// result in compilation errors.
type UnsafeMessengerServiceServer interface {
	mustEmbedUnimplementedMessengerServiceServer()
}

func RegisterMessengerServiceServer(s grpc.ServiceRegistrar, srv MessengerServiceServer) {
	s.RegisterService(&MessengerService_ServiceDesc, srv)
}

func _MessengerService_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.apigee.examples.grpc.server.grpcserver.MessengerService/getGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetGreeting(ctx, req.(*MessengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetPirateGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetPirateGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.apigee.examples.grpc.server.grpcserver.MessengerService/getPirateGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetPirateGreeting(ctx, req.(*MessengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessengerService_ServiceDesc is the grpc.ServiceDesc for MessengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.apigee.examples.grpc.server.grpcserver.MessengerService",
	HandlerType: (*MessengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getGreeting",
			Handler:    _MessengerService_GetGreeting_Handler,
		},
		{
			MethodName: "getPirateGreeting",
			Handler:    _MessengerService_GetPirateGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/apigee/examples/grpcserver/Messenger.proto",
}