// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: app.proto

package app

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

// DictServiceClient is the client API for DictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictServiceClient interface {
	AddPair(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*StatusResponse, error)
	GetPairs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPairs, error)
}

type dictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictServiceClient(cc grpc.ClientConnInterface) DictServiceClient {
	return &dictServiceClient{cc}
}

func (c *dictServiceClient) AddPair(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/DictService/AddPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) GetPairs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPairs, error) {
	out := new(ListPairs)
	err := c.cc.Invoke(ctx, "/DictService/GetPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServiceServer is the server API for DictService service.
// All implementations must embed UnimplementedDictServiceServer
// for forward compatibility
type DictServiceServer interface {
	AddPair(context.Context, *KeyValuePair) (*StatusResponse, error)
	GetPairs(context.Context, *Empty) (*ListPairs, error)
	mustEmbedUnimplementedDictServiceServer()
}

// UnimplementedDictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictServiceServer struct {
}

func (UnimplementedDictServiceServer) AddPair(context.Context, *KeyValuePair) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPair not implemented")
}
func (UnimplementedDictServiceServer) GetPairs(context.Context, *Empty) (*ListPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPairs not implemented")
}
func (UnimplementedDictServiceServer) mustEmbedUnimplementedDictServiceServer() {}

// UnsafeDictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServiceServer will
// result in compilation errors.
type UnsafeDictServiceServer interface {
	mustEmbedUnimplementedDictServiceServer()
}

func RegisterDictServiceServer(s grpc.ServiceRegistrar, srv DictServiceServer) {
	s.RegisterService(&DictService_ServiceDesc, srv)
}

func _DictService_AddPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).AddPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DictService/AddPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).AddPair(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_GetPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).GetPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DictService/GetPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).GetPairs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DictService_ServiceDesc is the grpc.ServiceDesc for DictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DictService",
	HandlerType: (*DictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPair",
			Handler:    _DictService_AddPair_Handler,
		},
		{
			MethodName: "GetPairs",
			Handler:    _DictService_GetPairs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
