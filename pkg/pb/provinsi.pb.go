// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: provinsi.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Provinsi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdWil     string `protobuf:"bytes,2,opt,name=id_wil,json=idWil,proto3" json:"id_wil,omitempty"`
	Nama      string `protobuf:"bytes,3,opt,name=nama,proto3" json:"nama,omitempty"`
	Ump       uint64 `protobuf:"varint,4,opt,name=ump,proto3" json:"ump,omitempty"`
	UmpPkts   uint64 `protobuf:"varint,5,opt,name=ump_pkts,json=umpPkts,proto3" json:"ump_pkts,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Provinsi) Reset() {
	*x = Provinsi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provinsi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provinsi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provinsi) ProtoMessage() {}

func (x *Provinsi) ProtoReflect() protoreflect.Message {
	mi := &file_provinsi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provinsi.ProtoReflect.Descriptor instead.
func (*Provinsi) Descriptor() ([]byte, []int) {
	return file_provinsi_proto_rawDescGZIP(), []int{0}
}

func (x *Provinsi) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Provinsi) GetIdWil() string {
	if x != nil {
		return x.IdWil
	}
	return ""
}

func (x *Provinsi) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *Provinsi) GetUmp() uint64 {
	if x != nil {
		return x.Ump
	}
	return 0
}

func (x *Provinsi) GetUmpPkts() uint64 {
	if x != nil {
		return x.UmpPkts
	}
	return 0
}

func (x *Provinsi) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Provinsi) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetAllProvinsiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Provinsi `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllProvinsiResponse) Reset() {
	*x = GetAllProvinsiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provinsi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProvinsiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProvinsiResponse) ProtoMessage() {}

