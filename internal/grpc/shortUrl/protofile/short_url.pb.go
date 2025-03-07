// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: shortUrl/short_url.proto

package shortUrl

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=ShortUrl,proto3" json:"ShortUrl,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_shortUrl_short_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortUrl_short_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_shortUrl_short_url_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_shortUrl_short_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortUrl_short_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_shortUrl_short_url_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=ShortUrl,proto3" json:"ShortUrl,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	mi := &file_shortUrl_short_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortUrl_short_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_shortUrl_short_url_proto_rawDescGZIP(), []int{2}
}

func (x *PostResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	mi := &file_shortUrl_short_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortUrl_short_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_shortUrl_short_url_proto_rawDescGZIP(), []int{3}
}

func (x *PostRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_shortUrl_short_url_proto protoreflect.FileDescriptor

var file_shortUrl_short_url_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22,
	0x28, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x22, 0x2a, 0x0a, 0x0c, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x32, 0x61, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortUrl_short_url_proto_rawDescOnce sync.Once
	file_shortUrl_short_url_proto_rawDescData = file_shortUrl_short_url_proto_rawDesc
)

func file_shortUrl_short_url_proto_rawDescGZIP() []byte {
	file_shortUrl_short_url_proto_rawDescOnce.Do(func() {
		file_shortUrl_short_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortUrl_short_url_proto_rawDescData)
	})
	return file_shortUrl_short_url_proto_rawDescData
}

var file_shortUrl_short_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortUrl_short_url_proto_goTypes = []any{
	(*GetRequest)(nil),   // 0: api.GetRequest
	(*GetResponse)(nil),  // 1: api.GetResponse
	(*PostResponse)(nil), // 2: api.PostResponse
	(*PostRequest)(nil),  // 3: api.PostRequest
}
var file_shortUrl_short_url_proto_depIdxs = []int32{
	0, // 0: api.ShortUrl.GET:input_type -> api.GetRequest
	3, // 1: api.ShortUrl.Post:input_type -> api.PostRequest
	1, // 2: api.ShortUrl.GET:output_type -> api.GetResponse
	2, // 3: api.ShortUrl.Post:output_type -> api.PostResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortUrl_short_url_proto_init() }
func file_shortUrl_short_url_proto_init() {
	if File_shortUrl_short_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortUrl_short_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortUrl_short_url_proto_goTypes,
		DependencyIndexes: file_shortUrl_short_url_proto_depIdxs,
		MessageInfos:      file_shortUrl_short_url_proto_msgTypes,
	}.Build()
	File_shortUrl_short_url_proto = out.File
	file_shortUrl_short_url_proto_rawDesc = nil
	file_shortUrl_short_url_proto_goTypes = nil
	file_shortUrl_short_url_proto_depIdxs = nil
}
