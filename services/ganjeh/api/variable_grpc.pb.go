// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: services/ganjeh/api/src/variable.proto

package api

import (
	context "context"
	api "github.com/mrNobody95/Gate/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VariableServiceClient is the client API for VariableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VariableServiceClient interface {
	Set(ctx context.Context, in *SetVariableRequest, opts ...grpc.CallOption) (*api.Void, error)
	Get(ctx context.Context, in *GetVariableRequest, opts ...grpc.CallOption) (*Variable, error)
	List(ctx context.Context, in *GetVariablesRequest, opts ...grpc.CallOption) (*Variables, error)
}

type variableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVariableServiceClient(cc grpc.ClientConnInterface) VariableServiceClient {
	return &variableServiceClient{cc}
}

func (c *variableServiceClient) Set(ctx context.Context, in *SetVariableRequest, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/networkApi.VariableService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) Get(ctx context.Context, in *GetVariableRequest, opts ...grpc.CallOption) (*Variable, error) {
	out := new(Variable)
	err := c.cc.Invoke(ctx, "/networkApi.VariableService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableServiceClient) List(ctx context.Context, in *GetVariablesRequest, opts ...grpc.CallOption) (*Variables, error) {
	out := new(Variables)
	err := c.cc.Invoke(ctx, "/networkApi.VariableService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VariableServiceServer is the server API for VariableService service.
// All implementations should embed UnimplementedVariableServiceServer
// for forward compatibility
type VariableServiceServer interface {
	Set(context.Context, *SetVariableRequest) (*api.Void, error)
	Get(context.Context, *GetVariableRequest) (*Variable, error)
	List(context.Context, *GetVariablesRequest) (*Variables, error)
}

// UnimplementedVariableServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVariableServiceServer struct {
}

func (UnimplementedVariableServiceServer) Set(context.Context, *SetVariableRequest) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedVariableServiceServer) Get(context.Context, *GetVariableRequest) (*Variable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedVariableServiceServer) List(context.Context, *GetVariablesRequest) (*Variables, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeVariableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VariableServiceServer will
// result in compilation errors.
type UnsafeVariableServiceServer interface {
	mustEmbedUnimplementedVariableServiceServer()
}

func RegisterVariableServiceServer(s grpc.ServiceRegistrar, srv VariableServiceServer) {
	s.RegisterService(&VariableService_ServiceDesc, srv)
}

func _VariableService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkApi.VariableService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).Set(ctx, req.(*SetVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkApi.VariableService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).Get(ctx, req.(*GetVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VariableService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkApi.VariableService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServiceServer).List(ctx, req.(*GetVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VariableService_ServiceDesc is the grpc.ServiceDesc for VariableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VariableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networkApi.VariableService",
	HandlerType: (*VariableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _VariableService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _VariableService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VariableService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/ganjeh/api/src/variable.proto",
}
