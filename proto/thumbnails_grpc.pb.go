// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/thumbnails.proto

package proto

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
	YoutubeThumbnails_GetThumbnail_FullMethodName = "/YoutubeThumbnails/GetThumbnail"
)

// YoutubeThumbnailsClient is the client API for YoutubeThumbnails service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YoutubeThumbnailsClient interface {
	GetThumbnail(ctx context.Context, in *GetThumbnailRequest, opts ...grpc.CallOption) (*GetThumbnailResponse, error)
}

type youtubeThumbnailsClient struct {
	cc grpc.ClientConnInterface
}

func NewYoutubeThumbnailsClient(cc grpc.ClientConnInterface) YoutubeThumbnailsClient {
	return &youtubeThumbnailsClient{cc}
}

func (c *youtubeThumbnailsClient) GetThumbnail(ctx context.Context, in *GetThumbnailRequest, opts ...grpc.CallOption) (*GetThumbnailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetThumbnailResponse)
	err := c.cc.Invoke(ctx, YoutubeThumbnails_GetThumbnail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeThumbnailsServer is the server API for YoutubeThumbnails service.
// All implementations must embed UnimplementedYoutubeThumbnailsServer
// for forward compatibility.
type YoutubeThumbnailsServer interface {
	GetThumbnail(context.Context, *GetThumbnailRequest) (*GetThumbnailResponse, error)
	mustEmbedUnimplementedYoutubeThumbnailsServer()
}

// UnimplementedYoutubeThumbnailsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedYoutubeThumbnailsServer struct{}

func (UnimplementedYoutubeThumbnailsServer) GetThumbnail(context.Context, *GetThumbnailRequest) (*GetThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThumbnail not implemented")
}
func (UnimplementedYoutubeThumbnailsServer) mustEmbedUnimplementedYoutubeThumbnailsServer() {}
func (UnimplementedYoutubeThumbnailsServer) testEmbeddedByValue()                           {}

// UnsafeYoutubeThumbnailsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YoutubeThumbnailsServer will
// result in compilation errors.
type UnsafeYoutubeThumbnailsServer interface {
	mustEmbedUnimplementedYoutubeThumbnailsServer()
}

func RegisterYoutubeThumbnailsServer(s grpc.ServiceRegistrar, srv YoutubeThumbnailsServer) {
	// If the following call pancis, it indicates UnimplementedYoutubeThumbnailsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&YoutubeThumbnails_ServiceDesc, srv)
}

func _YoutubeThumbnails_GetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeThumbnailsServer).GetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeThumbnails_GetThumbnail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeThumbnailsServer).GetThumbnail(ctx, req.(*GetThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YoutubeThumbnails_ServiceDesc is the grpc.ServiceDesc for YoutubeThumbnails service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YoutubeThumbnails_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YoutubeThumbnails",
	HandlerType: (*YoutubeThumbnailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThumbnail",
			Handler:    _YoutubeThumbnails_GetThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/thumbnails.proto",
}
