// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: cli_download.proto

package dldpb

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

type DownloadVideoStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadVideoStreamRequest) Reset() {
	*x = DownloadVideoStreamRequest{}
	mi := &file_cli_download_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadVideoStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadVideoStreamRequest) ProtoMessage() {}

func (x *DownloadVideoStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_download_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadVideoStreamRequest.ProtoReflect.Descriptor instead.
func (*DownloadVideoStreamRequest) Descriptor() ([]byte, []int) {
	return file_cli_download_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadVideoStreamRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DownloadVideoStreamResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*DownloadVideoStreamResponse_Filename
	//	*DownloadVideoStreamResponse_Chunk
	Data          isDownloadVideoStreamResponse_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadVideoStreamResponse) Reset() {
	*x = DownloadVideoStreamResponse{}
	mi := &file_cli_download_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadVideoStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadVideoStreamResponse) ProtoMessage() {}

func (x *DownloadVideoStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_download_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadVideoStreamResponse.ProtoReflect.Descriptor instead.
func (*DownloadVideoStreamResponse) Descriptor() ([]byte, []int) {
	return file_cli_download_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadVideoStreamResponse) GetData() isDownloadVideoStreamResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DownloadVideoStreamResponse) GetFilename() string {
	if x != nil {
		if x, ok := x.Data.(*DownloadVideoStreamResponse_Filename); ok {
			return x.Filename
		}
	}
	return ""
}

func (x *DownloadVideoStreamResponse) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*DownloadVideoStreamResponse_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

type isDownloadVideoStreamResponse_Data interface {
	isDownloadVideoStreamResponse_Data()
}

type DownloadVideoStreamResponse_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3,oneof"`
}

type DownloadVideoStreamResponse_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*DownloadVideoStreamResponse_Filename) isDownloadVideoStreamResponse_Data() {}

func (*DownloadVideoStreamResponse_Chunk) isDownloadVideoStreamResponse_Data() {}

var File_cli_download_proto protoreflect.FileDescriptor

var file_cli_download_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x63, 0x6c, 0x69, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x6c, 0x64, 0x22, 0x2e, 0x0a, 0x1a, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x1b, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x70, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x13,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x64, 0x6c, 0x64, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x6c, 0x64, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f,
	0x64, 0x6c, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_cli_download_proto_rawDescOnce sync.Once
	file_cli_download_proto_rawDescData []byte
)

func file_cli_download_proto_rawDescGZIP() []byte {
	file_cli_download_proto_rawDescOnce.Do(func() {
		file_cli_download_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cli_download_proto_rawDesc), len(file_cli_download_proto_rawDesc)))
	})
	return file_cli_download_proto_rawDescData
}

var file_cli_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cli_download_proto_goTypes = []any{
	(*DownloadVideoStreamRequest)(nil),  // 0: dld.DownloadVideoStreamRequest
	(*DownloadVideoStreamResponse)(nil), // 1: dld.DownloadVideoStreamResponse
}
var file_cli_download_proto_depIdxs = []int32{
	0, // 0: dld.CliDownloadService.DownloadVideoStream:input_type -> dld.DownloadVideoStreamRequest
	1, // 1: dld.CliDownloadService.DownloadVideoStream:output_type -> dld.DownloadVideoStreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cli_download_proto_init() }
func file_cli_download_proto_init() {
	if File_cli_download_proto != nil {
		return
	}
	file_cli_download_proto_msgTypes[1].OneofWrappers = []any{
		(*DownloadVideoStreamResponse_Filename)(nil),
		(*DownloadVideoStreamResponse_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cli_download_proto_rawDesc), len(file_cli_download_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cli_download_proto_goTypes,
		DependencyIndexes: file_cli_download_proto_depIdxs,
		MessageInfos:      file_cli_download_proto_msgTypes,
	}.Build()
	File_cli_download_proto = out.File
	file_cli_download_proto_goTypes = nil
	file_cli_download_proto_depIdxs = nil
}
