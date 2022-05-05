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

// CandleServiceClient is the client API for CandleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandleServiceClient interface {
	AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*api.Void, error)
	ReturnLastNCandles(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*Candles, error)
	CancelWorker(ctx context.Context, in *CancelWorkerRequest, opts ...grpc.CallOption) (*api.Void, error)
}

type candleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCandleServiceClient(cc grpc.ClientConnInterface) CandleServiceClient {
	return &candleServiceClient{cc}
}

func (c *candleServiceClient) AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.CandleService/AddMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candleServiceClient) ReturnLastNCandles(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*Candles, error) {
	out := new(Candles)
	err := c.cc.Invoke(ctx, "/chipmunkApi.CandleService/ReturnLastNCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candleServiceClient) CancelWorker(ctx context.Context, in *CancelWorkerRequest, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.CandleService/CancelWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandleServiceServer is the server API for CandleService service.
// All implementations should embed UnimplementedCandleServiceServer
// for forward compatibility
type CandleServiceServer interface {
	AddMarket(context.Context, *AddMarketRequest) (*api.Void, error)
	ReturnLastNCandles(context.Context, *BufferedCandlesRequest) (*Candles, error)
	CancelWorker(context.Context, *CancelWorkerRequest) (*api.Void, error)
}

// UnimplementedCandleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCandleServiceServer struct {
}

func (UnimplementedCandleServiceServer) AddMarket(context.Context, *AddMarketRequest) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMarket not implemented")
}
func (UnimplementedCandleServiceServer) ReturnLastNCandles(context.Context, *BufferedCandlesRequest) (*Candles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnLastNCandles not implemented")
}
func (UnimplementedCandleServiceServer) CancelWorker(context.Context, *CancelWorkerRequest) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelWorker not implemented")
}

// UnsafeCandleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandleServiceServer will
// result in compilation errors.
type UnsafeCandleServiceServer interface {
	mustEmbedUnimplementedCandleServiceServer()
}

func RegisterCandleServiceServer(s grpc.ServiceRegistrar, srv CandleServiceServer) {
	s.RegisterService(&CandleService_ServiceDesc, srv)
}

func _CandleService_AddMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandleServiceServer).AddMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.CandleService/AddMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandleServiceServer).AddMarket(ctx, req.(*AddMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandleService_ReturnLastNCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferedCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandleServiceServer).ReturnLastNCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.CandleService/ReturnLastNCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandleServiceServer).ReturnLastNCandles(ctx, req.(*BufferedCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandleService_CancelWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandleServiceServer).CancelWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.CandleService/CancelWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandleServiceServer).CancelWorker(ctx, req.(*CancelWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CandleService_ServiceDesc is the grpc.ServiceDesc for CandleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CandleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chipmunkApi.CandleService",
	HandlerType: (*CandleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMarket",
			Handler:    _CandleService_AddMarket_Handler,
		},
		{
			MethodName: "ReturnLastNCandles",
			Handler:    _CandleService_ReturnLastNCandles_Handler,
		},
		{
			MethodName: "CancelWorker",
			Handler:    _CandleService_CancelWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/src/candle.proto",
}
