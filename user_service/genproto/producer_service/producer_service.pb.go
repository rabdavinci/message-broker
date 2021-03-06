// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: producer_service.proto

package producer_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_producer_service_proto protoreflect.FileDescriptor

var file_producer_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x8d, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_producer_service_proto_goTypes = []interface{}{
	(*Topic)(nil),              // 0: genproto.Topic
	(*SendMessageRequest)(nil), // 1: genproto.SendMessageRequest
	(*TopicId)(nil),            // 2: genproto.TopicId
	(*emptypb.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_producer_service_proto_depIdxs = []int32{
	0, // 0: genproto.ProducerService.CreateTopic:input_type -> genproto.Topic
	1, // 1: genproto.ProducerService.SendMessage:input_type -> genproto.SendMessageRequest
	2, // 2: genproto.ProducerService.CreateTopic:output_type -> genproto.TopicId
	3, // 3: genproto.ProducerService.SendMessage:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_producer_service_proto_init() }
func file_producer_service_proto_init() {
	if File_producer_service_proto != nil {
		return
	}
	file_producer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_producer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_producer_service_proto_goTypes,
		DependencyIndexes: file_producer_service_proto_depIdxs,
	}.Build()
	File_producer_service_proto = out.File
	file_producer_service_proto_rawDesc = nil
	file_producer_service_proto_goTypes = nil
	file_producer_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProducerServiceClient is the client API for ProducerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProducerServiceClient interface {
	CreateTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicId, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type producerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerServiceClient(cc grpc.ClientConnInterface) ProducerServiceClient {
	return &producerServiceClient{cc}
}

func (c *producerServiceClient) CreateTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicId, error) {
	out := new(TopicId)
	err := c.cc.Invoke(ctx, "/genproto.ProducerService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ProducerService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServiceServer is the server API for ProducerService service.
type ProducerServiceServer interface {
	CreateTopic(context.Context, *Topic) (*TopicId, error)
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
}

// UnimplementedProducerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProducerServiceServer struct {
}

func (*UnimplementedProducerServiceServer) CreateTopic(context.Context, *Topic) (*TopicId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (*UnimplementedProducerServiceServer) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterProducerServiceServer(s *grpc.Server, srv ProducerServiceServer) {
	s.RegisterService(&_ProducerService_serviceDesc, srv)
}

func _ProducerService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProducerService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).CreateTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProducerService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProducerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ProducerService",
	HandlerType: (*ProducerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _ProducerService_CreateTopic_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ProducerService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "producer_service.proto",
}
