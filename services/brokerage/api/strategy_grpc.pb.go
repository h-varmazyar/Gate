// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// StrategyServiceClient is the client API for StrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyServiceClient interface {
	Create(ctx context.Context, in *CreateStrategyReq, opts ...grpc.CallOption) (*Strategy, error)
	Return(ctx context.Context, in *ReturnStrategyReq, opts ...grpc.CallOption) (*Strategy, error)
}

type strategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyServiceClient(cc grpc.ClientConnInterface) StrategyServiceClient {
	return &strategyServiceClient{cc}
}

func (c *strategyServiceClient) Create(ctx context.Context, in *CreateStrategyReq, opts ...grpc.CallOption) (*Strategy, error) {
	out := new(Strategy)
	err := c.cc.Invoke(ctx, "/brokerageApi.StrategyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyServiceClient) Return(ctx context.Context, in *ReturnStrategyReq, opts ...grpc.CallOption) (*Strategy, error) {
	out := new(Strategy)
	err := c.cc.Invoke(ctx, "/brokerageApi.StrategyService/Return", in, out, opts...)
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
		FullMethod: "/brokerageApi.StrategyService/Create",
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
		FullMethod: "/brokerageApi.StrategyService/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServiceServer).Return(ctx, req.(*ReturnStrategyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StrategyService_ServiceDesc is the grpc.ServiceDesc for StrategyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StrategyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brokerageApi.StrategyService",
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/brokerage/api/src/strategy.proto",
}
