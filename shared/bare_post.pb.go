// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: shared/bare_post.proto

package shared

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

type BarePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uuid of the bare post
	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The audit fields of the post
	AuditFields *AuditFields `protobuf:"bytes,2,opt,name=audit_fields,json=auditFields,proto3" json:"audit_fields,omitempty"`
	// The main comment of the post
	Comment *BarePost_Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	// The link provided in the post
	Link string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *BarePost) Reset() {
	*x = BarePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_bare_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarePost) ProtoMessage() {}

func (x *BarePost) ProtoReflect() protoreflect.Message {
	mi := &file_shared_bare_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarePost.ProtoReflect.Descriptor instead.
func (*BarePost) Descriptor() ([]byte, []int) {
	return file_shared_bare_post_proto_rawDescGZIP(), []int{0}
}

func (x *BarePost) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *BarePost) GetAuditFields() *AuditFields {
	if x != nil {
		return x.AuditFields
	}
	return nil
}

func (x *BarePost) GetComment() *BarePost_Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *BarePost) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type BarePost_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The primary content field of the post
	PrimaryContent string `protobuf:"bytes,1,opt,name=primary_content,json=primaryContent,proto3" json:"primary_content,omitempty"`
}

func (x *BarePost_Comment) Reset() {
	*x = BarePost_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_bare_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarePost_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarePost_Comment) ProtoMessage() {}

func (x *BarePost_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_shared_bare_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarePost_Comment.ProtoReflect.Descriptor instead.
func (*BarePost_Comment) Descriptor() ([]byte, []int) {
	return file_shared_bare_post_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BarePost_Comment) GetPrimaryContent() string {
	if x != nil {
		return x.PrimaryContent
	}
	return ""
}

var File_shared_bare_post_proto protoreflect.FileDescriptor

var file_shared_bare_post_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x72, 0x63, 0x61, 0x62, 0x6c,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x19, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x72, 0x63, 0x61,
	0x62, 0x6c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x72, 0x63, 0x61, 0x62, 0x6c, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x42, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x1a, 0x32, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x61, 0x62, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_bare_post_proto_rawDescOnce sync.Once
	file_shared_bare_post_proto_rawDescData = file_shared_bare_post_proto_rawDesc
)

func file_shared_bare_post_proto_rawDescGZIP() []byte {
	file_shared_bare_post_proto_rawDescOnce.Do(func() {
		file_shared_bare_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_bare_post_proto_rawDescData)
	})
	return file_shared_bare_post_proto_rawDescData
}

var file_shared_bare_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_bare_post_proto_goTypes = []interface{}{
	(*BarePost)(nil),         // 0: srcabl.shared.BarePost
	(*BarePost_Comment)(nil), // 1: srcabl.shared.BarePost.Comment
	(*AuditFields)(nil),      // 2: srcabl.shared.AuditFields
}
var file_shared_bare_post_proto_depIdxs = []int32{
	2, // 0: srcabl.shared.BarePost.audit_fields:type_name -> srcabl.shared.AuditFields
	1, // 1: srcabl.shared.BarePost.comment:type_name -> srcabl.shared.BarePost.Comment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shared_bare_post_proto_init() }
func file_shared_bare_post_proto_init() {
	if File_shared_bare_post_proto != nil {
		return
	}
	file_shared_audit_fields_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_bare_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarePost); i {
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
		file_shared_bare_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarePost_Comment); i {
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
			RawDescriptor: file_shared_bare_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_bare_post_proto_goTypes,
		DependencyIndexes: file_shared_bare_post_proto_depIdxs,
		MessageInfos:      file_shared_bare_post_proto_msgTypes,
	}.Build()
	File_shared_bare_post_proto = out.File
	file_shared_bare_post_proto_rawDesc = nil
	file_shared_bare_post_proto_goTypes = nil
	file_shared_bare_post_proto_depIdxs = nil
}
