// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: proto/ConsensusService.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_ConsensusService_proto protoreflect.FileDescriptor

var file_proto_ConsensusService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbf, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x67, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x52, 0x0a, 0x26, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6a, 0x61, 0x76, 0x61, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_ConsensusService_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*Query)(nil),               // 1: proto.Query
	(*TransactionResponse)(nil), // 2: proto.TransactionResponse
	(*Response)(nil),            // 3: proto.Response
}
var file_proto_ConsensusService_proto_depIdxs = []int32{
	0, // 0: proto.ConsensusService.createTopic:input_type -> proto.Transaction
	0, // 1: proto.ConsensusService.updateTopic:input_type -> proto.Transaction
	0, // 2: proto.ConsensusService.deleteTopic:input_type -> proto.Transaction
	1, // 3: proto.ConsensusService.getTopicInfo:input_type -> proto.Query
	0, // 4: proto.ConsensusService.submitMessage:input_type -> proto.Transaction
	2, // 5: proto.ConsensusService.createTopic:output_type -> proto.TransactionResponse
	2, // 6: proto.ConsensusService.updateTopic:output_type -> proto.TransactionResponse
	2, // 7: proto.ConsensusService.deleteTopic:output_type -> proto.TransactionResponse
	3, // 8: proto.ConsensusService.getTopicInfo:output_type -> proto.Response
	2, // 9: proto.ConsensusService.submitMessage:output_type -> proto.TransactionResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ConsensusService_proto_init() }
func file_proto_ConsensusService_proto_init() {
	if File_proto_ConsensusService_proto != nil {
		return
	}
	file_proto_Query_proto_init()
	file_proto_Response_proto_init()
	file_proto_TransactionResponse_proto_init()
	file_proto_Transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ConsensusService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ConsensusService_proto_goTypes,
		DependencyIndexes: file_proto_ConsensusService_proto_depIdxs,
	}.Build()
	File_proto_ConsensusService_proto = out.File
	file_proto_ConsensusService_proto_rawDesc = nil
	file_proto_ConsensusService_proto_goTypes = nil
	file_proto_ConsensusService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsensusServiceClient is the client API for ConsensusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusServiceClient interface {
	// Create a topic to be used for consensus.
	// If an autoRenewAccount is specified, that account must also sign this transaction.
	// If an adminKey is specified, the adminKey must sign the transaction.
	// On success, the resulting TransactionReceipt contains the newly created TopicId.
	// Request is [ConsensusCreateTopicTransactionBody](#proto.ConsensusCreateTopicTransactionBody)
	CreateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Update a topic.
	// If there is no adminKey, the only authorized update (available to anyone) is to extend the expirationTime.
	// Otherwise transaction must be signed by the adminKey.
	// If an adminKey is updated, the transaction must be signed by the pre-update adminKey and post-update adminKey.
	// If a new autoRenewAccount is specified (not just being removed), that account must also sign the transaction.
	// Request is [ConsensusUpdateTopicTransactionBody](#proto.ConsensusUpdateTopicTransactionBody)
	UpdateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Delete a topic. No more transactions or queries on the topic (via HAPI) will succeed.
	// If an adminKey is set, this transaction must be signed by that key.
	// If there is no adminKey, this transaction will fail UNAUTHORIZED.
	// Request is [ConsensusDeleteTopicTransactionBody](#proto.ConsensusDeleteTopicTransactionBody)
	DeleteTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Retrieve the latest state of a topic. This method is unrestricted and allowed on any topic by any payer account.
	// Deleted accounts will not be returned.
	// Request is [ConsensusGetTopicInfoQuery](#proto.ConsensusGetTopicInfoQuery)
	// Response is [ConsensusGetTopicInfoResponse](#proto.ConsensusGetTopicInfoResponse)
	GetTopicInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Submit a message for consensus.
	// Valid and authorized messages on valid topics will be ordered by the consensus service, gossipped to the
	// mirror net, and published (in order) to all subscribers (from the mirror net) on this topic.
	// The submitKey (if any) must sign this transaction.
	// On success, the resulting TransactionReceipt contains the topic's updated topicSequenceNumber and
	// topicRunningHash.
	// Request is [ConsensusSubmitMessageTransactionBody](#proto.ConsensusSubmitMessageTransactionBody)
	SubmitMessage(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type consensusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusServiceClient(cc grpc.ClientConnInterface) ConsensusServiceClient {
	return &consensusServiceClient{cc}
}

func (c *consensusServiceClient) CreateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/createTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) UpdateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/updateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) DeleteTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/deleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) GetTopicInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/getTopicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) SubmitMessage(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/submitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsensusServiceServer is the server API for ConsensusService service.
type ConsensusServiceServer interface {
	// Create a topic to be used for consensus.
	// If an autoRenewAccount is specified, that account must also sign this transaction.
	// If an adminKey is specified, the adminKey must sign the transaction.
	// On success, the resulting TransactionReceipt contains the newly created TopicId.
	// Request is [ConsensusCreateTopicTransactionBody](#proto.ConsensusCreateTopicTransactionBody)
	CreateTopic(context.Context, *Transaction) (*TransactionResponse, error)
	// Update a topic.
	// If there is no adminKey, the only authorized update (available to anyone) is to extend the expirationTime.
	// Otherwise transaction must be signed by the adminKey.
	// If an adminKey is updated, the transaction must be signed by the pre-update adminKey and post-update adminKey.
	// If a new autoRenewAccount is specified (not just being removed), that account must also sign the transaction.
	// Request is [ConsensusUpdateTopicTransactionBody](#proto.ConsensusUpdateTopicTransactionBody)
	UpdateTopic(context.Context, *Transaction) (*TransactionResponse, error)
	// Delete a topic. No more transactions or queries on the topic (via HAPI) will succeed.
	// If an adminKey is set, this transaction must be signed by that key.
	// If there is no adminKey, this transaction will fail UNAUTHORIZED.
	// Request is [ConsensusDeleteTopicTransactionBody](#proto.ConsensusDeleteTopicTransactionBody)
	DeleteTopic(context.Context, *Transaction) (*TransactionResponse, error)
	// Retrieve the latest state of a topic. This method is unrestricted and allowed on any topic by any payer account.
	// Deleted accounts will not be returned.
	// Request is [ConsensusGetTopicInfoQuery](#proto.ConsensusGetTopicInfoQuery)
	// Response is [ConsensusGetTopicInfoResponse](#proto.ConsensusGetTopicInfoResponse)
	GetTopicInfo(context.Context, *Query) (*Response, error)
	// Submit a message for consensus.
	// Valid and authorized messages on valid topics will be ordered by the consensus service, gossipped to the
	// mirror net, and published (in order) to all subscribers (from the mirror net) on this topic.
	// The submitKey (if any) must sign this transaction.
	// On success, the resulting TransactionReceipt contains the topic's updated topicSequenceNumber and
	// topicRunningHash.
	// Request is [ConsensusSubmitMessageTransactionBody](#proto.ConsensusSubmitMessageTransactionBody)
	SubmitMessage(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedConsensusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsensusServiceServer struct {
}

func (*UnimplementedConsensusServiceServer) CreateTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (*UnimplementedConsensusServiceServer) UpdateTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (*UnimplementedConsensusServiceServer) DeleteTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (*UnimplementedConsensusServiceServer) GetTopicInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicInfo not implemented")
}
func (*UnimplementedConsensusServiceServer) SubmitMessage(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMessage not implemented")
}

func RegisterConsensusServiceServer(s *grpc.Server, srv ConsensusServiceServer) {
	s.RegisterService(&_ConsensusService_serviceDesc, srv)
}

func _ConsensusService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).CreateTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).UpdateTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).DeleteTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_GetTopicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).GetTopicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/GetTopicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).GetTopicInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_SubmitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).SubmitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/SubmitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).SubmitMessage(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsensusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConsensusService",
	HandlerType: (*ConsensusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createTopic",
			Handler:    _ConsensusService_CreateTopic_Handler,
		},
		{
			MethodName: "updateTopic",
			Handler:    _ConsensusService_UpdateTopic_Handler,
		},
		{
			MethodName: "deleteTopic",
			Handler:    _ConsensusService_DeleteTopic_Handler,
		},
		{
			MethodName: "getTopicInfo",
			Handler:    _ConsensusService_GetTopicInfo_Handler,
		},
		{
			MethodName: "submitMessage",
			Handler:    _ConsensusService_SubmitMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ConsensusService.proto",
}
