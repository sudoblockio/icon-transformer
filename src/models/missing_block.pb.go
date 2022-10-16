// Code generated by protoc-gen-go. DO NOT EDIT.
// source: missing_block.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type MissingBlock struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number"`
}

func (m *MissingBlock) Reset()         { *m = MissingBlock{} }
func (m *MissingBlock) String() string { return proto.CompactTextString(m) }
func (*MissingBlock) ProtoMessage()    {}
func (*MissingBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd13a7fa3e989aad, []int{0}
}

func (m *MissingBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissingBlock.Unmarshal(m, b)
}
func (m *MissingBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissingBlock.Marshal(b, m, deterministic)
}
func (m *MissingBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissingBlock.Merge(m, src)
}
func (m *MissingBlock) XXX_Size() int {
	return xxx_messageInfo_MissingBlock.Size(m)
}
func (m *MissingBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MissingBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MissingBlock proto.InternalMessageInfo

func (m *MissingBlock) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*MissingBlock)(nil), "models.MissingBlock")
}

func init() {
	proto.RegisterFile("missing_block.proto", fileDescriptor_cd13a7fa3e989aad)
}

var fileDescriptor_cd13a7fa3e989aad = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcd, 0x2c, 0x2e,
	0xce, 0xcc, 0x4b, 0x8f, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xcb, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x96, 0x72, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcc, 0x4b, 0xcb, 0x4f, 0xca, 0xc9, 0xaf, 0xc8, 0x2f, 0x48, 0xcd,
	0xd3, 0x07, 0x2b, 0x4b, 0xd6, 0x4d, 0x4f, 0xcd, 0xd3, 0x4d, 0xcf, 0x2f, 0xca, 0x85, 0xf0, 0xf5,
	0xf3, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x41, 0x42, 0x10, 0x93, 0x94, 0x2c, 0xb8, 0x78,
	0x7c, 0x21, 0x16, 0x38, 0x81, 0xcc, 0x17, 0x52, 0xe0, 0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x76, 0xe2, 0xd8, 0xb5, 0x53, 0x92, 0x85, 0x8b, 0x49, 0x83,
	0x31, 0x08, 0x2a, 0x6e, 0xc5, 0xb6, 0x6b, 0xa7, 0x24, 0x13, 0x07, 0xa3, 0x13, 0x57, 0x14, 0x87,
	0x9e, 0x3e, 0xc4, 0x1d, 0x49, 0x6c, 0x60, 0xc3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x2c, 0x48, 0x66, 0xad, 0x00, 0x00, 0x00,
}
