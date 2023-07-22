// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: timens.proto

package timens

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

type Timespec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TvSec  *uint64 `protobuf:"varint,1,req,name=tv_sec,json=tvSec" json:"tv_sec,omitempty"`
	TvNsec *uint64 `protobuf:"varint,2,req,name=tv_nsec,json=tvNsec" json:"tv_nsec,omitempty"`
}

func (x *Timespec) Reset() {
	*x = Timespec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timespec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timespec) ProtoMessage() {}

func (x *Timespec) ProtoReflect() protoreflect.Message {
	mi := &file_timens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timespec.ProtoReflect.Descriptor instead.
func (*Timespec) Descriptor() ([]byte, []int) {
	return file_timens_proto_rawDescGZIP(), []int{0}
}

func (x *Timespec) GetTvSec() uint64 {
	if x != nil && x.TvSec != nil {
		return *x.TvSec
	}
	return 0
}

func (x *Timespec) GetTvNsec() uint64 {
	if x != nil && x.TvNsec != nil {
		return *x.TvNsec
	}
	return 0
}

type TimensEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monotonic *Timespec `protobuf:"bytes,1,req,name=monotonic" json:"monotonic,omitempty"`
	Boottime  *Timespec `protobuf:"bytes,2,req,name=boottime" json:"boottime,omitempty"`
}

func (x *TimensEntry) Reset() {
	*x = TimensEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimensEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimensEntry) ProtoMessage() {}

func (x *TimensEntry) ProtoReflect() protoreflect.Message {
	mi := &file_timens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimensEntry.ProtoReflect.Descriptor instead.
func (*TimensEntry) Descriptor() ([]byte, []int) {
	return file_timens_proto_rawDescGZIP(), []int{1}
}

func (x *TimensEntry) GetMonotonic() *Timespec {
	if x != nil {
		return x.Monotonic
	}
	return nil
}

func (x *TimensEntry) GetBoottime() *Timespec {
	if x != nil {
		return x.Boottime
	}
	return nil
}

var File_timens_proto protoreflect.FileDescriptor

var file_timens_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x65, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x76,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x74, 0x76, 0x53, 0x65,
	0x63, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x76, 0x5f, 0x6e, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x06, 0x74, 0x76, 0x4e, 0x73, 0x65, 0x63, 0x22, 0x5e, 0x0a, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x09, 0x6d, 0x6f,
	0x6e, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x6f, 0x74, 0x6f,
	0x6e, 0x69, 0x63, 0x12, 0x25, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x65, 0x63,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x74, 0x69, 0x6d, 0x65,
}

var (
	file_timens_proto_rawDescOnce sync.Once
	file_timens_proto_rawDescData = file_timens_proto_rawDesc
)

func file_timens_proto_rawDescGZIP() []byte {
	file_timens_proto_rawDescOnce.Do(func() {
		file_timens_proto_rawDescData = protoimpl.X.CompressGZIP(file_timens_proto_rawDescData)
	})
	return file_timens_proto_rawDescData
}

var file_timens_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_timens_proto_goTypes = []interface{}{
	(*Timespec)(nil),    // 0: timespec
	(*TimensEntry)(nil), // 1: timens_entry
}
var file_timens_proto_depIdxs = []int32{
	0, // 0: timens_entry.monotonic:type_name -> timespec
	0, // 1: timens_entry.boottime:type_name -> timespec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_timens_proto_init() }
func file_timens_proto_init() {
	if File_timens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timespec); i {
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
		file_timens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimensEntry); i {
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
			RawDescriptor: file_timens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timens_proto_goTypes,
		DependencyIndexes: file_timens_proto_depIdxs,
		MessageInfos:      file_timens_proto_msgTypes,
	}.Build()
	File_timens_proto = out.File
	file_timens_proto_rawDesc = nil
	file_timens_proto_goTypes = nil
	file_timens_proto_depIdxs = nil
}
