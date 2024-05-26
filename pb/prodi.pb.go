// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: prodi.proto

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

type Prodi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kode            string                 `protobuf:"bytes,2,opt,name=kode,proto3" json:"kode,omitempty"`
	KodeDikti       string                 `protobuf:"bytes,3,opt,name=kode_dikti,json=kodeDikti,proto3" json:"kode_dikti,omitempty"`
	KodeIntegrasi   string                 `protobuf:"bytes,4,opt,name=kode_integrasi,json=kodeIntegrasi,proto3" json:"kode_integrasi,omitempty"`
	Nama            string                 `protobuf:"bytes,5,opt,name=nama,proto3" json:"nama,omitempty"`
	Jenjang         string                 `protobuf:"bytes,6,opt,name=jenjang,proto3" json:"jenjang,omitempty"`
	KodeFakultas    string                 `protobuf:"bytes,7,opt,name=kode_fakultas,json=kodeFakultas,proto3" json:"kode_fakultas,omitempty"`
	NamaFakultas    string                 `protobuf:"bytes,8,opt,name=nama_fakultas,json=namaFakultas,proto3" json:"nama_fakultas,omitempty"`
	AkronimFakultas string                 `protobuf:"bytes,9,opt,name=akronim_fakultas,json=akronimFakultas,proto3" json:"akronim_fakultas,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Prodi) Reset() {
	*x = Prodi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prodi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prodi) ProtoMessage() {}

func (x *Prodi) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prodi.ProtoReflect.Descriptor instead.
func (*Prodi) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{0}
}

func (x *Prodi) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Prodi) GetKode() string {
	if x != nil {
		return x.Kode
	}
	return ""
}

func (x *Prodi) GetKodeDikti() string {
	if x != nil {
		return x.KodeDikti
	}
	return ""
}

func (x *Prodi) GetKodeIntegrasi() string {
	if x != nil {
		return x.KodeIntegrasi
	}
	return ""
}

func (x *Prodi) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *Prodi) GetJenjang() string {
	if x != nil {
		return x.Jenjang
	}
	return ""
}

func (x *Prodi) GetKodeFakultas() string {
	if x != nil {
		return x.KodeFakultas
	}
	return ""
}

func (x *Prodi) GetNamaFakultas() string {
	if x != nil {
		return x.NamaFakultas
	}
	return ""
}

func (x *Prodi) GetAkronimFakultas() string {
	if x != nil {
		return x.AkronimFakultas
	}
	return ""
}

func (x *Prodi) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Prodi) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Fakultas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kode    string `protobuf:"bytes,1,opt,name=kode,proto3" json:"kode,omitempty"`
	Nama    string `protobuf:"bytes,2,opt,name=nama,proto3" json:"nama,omitempty"`
	Akronim string `protobuf:"bytes,3,opt,name=akronim,proto3" json:"akronim,omitempty"`
}

func (x *Fakultas) Reset() {
	*x = Fakultas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fakultas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fakultas) ProtoMessage() {}

func (x *Fakultas) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fakultas.ProtoReflect.Descriptor instead.
func (*Fakultas) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{1}
}

func (x *Fakultas) GetKode() string {
	if x != nil {
		return x.Kode
	}
	return ""
}

func (x *Fakultas) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *Fakultas) GetAkronim() string {
	if x != nil {
		return x.Akronim
	}
	return ""
}

type GetAllProdiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Prodi `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllProdiResponse) Reset() {
	*x = GetAllProdiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProdiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProdiResponse) ProtoMessage() {}

func (x *GetAllProdiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProdiResponse.ProtoReflect.Descriptor instead.
func (*GetAllProdiResponse) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllProdiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllProdiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProdiResponse) GetData() []*Prodi {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllFakultasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Fakultas `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllFakultasResponse) Reset() {
	*x = GetAllFakultasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFakultasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFakultasResponse) ProtoMessage() {}

func (x *GetAllFakultasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFakultasResponse.ProtoReflect.Descriptor instead.
func (*GetAllFakultasResponse) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllFakultasResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllFakultasResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllFakultasResponse) GetData() []*Fakultas {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProdiByKodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kode string `protobuf:"bytes,1,opt,name=kode,proto3" json:"kode,omitempty"`
}

func (x *GetProdiByKodeRequest) Reset() {
	*x = GetProdiByKodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProdiByKodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProdiByKodeRequest) ProtoMessage() {}

func (x *GetProdiByKodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProdiByKodeRequest.ProtoReflect.Descriptor instead.
func (*GetProdiByKodeRequest) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{4}
}

func (x *GetProdiByKodeRequest) GetKode() string {
	if x != nil {
		return x.Kode
	}
	return ""
}

type GetProdiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Prodi `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProdiResponse) Reset() {
	*x = GetProdiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProdiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProdiResponse) ProtoMessage() {}

func (x *GetProdiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProdiResponse.ProtoReflect.Descriptor instead.
func (*GetProdiResponse) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{5}
}

func (x *GetProdiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetProdiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProdiResponse) GetData() *Prodi {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteProdiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteProdiResponse) Reset() {
	*x = DeleteProdiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProdiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProdiResponse) ProtoMessage() {}

