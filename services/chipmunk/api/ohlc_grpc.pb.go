// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: services/chipmunk/api/src/ohlc.proto

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

// OhlcServiceClient is the client API for OhlcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OhlcServiceClient interface {
	AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*api.Void, error)
	ReturnCandles(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*api.Candles, error)
	ReturnLastCandle(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*api.Candle, error)
	CancelWorker(ctx context.Context, in *CancelWorkerRequest, opts ...grpc.CallOption) (*api.Void, error)
}

type ohlcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOhlcServiceClient(cc grpc.ClientConnInterface) OhlcServiceClient {
	return &ohlcServiceClient{cc}
}

func (c *ohlcServiceClient) AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ohlcService/AddMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ohlcServiceClient) ReturnCandles(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*api.Candles, error) {
	out := new(api.Candles)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ohlcService/ReturnCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ohlcServiceClient) ReturnLastCandle(ctx context.Context, in *BufferedCandlesRequest, opts ...grpc.CallOption) (*api.Candle, error) {
	out := new(api.Candle)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ohlcService/ReturnLastCandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ohlcServiceClient) CancelWorker(ctx context.Context, in *CancelWorkerRequest, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/chipmunkApi.ohlcService/CancelWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OhlcServiceServer is the server API for OhlcService service.
// All implementations should embed UnimplementedOhlcServiceServer
// for forward compatibility
type OhlcServiceServer interface {
	AddMarket(context.Context, *AddMarketRequest) (*api.Void, error)
	ReturnCandles(context.Context, *BufferedCandlesRequest) (*api.Candles, error)
	ReturnLastCandle(context.Context, *BufferedCandlesRequest) (*api.Candle, error)
	CancelWorker(context.Context, *CancelWorkerRequest) (*api.Void, error)
}

// UnimplementedOhlcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOhlcServiceServer struct {
}

func (UnimplementedOhlcServiceServer) AddMarket(context.Context, *AddMarketRequest) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMarket not implemented")
}
func (UnimplementedOhlcServiceServer) ReturnCandles(context.Context, *BufferedCandlesRequest) (*api.Candles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnCandles not implemented")
}
func (UnimplementedOhlcServiceServer) ReturnLastCandle(context.Context, *BufferedCandlesRequest) (*api.Candle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnLastCandle not implemented")
}
func (UnimplementedOhlcServiceServer) CancelWorker(context.Context, *CancelWorkerRequest) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelWorker not implemented")
}

// UnsafeOhlcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OhlcServiceServer will
// result in compilation errors.
type UnsafeOhlcServiceServer interface {
	mustEmbedUnimplementedOhlcServiceServer()
}

func RegisterOhlcServiceServer(s grpc.ServiceRegistrar, srv OhlcServiceServer) {
	s.RegisterService(&OhlcService_ServiceDesc, srv)
}

func _OhlcService_AddMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OhlcServiceServer).AddMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ohlcService/AddMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OhlcServiceServer).AddMarket(ctx, req.(*AddMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OhlcService_ReturnCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferedCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OhlcServiceServer).ReturnCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ohlcService/ReturnCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OhlcServiceServer).ReturnCandles(ctx, req.(*BufferedCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OhlcService_ReturnLastCandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferedCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OhlcServiceServer).ReturnLastCandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ohlcService/ReturnLastCandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OhlcServiceServer).ReturnLastCandle(ctx, req.(*BufferedCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OhlcService_CancelWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OhlcServiceServer).CancelWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.ohlcService/CancelWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OhlcServiceServer).CancelWorker(ctx, req.(*CancelWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OhlcService_ServiceDesc is the grpc.ServiceDesc for OhlcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OhlcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chipmunkApi.ohlcService",
	HandlerType: (*OhlcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMarket",
			Handler:    _OhlcService_AddMarket_Handler,
		},
		{
			MethodName: "ReturnCandles",
			Handler:    _OhlcService_ReturnCandles_Handler,
		},
		{
			MethodName: "ReturnLastCandle",
			Handler:    _OhlcService_ReturnLastCandle_Handler,
		},
		{
			MethodName: "CancelWorker",
			Handler:    _OhlcService_CancelWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/src/ohlc.proto",
}
