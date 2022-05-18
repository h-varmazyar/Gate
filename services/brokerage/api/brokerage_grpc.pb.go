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

// BrokerageServiceClient is the client API for BrokerageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerageServiceClient interface {
	Create(ctx context.Context, in *CreateBrokerageReq, opts ...grpc.CallOption) (*Brokerage, error)
	Return(ctx context.Context, in *ReturnBrokerageReq, opts ...grpc.CallOption) (*Brokerage, error)
	List(ctx context.Context, in *api.Void, opts ...grpc.CallOption) (*Brokerages, error)
	Delete(ctx context.Context, in *DeleteBrokerageReq, opts ...grpc.CallOption) (*api.Void, error)
	ChangeStatus(ctx context.Context, in *StatusChangeRequest, opts ...grpc.CallOption) (*BrokerageStatus, error)
}

type brokerageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerageServiceClient(cc grpc.ClientConnInterface) BrokerageServiceClient {
	return &brokerageServiceClient{cc}
}

func (c *brokerageServiceClient) Create(ctx context.Context, in *CreateBrokerageReq, opts ...grpc.CallOption) (*Brokerage, error) {
	out := new(Brokerage)
	err := c.cc.Invoke(ctx, "/brokerageApi.BrokerageService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerageServiceClient) Return(ctx context.Context, in *ReturnBrokerageReq, opts ...grpc.CallOption) (*Brokerage, error) {
	out := new(Brokerage)
	err := c.cc.Invoke(ctx, "/brokerageApi.BrokerageService/Return", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerageServiceClient) List(ctx context.Context, in *api.Void, opts ...grpc.CallOption) (*Brokerages, error) {
	out := new(Brokerages)
	err := c.cc.Invoke(ctx, "/brokerageApi.BrokerageService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerageServiceClient) Delete(ctx context.Context, in *DeleteBrokerageReq, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/brokerageApi.BrokerageService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerageServiceClient) ChangeStatus(ctx context.Context, in *StatusChangeRequest, opts ...grpc.CallOption) (*BrokerageStatus, error) {
	out := new(BrokerageStatus)
	err := c.cc.Invoke(ctx, "/brokerageApi.BrokerageService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerageServiceServer is the server API for BrokerageService service.
// All implementations should embed UnimplementedBrokerageServiceServer
// for forward compatibility
type BrokerageServiceServer interface {
	Create(context.Context, *CreateBrokerageReq) (*Brokerage, error)
	Return(context.Context, *ReturnBrokerageReq) (*Brokerage, error)
	List(context.Context, *api.Void) (*Brokerages, error)
	Delete(context.Context, *DeleteBrokerageReq) (*api.Void, error)
	ChangeStatus(context.Context, *StatusChangeRequest) (*BrokerageStatus, error)
}

// UnimplementedBrokerageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBrokerageServiceServer struct {
}

func (UnimplementedBrokerageServiceServer) Create(context.Context, *CreateBrokerageReq) (*Brokerage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBrokerageServiceServer) Return(context.Context, *ReturnBrokerageReq) (*Brokerage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Return not implemented")
}
func (UnimplementedBrokerageServiceServer) List(context.Context, *api.Void) (*Brokerages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBrokerageServiceServer) Delete(context.Context, *DeleteBrokerageReq) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBrokerageServiceServer) ChangeStatus(context.Context, *StatusChangeRequest) (*BrokerageStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}

// UnsafeBrokerageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerageServiceServer will
// result in compilation errors.
type UnsafeBrokerageServiceServer interface {
	mustEmbedUnimplementedBrokerageServiceServer()
}

func RegisterBrokerageServiceServer(s grpc.ServiceRegistrar, srv BrokerageServiceServer) {
	s.RegisterService(&BrokerageService_ServiceDesc, srv)
}

func _BrokerageService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrokerageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerageServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.BrokerageService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerageServiceServer).Create(ctx, req.(*CreateBrokerageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerageService_Return_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnBrokerageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerageServiceServer).Return(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.BrokerageService/Return",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerageServiceServer).Return(ctx, req.(*ReturnBrokerageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerageService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerageServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.BrokerageService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerageServiceServer).List(ctx, req.(*api.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBrokerageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.BrokerageService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerageServiceServer).Delete(ctx, req.(*DeleteBrokerageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerageService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerageServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.BrokerageService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerageServiceServer).ChangeStatus(ctx, req.(*StatusChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerageService_ServiceDesc is the grpc.ServiceDesc for BrokerageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brokerageApi.BrokerageService",
	HandlerType: (*BrokerageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BrokerageService_Create_Handler,
		},
		{
			MethodName: "Return",
			Handler:    _BrokerageService_Return_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BrokerageService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BrokerageService_Delete_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _BrokerageService_ChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/brokerage/api/src/brokerage.proto",
}
