// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google_protobuf.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	google_protobuf.proto

It has these top-level messages:
	ProtoColorGroup
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoColorGroup struct {
	Id               *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Colors           []string `protobuf:"bytes,3,rep,name=colors" json:"colors,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoColorGroup) Reset()                    { *m = ProtoColorGroup{} }
func (m *ProtoColorGroup) String() string            { return proto.CompactTextString(m) }
func (*ProtoColorGroup) ProtoMessage()               {}
func (*ProtoColorGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoColorGroup) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ProtoColorGroup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProtoColorGroup) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoColorGroup)(nil), "main.ProtoColorGroup")
}

func init() { proto.RegisterFile("google_protobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0x8d, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x03, 0x33, 0x84, 0x58,
	0x72, 0x13, 0x33, 0xf3, 0x94, 0x7c, 0xb9, 0xf8, 0x03, 0x40, 0x5c, 0xe7, 0xfc, 0x9c, 0xfc, 0x22,
	0xf7, 0xa2, 0xfc, 0xd2, 0x02, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x26, 0x0d,
	0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05,
	0x26, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x8c, 0x8b, 0x2d, 0x19, 0xa4, 0xa3, 0x58, 0x82, 0x59,
	0x81, 0x59, 0x83, 0x33, 0x08, 0xca, 0x03, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x9b, 0x81, 0xd6,
	0x6c, 0x00, 0x00, 0x00,
}
