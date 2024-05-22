// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: kabkota.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KabKota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdWil      string                 `protobuf:"bytes,2,opt,name=id_wil,json=idWil,proto3" json:"id_wil,omitempty"`
	Nama       string                 `protobuf:"bytes,3,opt,name=nama,proto3" json:"nama,omitempty"`
	IdIndukWil string                 `protobuf:"bytes,4,opt,name=id_induk_wil,json=idIndukWil,proto3" json:"id_induk_wil,omitempty"`
	Provinsi   *Provinsi              `protobuf:"bytes,5,opt,name=provinsi,proto3" json:"provinsi,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *KabKota) Reset() {
	*x = KabKota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kabkota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KabKota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KabKota) ProtoMessage() {}

func (x *KabKota) ProtoReflect() protoreflect.Message {
	mi := &file_kabkota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KabKota.ProtoReflect.Descriptor instead.
func (*KabKota) Descriptor() ([]byte, []int) {
	return file_kabkota_proto_rawDescGZIP(), []int{0}
}

func (x *KabKota) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KabKota) GetIdWil() string {
	if x != nil {
		return x.IdWil
	}
	return ""
}

func (x *KabKota) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *KabKota) GetIdIndukWil() string {
	if x != nil {
		return x.IdIndukWil
	}
	return ""
}

func (x *KabKota) GetProvinsi() *Provinsi {
	if x != nil {
		return x.Provinsi
	}
	return nil
}

func (x *KabKota) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *KabKota) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetAllKabKotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*KabKota `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllKabKotaResponse) Reset() {
	*x = GetAllKabKotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kabkota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKabKotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKabKotaResponse) ProtoMessage() {}

