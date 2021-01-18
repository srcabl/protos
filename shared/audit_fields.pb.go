// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: shared/audit_fields.proto

package shared

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AuditFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user uuid that originally created the entity
	CreatedByUuid []byte `protobuf:"bytes,1,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	// The user uuid that last updated the entity
	UpdatedByUuid []byte `protobuf:"bytes,2,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	// The time the entity was created relative to epoc
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The time the entity was updated relative to epoc
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AuditFields) Reset() {
	*x = AuditFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_audit_fields_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditFields) ProtoMessage() {}

func (x *AuditFields) ProtoReflect() protoreflect.Message {
	mi := &file_shared_audit_fields_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditFields.ProtoReflect.Descriptor instead.
func (*AuditFields) Descriptor() ([]byte, []int) {
	return file_shared_audit_fields_proto_rawDescGZIP(), []int{0}
}

func (x *AuditFields) GetCreatedByUuid() []byte {
	if x != nil {
		return x.CreatedByUuid
	}
	return nil
}

func (x *AuditFields) GetUpdatedByUuid() []byte {
	if x != nil {
		return x.UpdatedByUuid
	}
	return nil
}

func (x *AuditFields) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AuditFields) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_shared_audit_fields_proto protoreflect.FileDescriptor

var file_shared_audit_fields_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x72, 0x63,
	0x61, 0x62, 0x6c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0b,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x72, 0x63, 0x61, 0x62, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_audit_fields_proto_rawDescOnce sync.Once
	file_shared_audit_fields_proto_rawDescData = file_shared_audit_fields_proto_rawDesc
)

func file_shared_audit_fields_proto_rawDescGZIP() []byte {
	file_shared_audit_fields_proto_rawDescOnce.Do(func() {
		file_shared_audit_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_audit_fields_proto_rawDescData)
	})
	return file_shared_audit_fields_proto_rawDescData
}

var file_shared_audit_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shared_audit_fields_proto_goTypes = []interface{}{
	(*AuditFields)(nil),           // 0: srcabl.shared.AuditFields
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_shared_audit_fields_proto_depIdxs = []int32{
	1, // 0: srcabl.shared.AuditFields.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: srcabl.shared.AuditFields.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shared_audit_fields_proto_init() }
func file_shared_audit_fields_proto_init() {
	if File_shared_audit_fields_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_audit_fields_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditFields); i {
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
			RawDescriptor: file_shared_audit_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_audit_fields_proto_goTypes,
		DependencyIndexes: file_shared_audit_fields_proto_depIdxs,
		MessageInfos:      file_shared_audit_fields_proto_msgTypes,
	}.Build()
	File_shared_audit_fields_proto = out.File
	file_shared_audit_fields_proto_rawDesc = nil
	file_shared_audit_fields_proto_goTypes = nil
	file_shared_audit_fields_proto_depIdxs = nil
}
