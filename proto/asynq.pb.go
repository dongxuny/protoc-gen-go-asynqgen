// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/asynq.proto

package asynq

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Typename string `protobuf:"bytes,1,opt,name=typename,proto3" json:"typename,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_asynq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asynq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_asynq_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetTypename() string {
	if x != nil {
		return x.Typename
	}
	return ""
}

var file_proto_asynq_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Task)(nil),
		Field:         567891,
		Name:          "asynq.task",
		Tag:           "bytes,567891,opt,name=task",
		Filename:      "proto/asynq.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional asynq.Task task = 567891;
	E_Task = &file_proto_asynq_proto_extTypes[0]
)

var File_proto_asynq_proto protoreflect.FileDescriptor

var file_proto_asynq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x3a, 0x41, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0xd4, 0x22, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x42, 0x41, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x50, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6e, 0x67, 0x78,
	0x75, 0x6e, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2d, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x61, 0x73, 0x79, 0x6e, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_asynq_proto_rawDescOnce sync.Once
	file_proto_asynq_proto_rawDescData = file_proto_asynq_proto_rawDesc
)

func file_proto_asynq_proto_rawDescGZIP() []byte {
	file_proto_asynq_proto_rawDescOnce.Do(func() {
		file_proto_asynq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_asynq_proto_rawDescData)
	})
	return file_proto_asynq_proto_rawDescData
}

var file_proto_asynq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_asynq_proto_goTypes = []interface{}{
	(*Task)(nil),                       // 0: asynq.Task
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_proto_asynq_proto_depIdxs = []int32{
	1, // 0: asynq.task:extendee -> google.protobuf.MethodOptions
	0, // 1: asynq.task:type_name -> asynq.Task
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_asynq_proto_init() }
func file_proto_asynq_proto_init() {
	if File_proto_asynq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_asynq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_proto_asynq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_proto_asynq_proto_goTypes,
		DependencyIndexes: file_proto_asynq_proto_depIdxs,
		MessageInfos:      file_proto_asynq_proto_msgTypes,
		ExtensionInfos:    file_proto_asynq_proto_extTypes,
	}.Build()
	File_proto_asynq_proto = out.File
	file_proto_asynq_proto_rawDesc = nil
	file_proto_asynq_proto_goTypes = nil
	file_proto_asynq_proto_depIdxs = nil
}
