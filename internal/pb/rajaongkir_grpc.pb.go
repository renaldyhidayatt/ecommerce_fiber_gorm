// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: rajaongkir.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RajaOngkirService_GetProvinsi_FullMethodName = "/pb.RajaOngkirService/GetProvinsi"
	RajaOngkirService_GetCities_FullMethodName   = "/pb.RajaOngkirService/GetCities"
	RajaOngkirService_GetCost_FullMethodName     = "/pb.RajaOngkirService/GetCost"
)

// RajaOngkirServiceClient is the client API for RajaOngkirService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RajaOngkirServiceClient interface {
	GetProvinsi(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RajaOngkirResponseProvinsi, error)
	GetCities(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*RajaOngkirCityResponse, error)
	GetCost(ctx context.Context, in *OngkosRequest, opts ...grpc.CallOption) (*RajaOngkirOngkosResponse, error)
}

type rajaOngkirServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRajaOngkirServiceClient(cc grpc.ClientConnInterface) RajaOngkirServiceClient {
	return &rajaOngkirServiceClient{cc}
}

func (c *rajaOngkirServiceClient) GetProvinsi(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RajaOngkirResponseProvinsi, error) {
	out := new(RajaOngkirResponseProvinsi)
	err := c.cc.Invoke(ctx, RajaOngkirService_GetProvinsi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rajaOngkirServiceClient) GetCities(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*RajaOngkirCityResponse, error) {
	out := new(RajaOngkirCityResponse)
	err := c.cc.Invoke(ctx, RajaOngkirService_GetCities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rajaOngkirServiceClient) GetCost(ctx context.Context, in *OngkosRequest, opts ...grpc.CallOption) (*RajaOngkirOngkosResponse, error) {
	out := new(RajaOngkirOngkosResponse)
	err := c.cc.Invoke(ctx, RajaOngkirService_GetCost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RajaOngkirServiceServer is the server API for RajaOngkirService service.
// All implementations must embed UnimplementedRajaOngkirServiceServer
// for forward compatibility
type RajaOngkirServiceServer interface {
	GetProvinsi(context.Context, *empty.Empty) (*RajaOngkirResponseProvinsi, error)
	GetCities(context.Context, *CityRequest) (*RajaOngkirCityResponse, error)
	GetCost(context.Context, *OngkosRequest) (*RajaOngkirOngkosResponse, error)
	mustEmbedUnimplementedRajaOngkirServiceServer()
}

// UnimplementedRajaOngkirServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRajaOngkirServiceServer struct {
}

func (UnimplementedRajaOngkirServiceServer) GetProvinsi(context.Context, *empty.Empty) (*RajaOngkirResponseProvinsi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvinsi not implemented")
}
func (UnimplementedRajaOngkirServiceServer) GetCities(context.Context, *CityRequest) (*RajaOngkirCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCities not implemented")
}
func (UnimplementedRajaOngkirServiceServer) GetCost(context.Context, *OngkosRequest) (*RajaOngkirOngkosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCost not implemented")
}
func (UnimplementedRajaOngkirServiceServer) mustEmbedUnimplementedRajaOngkirServiceServer() {}

// UnsafeRajaOngkirServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RajaOngkirServiceServer will
// result in compilation errors.
type UnsafeRajaOngkirServiceServer interface {
	mustEmbedUnimplementedRajaOngkirServiceServer()
}

func RegisterRajaOngkirServiceServer(s grpc.ServiceRegistrar, srv RajaOngkirServiceServer) {
	s.RegisterService(&RajaOngkirService_ServiceDesc, srv)
}

func _RajaOngkirService_GetProvinsi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RajaOngkirServiceServer).GetProvinsi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RajaOngkirService_GetProvinsi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RajaOngkirServiceServer).GetProvinsi(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RajaOngkirService_GetCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RajaOngkirServiceServer).GetCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RajaOngkirService_GetCities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RajaOngkirServiceServer).GetCities(ctx, req.(*CityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RajaOngkirService_GetCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OngkosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RajaOngkirServiceServer).GetCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RajaOngkirService_GetCost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RajaOngkirServiceServer).GetCost(ctx, req.(*OngkosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RajaOngkirService_ServiceDesc is the grpc.ServiceDesc for RajaOngkirService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RajaOngkirService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RajaOngkirService",
	HandlerType: (*RajaOngkirServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProvinsi",
			Handler:    _RajaOngkirService_GetProvinsi_Handler,
		},
		{
			MethodName: "GetCities",
			Handler:    _RajaOngkirService_GetCities_Handler,
		},
		{
			MethodName: "GetCost",
			Handler:    _RajaOngkirService_GetCost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rajaongkir.proto",
}