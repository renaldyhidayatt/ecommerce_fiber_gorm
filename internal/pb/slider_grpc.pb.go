// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: slider.proto

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
	SliderService_GetSliders_FullMethodName   = "/pb.SliderService/GetSliders"
	SliderService_GetSlider_FullMethodName    = "/pb.SliderService/GetSlider"
	SliderService_CreateSlider_FullMethodName = "/pb.SliderService/CreateSlider"
	SliderService_UpdateSlider_FullMethodName = "/pb.SliderService/UpdateSlider"
	SliderService_DeleteSlider_FullMethodName = "/pb.SliderService/DeleteSlider"
)

// SliderServiceClient is the client API for SliderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SliderServiceClient interface {
	GetSliders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SlidersResponse, error)
	GetSlider(ctx context.Context, in *SliderRequest, opts ...grpc.CallOption) (*SliderResponse, error)
	CreateSlider(ctx context.Context, in *CreateSliderRequest, opts ...grpc.CallOption) (*SliderResponse, error)
	UpdateSlider(ctx context.Context, in *UpdateSliderRequest, opts ...grpc.CallOption) (*SliderResponse, error)
	DeleteSlider(ctx context.Context, in *SliderRequest, opts ...grpc.CallOption) (*DeleteSliderResponse, error)
}

type sliderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSliderServiceClient(cc grpc.ClientConnInterface) SliderServiceClient {
	return &sliderServiceClient{cc}
}

func (c *sliderServiceClient) GetSliders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SlidersResponse, error) {
	out := new(SlidersResponse)
	err := c.cc.Invoke(ctx, SliderService_GetSliders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliderServiceClient) GetSlider(ctx context.Context, in *SliderRequest, opts ...grpc.CallOption) (*SliderResponse, error) {
	out := new(SliderResponse)
	err := c.cc.Invoke(ctx, SliderService_GetSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliderServiceClient) CreateSlider(ctx context.Context, in *CreateSliderRequest, opts ...grpc.CallOption) (*SliderResponse, error) {
	out := new(SliderResponse)
	err := c.cc.Invoke(ctx, SliderService_CreateSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliderServiceClient) UpdateSlider(ctx context.Context, in *UpdateSliderRequest, opts ...grpc.CallOption) (*SliderResponse, error) {
	out := new(SliderResponse)
	err := c.cc.Invoke(ctx, SliderService_UpdateSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliderServiceClient) DeleteSlider(ctx context.Context, in *SliderRequest, opts ...grpc.CallOption) (*DeleteSliderResponse, error) {
	out := new(DeleteSliderResponse)
	err := c.cc.Invoke(ctx, SliderService_DeleteSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliderServiceServer is the server API for SliderService service.
// All implementations must embed UnimplementedSliderServiceServer
// for forward compatibility
type SliderServiceServer interface {
	GetSliders(context.Context, *empty.Empty) (*SlidersResponse, error)
	GetSlider(context.Context, *SliderRequest) (*SliderResponse, error)
	CreateSlider(context.Context, *CreateSliderRequest) (*SliderResponse, error)
	UpdateSlider(context.Context, *UpdateSliderRequest) (*SliderResponse, error)
	DeleteSlider(context.Context, *SliderRequest) (*DeleteSliderResponse, error)
	mustEmbedUnimplementedSliderServiceServer()
}

// UnimplementedSliderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSliderServiceServer struct {
}

func (UnimplementedSliderServiceServer) GetSliders(context.Context, *empty.Empty) (*SlidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliders not implemented")
}
func (UnimplementedSliderServiceServer) GetSlider(context.Context, *SliderRequest) (*SliderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlider not implemented")
}
func (UnimplementedSliderServiceServer) CreateSlider(context.Context, *CreateSliderRequest) (*SliderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlider not implemented")
}
func (UnimplementedSliderServiceServer) UpdateSlider(context.Context, *UpdateSliderRequest) (*SliderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlider not implemented")
}
func (UnimplementedSliderServiceServer) DeleteSlider(context.Context, *SliderRequest) (*DeleteSliderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSlider not implemented")
}
func (UnimplementedSliderServiceServer) mustEmbedUnimplementedSliderServiceServer() {}

// UnsafeSliderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SliderServiceServer will
// result in compilation errors.
type UnsafeSliderServiceServer interface {
	mustEmbedUnimplementedSliderServiceServer()
}

func RegisterSliderServiceServer(s grpc.ServiceRegistrar, srv SliderServiceServer) {
	s.RegisterService(&SliderService_ServiceDesc, srv)
}

func _SliderService_GetSliders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServiceServer).GetSliders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliderService_GetSliders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServiceServer).GetSliders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliderService_GetSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServiceServer).GetSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliderService_GetSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServiceServer).GetSlider(ctx, req.(*SliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliderService_CreateSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServiceServer).CreateSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliderService_CreateSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServiceServer).CreateSlider(ctx, req.(*CreateSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliderService_UpdateSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServiceServer).UpdateSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliderService_UpdateSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServiceServer).UpdateSlider(ctx, req.(*UpdateSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliderService_DeleteSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServiceServer).DeleteSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliderService_DeleteSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServiceServer).DeleteSlider(ctx, req.(*SliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SliderService_ServiceDesc is the grpc.ServiceDesc for SliderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SliderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SliderService",
	HandlerType: (*SliderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSliders",
			Handler:    _SliderService_GetSliders_Handler,
		},
		{
			MethodName: "GetSlider",
			Handler:    _SliderService_GetSlider_Handler,
		},
		{
			MethodName: "CreateSlider",
			Handler:    _SliderService_CreateSlider_Handler,
		},
		{
			MethodName: "UpdateSlider",
			Handler:    _SliderService_UpdateSlider_Handler,
		},
		{
			MethodName: "DeleteSlider",
			Handler:    _SliderService_DeleteSlider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slider.proto",
}