// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/m_member_proto/hello.proto

package mytype

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloWorldRequest struct {
	Referer              string   `protobuf:"bytes,1,opt,name=referer,proto3" json:"referer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_900307a2e4697540, []int{0}
}

func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

type HelloWorldResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_900307a2e4697540, []int{1}
}

func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "proto.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "proto.HelloWorldResponse")
}

func init() { proto.RegisterFile("proto/m_member_proto/hello.proto", fileDescriptor_900307a2e4697540) }

var fileDescriptor_900307a2e4697540 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x8d, 0xcf, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0x87, 0x70, 0x33, 0x52, 0x73,
	0x72, 0xf2, 0xf5, 0xc0, 0x6c, 0x21, 0x56, 0x30, 0x25, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93,
	0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57,
	0x0c, 0x51, 0xa4, 0xa4, 0xcb, 0x25, 0xe8, 0x01, 0xd2, 0x13, 0x9e, 0x5f, 0x94, 0x93, 0x12, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc1, 0xc5, 0x5e, 0x94, 0x9a, 0x96, 0x5a, 0x94, 0x5a,
	0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x2a, 0xe9, 0x71, 0x09, 0x21, 0x2b, 0x2f,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0xa9, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x85, 0xa9,
	0x87, 0x72, 0x8d, 0xf2, 0x91, 0x8d, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x8a, 0xe2,
	0xe2, 0x0d, 0x4e, 0xac, 0x44, 0x88, 0x0b, 0x49, 0x40, 0x1c, 0xa3, 0x87, 0xe1, 0x12, 0x29, 0x49,
	0x2c, 0x32, 0x10, 0x4b, 0x95, 0x44, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0x27, 0xc4, 0x03, 0xf1,
	0x74, 0x7c, 0x39, 0x48, 0x36, 0x89, 0x0d, 0xac, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x26,
	0x32, 0xaf, 0xe0, 0x1f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/proto.HelloWorldService/SayHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	SayHelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

// UnimplementedHelloWorldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServiceServer struct {
}

func (*UnimplementedHelloWorldServiceServer) SayHelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloWorld not implemented")
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_SayHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HelloWorldService/SayHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWorld",
			Handler:    _HelloWorldService_SayHelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/m_member_proto/hello.proto",
}
