// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: find.proto

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

type FindMessage_MessageType int32

const (
	FindMessage_ERROR_RESPONSE          FindMessage_MessageType = 0
	FindMessage_FIND                    FindMessage_MessageType = 1
	FindMessage_FIND_RESPONSE           FindMessage_MessageType = 2
	FindMessage_LIST_PROVIDERS          FindMessage_MessageType = 3
	FindMessage_LIST_PROVIDERS_RESPONSE FindMessage_MessageType = 4
	FindMessage_GET_PROVIDER            FindMessage_MessageType = 5
	FindMessage_GET_PROVIDER_RESPONSE   FindMessage_MessageType = 6
	FindMessage_GET_STATS               FindMessage_MessageType = 7
	FindMessage_GET_STATS_RESPONSE      FindMessage_MessageType = 8
)

// Enum value maps for FindMessage_MessageType.
var (
	FindMessage_MessageType_name = map[int32]string{
		0: "ERROR_RESPONSE",
		1: "FIND",
		2: "FIND_RESPONSE",
		3: "LIST_PROVIDERS",
		4: "LIST_PROVIDERS_RESPONSE",
		5: "GET_PROVIDER",
		6: "GET_PROVIDER_RESPONSE",
		7: "GET_STATS",
		8: "GET_STATS_RESPONSE",
	}
	FindMessage_MessageType_value = map[string]int32{
		"ERROR_RESPONSE":          0,
		"FIND":                    1,
		"FIND_RESPONSE":           2,
		"LIST_PROVIDERS":          3,
		"LIST_PROVIDERS_RESPONSE": 4,
		"GET_PROVIDER":            5,
		"GET_PROVIDER_RESPONSE":   6,
		"GET_STATS":               7,
		"GET_STATS_RESPONSE":      8,
	}
)

func (x FindMessage_MessageType) Enum() *FindMessage_MessageType {
	p := new(FindMessage_MessageType)
	*p = x
	return p
}

func (x FindMessage_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FindMessage_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_find_proto_enumTypes[0].Descriptor()
}

func (FindMessage_MessageType) Type() protoreflect.EnumType {
	return &file_find_proto_enumTypes[0]
}

func (x FindMessage_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FindMessage_MessageType.Descriptor instead.
func (FindMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_find_proto_rawDescGZIP(), []int{0, 0}
}

type FindMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// defines what type of message it is.
	Type FindMessage_MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=reqresp.pb.FindMessage_MessageType" json:"type,omitempty"`
	// Value for the message
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindMessage) Reset() {
	*x = FindMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_find_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMessage) ProtoMessage() {}

func (x *FindMessage) ProtoReflect() protoreflect.Message {
	mi := &file_find_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMessage.ProtoReflect.Descriptor instead.
func (*FindMessage) Descriptor() ([]byte, []int) {
	return file_find_proto_rawDescGZIP(), []int{0}
}

func (x *FindMessage) GetType() FindMessage_MessageType {
	if x != nil {
		return x.Type
	}
	return FindMessage_ERROR_RESPONSE
}

func (x *FindMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_find_proto protoreflect.FileDescriptor

var file_find_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65,
	0x71, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x62, 0x22, 0xa0, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x71, 0x72, 0x65, 0x73, 0x70,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4e,
	0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x53, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x49,
	0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x53, 0x5f, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x54, 0x5f, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x53, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x08, 0x42, 0x24, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x6e, 0x69, 0x2f, 0x67,
	0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x69, 0x70, 0x6e, 0x69, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_find_proto_rawDescOnce sync.Once
	file_find_proto_rawDescData = file_find_proto_rawDesc
)

func file_find_proto_rawDescGZIP() []byte {
	file_find_proto_rawDescOnce.Do(func() {
		file_find_proto_rawDescData = protoimpl.X.CompressGZIP(file_find_proto_rawDescData)
	})
	return file_find_proto_rawDescData
}

var file_find_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_find_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_find_proto_goTypes = []interface{}{
	(FindMessage_MessageType)(0), // 0: reqresp.pb.FindMessage.MessageType
	(*FindMessage)(nil),          // 1: reqresp.pb.FindMessage
}
var file_find_proto_depIdxs = []int32{
	0, // 0: reqresp.pb.FindMessage.type:type_name -> reqresp.pb.FindMessage.MessageType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_find_proto_init() }
func file_find_proto_init() {
	if File_find_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_find_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMessage); i {
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
			RawDescriptor: file_find_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_find_proto_goTypes,
		DependencyIndexes: file_find_proto_depIdxs,
		EnumInfos:         file_find_proto_enumTypes,
		MessageInfos:      file_find_proto_msgTypes,
	}.Build()
	File_find_proto = out.File
	file_find_proto_rawDesc = nil
	file_find_proto_goTypes = nil
	file_find_proto_depIdxs = nil
}
