// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: download.proto

package filegrpc

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

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_download_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_download_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=file_chunk,json=fileChunk,proto3" json:"file_chunk,omitempty"`
	FileName  string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_download_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_download_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadResponse) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *DownloadResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_download_proto protoreflect.FileDescriptor

var file_download_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_download_proto_rawDescOnce sync.Once
	file_download_proto_rawDescData = file_download_proto_rawDesc
)

func file_download_proto_rawDescGZIP() []byte {
	file_download_proto_rawDescOnce.Do(func() {
		file_download_proto_rawDescData = protoimpl.X.CompressGZIP(file_download_proto_rawDescData)
	})
	return file_download_proto_rawDescData
}

var file_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_download_proto_goTypes = []any{
	(*DownloadRequest)(nil),  // 0: file.DownloadRequest
	(*DownloadResponse)(nil), // 1: file.DownloadResponse
}
var file_download_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_download_proto_init() }
func file_download_proto_init() {
	if File_download_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_download_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_download_proto_goTypes,
		DependencyIndexes: file_download_proto_depIdxs,
		MessageInfos:      file_download_proto_msgTypes,
	}.Build()
	File_download_proto = out.File
	file_download_proto_rawDesc = nil
	file_download_proto_goTypes = nil
	file_download_proto_depIdxs = nil
}
