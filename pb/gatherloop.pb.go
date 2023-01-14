// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.8
// source: proto/gatherloop.proto

package gatherloop

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

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gatherloop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gatherloop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_gatherloop_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BuatSeminarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pembicara string `protobuf:"bytes,1,opt,name=pembicara,proto3" json:"pembicara,omitempty"`
	Acara     string `protobuf:"bytes,2,opt,name=acara,proto3" json:"acara,omitempty"`
	Waktu     string `protobuf:"bytes,3,opt,name=waktu,proto3" json:"waktu,omitempty"`
	Lokasi    string `protobuf:"bytes,4,opt,name=lokasi,proto3" json:"lokasi,omitempty"`
	Berbayar  bool   `protobuf:"varint,5,opt,name=berbayar,proto3" json:"berbayar,omitempty"`
}

func (x *BuatSeminarRequest) Reset() {
	*x = BuatSeminarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gatherloop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuatSeminarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuatSeminarRequest) ProtoMessage() {}

func (x *BuatSeminarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gatherloop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuatSeminarRequest.ProtoReflect.Descriptor instead.
func (*BuatSeminarRequest) Descriptor() ([]byte, []int) {
	return file_proto_gatherloop_proto_rawDescGZIP(), []int{1}
}

func (x *BuatSeminarRequest) GetPembicara() string {
	if x != nil {
		return x.Pembicara
	}
	return ""
}

func (x *BuatSeminarRequest) GetAcara() string {
	if x != nil {
		return x.Acara
	}
	return ""
}

func (x *BuatSeminarRequest) GetWaktu() string {
	if x != nil {
		return x.Waktu
	}
	return ""
}

func (x *BuatSeminarRequest) GetLokasi() string {
	if x != nil {
		return x.Lokasi
	}
	return ""
}

func (x *BuatSeminarRequest) GetBerbayar() bool {
	if x != nil {
		return x.Berbayar
	}
	return false
}

type GetSeminarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lokasi string `protobuf:"bytes,1,opt,name=lokasi,proto3" json:"lokasi,omitempty"`
}

func (x *GetSeminarRequest) Reset() {
	*x = GetSeminarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gatherloop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeminarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeminarRequest) ProtoMessage() {}

func (x *GetSeminarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gatherloop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeminarRequest.ProtoReflect.Descriptor instead.
func (*GetSeminarRequest) Descriptor() ([]byte, []int) {
	return file_proto_gatherloop_proto_rawDescGZIP(), []int{2}
}

func (x *GetSeminarRequest) GetLokasi() string {
	if x != nil {
		return x.Lokasi
	}
	return ""
}

type Seminar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pembicara string `protobuf:"bytes,1,opt,name=pembicara,proto3" json:"pembicara,omitempty"`
	Acara     string `protobuf:"bytes,2,opt,name=acara,proto3" json:"acara,omitempty"`
	Waktu     string `protobuf:"bytes,3,opt,name=waktu,proto3" json:"waktu,omitempty"`
	Lokasi    string `protobuf:"bytes,4,opt,name=lokasi,proto3" json:"lokasi,omitempty"`
	Berbayar  bool   `protobuf:"varint,5,opt,name=berbayar,proto3" json:"berbayar,omitempty"`
}

func (x *Seminar) Reset() {
	*x = Seminar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gatherloop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seminar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seminar) ProtoMessage() {}

func (x *Seminar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gatherloop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seminar.ProtoReflect.Descriptor instead.
func (*Seminar) Descriptor() ([]byte, []int) {
	return file_proto_gatherloop_proto_rawDescGZIP(), []int{3}
}

func (x *Seminar) GetPembicara() string {
	if x != nil {
		return x.Pembicara
	}
	return ""
}

func (x *Seminar) GetAcara() string {
	if x != nil {
		return x.Acara
	}
	return ""
}

func (x *Seminar) GetWaktu() string {
	if x != nil {
		return x.Waktu
	}
	return ""
}

func (x *Seminar) GetLokasi() string {
	if x != nil {
		return x.Lokasi
	}
	return ""
}

func (x *Seminar) GetBerbayar() bool {
	if x != nil {
		return x.Berbayar
	}
	return false
}

type GetSeminarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seminars []*Seminar `protobuf:"bytes,1,rep,name=seminars,proto3" json:"seminars,omitempty"`
}

func (x *GetSeminarResponse) Reset() {
	*x = GetSeminarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gatherloop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeminarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeminarResponse) ProtoMessage() {}

