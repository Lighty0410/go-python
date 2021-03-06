// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: server.proto

package grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Topic) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Topics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *Topics) Reset() {
	*x = Topics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topics) ProtoMessage() {}

func (x *Topics) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topics.ProtoReflect.Descriptor instead.
func (*Topics) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *Topics) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=Word,proto3" json:"Word,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x2f, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x06, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12,
	0x23, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x57, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x57, 0x6f, 0x72, 0x64,
	0x32, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_proto_goTypes = []interface{}{
	(*Topic)(nil),  // 0: grpc.Topic
	(*Topics)(nil), // 1: grpc.Topics
	(*Word)(nil),   // 2: grpc.Word
}
var file_server_proto_depIdxs = []int32{
	0, // 0: grpc.Topics.topics:type_name -> grpc.Topic
	2, // 1: grpc.GetTopics.GetTopics:input_type -> grpc.Word
	1, // 2: grpc.GetTopics.GetTopics:output_type -> grpc.Topics
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topics); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GetTopicsClient is the client API for GetTopics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetTopicsClient interface {
	GetTopics(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Topics, error)
}

type getTopicsClient struct {
	cc grpc.ClientConnInterface
}

func NewGetTopicsClient(cc grpc.ClientConnInterface) GetTopicsClient {
	return &getTopicsClient{cc}
}

func (c *getTopicsClient) GetTopics(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Topics, error) {
	out := new(Topics)
	err := c.cc.Invoke(ctx, "/grpc.GetTopics/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTopicsServer is the server API for GetTopics service.
type GetTopicsServer interface {
	GetTopics(context.Context, *Word) (*Topics, error)
}

// UnimplementedGetTopicsServer can be embedded to have forward compatible implementations.
type UnimplementedGetTopicsServer struct {
}

func (*UnimplementedGetTopicsServer) GetTopics(context.Context, *Word) (*Topics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}

func RegisterGetTopicsServer(s *grpc.Server, srv GetTopicsServer) {
	s.RegisterService(&_GetTopics_serviceDesc, srv)
}

func _GetTopics_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTopicsServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GetTopics/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTopicsServer).GetTopics(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetTopics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.GetTopics",
	HandlerType: (*GetTopicsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopics",
			Handler:    _GetTopics_GetTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
