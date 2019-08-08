// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dto/AnyDTO.txt

package AnyDTO

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AnyDTO struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnyDTO) Reset()         { *m = AnyDTO{} }
func (m *AnyDTO) String() string { return proto.CompactTextString(m) }
func (*AnyDTO) ProtoMessage()    {}
func (*AnyDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64ac3bd14d3563d, []int{0}
}

func (m *AnyDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnyDTO.Unmarshal(m, b)
}
func (m *AnyDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnyDTO.Marshal(b, m, deterministic)
}
func (m *AnyDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyDTO.Merge(m, src)
}
func (m *AnyDTO) XXX_Size() int {
	return xxx_messageInfo_AnyDTO.Size(m)
}
func (m *AnyDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyDTO.DiscardUnknown(m)
}

var xxx_messageInfo_AnyDTO proto.InternalMessageInfo

func (m *AnyDTO) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AnyDTO) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AnyDTO)(nil), "AnyDTO")
}

func init() { proto.RegisterFile("proto/dto/AnyDTO.txt", fileDescriptor_e64ac3bd14d3563d) }

var fileDescriptor_e64ac3bd14d3563d = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x29, 0xc9, 0xd7, 0x77, 0xcc, 0xab, 0x74, 0x09, 0xf1, 0xd7, 0x2b, 0xa9, 0x28,
	0x51, 0x32, 0xe3, 0x62, 0x83, 0xf0, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4,
	0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0x9a, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xf8, 0x90, 0x68, 0x56, 0x00, 0x00, 0x00,
}
