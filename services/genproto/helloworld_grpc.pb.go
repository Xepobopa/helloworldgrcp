// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: helloworld.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SendBroadcast_Broadcast_FullMethodName = "/helloworld.SendBroadcast/Broadcast"
	SendBroadcast_Subscribe_FullMethodName = "/helloworld.SendBroadcast/Subscribe"
)

// SendBroadcastClient is the client API for SendBroadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendBroadcastClient interface {
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Subscribe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error)
}

type sendBroadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewSendBroadcastClient(cc grpc.ClientConnInterface) SendBroadcastClient {
	return &sendBroadcastClient{cc}
}

func (c *sendBroadcastClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SendBroadcast_Broadcast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendBroadcastClient) Subscribe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SendBroadcast_ServiceDesc.Streams[0], SendBroadcast_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, Message]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SendBroadcast_SubscribeClient = grpc.ServerStreamingClient[Message]

// SendBroadcastServer is the server API for SendBroadcast service.
// All implementations must embed UnimplementedSendBroadcastServer
// for forward compatibility.
type SendBroadcastServer interface {
	Broadcast(context.Context, *Message) (*emptypb.Empty, error)
	Subscribe(*emptypb.Empty, grpc.ServerStreamingServer[Message]) error
	mustEmbedUnimplementedSendBroadcastServer()
}

// UnimplementedSendBroadcastServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSendBroadcastServer struct{}

func (UnimplementedSendBroadcastServer) Broadcast(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedSendBroadcastServer) Subscribe(*emptypb.Empty, grpc.ServerStreamingServer[Message]) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSendBroadcastServer) mustEmbedUnimplementedSendBroadcastServer() {}
func (UnimplementedSendBroadcastServer) testEmbeddedByValue()                       {}

// UnsafeSendBroadcastServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendBroadcastServer will
// result in compilation errors.
type UnsafeSendBroadcastServer interface {
	mustEmbedUnimplementedSendBroadcastServer()
}

func RegisterSendBroadcastServer(s grpc.ServiceRegistrar, srv SendBroadcastServer) {
	// If the following call pancis, it indicates UnimplementedSendBroadcastServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SendBroadcast_ServiceDesc, srv)
}

func _SendBroadcast_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendBroadcastServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SendBroadcast_Broadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendBroadcastServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendBroadcast_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SendBroadcastServer).Subscribe(m, &grpc.GenericServerStream[emptypb.Empty, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SendBroadcast_SubscribeServer = grpc.ServerStreamingServer[Message]

// SendBroadcast_ServiceDesc is the grpc.ServiceDesc for SendBroadcast service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendBroadcast_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.SendBroadcast",
	HandlerType: (*SendBroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _SendBroadcast_Broadcast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SendBroadcast_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
