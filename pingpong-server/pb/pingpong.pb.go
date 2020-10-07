// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pingpong.proto

package pingpong

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PingRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PongResponse struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{1}
}

func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (m *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(m, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "pingpong.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "pingpong.PongResponse")
}

func init() { proto.RegisterFile("pingpong.proto", fileDescriptor_1cfbf639ab46154b) }

var fileDescriptor_1cfbf639ab46154b = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0xcc, 0x4b,
	0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x14,
	0xb9, 0xb8, 0x03, 0x32, 0xf3, 0xd2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8,
	0x58, 0x40, 0x52, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x12, 0x17, 0x4f,
	0x40, 0x3e, 0x48, 0x49, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x58, 0x4d, 0x3e, 0x92, 0x9a, 0xfc,
	0xbc, 0x74, 0x23, 0x1f, 0x2e, 0x7e, 0x90, 0x31, 0x20, 0x75, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9,
	0xa9, 0x42, 0x96, 0x5c, 0x60, 0x5b, 0x40, 0x42, 0x42, 0xa2, 0x7a, 0x70, 0x07, 0x20, 0xd9, 0x26,
	0x25, 0x86, 0x24, 0x8c, 0x64, 0x43, 0x12, 0x1b, 0xd8, 0x95, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0xf0, 0x3a, 0xda, 0xb7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingPongServiceClient is the client API for PingPongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongServiceClient interface {
	PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type pingPongServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingPongServiceClient(cc *grpc.ClientConn) PingPongServiceClient {
	return &pingPongServiceClient{cc}
}

func (c *pingPongServiceClient) PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/pingpong.PingPongService/pingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServiceServer is the server API for PingPongService service.
type PingPongServiceServer interface {
	PingPong(context.Context, *PingRequest) (*PongResponse, error)
}

// UnimplementedPingPongServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPingPongServiceServer struct {
}

func (*UnimplementedPingPongServiceServer) PingPong(ctx context.Context, req *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterPingPongServiceServer(s *grpc.Server, srv PingPongServiceServer) {
	s.RegisterService(&_PingPongService_serviceDesc, srv)
}

func _PingPongService_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServiceServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPongService/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServiceServer).PingPong(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPongService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPongService",
	HandlerType: (*PingPongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pingPong",
			Handler:    _PingPongService_PingPong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong.proto",
}
