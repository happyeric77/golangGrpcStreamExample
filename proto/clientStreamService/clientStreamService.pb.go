// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/clientStreamService/clientStreamService.proto

package clientStreamService

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientStreamService_clientStreamService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientStreamService_clientStreamService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_clientStreamService_clientStreamService_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clientStreamService_clientStreamService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clientStreamService_clientStreamService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_clientStreamService_clientStreamService_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_clientStreamService_clientStreamService_proto protoreflect.FileDescriptor

var file_proto_clientStreamService_clientStreamService_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x33, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_clientStreamService_clientStreamService_proto_rawDescOnce sync.Once
	file_proto_clientStreamService_clientStreamService_proto_rawDescData = file_proto_clientStreamService_clientStreamService_proto_rawDesc
)

func file_proto_clientStreamService_clientStreamService_proto_rawDescGZIP() []byte {
	file_proto_clientStreamService_clientStreamService_proto_rawDescOnce.Do(func() {
		file_proto_clientStreamService_clientStreamService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clientStreamService_clientStreamService_proto_rawDescData)
	})
	return file_proto_clientStreamService_clientStreamService_proto_rawDescData
}

var file_proto_clientStreamService_clientStreamService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_clientStreamService_clientStreamService_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
}
var file_proto_clientStreamService_clientStreamService_proto_depIdxs = []int32{
	0, // 0: ClientStream.DataUpload:input_type -> Request
	1, // 1: ClientStream.DataUpload:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_clientStreamService_clientStreamService_proto_init() }
func file_proto_clientStreamService_clientStreamService_proto_init() {
	if File_proto_clientStreamService_clientStreamService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clientStreamService_clientStreamService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_clientStreamService_clientStreamService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_clientStreamService_clientStreamService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clientStreamService_clientStreamService_proto_goTypes,
		DependencyIndexes: file_proto_clientStreamService_clientStreamService_proto_depIdxs,
		MessageInfos:      file_proto_clientStreamService_clientStreamService_proto_msgTypes,
	}.Build()
	File_proto_clientStreamService_clientStreamService_proto = out.File
	file_proto_clientStreamService_clientStreamService_proto_rawDesc = nil
	file_proto_clientStreamService_clientStreamService_proto_goTypes = nil
	file_proto_clientStreamService_clientStreamService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientStreamClient is the client API for ClientStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientStreamClient interface {
	DataUpload(ctx context.Context, opts ...grpc.CallOption) (ClientStream_DataUploadClient, error)
}

type clientStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamClient(cc grpc.ClientConnInterface) ClientStreamClient {
	return &clientStreamClient{cc}
}

func (c *clientStreamClient) DataUpload(ctx context.Context, opts ...grpc.CallOption) (ClientStream_DataUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientStream_serviceDesc.Streams[0], "/ClientStream/DataUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamDataUploadClient{stream}
	return x, nil
}

type ClientStream_DataUploadClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type clientStreamDataUploadClient struct {
	grpc.ClientStream
}

func (x *clientStreamDataUploadClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamDataUploadClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamServer is the server API for ClientStream service.
type ClientStreamServer interface {
	DataUpload(ClientStream_DataUploadServer) error
}

// UnimplementedClientStreamServer can be embedded to have forward compatible implementations.
type UnimplementedClientStreamServer struct {
}

func (*UnimplementedClientStreamServer) DataUpload(ClientStream_DataUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method DataUpload not implemented")
}

func RegisterClientStreamServer(s *grpc.Server, srv ClientStreamServer) {
	s.RegisterService(&_ClientStream_serviceDesc, srv)
}

func _ClientStream_DataUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamServer).DataUpload(&clientStreamDataUploadServer{stream})
}

type ClientStream_DataUploadServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type clientStreamDataUploadServer struct {
	grpc.ServerStream
}

func (x *clientStreamDataUploadServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamDataUploadServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClientStream",
	HandlerType: (*ClientStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataUpload",
			Handler:       _ClientStream_DataUpload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/clientStreamService/clientStreamService.proto",
}
