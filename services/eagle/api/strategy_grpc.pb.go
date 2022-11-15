// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services/eagle/api/src/strategy.proto

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

// StrategyServiceClient is the client API for StrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyServiceClient interface {
	Create(ctx context.Context, in *CreateStrategyReq, opts ...grpc.CallOption) (*Strategy, error)
	Return(ctx context.Context, in *ReturnStrategyReq, opts ...grpc.CallOption) (*Strategy, error)
	List(ctx context.Context, in *api.Void, opts ...grpc.CallOption) (*Strategies, error)
	Indicators(ctx context.Context, in *StrategyIndicatorReq, opts ...grpc.CallOption) (*StrategyIndicators, error)
}

type strategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyServiceClient(cc grpc.ClientConnInterface) StrategyServiceClient {
	return &strategyServiceClient{cc}
}

func (c *strategyServiceClient) Create(ctx context.Context, in *CreateStrategyReq, opts ...grpc.CallOption) (*Strategy, error) {
	out := new(Strategy)
	err := c.cc.Invoke(ctx, "/eagleApi.StrategyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyServiceClient) Return(ctx context.Context, in *ReturnStrategyReq, opts ...grpc.CallOption) (*Strategy, error) {
	out := new(Strategy)
	err := c.cc.Invoke(ctx, "/eagleApi.StrategyService/Return", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyServiceClient) List(ctx context.Context, in *api.Void, opts ...grpc.CallOption) (*Strategies, error) {
	out := new(Strategies)
	err := c.cc.Invoke(ctx, "/eagleApi.StrategyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyServiceClient) Indicators(ctx context.Context, in *StrategyIndicatorReq, opts ...grpc.CallOption) (*StrategyIndicators, error) {
	out := new(StrategyIndicators)
	err := c.cc.Invoke(ctx, "/eagleApi.StrategyService/Indicators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyServiceServer is the server API for StrategyService service.
// All implementations should embed UnimplementedStrategyServiceServer
// for forward compatibility
type StrategyServiceServer interface {
	Create(context.Context, *CreateStrategyReq) (*Strategy, error)
	Return(context.Context, *ReturnStrategyReq) (*Strategy, error)
	List(context.Context, *api.Void) (*Strategies, error)
	Indicators(context.Context, *StrategyIndicatorReq) (*StrategyIndicators, error)
}

// UnimplementedStrategyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStrategyServiceServer struct {
}

func (UnimplementedStrategyServiceServer) Create(context.Context, *CreateStrategyReq) (*Strategy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStrategyServiceServer) Return(context.Context, *ReturnStrategyReq) (*Strategy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}
func (UnimplementedStrategyServiceServer) List(context.Context, *api.Void) (*Strategies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStrategyServiceServer) Indicators(context.Context, *StrategyIndicatorReq) (*StrategyIndicators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Indicators not implemented")
}

// UnsafeStrategyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyServiceServer will
// result in compilation errors.
type UnsafeStrategyServiceServer interface {
	mustEmbedUnimplementedStrategyServiceServer()
}

func RegisterStrategyServiceServer(s grpc.ServiceRegistrar, srv StrategyServiceServer) {
	s.RegisterService(&StrategyService_ServiceDesc, srv)
}

func _StrategyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStrategyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eagleApi.StrategyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServiceServer).Create(ctx, req.(*CreateStrategyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyService_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnStrategyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServiceServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eagleApi.StrategyService/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServiceServer).Return(ctx, req.(*ReturnStrategyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eagleApi.StrategyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServiceServer).List(ctx, req.(*api.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyService_Indicators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrategyIndicatorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServiceServer).Indicators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eagleApi.StrategyService/Indicators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServiceServer).Indicators(ctx, req.(*StrategyIndicatorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StrategyService_ServiceDesc is the grpc.ServiceDesc for StrategyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StrategyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eagleApi.StrategyService",
	HandlerType: (*StrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StrategyService_Create_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _StrategyService_Return_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StrategyService_List_Handler,
		},
		{
			MethodName: "Indicators",
			Handler:    _StrategyService_Indicators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/eagle/api/src/strategy.proto",
}
