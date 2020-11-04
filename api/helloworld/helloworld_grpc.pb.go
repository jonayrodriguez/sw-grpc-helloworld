// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloworldClient interface {
	// Sends a HelloRequest
	SayHelloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldReply, error)
}

type helloworldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldClient(cc grpc.ClientConnInterface) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) SayHelloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldReply, error) {
	out := new(HelloworldReply)
	err := c.cc.Invoke(ctx, "/helloworld.Helloworld/SayHelloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServer is the server API for Helloworld service.
// All implementations must embed UnimplementedHelloworldServer
// for forward compatibility
type HelloworldServer interface {
	// Sends a HelloRequest
	SayHelloworld(context.Context, *HelloworldRequest) (*HelloworldReply, error)
	mustEmbedUnimplementedHelloworldServer()
}

// UnimplementedHelloworldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (UnimplementedHelloworldServer) SayHelloworld(context.Context, *HelloworldRequest) (*HelloworldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloworld not implemented")
}
func (UnimplementedHelloworldServer) mustEmbedUnimplementedHelloworldServer() {}

// UnsafeHelloworldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloworldServer will
// result in compilation errors.
type UnsafeHelloworldServer interface {
	mustEmbedUnimplementedHelloworldServer()
}

func RegisterHelloworldServer(s grpc.ServiceRegistrar, srv HelloworldServer) {
	s.RegisterService(&_Helloworld_serviceDesc, srv)
}

func _Helloworld_SayHelloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).SayHelloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Helloworld/SayHelloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).SayHelloworld(ctx, req.(*HelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloworld",
			Handler:    _Helloworld_SayHelloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/helloworld/helloworld.proto",
}
