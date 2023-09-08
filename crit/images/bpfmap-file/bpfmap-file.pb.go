// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: bpfmap-file.proto

package bpfmap_file

import (
	fown "github.com/checkpoint-restore/go-criu/v7/crit/images/fown"
	_ "github.com/checkpoint-restore/go-criu/v7/crit/images/opts"
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

type BpfmapFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint32         `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags      *uint32         `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Pos        *uint64         `protobuf:"varint,3,req,name=pos" json:"pos,omitempty"`
	Fown       *fown.FownEntry `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
	MapType    *uint32         `protobuf:"varint,5,req,name=map_type,json=mapType" json:"map_type,omitempty"`
	KeySize    *uint32         `protobuf:"varint,6,req,name=key_size,json=keySize" json:"key_size,omitempty"`
	ValueSize  *uint32         `protobuf:"varint,7,req,name=value_size,json=valueSize" json:"value_size,omitempty"`
	MapId      *uint32         `protobuf:"varint,8,req,name=map_id,json=mapId" json:"map_id,omitempty"`
	MaxEntries *uint32         `protobuf:"varint,9,req,name=max_entries,json=maxEntries" json:"max_entries,omitempty"`
	MapFlags   *uint32         `protobuf:"varint,10,req,name=map_flags,json=mapFlags" json:"map_flags,omitempty"`
	Memlock    *uint64         `protobuf:"varint,11,req,name=memlock" json:"memlock,omitempty"`
	Frozen     *bool           `protobuf:"varint,12,req,name=frozen,def=0" json:"frozen,omitempty"`
	MapName    *string         `protobuf:"bytes,13,req,name=map_name,json=mapName" json:"map_name,omitempty"`
	Ifindex    *uint32         `protobuf:"varint,14,req,name=ifindex,def=0" json:"ifindex,omitempty"`
	MntId      *int32          `protobuf:"zigzag32,15,opt,name=mnt_id,json=mntId,def=-1" json:"mnt_id,omitempty"`
	MapExtra   *uint64         `protobuf:"varint,16,opt,name=map_extra,json=mapExtra" json:"map_extra,omitempty"`
}

// Default values for BpfmapFileEntry fields.
const (
	Default_BpfmapFileEntry_Frozen  = bool(false)
	Default_BpfmapFileEntry_Ifindex = uint32(0)
	Default_BpfmapFileEntry_MntId   = int32(-1)
)

func (x *BpfmapFileEntry) Reset() {
	*x = BpfmapFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfmap_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BpfmapFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BpfmapFileEntry) ProtoMessage() {}

func (x *BpfmapFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_bpfmap_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BpfmapFileEntry.ProtoReflect.Descriptor instead.
func (*BpfmapFileEntry) Descriptor() ([]byte, []int) {
	return file_bpfmap_file_proto_rawDescGZIP(), []int{0}
}

func (x *BpfmapFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BpfmapFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *BpfmapFileEntry) GetPos() uint64 {
	if x != nil && x.Pos != nil {
		return *x.Pos
	}
	return 0
}

func (x *BpfmapFileEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *BpfmapFileEntry) GetMapType() uint32 {
	if x != nil && x.MapType != nil {
		return *x.MapType
	}
	return 0
}

func (x *BpfmapFileEntry) GetKeySize() uint32 {
	if x != nil && x.KeySize != nil {
		return *x.KeySize
	}
	return 0
}

func (x *BpfmapFileEntry) GetValueSize() uint32 {
	if x != nil && x.ValueSize != nil {
		return *x.ValueSize
	}
	return 0
}

func (x *BpfmapFileEntry) GetMapId() uint32 {
	if x != nil && x.MapId != nil {
		return *x.MapId
	}
	return 0
}

func (x *BpfmapFileEntry) GetMaxEntries() uint32 {
	if x != nil && x.MaxEntries != nil {
		return *x.MaxEntries
	}
	return 0
}

func (x *BpfmapFileEntry) GetMapFlags() uint32 {
	if x != nil && x.MapFlags != nil {
		return *x.MapFlags
	}
	return 0
}

func (x *BpfmapFileEntry) GetMemlock() uint64 {
	if x != nil && x.Memlock != nil {
		return *x.Memlock
	}
	return 0
}

func (x *BpfmapFileEntry) GetFrozen() bool {
	if x != nil && x.Frozen != nil {
		return *x.Frozen
	}
	return Default_BpfmapFileEntry_Frozen
}

func (x *BpfmapFileEntry) GetMapName() string {
	if x != nil && x.MapName != nil {
		return *x.MapName
	}
	return ""
}

func (x *BpfmapFileEntry) GetIfindex() uint32 {
	if x != nil && x.Ifindex != nil {
		return *x.Ifindex
	}
	return Default_BpfmapFileEntry_Ifindex
}

func (x *BpfmapFileEntry) GetMntId() int32 {
	if x != nil && x.MntId != nil {
		return *x.MntId
	}
	return Default_BpfmapFileEntry_MntId
}

func (x *BpfmapFileEntry) GetMapExtra() uint64 {
	if x != nil && x.MapExtra != nil {
		return *x.MapExtra
	}
	return 0
}

var File_bpfmap_file_proto protoreflect.FileDescriptor

var file_bpfmap_file_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x70, 0x66, 0x6d, 0x61, 0x70, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x03, 0x0a, 0x11,
	0x62, 0x70, 0x66, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x42, 0x10, 0xd2, 0x3f, 0x0d, 0x1a, 0x0b, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x66,
	0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6f, 0x77, 0x6e,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x61, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07,
	0x6d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61,
	0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x1d, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x08,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x69, 0x66,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x02, 0x28, 0x0d, 0x3a, 0x01, 0x30, 0x52, 0x07,
	0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x06, 0x6d, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x11, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x05, 0x6d, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61,
}

var (
	file_bpfmap_file_proto_rawDescOnce sync.Once
	file_bpfmap_file_proto_rawDescData = file_bpfmap_file_proto_rawDesc
)

func file_bpfmap_file_proto_rawDescGZIP() []byte {
	file_bpfmap_file_proto_rawDescOnce.Do(func() {
		file_bpfmap_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_bpfmap_file_proto_rawDescData)
	})
	return file_bpfmap_file_proto_rawDescData
}

var file_bpfmap_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bpfmap_file_proto_goTypes = []interface{}{
	(*BpfmapFileEntry)(nil), // 0: bpfmap_file_entry
	(*fown.FownEntry)(nil),  // 1: fown_entry
}
var file_bpfmap_file_proto_depIdxs = []int32{
	1, // 0: bpfmap_file_entry.fown:type_name -> fown_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bpfmap_file_proto_init() }
func file_bpfmap_file_proto_init() {
	if File_bpfmap_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bpfmap_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BpfmapFileEntry); i {
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
			RawDescriptor: file_bpfmap_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bpfmap_file_proto_goTypes,
		DependencyIndexes: file_bpfmap_file_proto_depIdxs,
		MessageInfos:      file_bpfmap_file_proto_msgTypes,
	}.Build()
	File_bpfmap_file_proto = out.File
	file_bpfmap_file_proto_rawDesc = nil
	file_bpfmap_file_proto_goTypes = nil
	file_bpfmap_file_proto_depIdxs = nil
}
