// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package model

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

// SampleAPIClient is the client API for SampleAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleAPIClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	DeleteOrder(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	ListOrders(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Orders, error)
	UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	ReadOrder(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Order, error)
}

type sampleAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleAPIClient(cc grpc.ClientConnInterface) SampleAPIClient {
	return &sampleAPIClient{cc}
}

func (c *sampleAPIClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleAPIClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleAPIClient) DeleteOrder(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleAPIClient) ListOrders(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleAPIClient) UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleAPIClient) ReadOrder(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/model.SampleAPI/ReadOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleAPIServer is the server API for SampleAPI service.
// All implementations must embed UnimplementedSampleAPIServer
// for forward compatibility
type SampleAPIServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	CreateOrder(context.Context, *Order) (*Order, error)
	DeleteOrder(context.Context, *Id) (*Empty, error)
	ListOrders(context.Context, *Id) (*Orders, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
	ReadOrder(context.Context, *Id) (*Order, error)
	mustEmbedUnimplementedSampleAPIServer()
}

// UnimplementedSampleAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSampleAPIServer struct {
}

func (UnimplementedSampleAPIServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSampleAPIServer) CreateOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedSampleAPIServer) DeleteOrder(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedSampleAPIServer) ListOrders(context.Context, *Id) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedSampleAPIServer) UpdateOrder(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedSampleAPIServer) ReadOrder(context.Context, *Id) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOrder not implemented")
}
func (UnimplementedSampleAPIServer) mustEmbedUnimplementedSampleAPIServer() {}

// UnsafeSampleAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleAPIServer will
// result in compilation errors.
type UnsafeSampleAPIServer interface {
	mustEmbedUnimplementedSampleAPIServer()
}

func RegisterSampleAPIServer(s grpc.ServiceRegistrar, srv SampleAPIServer) {
	s.RegisterService(&SampleAPI_ServiceDesc, srv)
}

func _SampleAPI_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleAPI_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleAPI_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).DeleteOrder(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleAPI_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).ListOrders(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleAPI_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).UpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleAPI_ReadOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleAPIServer).ReadOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SampleAPI/ReadOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleAPIServer).ReadOrder(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleAPI_ServiceDesc is the grpc.ServiceDesc for SampleAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.SampleAPI",
	HandlerType: (*SampleAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SampleAPI_SayHello_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _SampleAPI_CreateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _SampleAPI_DeleteOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _SampleAPI_ListOrders_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _SampleAPI_UpdateOrder_Handler,
		},
		{
			MethodName: "ReadOrder",
			Handler:    _SampleAPI_ReadOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/sample.proto",
}
