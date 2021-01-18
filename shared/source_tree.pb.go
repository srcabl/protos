// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: shared/source_tree.proto

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

type SourceTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceRoot *SourceNode `protobuf:"bytes,1,opt,name=source_root,json=sourceRoot,proto3" json:"source_root,omitempty"`
}

func (x *SourceTree) Reset() {
	*x = SourceTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_source_tree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTree) ProtoMessage() {}

func (x *SourceTree) ProtoReflect() protoreflect.Message {
	mi := &file_shared_source_tree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTree.ProtoReflect.Descriptor instead.
func (*SourceTree) Descriptor() ([]byte, []int) {
	return file_shared_source_tree_proto_rawDescGZIP(), []int{0}
}

func (x *SourceTree) GetSourceRoot() *SourceNode {
	if x != nil {
		return x.SourceRoot
	}
	return nil
}

var File_shared_source_tree_proto protoreflect.FileDescriptor

var file_shared_source_tree_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x72, 0x63, 0x61,
	0x62, 0x6c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x18, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x63, 0x61, 0x62, 0x6c, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x61,
	0x62, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_source_tree_proto_rawDescOnce sync.Once
	file_shared_source_tree_proto_rawDescData = file_shared_source_tree_proto_rawDesc
)

func file_shared_source_tree_proto_rawDescGZIP() []byte {
	file_shared_source_tree_proto_rawDescOnce.Do(func() {
		file_shared_source_tree_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_source_tree_proto_rawDescData)
	})
	return file_shared_source_tree_proto_rawDescData
}

var file_shared_source_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shared_source_tree_proto_goTypes = []interface{}{
	(*SourceTree)(nil), // 0: srcabl.shared.SourceTree
	(*SourceNode)(nil), // 1: srcabl.shared.SourceNode
}
var file_shared_source_tree_proto_depIdxs = []int32{
	1, // 0: srcabl.shared.SourceTree.source_root:type_name -> srcabl.shared.SourceNode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_source_tree_proto_init() }
func file_shared_source_tree_proto_init() {
	if File_shared_source_tree_proto != nil {
		return
	}
	file_shared_source_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_source_tree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceTree); i {
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
			RawDescriptor: file_shared_source_tree_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_source_tree_proto_goTypes,
		DependencyIndexes: file_shared_source_tree_proto_depIdxs,
		MessageInfos:      file_shared_source_tree_proto_msgTypes,
	}.Build()
	File_shared_source_tree_proto = out.File
	file_shared_source_tree_proto_rawDesc = nil
	file_shared_source_tree_proto_goTypes = nil
	file_shared_source_tree_proto_depIdxs = nil
}