// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// CreatePetServiceClient is the client API for CreatePetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreatePetServiceClient interface {
	CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetReply, error)
}

type createPetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreatePetServiceClient(cc grpc.ClientConnInterface) CreatePetServiceClient {
	return &createPetServiceClient{cc}
}

func (c *createPetServiceClient) CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetReply, error) {
	out := new(CreatePetReply)
	err := c.cc.Invoke(ctx, "/v1.CreatePetService/CreatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreatePetServiceServer is the server API for CreatePetService service.
// All implementations must embed UnimplementedCreatePetServiceServer
// for forward compatibility
type CreatePetServiceServer interface {
	CreatePet(context.Context, *CreatePetRequest) (*CreatePetReply, error)
	mustEmbedUnimplementedCreatePetServiceServer()
}

// UnimplementedCreatePetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreatePetServiceServer struct {
}

func (UnimplementedCreatePetServiceServer) CreatePet(context.Context, *CreatePetRequest) (*CreatePetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePet not implemented")
}
func (UnimplementedCreatePetServiceServer) mustEmbedUnimplementedCreatePetServiceServer() {}

// UnsafeCreatePetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreatePetServiceServer will
// result in compilation errors.
type UnsafeCreatePetServiceServer interface {
	mustEmbedUnimplementedCreatePetServiceServer()
}

func RegisterCreatePetServiceServer(s grpc.ServiceRegistrar, srv CreatePetServiceServer) {
	s.RegisterService(&CreatePetService_ServiceDesc, srv)
}

func _CreatePetService_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreatePetServiceServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CreatePetService/CreatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreatePetServiceServer).CreatePet(ctx, req.(*CreatePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreatePetService_ServiceDesc is the grpc.ServiceDesc for CreatePetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreatePetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CreatePetService",
	HandlerType: (*CreatePetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePet",
			Handler:    _CreatePetService_CreatePet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/challengeIskayPet.proto",
}

// GetStatisticsServiceClient is the client API for GetStatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetStatisticsServiceClient interface {
	GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsReply, error)
}

type getStatisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetStatisticsServiceClient(cc grpc.ClientConnInterface) GetStatisticsServiceClient {
	return &getStatisticsServiceClient{cc}
}

func (c *getStatisticsServiceClient) GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsReply, error) {
	out := new(GetStatisticsReply)
	err := c.cc.Invoke(ctx, "/v1.GetStatisticsService/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetStatisticsServiceServer is the server API for GetStatisticsService service.
// All implementations must embed UnimplementedGetStatisticsServiceServer
// for forward compatibility
type GetStatisticsServiceServer interface {
	GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsReply, error)
	mustEmbedUnimplementedGetStatisticsServiceServer()
}

// UnimplementedGetStatisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetStatisticsServiceServer struct {
}

func (UnimplementedGetStatisticsServiceServer) GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedGetStatisticsServiceServer) mustEmbedUnimplementedGetStatisticsServiceServer() {}

// UnsafeGetStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetStatisticsServiceServer will
// result in compilation errors.
type UnsafeGetStatisticsServiceServer interface {
	mustEmbedUnimplementedGetStatisticsServiceServer()
}

func RegisterGetStatisticsServiceServer(s grpc.ServiceRegistrar, srv GetStatisticsServiceServer) {
	s.RegisterService(&GetStatisticsService_ServiceDesc, srv)
}

func _GetStatisticsService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetStatisticsServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.GetStatisticsService/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetStatisticsServiceServer).GetStatistics(ctx, req.(*GetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetStatisticsService_ServiceDesc is the grpc.ServiceDesc for GetStatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetStatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GetStatisticsService",
	HandlerType: (*GetStatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatistics",
			Handler:    _GetStatisticsService_GetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/challengeIskayPet.proto",
}

// GetPetsServiceClient is the client API for GetPetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetPetsServiceClient interface {
	GetPets(ctx context.Context, in *GetPetsRequest, opts ...grpc.CallOption) (*GetPetsReply, error)
}

type getPetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetPetsServiceClient(cc grpc.ClientConnInterface) GetPetsServiceClient {
	return &getPetsServiceClient{cc}
}

func (c *getPetsServiceClient) GetPets(ctx context.Context, in *GetPetsRequest, opts ...grpc.CallOption) (*GetPetsReply, error) {
	out := new(GetPetsReply)
	err := c.cc.Invoke(ctx, "/v1.GetPetsService/GetPets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetPetsServiceServer is the server API for GetPetsService service.
// All implementations must embed UnimplementedGetPetsServiceServer
// for forward compatibility
type GetPetsServiceServer interface {
	GetPets(context.Context, *GetPetsRequest) (*GetPetsReply, error)
	mustEmbedUnimplementedGetPetsServiceServer()
}

// UnimplementedGetPetsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetPetsServiceServer struct {
}

func (UnimplementedGetPetsServiceServer) GetPets(context.Context, *GetPetsRequest) (*GetPetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPets not implemented")
}
func (UnimplementedGetPetsServiceServer) mustEmbedUnimplementedGetPetsServiceServer() {}

// UnsafeGetPetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetPetsServiceServer will
// result in compilation errors.
type UnsafeGetPetsServiceServer interface {
	mustEmbedUnimplementedGetPetsServiceServer()
}

func RegisterGetPetsServiceServer(s grpc.ServiceRegistrar, srv GetPetsServiceServer) {
	s.RegisterService(&GetPetsService_ServiceDesc, srv)
}

func _GetPetsService_GetPets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetPetsServiceServer).GetPets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.GetPetsService/GetPets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetPetsServiceServer).GetPets(ctx, req.(*GetPetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetPetsService_ServiceDesc is the grpc.ServiceDesc for GetPetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetPetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GetPetsService",
	HandlerType: (*GetPetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPets",
			Handler:    _GetPetsService_GetPets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/challengeIskayPet.proto",
}