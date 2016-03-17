// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	Bar
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HelloRequest struct {
	Greeting         *string `protobuf:"bytes,1,req,name=greeting" json:"greeting,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetGreeting() string {
	if m != nil && m.Greeting != nil {
		return *m.Greeting
	}
	return ""
}

type HelloResponse struct {
	Reply            *string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetReply() string {
	if m != nil && m.Reply != nil {
		return *m.Reply
	}
	return ""
}

type Bar struct {
	Id  *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Apa []string `protobuf:"bytes,2,rep,name=apa" json:"apa,omitempty"`
	// Types that are valid to be assigned to One:
	//	*Bar_A
	//	*Bar_B
	One              isBar_One `protobuf_oneof:"one"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isBar_One interface {
	isBar_One()
}

type Bar_A struct {
	A string `protobuf:"bytes,3,opt,name=a,oneof"`
}
type Bar_B struct {
	B int32 `protobuf:"varint,4,opt,name=b,oneof"`
}

func (*Bar_A) isBar_One() {}
func (*Bar_B) isBar_One() {}

func (m *Bar) GetOne() isBar_One {
	if m != nil {
		return m.One
	}
	return nil
}

func (m *Bar) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Bar) GetApa() []string {
	if m != nil {
		return m.Apa
	}
	return nil
}

func (m *Bar) GetA() string {
	if x, ok := m.GetOne().(*Bar_A); ok {
		return x.A
	}
	return ""
}

func (m *Bar) GetB() int32 {
	if x, ok := m.GetOne().(*Bar_B); ok {
		return x.B
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Bar) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Bar_OneofMarshaler, _Bar_OneofUnmarshaler, _Bar_OneofSizer, []interface{}{
		(*Bar_A)(nil),
		(*Bar_B)(nil),
	}
}

func _Bar_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Bar)
	// one
	switch x := m.One.(type) {
	case *Bar_A:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.A)
	case *Bar_B:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.B))
	case nil:
	default:
		return fmt.Errorf("Bar.One has unexpected type %T", x)
	}
	return nil
}

func _Bar_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Bar)
	switch tag {
	case 3: // one.a
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.One = &Bar_A{x}
		return true, err
	case 4: // one.b
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.One = &Bar_B{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Bar_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Bar)
	// one
	switch x := m.One.(type) {
	case *Bar_A:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.A)))
		n += len(x.A)
	case *Bar_B:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.B))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
	proto.RegisterType((*Bar)(nil), "Bar")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for HelloService service

type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/HelloService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServiceServer).SayHello(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8e, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x46, 0xd5, 0x6d, 0x41, 0xa7, 0x8c, 0x98, 0xd3, 0xd2, 0x21, 0xc4, 0x93, 0x5d, 0x3c, 0x74,
	0xeb, 0x14, 0x78, 0xf2, 0x9c, 0xbf, 0x60, 0xb3, 0x41, 0x04, 0x71, 0xb7, 0xd1, 0x02, 0xff, 0x7d,
	0xab, 0x78, 0xa8, 0xe3, 0x9b, 0xef, 0x0d, 0x3c, 0x88, 0xd8, 0xd6, 0xb9, 0x65, 0x33, 0x9a, 0x34,
	0x81, 0x5d, 0x49, 0x5d, 0x67, 0xee, 0xf4, 0x7a, 0xd3, 0x30, 0xe2, 0x01, 0xc2, 0x86, 0x89, 0xc6,
	0xb6, 0x6f, 0x94, 0x9f, 0x04, 0x59, 0x94, 0x9e, 0x20, 0x5e, 0x8d, 0xc1, 0x9a, 0x7e, 0x20, 0x8c,
	0x41, 0x32, 0xd9, 0x6e, 0x72, 0xbb, 0xef, 0xf6, 0x1b, 0x88, 0x42, 0x33, 0x02, 0x04, 0xed, 0x73,
	0x79, 0x91, 0xb8, 0x05, 0xa1, 0xad, 0x56, 0x41, 0x22, 0xb2, 0xc8, 0x81, 0xaf, 0x95, 0x98, 0xd5,
	0xd2, 0x9b, 0xe1, 0xa1, 0x36, 0x0e, 0x64, 0xe9, 0x15, 0x12, 0x84, 0xe9, 0xe9, 0x72, 0x5d, 0x13,
	0x2a, 0xe2, 0x4f, 0x5b, 0x13, 0x9e, 0x21, 0xac, 0xf4, 0xb4, 0x9c, 0x30, 0xce, 0x7f, 0xeb, 0x8e,
	0xfb, 0xfc, 0x2f, 0xe5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x0f, 0xaf, 0xd6, 0xc9, 0x00, 0x00,
	0x00,
}
