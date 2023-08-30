// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/ent_server_api.proto

package ent

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

// EntClient is the client API for Ent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntClient interface {
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error)
	SetTag(ctx context.Context, in *SetTagRequest, opts ...grpc.CallOption) (*SetTagResponse, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (Ent_GetEntryClient, error)
	GetEntryMetadata(ctx context.Context, in *GetEntryMetadataRequest, opts ...grpc.CallOption) (*GetEntryMetadataResponse, error)
	PutEntry(ctx context.Context, opts ...grpc.CallOption) (Ent_PutEntryClient, error)
}

type entClient struct {
	cc grpc.ClientConnInterface
}

func NewEntClient(cc grpc.ClientConnInterface) EntClient {
	return &entClient{cc}
}

func (c *entClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error) {
	out := new(GetTagResponse)
	err := c.cc.Invoke(ctx, "/ent.server.api.Ent/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entClient) SetTag(ctx context.Context, in *SetTagRequest, opts ...grpc.CallOption) (*SetTagResponse, error) {
	out := new(SetTagResponse)
	err := c.cc.Invoke(ctx, "/ent.server.api.Ent/SetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (Ent_GetEntryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ent_ServiceDesc.Streams[0], "/ent.server.api.Ent/GetEntry", opts...)
	if err != nil {
		return nil, err
	}
	x := &entGetEntryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ent_GetEntryClient interface {
	Recv() (*GetEntryResponse, error)
	grpc.ClientStream
}

type entGetEntryClient struct {
	grpc.ClientStream
}

func (x *entGetEntryClient) Recv() (*GetEntryResponse, error) {
	m := new(GetEntryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *entClient) GetEntryMetadata(ctx context.Context, in *GetEntryMetadataRequest, opts ...grpc.CallOption) (*GetEntryMetadataResponse, error) {
	out := new(GetEntryMetadataResponse)
	err := c.cc.Invoke(ctx, "/ent.server.api.Ent/GetEntryMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entClient) PutEntry(ctx context.Context, opts ...grpc.CallOption) (Ent_PutEntryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ent_ServiceDesc.Streams[1], "/ent.server.api.Ent/PutEntry", opts...)
	if err != nil {
		return nil, err
	}
	x := &entPutEntryClient{stream}
	return x, nil
}

type Ent_PutEntryClient interface {
	Send(*PutEntryRequest) error
	CloseAndRecv() (*PutEntryResponse, error)
	grpc.ClientStream
}

type entPutEntryClient struct {
	grpc.ClientStream
}

func (x *entPutEntryClient) Send(m *PutEntryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *entPutEntryClient) CloseAndRecv() (*PutEntryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutEntryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EntServer is the server API for Ent service.
// All implementations must embed UnimplementedEntServer
// for forward compatibility
type EntServer interface {
	GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error)
	SetTag(context.Context, *SetTagRequest) (*SetTagResponse, error)
	GetEntry(*GetEntryRequest, Ent_GetEntryServer) error
	GetEntryMetadata(context.Context, *GetEntryMetadataRequest) (*GetEntryMetadataResponse, error)
	PutEntry(Ent_PutEntryServer) error
	mustEmbedUnimplementedEntServer()
}

// UnimplementedEntServer must be embedded to have forward compatible implementations.
type UnimplementedEntServer struct {
}

func (UnimplementedEntServer) GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedEntServer) SetTag(context.Context, *SetTagRequest) (*SetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTag not implemented")
}
func (UnimplementedEntServer) GetEntry(*GetEntryRequest, Ent_GetEntryServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedEntServer) GetEntryMetadata(context.Context, *GetEntryMetadataRequest) (*GetEntryMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntryMetadata not implemented")
}
func (UnimplementedEntServer) PutEntry(Ent_PutEntryServer) error {
	return status.Errorf(codes.Unimplemented, "method PutEntry not implemented")
}
func (UnimplementedEntServer) mustEmbedUnimplementedEntServer() {}

// UnsafeEntServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntServer will
// result in compilation errors.
type UnsafeEntServer interface {
	mustEmbedUnimplementedEntServer()
}

func RegisterEntServer(s grpc.ServiceRegistrar, srv EntServer) {
	s.RegisterService(&Ent_ServiceDesc, srv)
}

func _Ent_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ent.server.api.Ent/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ent_SetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntServer).SetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ent.server.api.Ent/SetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntServer).SetTag(ctx, req.(*SetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ent_GetEntry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEntryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntServer).GetEntry(m, &entGetEntryServer{stream})
}

type Ent_GetEntryServer interface {
	Send(*GetEntryResponse) error
	grpc.ServerStream
}

type entGetEntryServer struct {
	grpc.ServerStream
}

func (x *entGetEntryServer) Send(m *GetEntryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ent_GetEntryMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntServer).GetEntryMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ent.server.api.Ent/GetEntryMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntServer).GetEntryMetadata(ctx, req.(*GetEntryMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ent_PutEntry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EntServer).PutEntry(&entPutEntryServer{stream})
}

type Ent_PutEntryServer interface {
	SendAndClose(*PutEntryResponse) error
	Recv() (*PutEntryRequest, error)
	grpc.ServerStream
}

type entPutEntryServer struct {
	grpc.ServerStream
}

func (x *entPutEntryServer) SendAndClose(m *PutEntryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *entPutEntryServer) Recv() (*PutEntryRequest, error) {
	m := new(PutEntryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Ent_ServiceDesc is the grpc.ServiceDesc for Ent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ent.server.api.Ent",
	HandlerType: (*EntServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTag",
			Handler:    _Ent_GetTag_Handler,
		},
		{
			MethodName: "SetTag",
			Handler:    _Ent_SetTag_Handler,
		},
		{
			MethodName: "GetEntryMetadata",
			Handler:    _Ent_GetEntryMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEntry",
			Handler:       _Ent_GetEntry_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutEntry",
			Handler:       _Ent_PutEntry_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/ent_server_api.proto",
}
