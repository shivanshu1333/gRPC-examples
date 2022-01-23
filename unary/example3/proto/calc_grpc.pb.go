// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: unary/example3/proto/calc.proto

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

// CalServiceClient is the client API for CalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalServiceClient interface {
	AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type calServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalServiceClient(cc grpc.ClientConnInterface) CalServiceClient {
	return &calServiceClient{cc}
}

func (c *calServiceClient) AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/unary.CalService/AddNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalServiceServer is the server API for CalService service.
// All implementations must embed UnimplementedCalServiceServer
// for forward compatibility
type CalServiceServer interface {
	AddNumbers(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCalServiceServer()
}

// UnimplementedCalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalServiceServer struct {
}

func (UnimplementedCalServiceServer) AddNumbers(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumbers not implemented")
}
func (UnimplementedCalServiceServer) mustEmbedUnimplementedCalServiceServer() {}

// UnsafeCalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalServiceServer will
// result in compilation errors.
type UnsafeCalServiceServer interface {
	mustEmbedUnimplementedCalServiceServer()
}

func RegisterCalServiceServer(s grpc.ServiceRegistrar, srv CalServiceServer) {
	s.RegisterService(&CalService_ServiceDesc, srv)
}

func _CalService_AddNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalServiceServer).AddNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unary.CalService/AddNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalServiceServer).AddNumbers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CalService_ServiceDesc is the grpc.ServiceDesc for CalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "unary.CalService",
	HandlerType: (*CalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumbers",
			Handler:    _CalService_AddNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unary/example3/proto/calc.proto",
}
