// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FightTypes.txt

package DTO

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

type FightTypes int32

const (
	FightTypes_MOVE_CREQ   FightTypes = 0
	FightTypes_SKILL_CREQ  FightTypes = 1
	FightTypes_INFORM_SRES FightTypes = 2
)

var FightTypes_name = map[int32]string{
	0: "MOVE_CREQ",
	1: "SKILL_CREQ",
	2: "INFORM_SRES",
}

var FightTypes_value = map[string]int32{
	"MOVE_CREQ":   0,
	"SKILL_CREQ":  1,
	"INFORM_SRES": 2,
}

func (x FightTypes) String() string {
	return proto.EnumName(FightTypes_name, int32(x))
}

func (FightTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e0ea92a0d687334, []int{0}
}

func init() {
	proto.RegisterEnum("FightTypes", FightTypes_name, FightTypes_value)
}

func init() { proto.RegisterFile("FightTypes.txt", fileDescriptor_7e0ea92a0d687334) }

var fileDescriptor_7e0ea92a0d687334 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0xcb, 0x4c, 0xcf,
	0x28, 0x09, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0xa9, 0x28, 0xd1, 0xb2, 0xe1, 0xe2, 0x42, 0x88,
	0x08, 0xf1, 0x72, 0x71, 0xfa, 0xfa, 0x87, 0xb9, 0xc6, 0x3b, 0x07, 0xb9, 0x06, 0x0a, 0x30, 0x08,
	0xf1, 0x71, 0x71, 0x05, 0x7b, 0x7b, 0xfa, 0xf8, 0x40, 0xf8, 0x8c, 0x42, 0xfc, 0x5c, 0xdc, 0x9e,
	0x7e, 0x6e, 0xfe, 0x41, 0xbe, 0xf1, 0xc1, 0x41, 0xae, 0xc1, 0x02, 0x4c, 0x49, 0x6c, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x50, 0xb9, 0x0d, 0x56, 0x00,
	0x00, 0x00,
}