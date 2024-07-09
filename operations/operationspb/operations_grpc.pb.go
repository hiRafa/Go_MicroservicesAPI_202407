// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: operations.proto

package operationspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OperationsService_GetOperationsById_FullMethodName = "/operations.OperationsService/GetOperationsById"
	OperationsService_CreateOperations_FullMethodName  = "/operations.OperationsService/CreateOperations"
)

// OperationsServiceClient is the client API for OperationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsServiceClient interface {
	GetOperationsById(ctx context.Context, in *OperationsByIdRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
	CreateOperations(ctx context.Context, in *CreateOperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
}

type operationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsServiceClient(cc grpc.ClientConnInterface) OperationsServiceClient {
	return &operationsServiceClient{cc}
}

func (c *operationsServiceClient) GetOperationsById(ctx context.Context, in *OperationsByIdRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, OperationsService_GetOperationsById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) CreateOperations(ctx context.Context, in *CreateOperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, OperationsService_CreateOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServiceServer is the server API for OperationsService service.
// All implementations must embed UnimplementedOperationsServiceServer
// for forward compatibility
type OperationsServiceServer interface {
	GetOperationsById(context.Context, *OperationsByIdRequest) (*OperationsResponse, error)
	CreateOperations(context.Context, *CreateOperationsRequest) (*OperationsResponse, error)
	mustEmbedUnimplementedOperationsServiceServer()
}

// UnimplementedOperationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsServiceServer struct {
}

func (UnimplementedOperationsServiceServer) GetOperationsById(context.Context, *OperationsByIdRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperationsById not implemented")
}
func (UnimplementedOperationsServiceServer) CreateOperations(context.Context, *CreateOperationsRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOperations not implemented")
}
func (UnimplementedOperationsServiceServer) mustEmbedUnimplementedOperationsServiceServer() {}

// UnsafeOperationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsServiceServer will
// result in compilation errors.
type UnsafeOperationsServiceServer interface {
	mustEmbedUnimplementedOperationsServiceServer()
}

func RegisterOperationsServiceServer(s grpc.ServiceRegistrar, srv OperationsServiceServer) {
	s.RegisterService(&OperationsService_ServiceDesc, srv)
}

func _OperationsService_GetOperationsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationsByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetOperationsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationsService_GetOperationsById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetOperationsById(ctx, req.(*OperationsByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_CreateOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).CreateOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationsService_CreateOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).CreateOperations(ctx, req.(*CreateOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationsService_ServiceDesc is the grpc.ServiceDesc for OperationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "operations.OperationsService",
	HandlerType: (*OperationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperationsById",
			Handler:    _OperationsService_GetOperationsById_Handler,
		},
		{
			MethodName: "CreateOperations",
			Handler:    _OperationsService_CreateOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations.proto",
}