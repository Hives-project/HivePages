// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: page.proto

package pb

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

// PageServiceClient is the client API for PageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PageServiceClient interface {
	GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageResponse, error)
}

type pageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPageServiceClient(cc grpc.ClientConnInterface) PageServiceClient {
	return &pageServiceClient{cc}
}

func (c *pageServiceClient) GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageResponse, error) {
	out := new(PageResponse)
	err := c.cc.Invoke(ctx, "/pb.PageService/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PageServiceServer is the server API for PageService service.
// All implementations must embed UnimplementedPageServiceServer
// for forward compatibility
type PageServiceServer interface {
	GetPage(context.Context, *PageRequest) (*PageResponse, error)
	mustEmbedUnimplementedPageServiceServer()
}

// UnimplementedPageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPageServiceServer struct {
}

func (UnimplementedPageServiceServer) GetPage(context.Context, *PageRequest) (*PageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedPageServiceServer) mustEmbedUnimplementedPageServiceServer() {}

// UnsafePageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PageServiceServer will
// result in compilation errors.
type UnsafePageServiceServer interface {
	mustEmbedUnimplementedPageServiceServer()
}

func RegisterPageServiceServer(s grpc.ServiceRegistrar, srv PageServiceServer) {
	s.RegisterService(&PageService_ServiceDesc, srv)
}

func _PageService_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageServiceServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PageService/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageServiceServer).GetPage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PageService_ServiceDesc is the grpc.ServiceDesc for PageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PageService",
	HandlerType: (*PageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPage",
			Handler:    _PageService_GetPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "page.proto",
}
