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

// IndicatorServiceClient is the client API for IndicatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndicatorServiceClient interface {
	Return(ctx context.Context, in *IndicatorReturnReq, opts ...grpc.CallOption) (*Indicator, error)
}

type indicatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndicatorServiceClient(cc grpc.ClientConnInterface) IndicatorServiceClient {
	return &indicatorServiceClient{cc}
}

func (c *indicatorServiceClient) Return(ctx context.Context, in *IndicatorReturnReq, opts ...grpc.CallOption) (*Indicator, error) {
	out := new(Indicator)
	err := c.cc.Invoke(ctx, "/chipmunkApi.indicatorService/Return", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndicatorServiceServer is the server API for IndicatorService service.
// All implementations should embed UnimplementedIndicatorServiceServer
// for forward compatibility
type IndicatorServiceServer interface {
	Return(context.Context, *IndicatorReturnReq) (*Indicator, error)
}

// UnimplementedIndicatorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIndicatorServiceServer struct {
}

func (UnimplementedIndicatorServiceServer) Return(context.Context, *IndicatorReturnReq) (*Indicator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}

// UnsafeIndicatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndicatorServiceServer will
// result in compilation errors.
type UnsafeIndicatorServiceServer interface {
	mustEmbedUnimplementedIndicatorServiceServer()
}

func RegisterIndicatorServiceServer(s grpc.ServiceRegistrar, srv IndicatorServiceServer) {
	s.RegisterService(&IndicatorService_ServiceDesc, srv)
}

func _IndicatorService_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndicatorReturnReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndicatorServiceServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chipmunkApi.indicatorService/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndicatorServiceServer).Return(ctx, req.(*IndicatorReturnReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IndicatorService_ServiceDesc is the grpc.ServiceDesc for IndicatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndicatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chipmunkApi.indicatorService",
	HandlerType: (*IndicatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Return",
			Handler:    _IndicatorService_Return_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chipmunk/api/src/indicator.proto",
}