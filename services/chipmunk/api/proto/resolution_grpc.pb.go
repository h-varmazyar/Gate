// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services/chipmunk/api/proto/src/resolution.proto

package proto

import (
	context "context"
	proto "github.com/h-varmazyar/Gate/api/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResolutionServiceClient is the client API for ResolutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolutionServiceClient interface {
	Set(ctx context.Context, in *Resolution, opts ...grpc.CallOption) (*proto.Void, error)
	ReturnByID(ctx context.Context, in *ResolutionReturnByIDReq, opts ...grpc.CallOption) (*Resolution, error)
	ReturnByDuration(ctx context.Context, in *ResolutionReturnByDurationReq, opts ...grpc.CallOption) (*Resolution, error)
	List(ctx context.Context, in *ResolutionListReq, opts ...grpc.CallOption) (*Resolutions, error)
}

type resolutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResolutionServiceClient(cc grpc.ClientConnInterface) ResolutionServiceClient {
	return &resolutionServiceClient{cc}
}

func (c *resolutionServiceClient) Set(ctx context.Context, in *Resolution, opts ...grpc.CallOption) (*proto.Void, error) {
	out := new(proto.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) ReturnByID(ctx context.Context, in *ResolutionReturnByIDReq, opts ...grpc.CallOption) (*Resolution, error) {
	out := new(Resolution)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/ReturnByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) ReturnByDuration(ctx context.Context, in *ResolutionReturnByDurationReq, opts ...grpc.CallOption) (*Resolution, error) {
	out := new(Resolution)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/ReturnByDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) List(ctx context.Context, in *ResolutionListReq, opts ...grpc.CallOption) (*Resolutions, error) {
	out := new(Resolutions)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolutionServiceServer is the server API for ResolutionService service.
// All implementations should embed UnimplementedResolutionServiceServer
// for forward compatibility
type ResolutionServiceServer interface {
	Set(context.Context, *Resolution) (*proto.Void, error)
	ReturnByID(context.Context, *ResolutionReturnByIDReq) (*Resolution, error)
	ReturnByDuration(context.Context, *ResolutionReturnByDurationReq) (*Resolution, error)
	List(context.Context, *ResolutionListReq) (*Resolutions, error)
}

// UnimplementedResolutionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedResolutionServiceServer struct {
}

func (UnimplementedResolutionServiceServer) Set(context.Context, *Resolution) (*proto.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedResolutionServiceServer) ReturnByID(context.Context, *ResolutionReturnByIDReq) (*Resolution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnByID not implemented")
}
func (UnimplementedResolutionServiceServer) ReturnByDuration(context.Context, *ResolutionReturnByDurationReq) (*Resolution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnByDuration not implemented")
}
func (UnimplementedResolutionServiceServer) List(context.Context, *ResolutionListReq) (*Resolutions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeResolutionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolutionServiceServer will
// result in compilation errors.
type UnsafeResolutionServiceServer interface {
	mustEmbedUnimplementedResolutionServiceServer()
}

func RegisterResolutionServiceServer(s grpc.ServiceRegistrar, srv ResolutionServiceServer) {
	s.RegisterService(&ResolutionService_ServiceDesc, srv)
}

func _ResolutionService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resolution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).Set(ctx, req.(*Resolution))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolutionService_ReturnByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolutionReturnByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).ReturnByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/ReturnByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).ReturnByID(ctx, req.(*ResolutionReturnByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolutionService_ReturnByDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolutionReturnByDurationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).ReturnByDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/ReturnByDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).ReturnByDuration(ctx, req.(*ResolutionReturnByDurationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolutionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolutionListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).List(ctx, req.(*ResolutionListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ResolutionService_ServiceDesc is the grpc.ServiceDesc for ResolutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResolutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chipmunkApi.ResolutionService",
	HandlerType: (*ResolutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _ResolutionService_Set_Handler,
		},
		{
			MethodName: "ReturnByID",
			Handler:    _ResolutionService_ReturnByID_Handler,
		},
		{
			MethodName: "ReturnByDuration",
			Handler:    _ResolutionService_ReturnByDuration_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResolutionService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/proto/src/resolution.proto",
}
