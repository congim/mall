// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: home.proto

package home

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

// HomeClient is the client API for Home service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeClient interface {
	Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*BannerRes, error)
}

type homeClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeClient(cc grpc.ClientConnInterface) HomeClient {
	return &homeClient{cc}
}

func (c *homeClient) Banner(ctx context.Context, in *BannerReq, opts ...grpc.CallOption) (*BannerRes, error) {
	out := new(BannerRes)
	err := c.cc.Invoke(ctx, "/home.Home/Banner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServer is the server API for Home service.
// All implementations must embed UnimplementedHomeServer
// for forward compatibility
type HomeServer interface {
	Banner(context.Context, *BannerReq) (*BannerRes, error)
	mustEmbedUnimplementedHomeServer()
}

// UnimplementedHomeServer must be embedded to have forward compatible implementations.
type UnimplementedHomeServer struct {
}

func (UnimplementedHomeServer) Banner(context.Context, *BannerReq) (*BannerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Banner not implemented")
}
func (UnimplementedHomeServer) mustEmbedUnimplementedHomeServer() {}

// UnsafeHomeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServer will
// result in compilation errors.
type UnsafeHomeServer interface {
	mustEmbedUnimplementedHomeServer()
}

func RegisterHomeServer(s grpc.ServiceRegistrar, srv HomeServer) {
	s.RegisterService(&Home_ServiceDesc, srv)
}

func _Home_Banner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServer).Banner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/home.Home/Banner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServer).Banner(ctx, req.(*BannerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Home_ServiceDesc is the grpc.ServiceDesc for Home service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Home_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "home.Home",
	HandlerType: (*HomeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Banner",
			Handler:    _Home_Banner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home.proto",
}
