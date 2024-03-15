// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: shop.proto

package shop

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

const (
	Shop_CreateGoods_FullMethodName  = "/shop.Shop/CreateGoods"
	Shop_CreateNorm_FullMethodName   = "/shop.Shop/CreateNorm"
	Shop_CreateSku_FullMethodName    = "/shop.Shop/CreateSku"
	Shop_GetGoods_FullMethodName     = "/shop.Shop/GetGoods"
	Shop_GetGoodsInfo_FullMethodName = "/shop.Shop/GetGoodsInfo"
	Shop_DeleteGoods_FullMethodName  = "/shop.Shop/DeleteGoods"
)

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopClient interface {
	CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error)
	CreateNorm(ctx context.Context, in *CreateNormGoodsRequest, opts ...grpc.CallOption) (*CreateNormGoodsResponse, error)
	CreateSku(ctx context.Context, in *CreateSkuRequest, opts ...grpc.CallOption) (*CreateSkuResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoResponse, error)
	DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsResponse, error)
}

type shopClient struct {
	cc grpc.ClientConnInterface
}

func NewShopClient(cc grpc.ClientConnInterface) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error) {
	out := new(CreateGoodsResponse)
	err := c.cc.Invoke(ctx, Shop_CreateGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) CreateNorm(ctx context.Context, in *CreateNormGoodsRequest, opts ...grpc.CallOption) (*CreateNormGoodsResponse, error) {
	out := new(CreateNormGoodsResponse)
	err := c.cc.Invoke(ctx, Shop_CreateNorm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) CreateSku(ctx context.Context, in *CreateSkuRequest, opts ...grpc.CallOption) (*CreateSkuResponse, error) {
	out := new(CreateSkuResponse)
	err := c.cc.Invoke(ctx, Shop_CreateSku_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Shop_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoResponse, error) {
	out := new(GetGoodsInfoResponse)
	err := c.cc.Invoke(ctx, Shop_GetGoodsInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsResponse, error) {
	out := new(DeleteGoodsResponse)
	err := c.cc.Invoke(ctx, Shop_DeleteGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
// All implementations must embed UnimplementedShopServer
// for forward compatibility
type ShopServer interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error)
	CreateNorm(context.Context, *CreateNormGoodsRequest) (*CreateNormGoodsResponse, error)
	CreateSku(context.Context, *CreateSkuRequest) (*CreateSkuResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoResponse, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsResponse, error)
	mustEmbedUnimplementedShopServer()
}

// UnimplementedShopServer must be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (UnimplementedShopServer) CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedShopServer) CreateNorm(context.Context, *CreateNormGoodsRequest) (*CreateNormGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNorm not implemented")
}
func (UnimplementedShopServer) CreateSku(context.Context, *CreateSkuRequest) (*CreateSkuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSku not implemented")
}
func (UnimplementedShopServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedShopServer) GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsInfo not implemented")
}
func (UnimplementedShopServer) DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedShopServer) mustEmbedUnimplementedShopServer() {}

// UnsafeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServer will
// result in compilation errors.
type UnsafeShopServer interface {
	mustEmbedUnimplementedShopServer()
}

func RegisterShopServer(s grpc.ServiceRegistrar, srv ShopServer) {
	s.RegisterService(&Shop_ServiceDesc, srv)
}

func _Shop_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_CreateGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).CreateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_CreateNorm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNormGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).CreateNorm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_CreateNorm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).CreateNorm(ctx, req.(*CreateNormGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_CreateSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSkuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).CreateSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_CreateSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).CreateSku(ctx, req.(*CreateSkuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_GetGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_GetGoodsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetGoodsInfo(ctx, req.(*GetGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shop_DeleteGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).DeleteGoods(ctx, req.(*DeleteGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shop_ServiceDesc is the grpc.ServiceDesc for Shop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoods",
			Handler:    _Shop_CreateGoods_Handler,
		},
		{
			MethodName: "CreateNorm",
			Handler:    _Shop_CreateNorm_Handler,
		},
		{
			MethodName: "CreateSku",
			Handler:    _Shop_CreateSku_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Shop_GetGoods_Handler,
		},
		{
			MethodName: "GetGoodsInfo",
			Handler:    _Shop_GetGoodsInfo_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _Shop_DeleteGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
