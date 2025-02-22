// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pkg/proto/basket.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BasketService_New_FullMethodName           = "/BasketService/New"
	BasketService_AddProduct_FullMethodName    = "/BasketService/AddProduct"
	BasketService_Get_FullMethodName           = "/BasketService/Get"
	BasketService_DeleteProduct_FullMethodName = "/BasketService/DeleteProduct"
)

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasketServiceClient interface {
	New(ctx context.Context, in *NewBasketReq, opts ...grpc.CallOption) (*Response, error)
	AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetBasketReq, opts ...grpc.CallOption) (*GetBasketResp, error)
	DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*Response, error)
}

type basketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasketServiceClient(cc grpc.ClientConnInterface) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) New(ctx context.Context, in *NewBasketReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, BasketService_New_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, BasketService_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) Get(ctx context.Context, in *GetBasketReq, opts ...grpc.CallOption) (*GetBasketResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBasketResp)
	err := c.cc.Invoke(ctx, BasketService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, BasketService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
// All implementations must embed UnimplementedBasketServiceServer
// for forward compatibility.
type BasketServiceServer interface {
	New(context.Context, *NewBasketReq) (*Response, error)
	AddProduct(context.Context, *AddProductReq) (*Response, error)
	Get(context.Context, *GetBasketReq) (*GetBasketResp, error)
	DeleteProduct(context.Context, *DeleteProductReq) (*Response, error)
	mustEmbedUnimplementedBasketServiceServer()
}

// UnimplementedBasketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBasketServiceServer struct{}

func (UnimplementedBasketServiceServer) New(context.Context, *NewBasketReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedBasketServiceServer) AddProduct(context.Context, *AddProductReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedBasketServiceServer) Get(context.Context, *GetBasketReq) (*GetBasketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBasketServiceServer) DeleteProduct(context.Context, *DeleteProductReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedBasketServiceServer) mustEmbedUnimplementedBasketServiceServer() {}
func (UnimplementedBasketServiceServer) testEmbeddedByValue()                       {}

// UnsafeBasketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasketServiceServer will
// result in compilation errors.
type UnsafeBasketServiceServer interface {
	mustEmbedUnimplementedBasketServiceServer()
}

func RegisterBasketServiceServer(s grpc.ServiceRegistrar, srv BasketServiceServer) {
	// If the following call pancis, it indicates UnimplementedBasketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BasketService_ServiceDesc, srv)
}

func _BasketService_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBasketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_New_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).New(ctx, req.(*NewBasketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).AddProduct(ctx, req.(*AddProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).Get(ctx, req.(*GetBasketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).DeleteProduct(ctx, req.(*DeleteProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BasketService_ServiceDesc is the grpc.ServiceDesc for BasketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _BasketService_New_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _BasketService_AddProduct_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BasketService_Get_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _BasketService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/basket.proto",
}
