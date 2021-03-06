// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ogay.proto

package ogay_grpc

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

type TranslationRequest struct {
	EnglishWord          string   `protobuf:"bytes,1,opt,name=englishWord,proto3" json:"englishWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationRequest) Reset()         { *m = TranslationRequest{} }
func (m *TranslationRequest) String() string { return proto.CompactTextString(m) }
func (*TranslationRequest) ProtoMessage()    {}
func (*TranslationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ogay_1adf9f52546c1375, []int{0}
}
func (m *TranslationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationRequest.Unmarshal(m, b)
}
func (m *TranslationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationRequest.Marshal(b, m, deterministic)
}
func (dst *TranslationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationRequest.Merge(dst, src)
}
func (m *TranslationRequest) XXX_Size() int {
	return xxx_messageInfo_TranslationRequest.Size(m)
}
func (m *TranslationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationRequest proto.InternalMessageInfo

func (m *TranslationRequest) GetEnglishWord() string {
	if m != nil {
		return m.EnglishWord
	}
	return ""
}

type TranslationResponse struct {
	PigLatinWord         string   `protobuf:"bytes,1,opt,name=pigLatinWord,proto3" json:"pigLatinWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslationResponse) Reset()         { *m = TranslationResponse{} }
func (m *TranslationResponse) String() string { return proto.CompactTextString(m) }
func (*TranslationResponse) ProtoMessage()    {}
func (*TranslationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ogay_1adf9f52546c1375, []int{1}
}
func (m *TranslationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslationResponse.Unmarshal(m, b)
}
func (m *TranslationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslationResponse.Marshal(b, m, deterministic)
}
func (dst *TranslationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslationResponse.Merge(dst, src)
}
func (m *TranslationResponse) XXX_Size() int {
	return xxx_messageInfo_TranslationResponse.Size(m)
}
func (m *TranslationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TranslationResponse proto.InternalMessageInfo

func (m *TranslationResponse) GetPigLatinWord() string {
	if m != nil {
		return m.PigLatinWord
	}
	return ""
}

func init() {
	proto.RegisterType((*TranslationRequest)(nil), "ogay_grpc.TranslationRequest")
	proto.RegisterType((*TranslationResponse)(nil), "ogay_grpc.TranslationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslatorClient interface {
	TranslateWord(ctx context.Context, in *TranslationRequest, opts ...grpc.CallOption) (*TranslationResponse, error)
}

type translatorClient struct {
	cc *grpc.ClientConn
}

func NewTranslatorClient(cc *grpc.ClientConn) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) TranslateWord(ctx context.Context, in *TranslationRequest, opts ...grpc.CallOption) (*TranslationResponse, error) {
	out := new(TranslationResponse)
	err := c.cc.Invoke(ctx, "/ogay_grpc.Translator/TranslateWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
type TranslatorServer interface {
	TranslateWord(context.Context, *TranslationRequest) (*TranslationResponse, error)
}

func RegisterTranslatorServer(s *grpc.Server, srv TranslatorServer) {
	s.RegisterService(&_Translator_serviceDesc, srv)
}

func _Translator_TranslateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).TranslateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ogay_grpc.Translator/TranslateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).TranslateWord(ctx, req.(*TranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Translator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ogay_grpc.Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranslateWord",
			Handler:    _Translator_TranslateWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ogay.proto",
}

func init() { proto.RegisterFile("ogay.proto", fileDescriptor_ogay_1adf9f52546c1375) }

var fileDescriptor_ogay_1adf9f52546c1375 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4f, 0x4f, 0xac,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xe3, 0xd3, 0x8b, 0x0a, 0x92, 0x95,
	0xcc, 0xb8, 0x84, 0x42, 0x8a, 0x12, 0xf3, 0x8a, 0x73, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x14, 0xb8, 0xb8, 0x53, 0xf3, 0xd2, 0x73, 0x32, 0x8b, 0x33,
	0xc2, 0xf3, 0x8b, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x85, 0x94, 0x2c, 0xb9,
	0x84, 0x51, 0xf4, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29, 0x71, 0xf1, 0x14, 0x64, 0xa6,
	0xfb, 0x24, 0x96, 0x64, 0xe6, 0x21, 0xe9, 0x44, 0x11, 0x33, 0x8a, 0xe3, 0xe2, 0x82, 0x69, 0xcd,
	0x2f, 0x12, 0x0a, 0xe0, 0xe2, 0x85, 0xf1, 0x52, 0x41, 0xd2, 0x42, 0xb2, 0x7a, 0x70, 0xd7, 0xe9,
	0x61, 0x3a, 0x4d, 0x4a, 0x0e, 0x97, 0x34, 0xc4, 0x05, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x4f, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x6c, 0x69, 0x8e, 0xf2, 0x00, 0x00, 0x00,
}
