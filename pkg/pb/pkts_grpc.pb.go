// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: pkts.proto

package pb

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
	PKTSService_GetAllPKTS_FullMethodName          = "/tracer_study_grpc.PKTSService/GetAllPKTS"
	PKTSService_GetPKTSByNim_FullMethodName        = "/tracer_study_grpc.PKTSService/GetPKTSByNim"
	PKTSService_CreatePKTS_FullMethodName          = "/tracer_study_grpc.PKTSService/CreatePKTS"
	PKTSService_UpdatePKTS_FullMethodName          = "/tracer_study_grpc.PKTSService/UpdatePKTS"
	PKTSService_GetNimByDataAtasan_FullMethodName  = "/tracer_study_grpc.PKTSService/GetNimByDataAtasan"
	PKTSService_ExportPKTSReport_FullMethodName    = "/tracer_study_grpc.PKTSService/ExportPKTSReport"
	PKTSService_GetPKTSRekapByProdi_FullMethodName = "/tracer_study_grpc.PKTSService/GetPKTSRekapByProdi"
	PKTSService_GetPKTSRekapByYear_FullMethodName  = "/tracer_study_grpc.PKTSService/GetPKTSRekapByYear"
)

// PKTSServiceClient is the client API for PKTSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PKTSServiceClient interface {
	GetAllPKTS(ctx context.Context, in *GetAllPKTSRequest, opts ...grpc.CallOption) (*GetAllPKTSResponse, error)
	GetPKTSByNim(ctx context.Context, in *GetPKTSByNimRequest, opts ...grpc.CallOption) (*GetPKTSResponse, error)
	CreatePKTS(ctx context.Context, in *PKTS, opts ...grpc.CallOption) (*GetPKTSResponse, error)
	UpdatePKTS(ctx context.Context, in *PKTS, opts ...grpc.CallOption) (*GetPKTSResponse, error)
	GetNimByDataAtasan(ctx context.Context, in *GetNimByDataAtasanRequest, opts ...grpc.CallOption) (*GetNimByDataAtasanResponse, error)
	ExportPKTSReport(ctx context.Context, in *ExportPKTSReportRequest, opts ...grpc.CallOption) (*ExportPKTSReportResponse, error)
	GetPKTSRekapByProdi(ctx context.Context, in *GetPKTSRekapByProdiRequest, opts ...grpc.CallOption) (*GetPKTSRekapByProdiResponse, error)
	GetPKTSRekapByYear(ctx context.Context, in *GetPKTSRekapByYearRequest, opts ...grpc.CallOption) (*GetPKTSRekapByYearResponse, error)
}

type pKTSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPKTSServiceClient(cc grpc.ClientConnInterface) PKTSServiceClient {
	return &pKTSServiceClient{cc}
}

