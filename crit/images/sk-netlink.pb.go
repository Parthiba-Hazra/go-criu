// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: sk-netlink.proto

package images

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

type NetlinkSkEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint32      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Ino       *uint32      `protobuf:"varint,2,req,name=ino" json:"ino,omitempty"`
	Protocol  *uint32      `protobuf:"varint,3,req,name=protocol" json:"protocol,omitempty"`
	State     *uint32      `protobuf:"varint,4,req,name=state" json:"state,omitempty"`
	Flags     *uint32      `protobuf:"varint,6,req,name=flags" json:"flags,omitempty"`
	Portid    *uint32      `protobuf:"varint,7,req,name=portid" json:"portid,omitempty"`
	Groups    []uint32     `protobuf:"varint,8,rep,name=groups" json:"groups,omitempty"`
	DstPortid *uint32      `protobuf:"varint,9,req,name=dst_portid,json=dstPortid" json:"dst_portid,omitempty"`
	DstGroup  *uint32      `protobuf:"varint,10,req,name=dst_group,json=dstGroup" json:"dst_group,omitempty"`
	Fown      *FownEntry   `protobuf:"bytes,11,req,name=fown" json:"fown,omitempty"`
	Opts      *SkOptsEntry `protobuf:"bytes,12,req,name=opts" json:"opts,omitempty"`
	NsId      *uint32      `protobuf:"varint,13,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
}

func (x *NetlinkSkEntry) Reset() {
	*x = NetlinkSkEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sk_netlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetlinkSkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetlinkSkEntry) ProtoMessage() {}

func (x *NetlinkSkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_netlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetlinkSkEntry.ProtoReflect.Descriptor instead.
func (*NetlinkSkEntry) Descriptor() ([]byte, []int) {
	return file_sk_netlink_proto_rawDescGZIP(), []int{0}
}

func (x *NetlinkSkEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *NetlinkSkEntry) GetIno() uint32 {
	if x != nil && x.Ino != nil {
		return *x.Ino
	}
	return 0
}

func (x *NetlinkSkEntry) GetProtocol() uint32 {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return 0
}

func (x *NetlinkSkEntry) GetState() uint32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

func (x *NetlinkSkEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *NetlinkSkEntry) GetPortid() uint32 {
	if x != nil && x.Portid != nil {
		return *x.Portid
	}
	return 0
}

func (x *NetlinkSkEntry) GetGroups() []uint32 {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *NetlinkSkEntry) GetDstPortid() uint32 {
	if x != nil && x.DstPortid != nil {
		return *x.DstPortid
	}
	return 0
}

func (x *NetlinkSkEntry) GetDstGroup() uint32 {
	if x != nil && x.DstGroup != nil {
		return *x.DstGroup
	}
	return 0
}

func (x *NetlinkSkEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *NetlinkSkEntry) GetOpts() *SkOptsEntry {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *NetlinkSkEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

var File_sk_netlink_proto protoreflect.FileDescriptor

var file_sk_netlink_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6b, 0x2d, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x72, 0x69, 0x75, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x73, 0x6b, 0x2d, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd3, 0x02, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x6b, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x03, 0x69, 0x6e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x64, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x73, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x24, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x72, 0x69, 0x75, 0x2e, 0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18,
	0x0c, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x69, 0x75, 0x2e, 0x73, 0x6b, 0x5f,
	0x6f, 0x70, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73,
	0x12, 0x13, 0x0a, 0x05, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x6e, 0x73, 0x49, 0x64,
}

var (
	file_sk_netlink_proto_rawDescOnce sync.Once
	file_sk_netlink_proto_rawDescData = file_sk_netlink_proto_rawDesc
)

func file_sk_netlink_proto_rawDescGZIP() []byte {
	file_sk_netlink_proto_rawDescOnce.Do(func() {
		file_sk_netlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_sk_netlink_proto_rawDescData)
	})
	return file_sk_netlink_proto_rawDescData
}

var file_sk_netlink_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sk_netlink_proto_goTypes = []interface{}{
	(*NetlinkSkEntry)(nil), // 0: criu.netlink_sk_entry
	(*FownEntry)(nil),      // 1: criu.fown_entry
	(*SkOptsEntry)(nil),    // 2: criu.sk_opts_entry
}
var file_sk_netlink_proto_depIdxs = []int32{
	1, // 0: criu.netlink_sk_entry.fown:type_name -> criu.fown_entry
	2, // 1: criu.netlink_sk_entry.opts:type_name -> criu.sk_opts_entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sk_netlink_proto_init() }
func file_sk_netlink_proto_init() {
	if File_sk_netlink_proto != nil {
		return
	}
	file_opts_proto_init()
	file_fown_proto_init()
	file_sk_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sk_netlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetlinkSkEntry); i {
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
			RawDescriptor: file_sk_netlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sk_netlink_proto_goTypes,
		DependencyIndexes: file_sk_netlink_proto_depIdxs,
		MessageInfos:      file_sk_netlink_proto_msgTypes,
	}.Build()
	File_sk_netlink_proto = out.File
	file_sk_netlink_proto_rawDesc = nil
	file_sk_netlink_proto_goTypes = nil
	file_sk_netlink_proto_depIdxs = nil
}
