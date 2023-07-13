// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services/network/api/proto/src/request.proto

package proto

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

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	DoAsync(ctx context.Context, in *DoAsyncReq, opts ...grpc.CallOption) (*DoAsyncResp, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/networkAPI.RequestService/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) DoAsync(ctx context.Context, in *DoAsyncReq, opts ...grpc.CallOption) (*DoAsyncResp, error) {
	out := new(DoAsyncResp)
	err := c.cc.Invoke(ctx, "/networkAPI.RequestService/DoAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations should embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	Do(context.Context, *Request) (*Response, error)
	DoAsync(context.Context, *DoAsyncReq) (*DoAsyncResp, error)
}

// UnimplementedRequestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) Do(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedRequestServiceServer) DoAsync(context.Context, *DoAsyncReq) (*DoAsyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoAsync not implemented")
}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkAPI.RequestService/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).Do(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_DoAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoAsyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).DoAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkAPI.RequestService/DoAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).DoAsync(ctx, req.(*DoAsyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networkAPI.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _RequestService_Do_Handler,
		},
		{
			MethodName: "DoAsync",
			Handler:    _RequestService_DoAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/network/api/proto/src/request.proto",
}
