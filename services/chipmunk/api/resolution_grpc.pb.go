// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	api "github.com/h-varmazyar/Gate/api"
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
	Set(ctx context.Context, in *Resolution, opts ...grpc.CallOption) (*api.Void, error)
	GetByID(ctx context.Context, in *GetResolutionByIDRequest, opts ...grpc.CallOption) (*Resolution, error)
	GetByDuration(ctx context.Context, in *GetResolutionByDurationRequest, opts ...grpc.CallOption) (*Resolution, error)
	List(ctx context.Context, in *GetResolutionListRequest, opts ...grpc.CallOption) (*Resolutions, error)
}

type resolutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResolutionServiceClient(cc grpc.ClientConnInterface) ResolutionServiceClient {
	return &resolutionServiceClient{cc}
}

func (c *resolutionServiceClient) Set(ctx context.Context, in *Resolution, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) GetByID(ctx context.Context, in *GetResolutionByIDRequest, opts ...grpc.CallOption) (*Resolution, error) {
	out := new(Resolution)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) GetByDuration(ctx context.Context, in *GetResolutionByDurationRequest, opts ...grpc.CallOption) (*Resolution, error) {
	out := new(Resolution)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ResolutionService/GetByDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolutionServiceClient) List(ctx context.Context, in *GetResolutionListRequest, opts ...grpc.CallOption) (*Resolutions, error) {
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
	Set(context.Context, *Resolution) (*api.Void, error)
	GetByID(context.Context, *GetResolutionByIDRequest) (*Resolution, error)
	GetByDuration(context.Context, *GetResolutionByDurationRequest) (*Resolution, error)
	List(context.Context, *GetResolutionListRequest) (*Resolutions, error)
}

// UnimplementedResolutionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedResolutionServiceServer struct {
}

func (UnimplementedResolutionServiceServer) Set(context.Context, *Resolution) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedResolutionServiceServer) GetByID(context.Context, *GetResolutionByIDRequest) (*Resolution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedResolutionServiceServer) GetByDuration(context.Context, *GetResolutionByDurationRequest) (*Resolution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByDuration not implemented")
}
func (UnimplementedResolutionServiceServer) List(context.Context, *GetResolutionListRequest) (*Resolutions, error) {
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

func _ResolutionService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResolutionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).GetByID(ctx, req.(*GetResolutionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolutionService_GetByDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResolutionByDurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolutionServiceServer).GetByDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ResolutionService/GetByDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolutionServiceServer).GetByDuration(ctx, req.(*GetResolutionByDurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolutionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResolutionListRequest)
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
		return srv.(ResolutionServiceServer).List(ctx, req.(*GetResolutionListRequest))
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
			MethodName: "GetByID",
			Handler:    _ResolutionService_GetByID_Handler,
		},
		{
			MethodName: "GetByDuration",
			Handler:    _ResolutionService_GetByDuration_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResolutionService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/src/resolution.proto",
}
