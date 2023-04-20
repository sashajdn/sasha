// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: service.openai/proto/openai.proto

package openaiproto

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

// OpenaiClient is the client API for Openai service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenaiClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type openaiClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenaiClient(cc grpc.ClientConnInterface) OpenaiClient {
	return &openaiClient{cc}
}

func (c *openaiClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/openai/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenaiServer is the server API for Openai service.
// All implementations must embed UnimplementedOpenaiServer
// for forward compatibility
type OpenaiServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedOpenaiServer()
}

// UnimplementedOpenaiServer must be embedded to have forward compatible implementations.
type UnimplementedOpenaiServer struct {
}

func (UnimplementedOpenaiServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOpenaiServer) mustEmbedUnimplementedOpenaiServer() {}

// UnsafeOpenaiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenaiServer will
// result in compilation errors.
type UnsafeOpenaiServer interface {
	mustEmbedUnimplementedOpenaiServer()
}

func RegisterOpenaiServer(s grpc.ServiceRegistrar, srv OpenaiServer) {
	s.RegisterService(&Openai_ServiceDesc, srv)
}

func _Openai_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenaiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openai/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenaiServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Openai_ServiceDesc is the grpc.ServiceDesc for Openai service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Openai_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openai",
	HandlerType: (*OpenaiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Openai_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.openai/proto/openai.proto",
}