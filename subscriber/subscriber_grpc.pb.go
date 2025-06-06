// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.0
// source: subscriber.proto

package subscriber

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PlatformSubscriber_Heartbeat_FullMethodName        = "/platform_subscriber.PlatformSubscriber/Heartbeat"
	PlatformSubscriber_MarkDisconnected_FullMethodName = "/platform_subscriber.PlatformSubscriber/MarkDisconnected"
)

// PlatformSubscriberClient is the client API for PlatformSubscriber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlatformSubscriberClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	MarkDisconnected(ctx context.Context, in *MarkDisconnectedRequest, opts ...grpc.CallOption) (*MarkDisconnectedResponse, error)
}

type platformSubscriberClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformSubscriberClient(cc grpc.ClientConnInterface) PlatformSubscriberClient {
	return &platformSubscriberClient{cc}
}

func (c *platformSubscriberClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, PlatformSubscriber_Heartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformSubscriberClient) MarkDisconnected(ctx context.Context, in *MarkDisconnectedRequest, opts ...grpc.CallOption) (*MarkDisconnectedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkDisconnectedResponse)
	err := c.cc.Invoke(ctx, PlatformSubscriber_MarkDisconnected_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformSubscriberServer is the server API for PlatformSubscriber service.
// All implementations must embed UnimplementedPlatformSubscriberServer
// for forward compatibility.
type PlatformSubscriberServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	MarkDisconnected(context.Context, *MarkDisconnectedRequest) (*MarkDisconnectedResponse, error)
	mustEmbedUnimplementedPlatformSubscriberServer()
}

// UnimplementedPlatformSubscriberServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlatformSubscriberServer struct{}

func (UnimplementedPlatformSubscriberServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedPlatformSubscriberServer) MarkDisconnected(context.Context, *MarkDisconnectedRequest) (*MarkDisconnectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkDisconnected not implemented")
}
func (UnimplementedPlatformSubscriberServer) mustEmbedUnimplementedPlatformSubscriberServer() {}
func (UnimplementedPlatformSubscriberServer) testEmbeddedByValue()                            {}

// UnsafePlatformSubscriberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformSubscriberServer will
// result in compilation errors.
type UnsafePlatformSubscriberServer interface {
	mustEmbedUnimplementedPlatformSubscriberServer()
}

func RegisterPlatformSubscriberServer(s grpc.ServiceRegistrar, srv PlatformSubscriberServer) {
	// If the following call pancis, it indicates UnimplementedPlatformSubscriberServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PlatformSubscriber_ServiceDesc, srv)
}

func _PlatformSubscriber_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformSubscriberServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformSubscriber_Heartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformSubscriberServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformSubscriber_MarkDisconnected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkDisconnectedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformSubscriberServer).MarkDisconnected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformSubscriber_MarkDisconnected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformSubscriberServer).MarkDisconnected(ctx, req.(*MarkDisconnectedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlatformSubscriber_ServiceDesc is the grpc.ServiceDesc for PlatformSubscriber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlatformSubscriber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platform_subscriber.PlatformSubscriber",
	HandlerType: (*PlatformSubscriberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _PlatformSubscriber_Heartbeat_Handler,
		},
		{
			MethodName: "MarkDisconnected",
			Handler:    _PlatformSubscriber_MarkDisconnected_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscriber.proto",
}
