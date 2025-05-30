// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: kafkalib/msg/tenant_meta.proto

package msg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TenantMeta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	TenantName    string                 `protobuf:"bytes,2,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TenantMeta) Reset() {
	*x = TenantMeta{}
	mi := &file_kafkalib_msg_tenant_meta_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantMeta) ProtoMessage() {}

func (x *TenantMeta) ProtoReflect() protoreflect.Message {
	mi := &file_kafkalib_msg_tenant_meta_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantMeta.ProtoReflect.Descriptor instead.
func (*TenantMeta) Descriptor() ([]byte, []int) {
	return file_kafkalib_msg_tenant_meta_proto_rawDescGZIP(), []int{0}
}

func (x *TenantMeta) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *TenantMeta) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

var File_kafkalib_msg_tenant_meta_proto protoreflect.FileDescriptor

const file_kafkalib_msg_tenant_meta_proto_rawDesc = "" +
	"\n" +
	"\x1ekafkalib/msg/tenant_meta.proto\x12\fkafkalib.msg\"J\n" +
	"\n" +
	"TenantMeta\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12\x1f\n" +
	"\vtenant_name\x18\x02 \x01(\tR\n" +
	"tenantNameB\xa6\x01\n" +
	"\x10com.kafkalib.msgB\x0fTenantMetaProtoP\x01Z0github.com/sambatechno/kafkalib/gen/kafkalib/msg\xa2\x02\x03KMX\xaa\x02\fKafkalib.Msg\xca\x02\fKafkalib\\Msg\xe2\x02\x18Kafkalib\\Msg\\GPBMetadata\xea\x02\rKafkalib::Msgb\x06proto3"

var (
	file_kafkalib_msg_tenant_meta_proto_rawDescOnce sync.Once
	file_kafkalib_msg_tenant_meta_proto_rawDescData []byte
)

func file_kafkalib_msg_tenant_meta_proto_rawDescGZIP() []byte {
	file_kafkalib_msg_tenant_meta_proto_rawDescOnce.Do(func() {
		file_kafkalib_msg_tenant_meta_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kafkalib_msg_tenant_meta_proto_rawDesc), len(file_kafkalib_msg_tenant_meta_proto_rawDesc)))
	})
	return file_kafkalib_msg_tenant_meta_proto_rawDescData
}

var file_kafkalib_msg_tenant_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kafkalib_msg_tenant_meta_proto_goTypes = []any{
	(*TenantMeta)(nil), // 0: kafkalib.msg.TenantMeta
}
var file_kafkalib_msg_tenant_meta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kafkalib_msg_tenant_meta_proto_init() }
func file_kafkalib_msg_tenant_meta_proto_init() {
	if File_kafkalib_msg_tenant_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kafkalib_msg_tenant_meta_proto_rawDesc), len(file_kafkalib_msg_tenant_meta_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafkalib_msg_tenant_meta_proto_goTypes,
		DependencyIndexes: file_kafkalib_msg_tenant_meta_proto_depIdxs,
		MessageInfos:      file_kafkalib_msg_tenant_meta_proto_msgTypes,
	}.Build()
	File_kafkalib_msg_tenant_meta_proto = out.File
	file_kafkalib_msg_tenant_meta_proto_goTypes = nil
	file_kafkalib_msg_tenant_meta_proto_depIdxs = nil
}
