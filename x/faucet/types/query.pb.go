// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucet/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("faucet/query.proto", fileDescriptor_32e01ab1e3e8ff22) }

var fileDescriptor_32e01ab1e3e8ff22 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x31, 0x6e, 0xc2, 0x30,
	0x14, 0xc6, 0xf1, 0x64, 0x68, 0x2b, 0x65, 0xcc, 0xd0, 0x21, 0xaa, 0xbc, 0xb7, 0x83, 0x9f, 0xd2,
	0xde, 0xa0, 0x9c, 0x80, 0x95, 0xcd, 0xb6, 0x1e, 0xc6, 0x92, 0xe3, 0x67, 0x62, 0x1b, 0x11, 0x4e,
	0xc1, 0xb1, 0x18, 0x33, 0x32, 0xa2, 0xe4, 0x22, 0x88, 0x18, 0xf6, 0xdf, 0xf7, 0xe9, 0x5f, 0xd5,
	0x5b, 0x91, 0x14, 0x46, 0xd8, 0x27, 0xec, 0x07, 0xee, 0x7b, 0x8a, 0x54, 0x7f, 0x4a, 0x9b, 0x4e,
	0x68, 0x2d, 0x72, 0x95, 0x7a, 0x93, 0x3a, 0x9e, 0x4d, 0xf3, 0xa5, 0x89, 0xb4, 0x45, 0x10, 0xde,
	0x80, 0x70, 0x8e, 0xa2, 0x88, 0x86, 0x5c, 0xc8, 0xab, 0xe6, 0x47, 0x51, 0xe8, 0x28, 0x80, 0x14,
	0x01, 0xf3, 0x1d, 0x1c, 0x5a, 0x89, 0x51, 0xb4, 0xe0, 0x85, 0x36, 0x6e, 0xc1, 0xd9, 0xfe, 0x7e,
	0x54, 0x6f, 0xeb, 0x87, 0xf8, 0x5f, 0x5d, 0x26, 0x56, 0x8e, 0x13, 0x2b, 0x6f, 0x13, 0x2b, 0xcf,
	0x33, 0x2b, 0xc6, 0x99, 0x15, 0xd7, 0x99, 0x15, 0x9b, 0x6f, 0x6d, 0xe2, 0x2e, 0x49, 0xae, 0xa8,
	0x83, 0x57, 0x0f, 0xe4, 0x1e, 0x38, 0xc2, 0xb3, 0x3a, 0x0e, 0x1e, 0x83, 0x7c, 0x5f, 0x4e, 0xff,
	0xee, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0xac, 0xc7, 0x8c, 0xcc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bluzelle.curium.faucet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "faucet/query.proto",
}
