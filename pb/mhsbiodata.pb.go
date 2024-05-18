// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: mhsbiodata.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MhsBiodata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NIM        string `protobuf:"bytes,1,opt,name=NIM,proto3" json:"NIM,omitempty"`
	NAMA       string `protobuf:"bytes,2,opt,name=NAMA,proto3" json:"NAMA,omitempty"`
	IPK        string `protobuf:"bytes,3,opt,name=IPK,proto3" json:"IPK,omitempty"`
	HP         string `protobuf:"bytes,4,opt,name=HP,proto3" json:"HP,omitempty"`
	EMAIL      string `protobuf:"bytes,5,opt,name=EMAIL,proto3" json:"EMAIL,omitempty"`
	KODEPSTD   string `protobuf:"bytes,6,opt,name=KODEPSTD,proto3" json:"KODEPSTD,omitempty"`
	JENJANG    string `protobuf:"bytes,7,opt,name=JENJANG,proto3" json:"JENJANG,omitempty"`
	PRODI      string `protobuf:"bytes,8,opt,name=PRODI,proto3" json:"PRODI,omitempty"`
	NAMAPST    string `protobuf:"bytes,9,opt,name=NAMAPST,proto3" json:"NAMAPST,omitempty"`
	KODEPST    string `protobuf:"bytes,10,opt,name=KODEPST,proto3" json:"KODEPST,omitempty"`
	KODEFAK    string `protobuf:"bytes,11,opt,name=KODEFAK,proto3" json:"KODEFAK,omitempty"`
	NAMAFAK    string `protobuf:"bytes,12,opt,name=NAMAFAK,proto3" json:"NAMAFAK,omitempty"`
	JLRMASUK   string `protobuf:"bytes,13,opt,name=JLRMASUK,proto3" json:"JLRMASUK,omitempty"`
	THNMASUK   string `protobuf:"bytes,14,opt,name=THNMASUK,proto3" json:"THNMASUK,omitempty"`
	LAMASTD    string `protobuf:"bytes,15,opt,name=LAMASTD,proto3" json:"LAMASTD,omitempty"`
	TGLSIDANG  string `protobuf:"bytes,16,opt,name=TGLSIDANG,proto3" json:"TGLSIDANG,omitempty"`
	KODEJK     string `protobuf:"bytes,17,opt,name=KODEJK,proto3" json:"KODEJK,omitempty"`
	KODESTATUS string `protobuf:"bytes,18,opt,name=KODESTATUS,proto3" json:"KODESTATUS,omitempty"`
}

func (x *MhsBiodata) Reset() {
	*x = MhsBiodata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mhsbiodata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MhsBiodata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MhsBiodata) ProtoMessage() {}

func (x *MhsBiodata) ProtoReflect() protoreflect.Message {
	mi := &file_mhsbiodata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MhsBiodata.ProtoReflect.Descriptor instead.
func (*MhsBiodata) Descriptor() ([]byte, []int) {
	return file_mhsbiodata_proto_rawDescGZIP(), []int{0}
}

func (x *MhsBiodata) GetNIM() string {
	if x != nil {
		return x.NIM
	}
	return ""
}

func (x *MhsBiodata) GetNAMA() string {
	if x != nil {
		return x.NAMA
	}
	return ""
}

func (x *MhsBiodata) GetIPK() string {
	if x != nil {
		return x.IPK
	}
	return ""
}

func (x *MhsBiodata) GetHP() string {
	if x != nil {
		return x.HP
	}
	return ""
}

func (x *MhsBiodata) GetEMAIL() string {
	if x != nil {
		return x.EMAIL
	}
	return ""
}

func (x *MhsBiodata) GetKODEPSTD() string {
	if x != nil {
		return x.KODEPSTD
	}
	return ""
}

func (x *MhsBiodata) GetJENJANG() string {
	if x != nil {
		return x.JENJANG
	}
	return ""
}

func (x *MhsBiodata) GetPRODI() string {
	if x != nil {
		return x.PRODI
	}
	return ""
}

func (x *MhsBiodata) GetNAMAPST() string {
	if x != nil {
		return x.NAMAPST
	}
	return ""
}

func (x *MhsBiodata) GetKODEPST() string {
	if x != nil {
		return x.KODEPST
	}
	return ""
}

func (x *MhsBiodata) GetKODEFAK() string {
	if x != nil {
		return x.KODEFAK
	}
	return ""
}

func (x *MhsBiodata) GetNAMAFAK() string {
	if x != nil {
		return x.NAMAFAK
	}
	return ""
}

func (x *MhsBiodata) GetJLRMASUK() string {
	if x != nil {
		return x.JLRMASUK
	}
	return ""
}

func (x *MhsBiodata) GetTHNMASUK() string {
	if x != nil {
		return x.THNMASUK
	}
	return ""
}

func (x *MhsBiodata) GetLAMASTD() string {
	if x != nil {
		return x.LAMASTD
	}
	return ""
}

func (x *MhsBiodata) GetTGLSIDANG() string {
	if x != nil {
		return x.TGLSIDANG
	}
	return ""
}

func (x *MhsBiodata) GetKODEJK() string {
	if x != nil {
		return x.KODEJK
	}
	return ""
}

func (x *MhsBiodata) GetKODESTATUS() string {
	if x != nil {
		return x.KODESTATUS
	}
	return ""
}

type MhsBiodataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nim string `protobuf:"bytes,1,opt,name=nim,proto3" json:"nim,omitempty"`
}

