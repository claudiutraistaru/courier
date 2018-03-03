// Code generated by protoc-gen-go. DO NOT EDIT.
// source: courier.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	courier.proto

It has these top-level messages:
	SendBody
	Response
*/
package pb

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_Success Status = 0
	Status_Fail    Status = 1
)

var Status_name = map[int32]string{
	0: "Success",
	1: "Fail",
}
var Status_value = map[string]int32{
	"Success": 0,
	"Fail":    1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SendBody struct {
	Sender     string   `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Recipients []string `protobuf:"bytes,2,rep,name=recipients" json:"recipients,omitempty"`
	Subject    string   `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	BodyText   string   `protobuf:"bytes,4,opt,name=bodyText" json:"bodyText,omitempty"`
	BodyType   string   `protobuf:"bytes,5,opt,name=bodyType" json:"bodyType,omitempty"`
}

func (m *SendBody) Reset()                    { *m = SendBody{} }
func (m *SendBody) String() string            { return proto.CompactTextString(m) }
func (*SendBody) ProtoMessage()               {}
func (*SendBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendBody) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SendBody) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *SendBody) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendBody) GetBodyText() string {
	if m != nil {
		return m.BodyText
	}
	return ""
}

func (m *SendBody) GetBodyType() string {
	if m != nil {
		return m.BodyType
	}
	return ""
}

type Response struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=courier.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Success
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SendBody)(nil), "courier.SendBody")
	proto.RegisterType((*Response)(nil), "courier.Response")
	proto.RegisterEnum("courier.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mail service

type MailClient interface {
	Send(ctx context.Context, in *SendBody, opts ...grpc.CallOption) (*Response, error)
}

type mailClient struct {
	cc *grpc.ClientConn
}

func NewMailClient(cc *grpc.ClientConn) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) Send(ctx context.Context, in *SendBody, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/courier.Mail/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mail service

type MailServer interface {
	Send(context.Context, *SendBody) (*Response, error)
}

func RegisterMailServer(s *grpc.Server, srv MailServer) {
	s.RegisterService(&_Mail_serviceDesc, srv)
}

func _Mail_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Mail/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).Send(ctx, req.(*SendBody))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "courier.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Mail_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courier.proto",
}

func init() { proto.RegisterFile("courier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x49, 0x6b, 0x92, 0xf4, 0x21, 0xa0, 0xbc, 0x01, 0x59, 0x1d, 0xa0, 0xea, 0x42, 0x85,
	0x50, 0x87, 0xc2, 0x17, 0x74, 0x60, 0xeb, 0x92, 0x30, 0xb1, 0x25, 0xce, 0x13, 0x32, 0x2a, 0xb1,
	0xe5, 0xe7, 0x48, 0xe4, 0x3f, 0xf8, 0x60, 0x54, 0xd7, 0x69, 0x3b, 0x9e, 0x7b, 0x6d, 0xf9, 0xf8,
	0xc2, 0xb5, 0x32, 0x9d, 0xd3, 0xe4, 0x56, 0xd6, 0x19, 0x6f, 0x30, 0x8b, 0xb8, 0xf8, 0x4b, 0x20,
	0x2f, 0xa9, 0x6d, 0x36, 0xa6, 0xe9, 0xf1, 0x1e, 0x52, 0xa6, 0xb6, 0x21, 0x27, 0x93, 0x79, 0xb2,
	0x9c, 0x14, 0x91, 0xf0, 0x01, 0xc0, 0x91, 0xd2, 0x56, 0x53, 0xeb, 0x59, 0x8e, 0xe6, 0xe3, 0xe5,
	0xa4, 0x38, 0x4b, 0x50, 0x42, 0xc6, 0x5d, 0xfd, 0x4d, 0xca, 0xcb, 0x71, 0xb8, 0x38, 0x20, 0xce,
	0x20, 0xaf, 0x4d, 0xd3, 0x7f, 0xd0, 0xaf, 0x97, 0x22, 0x54, 0x47, 0x3e, 0x76, 0xbd, 0x25, 0x79,
	0x79, 0xd6, 0xf5, 0x96, 0x16, 0x5b, 0xc8, 0x0b, 0x62, 0x6b, 0x5a, 0x26, 0x7c, 0x82, 0x94, 0x7d,
	0xe5, 0x3b, 0x0e, 0x56, 0x37, 0xeb, 0xdb, 0xd5, 0xf0, 0x97, 0x32, 0xc4, 0x45, 0xac, 0xf7, 0x1a,
	0x3f, 0xc4, 0x5c, 0x7d, 0x91, 0x1c, 0x1d, 0x34, 0x22, 0x3e, 0x3f, 0x42, 0x7a, 0x38, 0x8b, 0x57,
	0x90, 0x95, 0x9d, 0x52, 0xc4, 0x3c, 0xbd, 0xc0, 0x1c, 0xc4, 0x7b, 0xa5, 0x77, 0xd3, 0x64, 0xfd,
	0x06, 0x62, 0x5b, 0xe9, 0x1d, 0xbe, 0x80, 0xd8, 0xaf, 0x81, 0x77, 0xa7, 0x37, 0xe2, 0x38, 0xb3,
	0x53, 0x34, 0x98, 0x6d, 0xc4, 0xe7, 0xc8, 0xd6, 0x75, 0x1a, 0x26, 0x7d, 0xfd, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0xbb, 0xef, 0xea, 0x63, 0x01, 0x00, 0x00,
}