func (x *DeleteProdiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProdiResponse.ProtoReflect.Descriptor instead.
func (*DeleteProdiResponse) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProdiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteProdiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_prodi_proto protoreflect.FileDescriptor

var file_prodi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x03, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6b, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6b, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x6b, 0x74, 0x69, 0x12, 0x25, 0x0a, 0x0e, 0x6b,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x73, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x73, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x65, 0x6e, 0x6a, 0x61, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x65, 0x6e, 0x6a, 0x61, 0x6e, 0x67,
	0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x6f, 0x64, 0x65, 0x46, 0x61, 0x6b,
	0x75, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x61, 0x5f, 0x66, 0x61,
	0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61,
	0x6d, 0x61, 0x46, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6b,
	0x72, 0x6f, 0x6e, 0x69, 0x6d, 0x5f, 0x66, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6b, 0x72, 0x6f, 0x6e, 0x69, 0x6d, 0x46, 0x61, 0x6b,
	0x75, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4c, 0x0a, 0x08, 0x46,
	0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x6b, 0x72, 0x6f, 0x6e, 0x69, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x6b, 0x72, 0x6f, 0x6e, 0x69, 0x6d, 0x22, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x69, 0x42, 0x79, 0x4b, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x6f,
	0x64, 0x65, 0x22, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9c, 0x04, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x46, 0x61, 0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61,
	0x6b, 0x75, 0x6c, 0x74, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x42, 0x79, 0x4b, 0x6f,
	0x64, 0x65, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x42,
	0x79, 0x4b, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x69, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x1a, 0x23, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x69, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x1a, 0x23, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x69, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x42,
	0x79, 0x4b, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodi_proto_rawDescOnce sync.Once
	file_prodi_proto_rawDescData = file_prodi_proto_rawDesc
)

func file_prodi_proto_rawDescGZIP() []byte {
	file_prodi_proto_rawDescOnce.Do(func() {
		file_prodi_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodi_proto_rawDescData)
	})
	return file_prodi_proto_rawDescData
}

var file_prodi_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_prodi_proto_goTypes = []interface{}{
	(*Prodi)(nil),                  // 0: tracer_study_grpc.Prodi
	(*Fakultas)(nil),               // 1: tracer_study_grpc.Fakultas
	(*GetAllProdiResponse)(nil),    // 2: tracer_study_grpc.GetAllProdiResponse
	(*GetAllFakultasResponse)(nil), // 3: tracer_study_grpc.GetAllFakultasResponse
	(*GetProdiByKodeRequest)(nil),  // 4: tracer_study_grpc.GetProdiByKodeRequest
	(*GetProdiResponse)(nil),       // 5: tracer_study_grpc.GetProdiResponse
	(*DeleteProdiResponse)(nil),    // 6: tracer_study_grpc.DeleteProdiResponse
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_prodi_proto_depIdxs = []int32{
	7,  // 0: tracer_study_grpc.Prodi.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: tracer_study_grpc.Prodi.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: tracer_study_grpc.GetAllProdiResponse.data:type_name -> tracer_study_grpc.Prodi
	1,  // 3: tracer_study_grpc.GetAllFakultasResponse.data:type_name -> tracer_study_grpc.Fakultas
	0,  // 4: tracer_study_grpc.GetProdiResponse.data:type_name -> tracer_study_grpc.Prodi
	8,  // 5: tracer_study_grpc.ProdiService.GetAllProdi:input_type -> google.protobuf.Empty
	8,  // 6: tracer_study_grpc.ProdiService.GetAllFakultas:input_type -> google.protobuf.Empty
	4,  // 7: tracer_study_grpc.ProdiService.GetProdiByKode:input_type -> tracer_study_grpc.GetProdiByKodeRequest
	0,  // 8: tracer_study_grpc.ProdiService.CreateProdi:input_type -> tracer_study_grpc.Prodi
	0,  // 9: tracer_study_grpc.ProdiService.UpdateProdi:input_type -> tracer_study_grpc.Prodi
	4,  // 10: tracer_study_grpc.ProdiService.DeleteProdi:input_type -> tracer_study_grpc.GetProdiByKodeRequest
	2,  // 11: tracer_study_grpc.ProdiService.GetAllProdi:output_type -> tracer_study_grpc.GetAllProdiResponse
	3,  // 12: tracer_study_grpc.ProdiService.GetAllFakultas:output_type -> tracer_study_grpc.GetAllFakultasResponse
	5,  // 13: tracer_study_grpc.ProdiService.GetProdiByKode:output_type -> tracer_study_grpc.GetProdiResponse
	5,  // 14: tracer_study_grpc.ProdiService.CreateProdi:output_type -> tracer_study_grpc.GetProdiResponse
	5,  // 15: tracer_study_grpc.ProdiService.UpdateProdi:output_type -> tracer_study_grpc.GetProdiResponse
	6,  // 16: tracer_study_grpc.ProdiService.DeleteProdi:output_type -> tracer_study_grpc.DeleteProdiResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_prodi_proto_init() }
func file_prodi_proto_init() {
	if File_prodi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prodi); i {
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
		file_prodi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fakultas); i {
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
		file_prodi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProdiResponse); i {
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
		file_prodi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFakultasResponse); i {
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
		file_prodi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProdiByKodeRequest); i {
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
		file_prodi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProdiResponse); i {
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
		file_prodi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProdiResponse); i {
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
			RawDescriptor: file_prodi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prodi_proto_goTypes,
		DependencyIndexes: file_prodi_proto_depIdxs,
		MessageInfos:      file_prodi_proto_msgTypes,
	}.Build()
	File_prodi_proto = out.File
	file_prodi_proto_rawDesc = nil
	file_prodi_proto_goTypes = nil
	file_prodi_proto_depIdxs = nil
}