func (x *MhsBiodataRequest) Reset() {
	*x = MhsBiodataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mhsbiodata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MhsBiodataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MhsBiodataRequest) ProtoMessage() {}

func (x *MhsBiodataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mhsbiodata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MhsBiodataRequest.ProtoReflect.Descriptor instead.
func (*MhsBiodataRequest) Descriptor() ([]byte, []int) {
	return file_mhsbiodata_proto_rawDescGZIP(), []int{1}
}

func (x *MhsBiodataRequest) GetNim() string {
	if x != nil {
		return x.Nim
	}
	return ""
}

type MhsBiodataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *MhsBiodata `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MhsBiodataResponse) Reset() {
	*x = MhsBiodataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mhsbiodata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MhsBiodataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MhsBiodataResponse) ProtoMessage() {}

func (x *MhsBiodataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mhsbiodata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MhsBiodataResponse.ProtoReflect.Descriptor instead.
func (*MhsBiodataResponse) Descriptor() ([]byte, []int) {
	return file_mhsbiodata_proto_rawDescGZIP(), []int{2}
}

func (x *MhsBiodataResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MhsBiodataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MhsBiodataResponse) GetData() *MhsBiodata {
	if x != nil {
		return x.Data
	}
	return nil
}

type CheckMhsAlumniResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	IsAlumni bool   `protobuf:"varint,3,opt,name=isAlumni,proto3" json:"isAlumni,omitempty"`
}

func (x *CheckMhsAlumniResponse) Reset() {
	*x = CheckMhsAlumniResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mhsbiodata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckMhsAlumniResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckMhsAlumniResponse) ProtoMessage() {}

func (x *CheckMhsAlumniResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mhsbiodata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckMhsAlumniResponse.ProtoReflect.Descriptor instead.
func (*CheckMhsAlumniResponse) Descriptor() ([]byte, []int) {
	return file_mhsbiodata_proto_rawDescGZIP(), []int{3}
}

func (x *CheckMhsAlumniResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CheckMhsAlumniResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CheckMhsAlumniResponse) GetIsAlumni() bool {
	if x != nil {
		return x.IsAlumni
	}
	return false
}

var File_mhsbiodata_proto protoreflect.FileDescriptor

var file_mhsbiodata_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x68, 0x73, 0x62, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0xc6, 0x03, 0x0a, 0x0a, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x49, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4e, 0x49, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x41, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x41, 0x4d, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x50,
	0x4b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x50, 0x4b, 0x12, 0x0e, 0x0a, 0x02,
	0x48, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x48, 0x50, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x4d, 0x41,
	0x49, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x4f, 0x44, 0x45, 0x50, 0x53, 0x54, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x4f, 0x44, 0x45, 0x50, 0x53, 0x54, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x4a, 0x45, 0x4e, 0x4a, 0x41, 0x4e, 0x47, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4a, 0x45, 0x4e, 0x4a, 0x41, 0x4e, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x44,
	0x49, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x52, 0x4f, 0x44, 0x49, 0x12, 0x18,
	0x0a, 0x07, 0x4e, 0x41, 0x4d, 0x41, 0x50, 0x53, 0x54, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4e, 0x41, 0x4d, 0x41, 0x50, 0x53, 0x54, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x4f, 0x44, 0x45,
	0x50, 0x53, 0x54, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x4f, 0x44, 0x45, 0x50,
	0x53, 0x54, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x4f, 0x44, 0x45, 0x46, 0x41, 0x4b, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x4f, 0x44, 0x45, 0x46, 0x41, 0x4b, 0x12, 0x18, 0x0a, 0x07,
	0x4e, 0x41, 0x4d, 0x41, 0x46, 0x41, 0x4b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e,
	0x41, 0x4d, 0x41, 0x46, 0x41, 0x4b, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x4c, 0x52, 0x4d, 0x41, 0x53,
	0x55, 0x4b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4a, 0x4c, 0x52, 0x4d, 0x41, 0x53,
	0x55, 0x4b, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x48, 0x4e, 0x4d, 0x41, 0x53, 0x55, 0x4b, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x48, 0x4e, 0x4d, 0x41, 0x53, 0x55, 0x4b, 0x12, 0x18,
	0x0a, 0x07, 0x4c, 0x41, 0x4d, 0x41, 0x53, 0x54, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4c, 0x41, 0x4d, 0x41, 0x53, 0x54, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x47, 0x4c, 0x53,
	0x49, 0x44, 0x41, 0x4e, 0x47, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x47, 0x4c,
	0x53, 0x49, 0x44, 0x41, 0x4e, 0x47, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x4f, 0x44, 0x45, 0x4a, 0x4b,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4b, 0x4f, 0x44, 0x45, 0x4a, 0x4b, 0x12, 0x1e,
	0x0a, 0x0a, 0x4b, 0x4f, 0x44, 0x45, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4b, 0x4f, 0x44, 0x45, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x22, 0x25,
	0x0a, 0x11, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6e, 0x69, 0x6d, 0x22, 0x75, 0x0a, 0x12, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x68, 0x73, 0x42,
	0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x16,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x68, 0x73, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69,
	0x32, 0xdf, 0x01, 0x0a, 0x11, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d,
	0x68, 0x73, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x79, 0x4e, 0x69, 0x6d, 0x12, 0x24,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x68, 0x73, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x12,
	0x24, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x68, 0x73, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d,
	0x68, 0x73, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mhsbiodata_proto_rawDescOnce sync.Once
	file_mhsbiodata_proto_rawDescData = file_mhsbiodata_proto_rawDesc
)

func file_mhsbiodata_proto_rawDescGZIP() []byte {
	file_mhsbiodata_proto_rawDescOnce.Do(func() {
		file_mhsbiodata_proto_rawDescData = protoimpl.X.CompressGZIP(file_mhsbiodata_proto_rawDescData)
	})
	return file_mhsbiodata_proto_rawDescData
}

var file_mhsbiodata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mhsbiodata_proto_goTypes = []interface{}{
	(*MhsBiodata)(nil),             // 0: tracer_study_grpc.MhsBiodata
	(*MhsBiodataRequest)(nil),      // 1: tracer_study_grpc.MhsBiodataRequest
	(*MhsBiodataResponse)(nil),     // 2: tracer_study_grpc.MhsBiodataResponse
	(*CheckMhsAlumniResponse)(nil), // 3: tracer_study_grpc.CheckMhsAlumniResponse
}
var file_mhsbiodata_proto_depIdxs = []int32{
	0, // 0: tracer_study_grpc.MhsBiodataResponse.data:type_name -> tracer_study_grpc.MhsBiodata
	1, // 1: tracer_study_grpc.MhsBiodataService.FetchMhsBiodataByNim:input_type -> tracer_study_grpc.MhsBiodataRequest
	1, // 2: tracer_study_grpc.MhsBiodataService.CheckMhsAlumni:input_type -> tracer_study_grpc.MhsBiodataRequest
	2, // 3: tracer_study_grpc.MhsBiodataService.FetchMhsBiodataByNim:output_type -> tracer_study_grpc.MhsBiodataResponse
	3, // 4: tracer_study_grpc.MhsBiodataService.CheckMhsAlumni:output_type -> tracer_study_grpc.CheckMhsAlumniResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mhsbiodata_proto_init() }
func file_mhsbiodata_proto_init() {
	if File_mhsbiodata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mhsbiodata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MhsBiodata); i {
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
		file_mhsbiodata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MhsBiodataRequest); i {
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
		file_mhsbiodata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MhsBiodataResponse); i {
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
		file_mhsbiodata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckMhsAlumniResponse); i {
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
			RawDescriptor: file_mhsbiodata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mhsbiodata_proto_goTypes,
		DependencyIndexes: file_mhsbiodata_proto_depIdxs,
		MessageInfos:      file_mhsbiodata_proto_msgTypes,
	}.Build()
	File_mhsbiodata_proto = out.File
	file_mhsbiodata_proto_rawDesc = nil
	file_mhsbiodata_proto_goTypes = nil
	file_mhsbiodata_proto_depIdxs = nil
}
