// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: kabkota.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KabKotaService_GetAllKabKota_FullMethodName        = "/tracer_study_grpc.KabKotaService/GetAllKabKota"
	KabKotaService_GetKabKotaByProvinsi_FullMethodName = "/tracer_study_grpc.KabKotaService/GetKabKotaByProvinsi"
	KabKotaService_GetKabKotaByIdWil_FullMethodName    = "/tracer_study_grpc.KabKotaService/GetKabKotaByIdWil"
	KabKotaService_CreateKabKota_FullMethodName        = "/tracer_study_grpc.KabKotaService/CreateKabKota"
	KabKotaService_UpdateKabKota_FullMethodName        = "/tracer_study_grpc.KabKotaService/UpdateKabKota"
	KabKotaService_DeleteKabKota_FullMethodName        = "/tracer_study_grpc.KabKotaService/DeleteKabKota"
)

// KabKotaServiceClient is the client API for KabKotaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KabKotaServiceClient interface {
	GetAllKabKota(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error)
	GetKabKotaByProvinsi(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error)
	GetKabKotaByIdWil(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*GetKabKotaResponse, error)
	CreateKabKota(ctx context.Context, in *KabKota, opts ...grpc.CallOption) (*GetKabKotaResponse, error)
	UpdateKabKota(ctx context.Context, in *KabKota, opts ...grpc.CallOption) (*GetKabKotaResponse, error)
	DeleteKabKota(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*DeleteKabKotaResponse, error)
}

type kabKotaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKabKotaServiceClient(cc grpc.ClientConnInterface) KabKotaServiceClient {
	return &kabKotaServiceClient{cc}
}

func (c *kabKotaServiceClient) GetAllKabKota(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error) {
	out := new(GetAllKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_GetAllKabKota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabKotaServiceClient) GetKabKotaByProvinsi(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error) {
	out := new(GetAllKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_GetKabKotaByProvinsi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabKotaServiceClient) GetKabKotaByIdWil(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*GetKabKotaResponse, error) {
	out := new(GetKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_GetKabKotaByIdWil_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabKotaServiceClient) CreateKabKota(ctx context.Context, in *KabKota, opts ...grpc.CallOption) (*GetKabKotaResponse, error) {
	out := new(GetKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_CreateKabKota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabKotaServiceClient) UpdateKabKota(ctx context.Context, in *KabKota, opts ...grpc.CallOption) (*GetKabKotaResponse, error) {
	out := new(GetKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_UpdateKabKota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kabKotaServiceClient) DeleteKabKota(ctx context.Context, in *GetKabKotaByIdWilRequest, opts ...grpc.CallOption) (*DeleteKabKotaResponse, error) {
	out := new(DeleteKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_DeleteKabKota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KabKotaServiceServer is the server API for KabKotaService service.
// All implementations must embed UnimplementedKabKotaServiceServer
// for forward compatibility
type KabKotaServiceServer interface {
	GetAllKabKota(context.Context, *emptypb.Empty) (*GetAllKabKotaResponse, error)
	GetKabKotaByProvinsi(context.Context, *GetKabKotaByIdWilRequest) (*GetAllKabKotaResponse, error)
	GetKabKotaByIdWil(context.Context, *GetKabKotaByIdWilRequest) (*GetKabKotaResponse, error)
	CreateKabKota(context.Context, *KabKota) (*GetKabKotaResponse, error)
	UpdateKabKota(context.Context, *KabKota) (*GetKabKotaResponse, error)
	DeleteKabKota(context.Context, *GetKabKotaByIdWilRequest) (*DeleteKabKotaResponse, error)
	mustEmbedUnimplementedKabKotaServiceServer()
}

// UnimplementedKabKotaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKabKotaServiceServer struct {
}

func (UnimplementedKabKotaServiceServer) GetAllKabKota(context.Context, *emptypb.Empty) (*GetAllKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllKabKota not implemented")
}
func (UnimplementedKabKotaServiceServer) GetKabKotaByProvinsi(context.Context, *GetKabKotaByIdWilRequest) (*GetAllKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKabKotaByProvinsi not implemented")
}
func (UnimplementedKabKotaServiceServer) GetKabKotaByIdWil(context.Context, *GetKabKotaByIdWilRequest) (*GetKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKabKotaByIdWil not implemented")
}
func (UnimplementedKabKotaServiceServer) CreateKabKota(context.Context, *KabKota) (*GetKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKabKota not implemented")
}
func (UnimplementedKabKotaServiceServer) UpdateKabKota(context.Context, *KabKota) (*GetKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKabKota not implemented")
}
func (UnimplementedKabKotaServiceServer) DeleteKabKota(context.Context, *GetKabKotaByIdWilRequest) (*DeleteKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKabKota not implemented")
}
func (UnimplementedKabKotaServiceServer) mustEmbedUnimplementedKabKotaServiceServer() {}

// UnsafeKabKotaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KabKotaServiceServer will
// result in compilation errors.
type UnsafeKabKotaServiceServer interface {
	mustEmbedUnimplementedKabKotaServiceServer()
}

func RegisterKabKotaServiceServer(s grpc.ServiceRegistrar, srv KabKotaServiceServer) {
	s.RegisterService(&KabKotaService_ServiceDesc, srv)
}

func _KabKotaService_GetAllKabKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).GetAllKabKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_GetAllKabKota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).GetAllKabKota(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabKotaService_GetKabKotaByProvinsi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKabKotaByIdWilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).GetKabKotaByProvinsi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_GetKabKotaByProvinsi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).GetKabKotaByProvinsi(ctx, req.(*GetKabKotaByIdWilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabKotaService_GetKabKotaByIdWil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKabKotaByIdWilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).GetKabKotaByIdWil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_GetKabKotaByIdWil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).GetKabKotaByIdWil(ctx, req.(*GetKabKotaByIdWilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabKotaService_CreateKabKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KabKota)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).CreateKabKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_CreateKabKota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).CreateKabKota(ctx, req.(*KabKota))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabKotaService_UpdateKabKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KabKota)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).UpdateKabKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_UpdateKabKota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).UpdateKabKota(ctx, req.(*KabKota))
	}
	return interceptor(ctx, in, info, handler)
}

func _KabKotaService_DeleteKabKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKabKotaByIdWilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).DeleteKabKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_DeleteKabKota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).DeleteKabKota(ctx, req.(*GetKabKotaByIdWilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KabKotaService_ServiceDesc is the grpc.ServiceDesc for KabKotaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KabKotaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.KabKotaService",
	HandlerType: (*KabKotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllKabKota",
			Handler:    _KabKotaService_GetAllKabKota_Handler,
		},
		{
			MethodName: "GetKabKotaByProvinsi",
			Handler:    _KabKotaService_GetKabKotaByProvinsi_Handler,
		},
		{
			MethodName: "GetKabKotaByIdWil",
			Handler:    _KabKotaService_GetKabKotaByIdWil_Handler,
		},
		{
			MethodName: "CreateKabKota",
			Handler:    _KabKotaService_CreateKabKota_Handler,
		},
		{
			MethodName: "UpdateKabKota",
			Handler:    _KabKotaService_UpdateKabKota_Handler,
		},
		{
			MethodName: "DeleteKabKota",
			Handler:    _KabKotaService_DeleteKabKota_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kabkota.proto",
}
