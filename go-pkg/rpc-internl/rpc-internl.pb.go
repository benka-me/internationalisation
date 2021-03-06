// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc-internl.proto

package rpc_internl

import (
	context "context"
	fmt "fmt"
	internl "github.com/benka-me/internationalisation/go-pkg/internl"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("rpc-internl.proto", fileDescriptor_e2fc9808231097fb) }

var fileDescriptor_e2fc9808231097fb = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2a, 0x48, 0xd6,
	0xcd, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0xcb, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0a,
	0x2a, 0x48, 0xf6, 0x84, 0x88, 0x48, 0xd9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x27, 0xa5, 0xe6, 0x65, 0x27, 0xea, 0xe6, 0xa6, 0xea, 0x43, 0x94, 0x27, 0x96, 0x64,
	0xe6, 0xe7, 0x25, 0xe6, 0x64, 0x16, 0x83, 0x19, 0xfa, 0x60, 0xbd, 0x49, 0xa5, 0x69, 0xfa, 0x28,
	0x86, 0x19, 0x75, 0x33, 0x72, 0xb1, 0x43, 0x0d, 0x13, 0xb2, 0xe7, 0xe2, 0x72, 0x4f, 0x2d, 0xf1,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xd7, 0x83, 0xa9, 0x84, 0x8a, 0x04, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x48, 0x49, 0x60, 0x4a, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xd9,
	0x73, 0x71, 0x07, 0x94, 0x16, 0x67, 0x60, 0x9a, 0x80, 0x24, 0x1a, 0x94, 0x5a, 0x28, 0x85, 0x43,
	0xa2, 0xd8, 0xa9, 0xf0, 0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63,
	0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x35, 0xd1, 0x81, 0x90,
	0x9e, 0xaf, 0x5b, 0x90, 0x9d, 0xae, 0x8f, 0x14, 0xa6, 0x49, 0x6c, 0xe0, 0x70, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x14, 0xbe, 0xb9, 0x24, 0x69, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternlClient is the client API for Internl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternlClient interface {
	GetMessage(ctx context.Context, in *internl.MessageRequest, opts ...grpc.CallOption) (*internl.MessageResponse, error)
	PushMessage(ctx context.Context, in *internl.PushMessageReq, opts ...grpc.CallOption) (*internl.PushMessageRes, error)
}

type internlClient struct {
	cc *grpc.ClientConn
}

func NewInternlClient(cc *grpc.ClientConn) InternlClient {
	return &internlClient{cc}
}

func (c *internlClient) GetMessage(ctx context.Context, in *internl.MessageRequest, opts ...grpc.CallOption) (*internl.MessageResponse, error) {
	out := new(internl.MessageResponse)
	err := c.cc.Invoke(ctx, "/RpcInternl.Internl/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internlClient) PushMessage(ctx context.Context, in *internl.PushMessageReq, opts ...grpc.CallOption) (*internl.PushMessageRes, error) {
	out := new(internl.PushMessageRes)
	err := c.cc.Invoke(ctx, "/RpcInternl.Internl/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternlServer is the server API for Internl service.
type InternlServer interface {
	GetMessage(context.Context, *internl.MessageRequest) (*internl.MessageResponse, error)
	PushMessage(context.Context, *internl.PushMessageReq) (*internl.PushMessageRes, error)
}

// UnimplementedInternlServer can be embedded to have forward compatible implementations.
type UnimplementedInternlServer struct {
}

func (*UnimplementedInternlServer) GetMessage(ctx context.Context, req *internl.MessageRequest) (*internl.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (*UnimplementedInternlServer) PushMessage(ctx context.Context, req *internl.PushMessageReq) (*internl.PushMessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}

func RegisterInternlServer(s *grpc.Server, srv InternlServer) {
	s.RegisterService(&_Internl_serviceDesc, srv)
}

func _Internl_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internl.MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternlServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcInternl.Internl/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternlServer).GetMessage(ctx, req.(*internl.MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internl_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internl.PushMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternlServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcInternl.Internl/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternlServer).PushMessage(ctx, req.(*internl.PushMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RpcInternl.Internl",
	HandlerType: (*InternlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _Internl_GetMessage_Handler,
		},
		{
			MethodName: "PushMessage",
			Handler:    _Internl_PushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc-internl.proto",
}
