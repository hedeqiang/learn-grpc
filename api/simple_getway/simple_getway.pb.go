// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: simple_getway/simple_getway.proto

package simple_getway

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_getway_simple_getway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_getway_simple_getway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_simple_getway_simple_getway_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Body    string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_getway_simple_getway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_getway_simple_getway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_simple_getway_simple_getway_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_getway_simple_getway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_getway_simple_getway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_simple_getway_simple_getway_proto_rawDescGZIP(), []int{2}
}

func (x *PostRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Body    string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_getway_simple_getway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_getway_simple_getway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_simple_getway_simple_getway_proto_rawDescGZIP(), []int{3}
}

func (x *PostResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PostResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_simple_getway_simple_getway_proto protoreflect.FileDescriptor

var file_simple_getway_simple_getway_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77,
	0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0xb2, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x19, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x21, 0x5a,
	0x1f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77,
	0x61, 0x79, 0x3b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x77, 0x61, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simple_getway_simple_getway_proto_rawDescOnce sync.Once
	file_simple_getway_simple_getway_proto_rawDescData = file_simple_getway_simple_getway_proto_rawDesc
)

func file_simple_getway_simple_getway_proto_rawDescGZIP() []byte {
	file_simple_getway_simple_getway_proto_rawDescOnce.Do(func() {
		file_simple_getway_simple_getway_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_getway_simple_getway_proto_rawDescData)
	})
	return file_simple_getway_simple_getway_proto_rawDescData
}

var file_simple_getway_simple_getway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_simple_getway_simple_getway_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: simple_getway.GetRequest
	(*GetResponse)(nil),  // 1: simple_getway.GetResponse
	(*PostRequest)(nil),  // 2: simple_getway.PostRequest
	(*PostResponse)(nil), // 3: simple_getway.PostResponse
}
var file_simple_getway_simple_getway_proto_depIdxs = []int32{
	0, // 0: simple_getway.HttpService.Get:input_type -> simple_getway.GetRequest
	2, // 1: simple_getway.HttpService.Post:input_type -> simple_getway.PostRequest
	1, // 2: simple_getway.HttpService.Get:output_type -> simple_getway.GetResponse
	3, // 3: simple_getway.HttpService.Post:output_type -> simple_getway.PostResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_simple_getway_simple_getway_proto_init() }
func file_simple_getway_simple_getway_proto_init() {
	if File_simple_getway_simple_getway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_getway_simple_getway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_simple_getway_simple_getway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_simple_getway_simple_getway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequest); i {
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
		file_simple_getway_simple_getway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
			RawDescriptor: file_simple_getway_simple_getway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simple_getway_simple_getway_proto_goTypes,
		DependencyIndexes: file_simple_getway_simple_getway_proto_depIdxs,
		MessageInfos:      file_simple_getway_simple_getway_proto_msgTypes,
	}.Build()
	File_simple_getway_simple_getway_proto = out.File
	file_simple_getway_simple_getway_proto_rawDesc = nil
	file_simple_getway_simple_getway_proto_goTypes = nil
	file_simple_getway_simple_getway_proto_depIdxs = nil
}