func (x *GetSeminarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gatherloop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeminarResponse.ProtoReflect.Descriptor instead.
func (*GetSeminarResponse) Descriptor() ([]byte, []int) {
	return file_proto_gatherloop_proto_rawDescGZIP(), []int{4}
}

func (x *GetSeminarResponse) GetSeminars() []*Seminar {
	if x != nil {
		return x.Seminars
	}
	return nil
}

var File_proto_gatherloop_proto protoreflect.FileDescriptor

var file_proto_gatherloop_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x6c, 0x6f, 0x6f, 0x70, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x92, 0x01, 0x0a, 0x12, 0x42, 0x75, 0x61, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x6d, 0x62, 0x69, 0x63,
	0x61, 0x72, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x6d, 0x62, 0x69,
	0x63, 0x61, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x61, 0x72, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x61, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61,
	0x6b, 0x74, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x61, 0x6b, 0x74, 0x75,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x6b, 0x61, 0x73, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x6b, 0x61, 0x73, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x72, 0x62,
	0x61, 0x79, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x65, 0x72, 0x62,
	0x61, 0x79, 0x61, 0x72, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x6b,
	0x61, 0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x6b, 0x61, 0x73,
	0x69, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x65, 0x6d, 0x62, 0x69, 0x63, 0x61, 0x72, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x65, 0x6d, 0x62, 0x69, 0x63, 0x61, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x63, 0x61, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x61, 0x72,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x6b, 0x74, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x77, 0x61, 0x6b, 0x74, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x6b, 0x61, 0x73,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x6b, 0x61, 0x73, 0x69, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x65, 0x72, 0x62, 0x61, 0x79, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x62, 0x65, 0x72, 0x62, 0x61, 0x79, 0x61, 0x72, 0x22, 0x45, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61,
	0x72, 0x73, 0x32, 0xac, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x42, 0x75, 0x61, 0x74,
	0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x75, 0x61, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e,
	0x61, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x3b, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6c,
	0x6f, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gatherloop_proto_rawDescOnce sync.Once
	file_proto_gatherloop_proto_rawDescData = file_proto_gatherloop_proto_rawDesc
)

func file_proto_gatherloop_proto_rawDescGZIP() []byte {
	file_proto_gatherloop_proto_rawDescOnce.Do(func() {
		file_proto_gatherloop_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gatherloop_proto_rawDescData)
	})
	return file_proto_gatherloop_proto_rawDescData
}

var file_proto_gatherloop_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_gatherloop_proto_goTypes = []interface{}{
	(*ServiceResponse)(nil),    // 0: gatherloop.ServiceResponse
	(*BuatSeminarRequest)(nil), // 1: gatherloop.BuatSeminarRequest
	(*GetSeminarRequest)(nil),  // 2: gatherloop.GetSeminarRequest
	(*Seminar)(nil),            // 3: gatherloop.Seminar
	(*GetSeminarResponse)(nil), // 4: gatherloop.GetSeminarResponse
}
var file_proto_gatherloop_proto_depIdxs = []int32{
	3, // 0: gatherloop.GetSeminarResponse.seminars:type_name -> gatherloop.Seminar
	1, // 1: gatherloop.GatherloopService.BuatSeminar:input_type -> gatherloop.BuatSeminarRequest
	2, // 2: gatherloop.GatherloopService.GetSeminar:input_type -> gatherloop.GetSeminarRequest
	0, // 3: gatherloop.GatherloopService.BuatSeminar:output_type -> gatherloop.ServiceResponse
	4, // 4: gatherloop.GatherloopService.GetSeminar:output_type -> gatherloop.GetSeminarResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_gatherloop_proto_init() }
func file_proto_gatherloop_proto_init() {
	if File_proto_gatherloop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gatherloop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_proto_gatherloop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuatSeminarRequest); i {
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
		file_proto_gatherloop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeminarRequest); i {
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
		file_proto_gatherloop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seminar); i {
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
		file_proto_gatherloop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeminarResponse); i {
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
			RawDescriptor: file_proto_gatherloop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gatherloop_proto_goTypes,
		DependencyIndexes: file_proto_gatherloop_proto_depIdxs,
		MessageInfos:      file_proto_gatherloop_proto_msgTypes,
	}.Build()
	File_proto_gatherloop_proto = out.File
	file_proto_gatherloop_proto_rawDesc = nil
	file_proto_gatherloop_proto_goTypes = nil
	file_proto_gatherloop_proto_depIdxs = nil
}
