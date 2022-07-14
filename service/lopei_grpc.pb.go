// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: model/lopei.proto

package service

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

// LopeiPaymentClient is the client API for LopeiPayment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LopeiPaymentClient interface {
	CheckBalance(ctx context.Context, in *CheckBalanceMessage, opts ...grpc.CallOption) (*ResultMessage, error)
	DoPayment(ctx context.Context, in *PaymentMessage, opts ...grpc.CallOption) (*ResultMessage, error)
}

type lopeiPaymentClient struct {
	cc grpc.ClientConnInterface
}

func NewLopeiPaymentClient(cc grpc.ClientConnInterface) LopeiPaymentClient {
	return &lopeiPaymentClient{cc}
}

func (c *lopeiPaymentClient) CheckBalance(ctx context.Context, in *CheckBalanceMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.LopeiPayment/CheckBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lopeiPaymentClient) DoPayment(ctx context.Context, in *PaymentMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/api.LopeiPayment/DoPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LopeiPaymentServer is the server API for LopeiPayment service.
// All implementations must embed UnimplementedLopeiPaymentServer
// for forward compatibility
type LopeiPaymentServer interface {
	CheckBalance(context.Context, *CheckBalanceMessage) (*ResultMessage, error)
	DoPayment(context.Context, *PaymentMessage) (*ResultMessage, error)
	mustEmbedUnimplementedLopeiPaymentServer()
}

// UnimplementedLopeiPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedLopeiPaymentServer struct {
}

func (UnimplementedLopeiPaymentServer) CheckBalance(context.Context, *CheckBalanceMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBalance not implemented")
}
func (UnimplementedLopeiPaymentServer) DoPayment(context.Context, *PaymentMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoPayment not implemented")
}
func (UnimplementedLopeiPaymentServer) mustEmbedUnimplementedLopeiPaymentServer() {}

// UnsafeLopeiPaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LopeiPaymentServer will
// result in compilation errors.
type UnsafeLopeiPaymentServer interface {
	mustEmbedUnimplementedLopeiPaymentServer()
}

func RegisterLopeiPaymentServer(s grpc.ServiceRegistrar, srv LopeiPaymentServer) {
	s.RegisterService(&LopeiPayment_ServiceDesc, srv)
}

func _LopeiPayment_CheckBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBalanceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LopeiPaymentServer).CheckBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LopeiPayment/CheckBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LopeiPaymentServer).CheckBalance(ctx, req.(*CheckBalanceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _LopeiPayment_DoPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LopeiPaymentServer).DoPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LopeiPayment/DoPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LopeiPaymentServer).DoPayment(ctx, req.(*PaymentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// LopeiPayment_ServiceDesc is the grpc.ServiceDesc for LopeiPayment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LopeiPayment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.LopeiPayment",
	HandlerType: (*LopeiPaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckBalance",
			Handler:    _LopeiPayment_CheckBalance_Handler,
		},
		{
			MethodName: "DoPayment",
			Handler:    _LopeiPayment_DoPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/lopei.proto",
}
