// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResponderClient is the client API for Responder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResponderClient interface {
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type responderClient struct {
	cc grpc.ClientConnInterface
}

func NewResponderClient(cc grpc.ClientConnInterface) ResponderClient {
	return &responderClient{cc}
}

func (c *responderClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/responder.Responder/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResponderServer is the server API for Responder service.
// All implementations should embed UnimplementedResponderServer
// for forward compatibility
type ResponderServer interface {
	GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error)
}

// UnimplementedResponderServer should be embedded to have forward compatible implementations.
type UnimplementedResponderServer struct {
}

func (UnimplementedResponderServer) GetVersion(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

// UnsafeResponderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResponderServer will
// result in compilation errors.
type UnsafeResponderServer interface {
	mustEmbedUnimplementedResponderServer()
}

func RegisterResponderServer(s grpc.ServiceRegistrar, srv ResponderServer) {
	s.RegisterService(&Responder_ServiceDesc, srv)
}

func _Responder_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponderServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/responder.Responder/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponderServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Responder_ServiceDesc is the grpc.ServiceDesc for Responder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Responder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "responder.Responder",
	HandlerType: (*ResponderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Responder_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro/responder/pkg/pb/service.proto",
}