func (c *pKTSServiceClient) GetAllPKTS(ctx context.Context, in *GetAllPKTSRequest, opts ...grpc.CallOption) (*GetAllPKTSResponse, error) {
	out := new(GetAllPKTSResponse)
	err := c.cc.Invoke(ctx, PKTSService_GetAllPKTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) GetPKTSByNim(ctx context.Context, in *GetPKTSByNimRequest, opts ...grpc.CallOption) (*GetPKTSResponse, error) {
	out := new(GetPKTSResponse)
	err := c.cc.Invoke(ctx, PKTSService_GetPKTSByNim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) CreatePKTS(ctx context.Context, in *PKTS, opts ...grpc.CallOption) (*GetPKTSResponse, error) {
	out := new(GetPKTSResponse)
	err := c.cc.Invoke(ctx, PKTSService_CreatePKTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) UpdatePKTS(ctx context.Context, in *PKTS, opts ...grpc.CallOption) (*GetPKTSResponse, error) {
	out := new(GetPKTSResponse)
	err := c.cc.Invoke(ctx, PKTSService_UpdatePKTS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) GetNimByDataAtasan(ctx context.Context, in *GetNimByDataAtasanRequest, opts ...grpc.CallOption) (*GetNimByDataAtasanResponse, error) {
	out := new(GetNimByDataAtasanResponse)
	err := c.cc.Invoke(ctx, PKTSService_GetNimByDataAtasan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) ExportPKTSReport(ctx context.Context, in *ExportPKTSReportRequest, opts ...grpc.CallOption) (*ExportPKTSReportResponse, error) {
	out := new(ExportPKTSReportResponse)
	err := c.cc.Invoke(ctx, PKTSService_ExportPKTSReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) GetPKTSRekapByProdi(ctx context.Context, in *GetPKTSRekapByProdiRequest, opts ...grpc.CallOption) (*GetPKTSRekapByProdiResponse, error) {
	out := new(GetPKTSRekapByProdiResponse)
	err := c.cc.Invoke(ctx, PKTSService_GetPKTSRekapByProdi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pKTSServiceClient) GetPKTSRekapByYear(ctx context.Context, in *GetPKTSRekapByYearRequest, opts ...grpc.CallOption) (*GetPKTSRekapByYearResponse, error) {
	out := new(GetPKTSRekapByYearResponse)
	err := c.cc.Invoke(ctx, PKTSService_GetPKTSRekapByYear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PKTSServiceServer is the server API for PKTSService service.
// All implementations must embed UnimplementedPKTSServiceServer
// for forward compatibility
type PKTSServiceServer interface {
	GetAllPKTS(context.Context, *GetAllPKTSRequest) (*GetAllPKTSResponse, error)
	GetPKTSByNim(context.Context, *GetPKTSByNimRequest) (*GetPKTSResponse, error)
	CreatePKTS(context.Context, *PKTS) (*GetPKTSResponse, error)
	UpdatePKTS(context.Context, *PKTS) (*GetPKTSResponse, error)
	GetNimByDataAtasan(context.Context, *GetNimByDataAtasanRequest) (*GetNimByDataAtasanResponse, error)
	ExportPKTSReport(context.Context, *ExportPKTSReportRequest) (*ExportPKTSReportResponse, error)
	GetPKTSRekapByProdi(context.Context, *GetPKTSRekapByProdiRequest) (*GetPKTSRekapByProdiResponse, error)
	GetPKTSRekapByYear(context.Context, *GetPKTSRekapByYearRequest) (*GetPKTSRekapByYearResponse, error)
	mustEmbedUnimplementedPKTSServiceServer()
}

// UnimplementedPKTSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPKTSServiceServer struct {
}

func (UnimplementedPKTSServiceServer) GetAllPKTS(context.Context, *GetAllPKTSRequest) (*GetAllPKTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPKTS not implemented")
}
func (UnimplementedPKTSServiceServer) GetPKTSByNim(context.Context, *GetPKTSByNimRequest) (*GetPKTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPKTSByNim not implemented")
}
func (UnimplementedPKTSServiceServer) CreatePKTS(context.Context, *PKTS) (*GetPKTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePKTS not implemented")
}
func (UnimplementedPKTSServiceServer) UpdatePKTS(context.Context, *PKTS) (*GetPKTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePKTS not implemented")
}
func (UnimplementedPKTSServiceServer) GetNimByDataAtasan(context.Context, *GetNimByDataAtasanRequest) (*GetNimByDataAtasanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNimByDataAtasan not implemented")
}
func (UnimplementedPKTSServiceServer) ExportPKTSReport(context.Context, *ExportPKTSReportRequest) (*ExportPKTSReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportPKTSReport not implemented")
}
func (UnimplementedPKTSServiceServer) GetPKTSRekapByProdi(context.Context, *GetPKTSRekapByProdiRequest) (*GetPKTSRekapByProdiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPKTSRekapByProdi not implemented")
}
func (UnimplementedPKTSServiceServer) GetPKTSRekapByYear(context.Context, *GetPKTSRekapByYearRequest) (*GetPKTSRekapByYearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPKTSRekapByYear not implemented")
}
func (UnimplementedPKTSServiceServer) mustEmbedUnimplementedPKTSServiceServer() {}

// UnsafePKTSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PKTSServiceServer will
// result in compilation errors.
type UnsafePKTSServiceServer interface {
	mustEmbedUnimplementedPKTSServiceServer()
}

func RegisterPKTSServiceServer(s grpc.ServiceRegistrar, srv PKTSServiceServer) {
	s.RegisterService(&PKTSService_ServiceDesc, srv)
}

func _PKTSService_GetAllPKTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPKTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).GetAllPKTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_GetAllPKTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).GetAllPKTS(ctx, req.(*GetAllPKTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_GetPKTSByNim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPKTSByNimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).GetPKTSByNim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_GetPKTSByNim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).GetPKTSByNim(ctx, req.(*GetPKTSByNimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_CreatePKTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PKTS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).CreatePKTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_CreatePKTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).CreatePKTS(ctx, req.(*PKTS))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_UpdatePKTS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PKTS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).UpdatePKTS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_UpdatePKTS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).UpdatePKTS(ctx, req.(*PKTS))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_GetNimByDataAtasan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNimByDataAtasanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).GetNimByDataAtasan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_GetNimByDataAtasan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).GetNimByDataAtasan(ctx, req.(*GetNimByDataAtasanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_ExportPKTSReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportPKTSReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).ExportPKTSReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_ExportPKTSReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).ExportPKTSReport(ctx, req.(*ExportPKTSReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_GetPKTSRekapByProdi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPKTSRekapByProdiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).GetPKTSRekapByProdi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_GetPKTSRekapByProdi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).GetPKTSRekapByProdi(ctx, req.(*GetPKTSRekapByProdiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PKTSService_GetPKTSRekapByYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPKTSRekapByYearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PKTSServiceServer).GetPKTSRekapByYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PKTSService_GetPKTSRekapByYear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PKTSServiceServer).GetPKTSRekapByYear(ctx, req.(*GetPKTSRekapByYearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PKTSService_ServiceDesc is the grpc.ServiceDesc for PKTSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PKTSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.PKTSService",
	HandlerType: (*PKTSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPKTS",
			Handler:    _PKTSService_GetAllPKTS_Handler,
		},
		{
			MethodName: "GetPKTSByNim",
			Handler:    _PKTSService_GetPKTSByNim_Handler,
		},
		{
			MethodName: "CreatePKTS",
			Handler:    _PKTSService_CreatePKTS_Handler,
		},
		{
			MethodName: "UpdatePKTS",
			Handler:    _PKTSService_UpdatePKTS_Handler,
		},
		{
			MethodName: "GetNimByDataAtasan",
			Handler:    _PKTSService_GetNimByDataAtasan_Handler,
		},
		{
			MethodName: "ExportPKTSReport",
			Handler:    _PKTSService_ExportPKTSReport_Handler,
		},
		{
			MethodName: "GetPKTSRekapByProdi",
			Handler:    _PKTSService_GetPKTSRekapByProdi_Handler,
		},
		{
			MethodName: "GetPKTSRekapByYear",
			Handler:    _PKTSService_GetPKTSRekapByYear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkts.proto",
}
