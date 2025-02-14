// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: daemon/proto/meshnet/v1beta1/meshnet.proto

package v1beta1

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

// LocalClient is the client API for Local service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalClient interface {
	Get(ctx context.Context, in *PodQuery, opts ...grpc.CallOption) (*Pod, error)
	SetAlive(ctx context.Context, in *Pod, opts ...grpc.CallOption) (*BoolResponse, error)
	SkipReverse(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error)
	Skip(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error)
	IsSkipped(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error)
	GRPCWireExists(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*WireCreateResponse, error)
	AddGRPCWireLocal(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*BoolResponse, error)
	RemGRPCWire(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*BoolResponse, error)
	// A node is going to hold multiple veth to connect to multiple containers.
	// Each veth name must be unique with in a node. Daemon generates an ID that
	// is unique in this node.
	GenerateNodeInterfaceName(ctx context.Context, in *GenerateNodeInterfaceNameRequest, opts ...grpc.CallOption) (*GenerateNodeInterfaceNameResponse, error)
}

type localClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalClient(cc grpc.ClientConnInterface) LocalClient {
	return &localClient{cc}
}

func (c *localClient) Get(ctx context.Context, in *PodQuery, opts ...grpc.CallOption) (*Pod, error) {
	out := new(Pod)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) SetAlive(ctx context.Context, in *Pod, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/SetAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) SkipReverse(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/SkipReverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Skip(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/Skip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) IsSkipped(ctx context.Context, in *SkipQuery, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/IsSkipped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) GRPCWireExists(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*WireCreateResponse, error) {
	out := new(WireCreateResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/GRPCWireExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) AddGRPCWireLocal(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/AddGRPCWireLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) RemGRPCWire(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/RemGRPCWire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) GenerateNodeInterfaceName(ctx context.Context, in *GenerateNodeInterfaceNameRequest, opts ...grpc.CallOption) (*GenerateNodeInterfaceNameResponse, error) {
	out := new(GenerateNodeInterfaceNameResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Local/GenerateNodeInterfaceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalServer is the server API for Local service.
// All implementations must embed UnimplementedLocalServer
// for forward compatibility
type LocalServer interface {
	Get(context.Context, *PodQuery) (*Pod, error)
	SetAlive(context.Context, *Pod) (*BoolResponse, error)
	SkipReverse(context.Context, *SkipQuery) (*BoolResponse, error)
	Skip(context.Context, *SkipQuery) (*BoolResponse, error)
	IsSkipped(context.Context, *SkipQuery) (*BoolResponse, error)
	GRPCWireExists(context.Context, *WireDef) (*WireCreateResponse, error)
	AddGRPCWireLocal(context.Context, *WireDef) (*BoolResponse, error)
	RemGRPCWire(context.Context, *WireDef) (*BoolResponse, error)
	// A node is going to hold multiple veth to connect to multiple containers.
	// Each veth name must be unique with in a node. Daemon generates an ID that
	// is unique in this node.
	GenerateNodeInterfaceName(context.Context, *GenerateNodeInterfaceNameRequest) (*GenerateNodeInterfaceNameResponse, error)
	mustEmbedUnimplementedLocalServer()
}

// UnimplementedLocalServer must be embedded to have forward compatible implementations.
type UnimplementedLocalServer struct {
}

func (UnimplementedLocalServer) Get(context.Context, *PodQuery) (*Pod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLocalServer) SetAlive(context.Context, *Pod) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAlive not implemented")
}
func (UnimplementedLocalServer) SkipReverse(context.Context, *SkipQuery) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkipReverse not implemented")
}
func (UnimplementedLocalServer) Skip(context.Context, *SkipQuery) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Skip not implemented")
}
func (UnimplementedLocalServer) IsSkipped(context.Context, *SkipQuery) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSkipped not implemented")
}
func (UnimplementedLocalServer) GRPCWireExists(context.Context, *WireDef) (*WireCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GRPCWireExists not implemented")
}
func (UnimplementedLocalServer) AddGRPCWireLocal(context.Context, *WireDef) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGRPCWireLocal not implemented")
}
func (UnimplementedLocalServer) RemGRPCWire(context.Context, *WireDef) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemGRPCWire not implemented")
}
func (UnimplementedLocalServer) GenerateNodeInterfaceName(context.Context, *GenerateNodeInterfaceNameRequest) (*GenerateNodeInterfaceNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateNodeInterfaceName not implemented")
}
func (UnimplementedLocalServer) mustEmbedUnimplementedLocalServer() {}

// UnsafeLocalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalServer will
// result in compilation errors.
type UnsafeLocalServer interface {
	mustEmbedUnimplementedLocalServer()
}

func RegisterLocalServer(s grpc.ServiceRegistrar, srv LocalServer) {
	s.RegisterService(&Local_ServiceDesc, srv)
}

func _Local_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).Get(ctx, req.(*PodQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_SetAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).SetAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/SetAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).SetAlive(ctx, req.(*Pod))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_SkipReverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkipQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).SkipReverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/SkipReverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).SkipReverse(ctx, req.(*SkipQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_Skip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkipQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).Skip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/Skip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).Skip(ctx, req.(*SkipQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_IsSkipped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkipQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).IsSkipped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/IsSkipped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).IsSkipped(ctx, req.(*SkipQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_GRPCWireExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WireDef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).GRPCWireExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/GRPCWireExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).GRPCWireExists(ctx, req.(*WireDef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_AddGRPCWireLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WireDef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).AddGRPCWireLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/AddGRPCWireLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).AddGRPCWireLocal(ctx, req.(*WireDef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_RemGRPCWire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WireDef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).RemGRPCWire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/RemGRPCWire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).RemGRPCWire(ctx, req.(*WireDef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Local_GenerateNodeInterfaceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateNodeInterfaceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServer).GenerateNodeInterfaceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Local/GenerateNodeInterfaceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServer).GenerateNodeInterfaceName(ctx, req.(*GenerateNodeInterfaceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Local_ServiceDesc is the grpc.ServiceDesc for Local service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Local_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meshnet.v1beta1.Local",
	HandlerType: (*LocalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Local_Get_Handler,
		},
		{
			MethodName: "SetAlive",
			Handler:    _Local_SetAlive_Handler,
		},
		{
			MethodName: "SkipReverse",
			Handler:    _Local_SkipReverse_Handler,
		},
		{
			MethodName: "Skip",
			Handler:    _Local_Skip_Handler,
		},
		{
			MethodName: "IsSkipped",
			Handler:    _Local_IsSkipped_Handler,
		},
		{
			MethodName: "GRPCWireExists",
			Handler:    _Local_GRPCWireExists_Handler,
		},
		{
			MethodName: "AddGRPCWireLocal",
			Handler:    _Local_AddGRPCWireLocal_Handler,
		},
		{
			MethodName: "RemGRPCWire",
			Handler:    _Local_RemGRPCWire_Handler,
		},
		{
			MethodName: "GenerateNodeInterfaceName",
			Handler:    _Local_GenerateNodeInterfaceName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon/proto/meshnet/v1beta1/meshnet.proto",
}

// RemoteClient is the client API for Remote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteClient interface {
	Update(ctx context.Context, in *RemotePod, opts ...grpc.CallOption) (*BoolResponse, error)
	AddGRPCWireRemote(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*WireCreateResponse, error)
}

type remoteClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteClient(cc grpc.ClientConnInterface) RemoteClient {
	return &remoteClient{cc}
}

func (c *remoteClient) Update(ctx context.Context, in *RemotePod, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Remote/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteClient) AddGRPCWireRemote(ctx context.Context, in *WireDef, opts ...grpc.CallOption) (*WireCreateResponse, error) {
	out := new(WireCreateResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.Remote/AddGRPCWireRemote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteServer is the server API for Remote service.
// All implementations must embed UnimplementedRemoteServer
// for forward compatibility
type RemoteServer interface {
	Update(context.Context, *RemotePod) (*BoolResponse, error)
	AddGRPCWireRemote(context.Context, *WireDef) (*WireCreateResponse, error)
	mustEmbedUnimplementedRemoteServer()
}

// UnimplementedRemoteServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteServer struct {
}

func (UnimplementedRemoteServer) Update(context.Context, *RemotePod) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRemoteServer) AddGRPCWireRemote(context.Context, *WireDef) (*WireCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGRPCWireRemote not implemented")
}
func (UnimplementedRemoteServer) mustEmbedUnimplementedRemoteServer() {}

// UnsafeRemoteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteServer will
// result in compilation errors.
type UnsafeRemoteServer interface {
	mustEmbedUnimplementedRemoteServer()
}

func RegisterRemoteServer(s grpc.ServiceRegistrar, srv RemoteServer) {
	s.RegisterService(&Remote_ServiceDesc, srv)
}

func _Remote_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemotePod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Remote/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServer).Update(ctx, req.(*RemotePod))
	}
	return interceptor(ctx, in, info, handler)
}

