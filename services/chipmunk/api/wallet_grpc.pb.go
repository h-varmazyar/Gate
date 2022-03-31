// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	api1 "github.com/h-varmazyar/Gate/api"
	api "github.com/h-varmazyar/Gate/services/brokerage/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletsServiceClient is the client API for WalletsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletsServiceClient interface {
	List(ctx context.Context, in *WalletListRequest, opts ...grpc.CallOption) (*api.Wallets, error)
	StartWorker(ctx context.Context, in *StartWorkerRequest, opts ...grpc.CallOption) (*api1.Void, error)
	CancelWorker(ctx context.Context, in *api1.Void, opts ...grpc.CallOption) (*api1.Void, error)
}

type walletsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletsServiceClient(cc grpc.ClientConnInterface) WalletsServiceClient {
	return &walletsServiceClient{cc}
}

func (c *walletsServiceClient) List(ctx context.Context, in *WalletListRequest, opts ...grpc.CallOption) (*api.Wallets, error) {
	out := new(api.Wallets)
	err := c.cc.Invoke(ctx, "/brokerageApi.WalletsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsServiceClient) StartWorker(ctx context.Context, in *StartWorkerRequest, opts ...grpc.CallOption) (*api1.Void, error) {
	out := new(api1.Void)
	err := c.cc.Invoke(ctx, "/brokerageApi.WalletsService/StartWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsServiceClient) CancelWorker(ctx context.Context, in *api1.Void, opts ...grpc.CallOption) (*api1.Void, error) {
	out := new(api1.Void)
	err := c.cc.Invoke(ctx, "/brokerageApi.WalletsService/CancelWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletsServiceServer is the server API for WalletsService service.
// All implementations should embed UnimplementedWalletsServiceServer
// for forward compatibility
type WalletsServiceServer interface {
	List(context.Context, *WalletListRequest) (*api.Wallets, error)
	StartWorker(context.Context, *StartWorkerRequest) (*api1.Void, error)
	CancelWorker(context.Context, *api1.Void) (*api1.Void, error)
}

// UnimplementedWalletsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWalletsServiceServer struct {
}

func (UnimplementedWalletsServiceServer) List(context.Context, *WalletListRequest) (*api.Wallets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWalletsServiceServer) StartWorker(context.Context, *StartWorkerRequest) (*api1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWorker not implemented")
}
func (UnimplementedWalletsServiceServer) CancelWorker(context.Context, *api1.Void) (*api1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelWorker not implemented")
}

// UnsafeWalletsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletsServiceServer will
// result in compilation errors.
type UnsafeWalletsServiceServer interface {
	mustEmbedUnimplementedWalletsServiceServer()
}

func RegisterWalletsServiceServer(s grpc.ServiceRegistrar, srv WalletsServiceServer) {
	s.RegisterService(&WalletsService_ServiceDesc, srv)
}

func _WalletsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.WalletsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServiceServer).List(ctx, req.(*WalletListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletsService_StartWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServiceServer).StartWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.WalletsService/StartWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServiceServer).StartWorker(ctx, req.(*StartWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletsService_CancelWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api1.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServiceServer).CancelWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.WalletsService/CancelWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServiceServer).CancelWorker(ctx, req.(*api1.Void))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletsService_ServiceDesc is the grpc.ServiceDesc for WalletsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brokerageApi.WalletsService",
	HandlerType: (*WalletsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _WalletsService_List_Handler,
		},
		{
			MethodName: "StartWorker",
			Handler:    _WalletsService_StartWorker_Handler,
		},
		{
			MethodName: "CancelWorker",
			Handler:    _WalletsService_CancelWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/src/wallet.proto",
}