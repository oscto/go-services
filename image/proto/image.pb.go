// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/image.proto

package image

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CallRequest) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *CallRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	X0  int64  `protobuf:"varint,2,opt,name=x0,proto3" json:"x0,omitempty"`
	Y0  int64  `protobuf:"varint,3,opt,name=y0,proto3" json:"y0,omitempty"`
	X1  int64  `protobuf:"varint,4,opt,name=x1,proto3" json:"x1,omitempty"`
	Y1  int64  `protobuf:"varint,5,opt,name=y1,proto3" json:"y1,omitempty"`
}

func (x *DrawRequest) Reset() {
	*x = DrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawRequest) ProtoMessage() {}

func (x *DrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawRequest.ProtoReflect.Descriptor instead.
func (*DrawRequest) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{2}
}

func (x *DrawRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DrawRequest) GetX0() int64 {
	if x != nil {
		return x.X0
	}
	return 0
}

func (x *DrawRequest) GetY0() int64 {
	if x != nil {
		return x.Y0
	}
	return 0
}

func (x *DrawRequest) GetX1() int64 {
	if x != nil {
		return x.X1
	}
	return 0
}

func (x *DrawRequest) GetY1() int64 {
	if x != nil {
		return x.Y1
	}
	return 0
}

type DrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DrawResponse) Reset() {
	*x = DrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawResponse) ProtoMessage() {}

func (x *DrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawResponse.ProtoReflect.Descriptor instead.
func (*DrawResponse) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{3}
}

func (x *DrawResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ToWebPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ToWebPRequest) Reset() {
	*x = ToWebPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToWebPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToWebPRequest) ProtoMessage() {}

func (x *ToWebPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToWebPRequest.ProtoReflect.Descriptor instead.
func (*ToWebPRequest) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{4}
}

func (x *ToWebPRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ToWebPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ToWebPResponse) Reset() {
	*x = ToWebPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToWebPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToWebPResponse) ProtoMessage() {}

func (x *ToWebPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToWebPResponse.ProtoReflect.Descriptor instead.
func (*ToWebPResponse) Descriptor() ([]byte, []int) {
	return file_proto_image_proto_rawDescGZIP(), []int{5}
}

func (x *ToWebPResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_proto_image_proto protoreflect.FileDescriptor

var file_proto_image_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x5f, 0x0a,
	0x0b, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x78, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x78, 0x30, 0x12, 0x0e,
	0x0a, 0x02, 0x79, 0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x30, 0x12, 0x0e,
	0x0a, 0x02, 0x78, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x78, 0x31, 0x12, 0x0e,
	0x0a, 0x02, 0x79, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x79, 0x31, 0x22, 0x22,
	0x0a, 0x0c, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x21, 0x0a, 0x0d, 0x54, 0x6f, 0x57, 0x65, 0x62, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x24, 0x0a, 0x0e, 0x54, 0x6f, 0x57, 0x65, 0x62, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0xa8, 0x01, 0x0a, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x44, 0x72,
	0x61, 0x77, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x06, 0x54, 0x6f, 0x57, 0x65, 0x62, 0x50, 0x12, 0x14, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x6f, 0x57, 0x65, 0x62, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x57, 0x65, 0x62, 0x50, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_image_proto_rawDescOnce sync.Once
	file_proto_image_proto_rawDescData = file_proto_image_proto_rawDesc
)

func file_proto_image_proto_rawDescGZIP() []byte {
	file_proto_image_proto_rawDescOnce.Do(func() {
		file_proto_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_image_proto_rawDescData)
	})
	return file_proto_image_proto_rawDescData
}

var file_proto_image_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_image_proto_goTypes = []interface{}{
	(*CallRequest)(nil),    // 0: image.CallRequest
	(*CallResponse)(nil),   // 1: image.CallResponse
	(*DrawRequest)(nil),    // 2: image.DrawRequest
	(*DrawResponse)(nil),   // 3: image.DrawResponse
	(*ToWebPRequest)(nil),  // 4: image.ToWebPRequest
	(*ToWebPResponse)(nil), // 5: image.ToWebPResponse
}
var file_proto_image_proto_depIdxs = []int32{
	0, // 0: image.Image.Resize:input_type -> image.CallRequest
	2, // 1: image.Image.Draw:input_type -> image.DrawRequest
	4, // 2: image.Image.ToWebP:input_type -> image.ToWebPRequest
	1, // 3: image.Image.Resize:output_type -> image.CallResponse
	3, // 4: image.Image.Draw:output_type -> image.DrawResponse
	5, // 5: image.Image.ToWebP:output_type -> image.ToWebPResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_image_proto_init() }
func file_proto_image_proto_init() {
	if File_proto_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_proto_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_proto_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawRequest); i {
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
		file_proto_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawResponse); i {
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
		file_proto_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToWebPRequest); i {
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
		file_proto_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToWebPResponse); i {
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
			RawDescriptor: file_proto_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_image_proto_goTypes,
		DependencyIndexes: file_proto_image_proto_depIdxs,
		MessageInfos:      file_proto_image_proto_msgTypes,
	}.Build()
	File_proto_image_proto = out.File
	file_proto_image_proto_rawDesc = nil
	file_proto_image_proto_goTypes = nil
	file_proto_image_proto_depIdxs = nil
}