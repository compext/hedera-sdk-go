// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.19.0-devel
// 	protoc        v3.11.3
// source: proto/FreezeService.proto

package proto

import (
	context "context"
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
	_ = protoimpl.EnforceVersion(19 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 19)
)

var File_proto_FreezeService_proto protoreflect.FileDescriptor

var file_proto_FreezeService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x49, 0x0a, 0x0d,
	0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FreezeService_proto_rawDescOnce sync.Once
	file_proto_FreezeService_proto_rawDescData = file_proto_FreezeService_proto_rawDesc
)

func file_proto_FreezeService_proto_rawDescGZIP() []byte {
	file_proto_FreezeService_proto_rawDescOnce.Do(func() {
		file_proto_FreezeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FreezeService_proto_rawDescData)
	})
	return file_proto_FreezeService_proto_rawDescData
}

var file_proto_FreezeService_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*TransactionResponse)(nil), // 1: proto.TransactionResponse
}
var file_proto_FreezeService_proto_depIdxs = []int32{
	0, // 0: proto.FreezeService.freeze:input_type -> proto.Transaction
	1, // 1: proto.FreezeService.freeze:output_type -> proto.TransactionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_FreezeService_proto_init() }
func file_proto_FreezeService_proto_init() {
	if File_proto_FreezeService_proto != nil {
		return
	}
	file_proto_TransactionResponse_proto_init()
	file_proto_Transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_FreezeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_FreezeService_proto_goTypes,
		DependencyIndexes: file_proto_FreezeService_proto_depIdxs,
	}.Build()
	File_proto_FreezeService_proto = out.File
	file_proto_FreezeService_proto_rawDesc = nil
	file_proto_FreezeService_proto_goTypes = nil
	file_proto_FreezeService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FreezeServiceClient is the client API for FreezeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FreezeServiceClient interface {
	Freeze(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type freezeServiceClient struct {
	cc *grpc.ClientConn
}

func NewFreezeServiceClient(cc *grpc.ClientConn) FreezeServiceClient {
	return &freezeServiceClient{cc}
}

func (c *freezeServiceClient) Freeze(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FreezeService/freeze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreezeServiceServer is the server API for FreezeService service.
type FreezeServiceServer interface {
	Freeze(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedFreezeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFreezeServiceServer struct {
}

func (*UnimplementedFreezeServiceServer) Freeze(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Freeze not implemented")
}

func RegisterFreezeServiceServer(s *grpc.Server, srv FreezeServiceServer) {
	s.RegisterService(&_FreezeService_serviceDesc, srv)
}

func _FreezeService_Freeze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreezeServiceServer).Freeze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FreezeService/Freeze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreezeServiceServer).Freeze(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _FreezeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FreezeService",
	HandlerType: (*FreezeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "freeze",
			Handler:    _FreezeService_Freeze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/FreezeService.proto",
}
