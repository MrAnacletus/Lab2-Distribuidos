// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// LiderServiceClient is the client API for LiderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	RequestPozo(ctx context.Context, in *RequestPozoActual, opts ...grpc.CallOption) (*ResponsePozoActual, error)
	SendJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Resultado, error)
	SendJugada2(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Resultado, error)
	SendJugada3(ctx context.Context, in *Jugada3, opts ...grpc.CallOption) (*Resultado, error)
}

type liderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderServiceClient(cc grpc.ClientConnInterface) LiderServiceClient {
	return &liderServiceClient{cc}
}

func (c *liderServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/grpc.LiderService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) RequestPozo(ctx context.Context, in *RequestPozoActual, opts ...grpc.CallOption) (*ResponsePozoActual, error) {
	out := new(ResponsePozoActual)
	err := c.cc.Invoke(ctx, "/grpc.LiderService/RequestPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) SendJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Resultado, error) {
	out := new(Resultado)
	err := c.cc.Invoke(ctx, "/grpc.LiderService/SendJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) SendJugada2(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Resultado, error) {
	out := new(Resultado)
	err := c.cc.Invoke(ctx, "/grpc.LiderService/SendJugada2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) SendJugada3(ctx context.Context, in *Jugada3, opts ...grpc.CallOption) (*Resultado, error) {
	out := new(Resultado)
	err := c.cc.Invoke(ctx, "/grpc.LiderService/SendJugada3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderServiceServer is the server API for LiderService service.
// All implementations must embed UnimplementedLiderServiceServer
// for forward compatibility
type LiderServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	RequestPozo(context.Context, *RequestPozoActual) (*ResponsePozoActual, error)
	SendJugada(context.Context, *Jugada) (*Resultado, error)
	SendJugada2(context.Context, *Jugada) (*Resultado, error)
	SendJugada3(context.Context, *Jugada3) (*Resultado, error)
	mustEmbedUnimplementedLiderServiceServer()
}

// UnimplementedLiderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiderServiceServer struct {
}

func (UnimplementedLiderServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedLiderServiceServer) RequestPozo(context.Context, *RequestPozoActual) (*ResponsePozoActual, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPozo not implemented")
}
func (UnimplementedLiderServiceServer) SendJugada(context.Context, *Jugada) (*Resultado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugada not implemented")
}
func (UnimplementedLiderServiceServer) SendJugada2(context.Context, *Jugada) (*Resultado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugada2 not implemented")
}
func (UnimplementedLiderServiceServer) SendJugada3(context.Context, *Jugada3) (*Resultado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugada3 not implemented")
}
func (UnimplementedLiderServiceServer) mustEmbedUnimplementedLiderServiceServer() {}

// UnsafeLiderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderServiceServer will
// result in compilation errors.
type UnsafeLiderServiceServer interface {
	mustEmbedUnimplementedLiderServiceServer()
}

func RegisterLiderServiceServer(s grpc.ServiceRegistrar, srv LiderServiceServer) {
	s.RegisterService(&LiderService_ServiceDesc, srv)
}

func _LiderService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_RequestPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPozoActual)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).RequestPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderService/RequestPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).RequestPozo(ctx, req.(*RequestPozoActual))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_SendJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).SendJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderService/SendJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).SendJugada(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_SendJugada2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).SendJugada2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderService/SendJugada2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).SendJugada2(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_SendJugada3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada3)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).SendJugada3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LiderService/SendJugada3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).SendJugada3(ctx, req.(*Jugada3))
	}
	return interceptor(ctx, in, info, handler)
}

// LiderService_ServiceDesc is the grpc.ServiceDesc for LiderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.LiderService",
	HandlerType: (*LiderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _LiderService_SayHello_Handler,
		},
		{
			MethodName: "RequestPozo",
			Handler:    _LiderService_RequestPozo_Handler,
		},
		{
			MethodName: "SendJugada",
			Handler:    _LiderService_SendJugada_Handler,
		},
		{
			MethodName: "SendJugada2",
			Handler:    _LiderService_SendJugada2_Handler,
		},
		{
			MethodName: "SendJugada3",
			Handler:    _LiderService_SendJugada3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comunicacion.proto",
}

// NameNodeServiceClient is the client API for NameNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeServiceClient interface {
	SendJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*HelloReply, error)
	GetJugadas(ctx context.Context, in *RequestJugadas, opts ...grpc.CallOption) (*HelloReply, error)
}

type nameNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeServiceClient(cc grpc.ClientConnInterface) NameNodeServiceClient {
	return &nameNodeServiceClient{cc}
}

func (c *nameNodeServiceClient) SendJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/grpc.NameNodeService/SendJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) GetJugadas(ctx context.Context, in *RequestJugadas, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/grpc.NameNodeService/GetJugadas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServiceServer is the server API for NameNodeService service.
// All implementations must embed UnimplementedNameNodeServiceServer
// for forward compatibility
type NameNodeServiceServer interface {
	SendJugada(context.Context, *Jugada) (*HelloReply, error)
	GetJugadas(context.Context, *RequestJugadas) (*HelloReply, error)
	mustEmbedUnimplementedNameNodeServiceServer()
}

// UnimplementedNameNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServiceServer struct {
}

func (UnimplementedNameNodeServiceServer) SendJugada(context.Context, *Jugada) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugada not implemented")
}
func (UnimplementedNameNodeServiceServer) GetJugadas(context.Context, *RequestJugadas) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJugadas not implemented")
}
func (UnimplementedNameNodeServiceServer) mustEmbedUnimplementedNameNodeServiceServer() {}

// UnsafeNameNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServiceServer will
// result in compilation errors.
type UnsafeNameNodeServiceServer interface {
	mustEmbedUnimplementedNameNodeServiceServer()
}

func RegisterNameNodeServiceServer(s grpc.ServiceRegistrar, srv NameNodeServiceServer) {
	s.RegisterService(&NameNodeService_ServiceDesc, srv)
}

func _NameNodeService_SendJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).SendJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NameNodeService/SendJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).SendJugada(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_GetJugadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestJugadas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).GetJugadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NameNodeService/GetJugadas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).GetJugadas(ctx, req.(*RequestJugadas))
	}
	return interceptor(ctx, in, info, handler)
}

// NameNodeService_ServiceDesc is the grpc.ServiceDesc for NameNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.NameNodeService",
	HandlerType: (*NameNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendJugada",
			Handler:    _NameNodeService_SendJugada_Handler,
		},
		{
			MethodName: "GetJugadas",
			Handler:    _NameNodeService_GetJugadas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comunicacion.proto",
}

// PozoServiceClient is the client API for PozoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PozoServiceClient interface {
	RequestPozo(ctx context.Context, in *RequestPozoActual, opts ...grpc.CallOption) (*ResponsePozoActual, error)
}

type pozoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPozoServiceClient(cc grpc.ClientConnInterface) PozoServiceClient {
	return &pozoServiceClient{cc}
}

func (c *pozoServiceClient) RequestPozo(ctx context.Context, in *RequestPozoActual, opts ...grpc.CallOption) (*ResponsePozoActual, error) {
	out := new(ResponsePozoActual)
	err := c.cc.Invoke(ctx, "/grpc.PozoService/RequestPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PozoServiceServer is the server API for PozoService service.
// All implementations must embed UnimplementedPozoServiceServer
// for forward compatibility
type PozoServiceServer interface {
	RequestPozo(context.Context, *RequestPozoActual) (*ResponsePozoActual, error)
	mustEmbedUnimplementedPozoServiceServer()
}

// UnimplementedPozoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPozoServiceServer struct {
}

func (UnimplementedPozoServiceServer) RequestPozo(context.Context, *RequestPozoActual) (*ResponsePozoActual, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPozo not implemented")
}
func (UnimplementedPozoServiceServer) mustEmbedUnimplementedPozoServiceServer() {}

// UnsafePozoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PozoServiceServer will
// result in compilation errors.
type UnsafePozoServiceServer interface {
	mustEmbedUnimplementedPozoServiceServer()
}

func RegisterPozoServiceServer(s grpc.ServiceRegistrar, srv PozoServiceServer) {
	s.RegisterService(&PozoService_ServiceDesc, srv)
}

func _PozoService_RequestPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPozoActual)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PozoServiceServer).RequestPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.PozoService/RequestPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PozoServiceServer).RequestPozo(ctx, req.(*RequestPozoActual))
	}
	return interceptor(ctx, in, info, handler)
}

// PozoService_ServiceDesc is the grpc.ServiceDesc for PozoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PozoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PozoService",
	HandlerType: (*PozoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestPozo",
			Handler:    _PozoService_RequestPozo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comunicacion.proto",
}
