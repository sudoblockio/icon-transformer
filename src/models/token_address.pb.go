// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token_address.proto

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

type TokenAddress struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	TokenContractAddress string   `protobuf:"bytes,2,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address"`
	Balance              float64  `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance"`
}

func (m *TokenAddress) Reset()         { *m = TokenAddress{} }
func (m *TokenAddress) String() string { return proto.CompactTextString(m) }
func (*TokenAddress) ProtoMessage()    {}
func (*TokenAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7e686e5317797f7, []int{0}
}

func (m *TokenAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenAddress.Unmarshal(m, b)
}
func (m *TokenAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenAddress.Marshal(b, m, deterministic)
}
func (m *TokenAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAddress.Merge(m, src)
}
func (m *TokenAddress) XXX_Size() int {
	return xxx_messageInfo_TokenAddress.Size(m)
}
func (m *TokenAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAddress.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAddress proto.InternalMessageInfo

func (m *TokenAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TokenAddress) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *TokenAddress) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenAddress)(nil), "models.TokenAddress")
}

func init() {
	proto.RegisterFile("token_address.proto", fileDescriptor_b7e686e5317797f7)
}

var fileDescriptor_b7e686e5317797f7 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4e, 0x84, 0x30,
	0x14, 0xc6, 0xf3, 0xd0, 0x30, 0xf8, 0xe2, 0x0a, 0x8d, 0xa1, 0x1a, 0x93, 0x71, 0x56, 0x5d, 0x38,
	0x34, 0xd1, 0x9d, 0x6e, 0x14, 0x6f, 0xd0, 0xb8, 0x72, 0x43, 0xa0, 0x54, 0x86, 0x08, 0x7d, 0x04,
	0x6a, 0x32, 0x57, 0x83, 0xdb, 0x78, 0x13, 0x33, 0x14, 0x16, 0xc6, 0x59, 0xf6, 0xeb, 0xf7, 0xe7,
	0xd7, 0xe2, 0x85, 0xa5, 0x2f, 0x6d, 0xd2, 0xac, 0x28, 0x3a, 0xdd, 0xf7, 0x71, 0xdb, 0x91, 0xa5,
	0xd0, 0x6f, 0xa8, 0xd0, 0x75, 0x7f, 0xfd, 0x52, 0x56, 0x76, 0xf7, 0x9d, 0xc7, 0x8a, 0x1a, 0x51,
	0x99, 0x4f, 0xca, 0x6b, 0xda, 0x53, 0xab, 0x8d, 0x98, 0x6c, 0x6a, 0x5b, 0x6a, 0xb3, 0x2d, 0xa9,
	0x6b, 0xdc, 0x59, 0x50, 0x6b, 0x2b, 0x32, 0xbd, 0x38, 0x48, 0xae, 0x69, 0xf3, 0x03, 0x78, 0xfe,
	0x7e, 0x58, 0x78, 0x75, 0x03, 0xe1, 0x06, 0x57, 0xf3, 0x56, 0x04, 0x6b, 0xe0, 0x67, 0x49, 0x30,
	0x0e, 0xec, 0x14, 0x3d, 0x0e, 0x72, 0xb9, 0x08, 0x77, 0x78, 0xe5, 0xa8, 0x14, 0x19, 0xdb, 0x65,
	0xca, 0x2e, 0x78, 0x91, 0x37, 0x45, 0x1e, 0xc6, 0x81, 0xc5, 0x78, 0xcf, 0x41, 0xf2, 0x3f, 0xf4,
	0x69, 0x55, 0xec, 0xd3, 0xe3, 0x49, 0x79, 0x39, 0xe9, 0x6f, 0xb3, 0xbc, 0xd0, 0x3c, 0xe3, 0x2a,
	0xcf, 0xea, 0xcc, 0x28, 0x1d, 0x9d, 0xac, 0x81, 0x43, 0x72, 0x37, 0x0e, 0xec, 0x16, 0x6f, 0x24,
	0xfb, 0x5f, 0x3c, 0x1b, 0xe5, 0x92, 0x78, 0xf2, 0xc7, 0x81, 0x79, 0x01, 0x24, 0xf8, 0x11, 0xc4,
	0xc2, 0xfd, 0x58, 0xee, 0x4f, 0xcf, 0x7e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x67, 0x4c, 0x2f,
	0x03, 0x57, 0x01, 0x00, 0x00,
}
