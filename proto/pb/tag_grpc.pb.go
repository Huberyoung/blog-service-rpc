// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: proto/tag.proto

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

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error) {
	out := new(GetTagListReply)
	err := c.cc.Invoke(ctx, "/TagService/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TagService/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTagList(ctx, req.(*GetTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagList",
			Handler:    _TagService_GetTagList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag.proto",
}