func (x *GetAllProvinsiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provinsi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProvinsiResponse.ProtoReflect.Descriptor instead.
func (*GetAllProvinsiResponse) Descriptor() ([]byte, []int) {
	return file_provinsi_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllProvinsiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllProvinsiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProvinsiResponse) GetData() []*Provinsi {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProvinsiByIdWilRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdWil string `protobuf:"bytes,1,opt,name=id_wil,json=idWil,proto3" json:"id_wil,omitempty"`
}

func (x *GetProvinsiByIdWilRequest) Reset() {
	*x = GetProvinsiByIdWilRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provinsi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvinsiByIdWilRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvinsiByIdWilRequest) ProtoMessage() {}

func (x *GetProvinsiByIdWilRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provinsi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvinsiByIdWilRequest.ProtoReflect.Descriptor instead.
func (*GetProvinsiByIdWilRequest) Descriptor() ([]byte, []int) {
	return file_provinsi_proto_rawDescGZIP(), []int{2}
}

func (x *GetProvinsiByIdWilRequest) GetIdWil() string {
	if x != nil {
		return x.IdWil
	}
	return ""
}

type GetProvinsiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Provinsi `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProvinsiResponse) Reset() {
	*x = GetProvinsiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provinsi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvinsiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvinsiResponse) ProtoMessage() {}

func (x *GetProvinsiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provinsi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvinsiResponse.ProtoReflect.Descriptor instead.
func (*GetProvinsiResponse) Descriptor() ([]byte, []int) {
	return file_provinsi_proto_rawDescGZIP(), []int{3}
}

func (x *GetProvinsiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetProvinsiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProvinsiResponse) GetData() *Provinsi {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteProvinsiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteProvinsiResponse) Reset() {
	*x = DeleteProvinsiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provinsi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProvinsiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProvinsiResponse) ProtoMessage() {}

func (x *DeleteProvinsiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provinsi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProvinsiResponse.ProtoReflect.Descriptor instead.
func (*DeleteProvinsiResponse) Descriptor() ([]byte, []int) {
	return file_provinsi_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProvinsiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteProvinsiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_provinsi_proto protoreflect.FileDescriptor

var file_provinsi_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb0, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x64, 0x5f, 0x77, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x64, 0x57, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x6d,
	0x70, 0x5f, 0x70, 0x6b, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x6d,
	0x70, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x77, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x42, 0x79, 0x49, 0x64, 0x57,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x64, 0x5f,
	0x77, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x57, 0x69, 0x6c,
	0x22, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf5,
	0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x73, 0x69, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x42, 0x79, 0x49, 0x64, 0x57, 0x69, 0x6c, 0x12,
	0x2c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x42,
	0x79, 0x49, 0x64, 0x57, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x73, 0x69, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x1a,
	0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x12, 0x2c, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x42, 0x79, 0x49, 0x64, 0x57,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provinsi_proto_rawDescOnce sync.Once
	file_provinsi_proto_rawDescData = file_provinsi_proto_rawDesc
)

func file_provinsi_proto_rawDescGZIP() []byte {
	file_provinsi_proto_rawDescOnce.Do(func() {
		file_provinsi_proto_rawDescData = protoimpl.X.CompressGZIP(file_provinsi_proto_rawDescData)
	})
	return file_provinsi_proto_rawDescData
}

var file_provinsi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_provinsi_proto_goTypes = []interface{}{
	(*Provinsi)(nil),                  // 0: tracer_study_grpc.Provinsi
	(*GetAllProvinsiResponse)(nil),    // 1: tracer_study_grpc.GetAllProvinsiResponse
	(*GetProvinsiByIdWilRequest)(nil), // 2: tracer_study_grpc.GetProvinsiByIdWilRequest
	(*GetProvinsiResponse)(nil),       // 3: tracer_study_grpc.GetProvinsiResponse
	(*DeleteProvinsiResponse)(nil),    // 4: tracer_study_grpc.DeleteProvinsiResponse
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_provinsi_proto_depIdxs = []int32{
	0, // 0: tracer_study_grpc.GetAllProvinsiResponse.data:type_name -> tracer_study_grpc.Provinsi
	0, // 1: tracer_study_grpc.GetProvinsiResponse.data:type_name -> tracer_study_grpc.Provinsi
	5, // 2: tracer_study_grpc.ProvinsiService.GetAllProvinsi:input_type -> google.protobuf.Empty
	2, // 3: tracer_study_grpc.ProvinsiService.GetProvinsiByIdWil:input_type -> tracer_study_grpc.GetProvinsiByIdWilRequest
	0, // 4: tracer_study_grpc.ProvinsiService.CreateProvinsi:input_type -> tracer_study_grpc.Provinsi
	0, // 5: tracer_study_grpc.ProvinsiService.UpdateProvinsi:input_type -> tracer_study_grpc.Provinsi
	2, // 6: tracer_study_grpc.ProvinsiService.DeleteProvinsi:input_type -> tracer_study_grpc.GetProvinsiByIdWilRequest
	1, // 7: tracer_study_grpc.ProvinsiService.GetAllProvinsi:output_type -> tracer_study_grpc.GetAllProvinsiResponse
	3, // 8: tracer_study_grpc.ProvinsiService.GetProvinsiByIdWil:output_type -> tracer_study_grpc.GetProvinsiResponse
	3, // 9: tracer_study_grpc.ProvinsiService.CreateProvinsi:output_type -> tracer_study_grpc.GetProvinsiResponse
	3, // 10: tracer_study_grpc.ProvinsiService.UpdateProvinsi:output_type -> tracer_study_grpc.GetProvinsiResponse
	4, // 11: tracer_study_grpc.ProvinsiService.DeleteProvinsi:output_type -> tracer_study_grpc.DeleteProvinsiResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_provinsi_proto_init() }
func file_provinsi_proto_init() {
	if File_provinsi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provinsi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provinsi); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_provinsi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProvinsiResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_provinsi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvinsiByIdWilRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_provinsi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvinsiResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_provinsi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProvinsiResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_provinsi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provinsi_proto_goTypes,
		DependencyIndexes: file_provinsi_proto_depIdxs,
		MessageInfos:      file_provinsi_proto_msgTypes,
	}.Build()
	File_provinsi_proto = out.File
	file_provinsi_proto_rawDesc = nil
	file_provinsi_proto_goTypes = nil
	file_provinsi_proto_depIdxs = nil
}