func (x *GetAllKabKotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kabkota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKabKotaResponse.ProtoReflect.Descriptor instead.
func (*GetAllKabKotaResponse) Descriptor() ([]byte, []int) {
	return file_kabkota_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllKabKotaResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllKabKotaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllKabKotaResponse) GetData() []*KabKota {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetKabKotaByIdWilRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdWil string `protobuf:"bytes,1,opt,name=id_wil,json=idWil,proto3" json:"id_wil,omitempty"`
}

func (x *GetKabKotaByIdWilRequest) Reset() {
	*x = GetKabKotaByIdWilRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kabkota_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKabKotaByIdWilRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKabKotaByIdWilRequest) ProtoMessage() {}

func (x *GetKabKotaByIdWilRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kabkota_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKabKotaByIdWilRequest.ProtoReflect.Descriptor instead.
func (*GetKabKotaByIdWilRequest) Descriptor() ([]byte, []int) {
	return file_kabkota_proto_rawDescGZIP(), []int{2}
}

func (x *GetKabKotaByIdWilRequest) GetIdWil() string {
	if x != nil {
		return x.IdWil
	}
	return ""
}

type GetKabKotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *KabKota `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetKabKotaResponse) Reset() {
	*x = GetKabKotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kabkota_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKabKotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKabKotaResponse) ProtoMessage() {}

func (x *GetKabKotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kabkota_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKabKotaResponse.ProtoReflect.Descriptor instead.
func (*GetKabKotaResponse) Descriptor() ([]byte, []int) {
	return file_kabkota_proto_rawDescGZIP(), []int{3}
}

func (x *GetKabKotaResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetKabKotaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetKabKotaResponse) GetData() *KabKota {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteKabKotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteKabKotaResponse) Reset() {
	*x = DeleteKabKotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kabkota_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKabKotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKabKotaResponse) ProtoMessage() {}

func (x *DeleteKabKotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kabkota_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKabKotaResponse.ProtoReflect.Descriptor instead.
func (*DeleteKabKotaResponse) Descriptor() ([]byte, []int) {
	return file_kabkota_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKabKotaResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteKabKotaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_kabkota_proto protoreflect.FileDescriptor

var file_kabkota_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x61, 0x62, 0x6b, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x02, 0x0a, 0x07, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x64, 0x5f, 0x77, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64,
	0x57, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x64, 0x5f, 0x69, 0x6e,
	0x64, 0x75, 0x6b, 0x5f, 0x77, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x49, 0x6e, 0x64, 0x75, 0x6b, 0x57, 0x69, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x73, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x73, 0x69, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x73, 0x69, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x31, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x49,
	0x64, 0x57, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x64, 0x5f, 0x77, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x57,
	0x69, 0x6c, 0x22, 0x72, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe6, 0x03,
	0x0a, 0x0e, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74,
	0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b,
	0x6f, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x57, 0x69, 0x6c, 0x12, 0x2b, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x57, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74,
	0x61, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x1a, 0x25, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x61, 0x62, 0x4b,
	0x6f, 0x74, 0x61, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x12, 0x2b, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64,
	0x57, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kabkota_proto_rawDescOnce sync.Once
	file_kabkota_proto_rawDescData = file_kabkota_proto_rawDesc
)

func file_kabkota_proto_rawDescGZIP() []byte {
	file_kabkota_proto_rawDescOnce.Do(func() {
		file_kabkota_proto_rawDescData = protoimpl.X.CompressGZIP(file_kabkota_proto_rawDescData)
	})
	return file_kabkota_proto_rawDescData
}

var file_kabkota_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kabkota_proto_goTypes = []interface{}{
	(*KabKota)(nil),                  // 0: tracer_study_grpc.KabKota
	(*GetAllKabKotaResponse)(nil),    // 1: tracer_study_grpc.GetAllKabKotaResponse
	(*GetKabKotaByIdWilRequest)(nil), // 2: tracer_study_grpc.GetKabKotaByIdWilRequest
	(*GetKabKotaResponse)(nil),       // 3: tracer_study_grpc.GetKabKotaResponse
	(*DeleteKabKotaResponse)(nil),    // 4: tracer_study_grpc.DeleteKabKotaResponse
	(*Provinsi)(nil),                 // 5: tracer_study_grpc.Provinsi
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_kabkota_proto_depIdxs = []int32{
	5,  // 0: tracer_study_grpc.KabKota.provinsi:type_name -> tracer_study_grpc.Provinsi
	6,  // 1: tracer_study_grpc.KabKota.created_at:type_name -> google.protobuf.Timestamp
	6,  // 2: tracer_study_grpc.KabKota.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: tracer_study_grpc.GetAllKabKotaResponse.data:type_name -> tracer_study_grpc.KabKota
	0,  // 4: tracer_study_grpc.GetKabKotaResponse.data:type_name -> tracer_study_grpc.KabKota
	7,  // 5: tracer_study_grpc.KabKotaService.GetAllKabKota:input_type -> google.protobuf.Empty
	2,  // 6: tracer_study_grpc.KabKotaService.GetKabKotaByIdWil:input_type -> tracer_study_grpc.GetKabKotaByIdWilRequest
	0,  // 7: tracer_study_grpc.KabKotaService.CreateKabKota:input_type -> tracer_study_grpc.KabKota
	0,  // 8: tracer_study_grpc.KabKotaService.UpdateKabKota:input_type -> tracer_study_grpc.KabKota
	2,  // 9: tracer_study_grpc.KabKotaService.DeleteKabKota:input_type -> tracer_study_grpc.GetKabKotaByIdWilRequest
	1,  // 10: tracer_study_grpc.KabKotaService.GetAllKabKota:output_type -> tracer_study_grpc.GetAllKabKotaResponse
	3,  // 11: tracer_study_grpc.KabKotaService.GetKabKotaByIdWil:output_type -> tracer_study_grpc.GetKabKotaResponse
	3,  // 12: tracer_study_grpc.KabKotaService.CreateKabKota:output_type -> tracer_study_grpc.GetKabKotaResponse
	3,  // 13: tracer_study_grpc.KabKotaService.UpdateKabKota:output_type -> tracer_study_grpc.GetKabKotaResponse
	4,  // 14: tracer_study_grpc.KabKotaService.DeleteKabKota:output_type -> tracer_study_grpc.DeleteKabKotaResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_kabkota_proto_init() }
func file_kabkota_proto_init() {
	if File_kabkota_proto != nil {
		return
	}
	file_provinsi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kabkota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KabKota); i {
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
		file_kabkota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKabKotaResponse); i {
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
		file_kabkota_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKabKotaByIdWilRequest); i {
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
		file_kabkota_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKabKotaResponse); i {
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
		file_kabkota_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKabKotaResponse); i {
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
			RawDescriptor: file_kabkota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kabkota_proto_goTypes,
		DependencyIndexes: file_kabkota_proto_depIdxs,
		MessageInfos:      file_kabkota_proto_msgTypes,
	}.Build()
	File_kabkota_proto = out.File
	file_kabkota_proto_rawDesc = nil
	file_kabkota_proto_goTypes = nil
	file_kabkota_proto_depIdxs = nil
}