func _Remote_AddGRPCWireRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WireDef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServer).AddGRPCWireRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.Remote/AddGRPCWireRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServer).AddGRPCWireRemote(ctx, req.(*WireDef))
	}
	return interceptor(ctx, in, info, handler)
}

// Remote_ServiceDesc is the grpc.ServiceDesc for Remote service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Remote_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meshnet.v1beta1.Remote",
	HandlerType: (*RemoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Remote_Update_Handler,
		},
		{
			MethodName: "AddGRPCWireRemote",
			Handler:    _Remote_AddGRPCWireRemote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon/proto/meshnet/v1beta1/meshnet.proto",
}

// WireProtocolClient is the client API for WireProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WireProtocolClient interface {
	SendToOnce(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*BoolResponse, error)
	SendToStream(ctx context.Context, opts ...grpc.CallOption) (WireProtocol_SendToStreamClient, error)
}

type wireProtocolClient struct {
	cc grpc.ClientConnInterface
}

func NewWireProtocolClient(cc grpc.ClientConnInterface) WireProtocolClient {
	return &wireProtocolClient{cc}
}

func (c *wireProtocolClient) SendToOnce(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/meshnet.v1beta1.WireProtocol/SendToOnce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireProtocolClient) SendToStream(ctx context.Context, opts ...grpc.CallOption) (WireProtocol_SendToStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WireProtocol_ServiceDesc.Streams[0], "/meshnet.v1beta1.WireProtocol/SendToStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &wireProtocolSendToStreamClient{stream}
	return x, nil
}

type WireProtocol_SendToStreamClient interface {
	Send(*Packet) error
	CloseAndRecv() (*BoolResponse, error)
	grpc.ClientStream
}

type wireProtocolSendToStreamClient struct {
	grpc.ClientStream
}

func (x *wireProtocolSendToStreamClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wireProtocolSendToStreamClient) CloseAndRecv() (*BoolResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BoolResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WireProtocolServer is the server API for WireProtocol service.
// All implementations must embed UnimplementedWireProtocolServer
// for forward compatibility
type WireProtocolServer interface {
	SendToOnce(context.Context, *Packet) (*BoolResponse, error)
	SendToStream(WireProtocol_SendToStreamServer) error
	mustEmbedUnimplementedWireProtocolServer()
}

// UnimplementedWireProtocolServer must be embedded to have forward compatible implementations.
type UnimplementedWireProtocolServer struct {
}

func (UnimplementedWireProtocolServer) SendToOnce(context.Context, *Packet) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToOnce not implemented")
}
func (UnimplementedWireProtocolServer) SendToStream(WireProtocol_SendToStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendToStream not implemented")
}
func (UnimplementedWireProtocolServer) mustEmbedUnimplementedWireProtocolServer() {}

// UnsafeWireProtocolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WireProtocolServer will
// result in compilation errors.
type UnsafeWireProtocolServer interface {
	mustEmbedUnimplementedWireProtocolServer()
}

func RegisterWireProtocolServer(s grpc.ServiceRegistrar, srv WireProtocolServer) {
	s.RegisterService(&WireProtocol_ServiceDesc, srv)
}

func _WireProtocol_SendToOnce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Packet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireProtocolServer).SendToOnce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshnet.v1beta1.WireProtocol/SendToOnce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireProtocolServer).SendToOnce(ctx, req.(*Packet))
	}
	return interceptor(ctx, in, info, handler)
}

func _WireProtocol_SendToStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WireProtocolServer).SendToStream(&wireProtocolSendToStreamServer{stream})
}

type WireProtocol_SendToStreamServer interface {
	SendAndClose(*BoolResponse) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type wireProtocolSendToStreamServer struct {
	grpc.ServerStream
}

func (x *wireProtocolSendToStreamServer) SendAndClose(m *BoolResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wireProtocolSendToStreamServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WireProtocol_ServiceDesc is the grpc.ServiceDesc for WireProtocol service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WireProtocol_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meshnet.v1beta1.WireProtocol",
	HandlerType: (*WireProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToOnce",
			Handler:    _WireProtocol_SendToOnce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendToStream",
			Handler:       _WireProtocol_SendToStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "daemon/proto/meshnet/v1beta1/meshnet.proto",
}
