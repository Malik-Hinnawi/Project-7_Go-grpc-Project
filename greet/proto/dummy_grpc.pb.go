// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: dummy.proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetManyTimes(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
	GreetWithDeadline(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetManyTimes(*GreetRequest, GreetService_GreetManyTimesServer) error
	LongGreet(GreetService_LongGreetServer) error
	GreetEveryone(GreetService_GreetEveryoneServer) error
	GreetWithDeadline(context.Context, *GreetRequest) (*GreetResponse, error)
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServiceServer) GreetManyTimes(*GreetRequest, GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (UnimplementedGreetServiceServer) LongGreet(GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}
func (UnimplementedGreetServiceServer) GreetEveryone(GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (UnimplementedGreetServiceServer) GreetWithDeadline(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetWithDeadline not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dummy.proto",
}
