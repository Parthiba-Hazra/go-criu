// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: binfmt-misc.proto

package binfmt_misc

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

type BinfmtMiscEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Enabled     *bool   `protobuf:"varint,2,req,name=enabled" json:"enabled,omitempty"`
	Interpreter *string `protobuf:"bytes,3,req,name=interpreter" json:"interpreter,omitempty"`
	Flags       *string `protobuf:"bytes,4,opt,name=flags" json:"flags,omitempty"`
	Extension   *string `protobuf:"bytes,5,opt,name=extension" json:"extension,omitempty"`
	Magic       *string `protobuf:"bytes,6,opt,name=magic" json:"magic,omitempty"`
	Mask        *string `protobuf:"bytes,7,opt,name=mask" json:"mask,omitempty"`
	Offset      *int32  `protobuf:"varint,8,opt,name=offset" json:"offset,omitempty"`
}

func (x *BinfmtMiscEntry) Reset() {
	*x = BinfmtMiscEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binfmt_misc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinfmtMiscEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinfmtMiscEntry) ProtoMessage() {}

func (x *BinfmtMiscEntry) ProtoReflect() protoreflect.Message {
	mi := &file_binfmt_misc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinfmtMiscEntry.ProtoReflect.Descriptor instead.
func (*BinfmtMiscEntry) Descriptor() ([]byte, []int) {
	return file_binfmt_misc_proto_rawDescGZIP(), []int{0}
}

func (x *BinfmtMiscEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BinfmtMiscEntry) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *BinfmtMiscEntry) GetInterpreter() string {
	if x != nil && x.Interpreter != nil {
		return *x.Interpreter
	}
	return ""
}

func (x *BinfmtMiscEntry) GetFlags() string {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return ""
}

func (x *BinfmtMiscEntry) GetExtension() string {
	if x != nil && x.Extension != nil {
		return *x.Extension
	}
	return ""
}

func (x *BinfmtMiscEntry) GetMagic() string {
	if x != nil && x.Magic != nil {
		return *x.Magic
	}
	return ""
}

func (x *BinfmtMiscEntry) GetMask() string {
	if x != nil && x.Mask != nil {
		return *x.Mask
	}
	return ""
}

func (x *BinfmtMiscEntry) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

var File_binfmt_misc_proto protoreflect.FileDescriptor

var file_binfmt_misc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x69, 0x6e, 0x66, 0x6d, 0x74, 0x2d, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x11, 0x62, 0x69, 0x6e, 0x66, 0x6d, 0x74, 0x5f, 0x6d,
	0x69, 0x73, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61,
	0x67, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
}

var (
	file_binfmt_misc_proto_rawDescOnce sync.Once
	file_binfmt_misc_proto_rawDescData = file_binfmt_misc_proto_rawDesc
)

func file_binfmt_misc_proto_rawDescGZIP() []byte {
	file_binfmt_misc_proto_rawDescOnce.Do(func() {
		file_binfmt_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_binfmt_misc_proto_rawDescData)
	})
	return file_binfmt_misc_proto_rawDescData
}

var file_binfmt_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_binfmt_misc_proto_goTypes = []interface{}{
	(*BinfmtMiscEntry)(nil), // 0: binfmt_misc_entry
}
var file_binfmt_misc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_binfmt_misc_proto_init() }
func file_binfmt_misc_proto_init() {
	if File_binfmt_misc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_binfmt_misc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinfmtMiscEntry); i {
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
			RawDescriptor: file_binfmt_misc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binfmt_misc_proto_goTypes,
		DependencyIndexes: file_binfmt_misc_proto_depIdxs,
		MessageInfos:      file_binfmt_misc_proto_msgTypes,
	}.Build()
	File_binfmt_misc_proto = out.File
	file_binfmt_misc_proto_rawDesc = nil
	file_binfmt_misc_proto_goTypes = nil
	file_binfmt_misc_proto_depIdxs = nil
}
